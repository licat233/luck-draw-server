package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"luck-draw/service/v1/luckdraw/rpc/internal/svc"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"
	"luck-draw/service/v1/model"
	"math/rand"
	"sort"
	"time"

	"luck-draw/common/errorx"
	"luck-draw/common/kqueue"
	"luck-draw/common/uniqueid"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type GoodluckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGoodluckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodluckLogic {
	return &GoodluckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 *需要知道哪个用户(user_id)在抽奖，验证该用户状态(status)以及是否已经被删除了(is_deleted)
 *参加的是哪个活动(activity_id)的抽奖
 *查找抽奖名单(lucky)，看其抽奖次数是否已经达到活动(activity)的规定上限(max_count)
 *取出该活动的奖品(prizes)，用于抽奖goodluck
 *抽取到奖品(prize)以后，建立中奖信息(lucky)
 *用户(user_id)积分(score)增删操作，如果没有中奖，则给予 10 个积分奖励
 *最后返回中奖信息(lucky)
 */
func (l *GoodluckLogic) Goodluck(in *luckdraw_pb.GoodluckReq) (*luckdraw_pb.GoodluckResp, error) {
	if !govalidator.IsIP(in.Ip) {
		return nil, errorx.ExternalError(errors.New("request ip validator failed"))
	}
	//获取用户信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("Goodluck UserModel.Exists error, err:%+v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(fmt.Errorf("user not found, id: %d", in.UserId))
	}
	err = model.VerifyUser(user)
	if err != nil {
		return nil, errorx.ExternalError(err)
	}

	//获取活动信息
	mdActivity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("Goodluck ActivityModel.FindOne error, err:%+v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(fmt.Errorf("activity not found, id: %d", in.ActivityId))
	}
	if mdActivity.Uuid != in.ActivityUuid {
		return nil, errorx.ExternalError(errors.New("Activity UUID does not match activity ID"))
	}
	err = model.VerifyActivity(mdActivity)
	if err != nil {
		return nil, errorx.ExternalError(err)
	}

	//检查该用户在该活动的抽奖次数
	luckCount, err := l.svcCtx.LuckerModel.FindCountByUserAndActivity(l.ctx, in.UserId, mdActivity.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("Goodluck LuckerModel.FindCountByUserAndActivity error, err:%+v", err))
	}
	if luckCount >= mdActivity.MaxCount {
		return nil, errorx.ExternalError(nil, fmt.Errorf("你在[%s]活动中的抽奖次数已达到上限！", mdActivity.Name))
	}

	//获取活动奖品列表
	mdPrizesPoolList, err := l.svcCtx.PrizesPoolModel.FindByActivityUuid(l.ctx, in.ActivityUuid)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("Goodluck PrizesPoolModel.FindByActivityUuid err:%+v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(nil, errors.New("该活动的奖品池还没有放入奖品"))
	}
	pbPrizesPoolList := convertPrizesPoolList(mdPrizesPoolList)

	//抽取奖品
	PrizesPool, err := randomPrize(pbPrizesPoolList)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("Goodluck randomPrize errors err:%+v", err))
	}

	//建立中奖信息
	mdLucker := &model.Lucker{
		Uuid:         uniqueid.GenerateUUID(),
		UserId:       user.Id,
		PrizesPoolId: PrizesPool.Id,
		Status:       0,
		LuckDate:     time.Now(),
		Ip:           in.Ip,
	}
	_, err = l.svcCtx.LuckerModel.Insert(l.ctx, mdLucker)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("Goodluck LuckerModel.Insert err:%+v", err))
	}

	//积分奖励
	// 可以交由mq去处理
	// if PrizesPool.IsWin == false {
	// 	user.Score += 10
	// 	if err = l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
	// 		return nil, errorx.InternalError(fmt.Errorf("Goodluck UserModel.Update err:%+v", err))
	// 	}
	// }

	//交给mq去处理
	if err := l.kqUpdateUserScore(user.Id, PrizesPool.IsWin); err != nil {
		logx.Errorf("l.kqUpdateUserScore(user.Id, PrizesPool.IsWin) error: %v", err)
	}

	//获取奖品信息
	mdPrize, err := l.svcCtx.PrizeModel.FindOne(l.ctx, PrizesPool.PrizeId)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("Goodluck PrizeModel.FindOne err:%+v", err))
	}
	if err == model.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("Goodluck PrizesPool.PrizeId:%d not found, err:%+v", PrizesPool.PrizeId, err))
	}

	//返回中奖信息
	PrizesPool.LuckValue = 0
	return &luckdraw.GoodluckResp{
		User: convertUser(user),
		LuckerDetail: &luckdraw_pb.LuckerDetail{
			Lucker:     convertLucker(mdLucker),
			PrizesPool: PrizesPool,
			Activity:   convertActivity(mdActivity),
			Prize:      convertPrize(mdPrize),
		},
	}, nil
}

func (l *GoodluckLogic) kqUpdateUserScore(userId int64, isWin bool) error {
	m := kqueue.GoodluckStatusNotifyMessage{
		IsWin:  isWin,
		UserId: userId,
	}
	body, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}
	return l.svcCtx.KqueueGoodluckStatusHandlerClient.Push(string(body))
}

func randomPrize(pbPrizes []*luckdraw_pb.PrizesPool) (*luckdraw_pb.PrizesPool, error) {
	//注意：考虑某个傻子设置幸运值为负数，负数或者0的prize不参与抽奖，因为逻辑上完全不可能抽中
	if len(pbPrizes) == 0 {
		return nil, errors.New("randomPrize Error, len(pbPrizes) == 0")
	}
	var totalLuckValue int64
	var prizes []*luckdraw_pb.PrizesPool
	for _, prize := range pbPrizes {
		//如果机会 <=0 ，则不参与抽奖
		if prize.LuckValue <= 0 || prize.Number <= 0 {
			continue
		}
		totalLuckValue += prize.LuckValue * prize.Number
		for i := 0; i < int(prize.Number); i++ {
			prizes = append(prizes, prize)
		}
	}

	//如果全都不可能抽中
	if totalLuckValue == 0 {
		return nil, errors.New("randomPrize Error, totalLuckValue = 0")
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Int63n(totalLuckValue)
	temp := random

	sortPrizes(prizes)

	//抽取算法
	for _, prize := range prizes {
		if random < prize.LuckValue {
			return prize, nil
		}
		random -= prize.LuckValue
	}
	return nil, fmt.Errorf("randomPrize Error, totalChance=%d, random=%d", totalLuckValue, temp)
}

//升序排序
func sortPrizes(pbPrizes []*luckdraw_pb.PrizesPool) {
	sort.Slice(pbPrizes, func(i, j int) bool {
		return pbPrizes[i].LuckValue < pbPrizes[j].LuckValue
	})
}
