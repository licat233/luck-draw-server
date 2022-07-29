package logic

import (
	"context"
	"errors"
	"fmt"
	"luck-draw/service/v1/luckdraw/rpc/internal/svc"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"
	model2 "luck-draw/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
	"luck-draw/common/errorx"
	"luck-draw/common/utils"
)

type SelectPrizesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSelectPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectPrizesLogic {
	return &SelectPrizesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/** SelectPrizes 用户选择要领取的奖品
 */
func (l *SelectPrizesLogic) SelectPrizes(in *luckdraw_pb.SelectPrizesReq) (*luckdraw_pb.SelectPrizesResp, error) {
	if len(in.LuckerUuids) == 0 {
		return nil, errorx.ExternalError(errors.New("LuckerUuidList len = 0"))
	}
	//查询是否存在该活动
	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil && err != model2.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("SelectPrizes ActivityModel.FindOneByUUID err:%+v", err))
	}
	if err == model2.ErrNotFound {
		return nil, errorx.ExternalError(errors.New("activity not found"))
	}
	if activity.Uuid != in.ActivityUuid {
		return nil, errorx.ExternalError(errors.New("Activity UUID does not match activity ID"))
	}

	//取出中奖名单
	luckerList, err := l.svcCtx.LuckerModel.FindByUserAndActivity(l.ctx, in.UserId, activity.Id)
	if err != nil && err != model2.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("SelectPrizes LuckerModel.FindByUserAndActivity err:%+v", err))
	}
	if err == model2.ErrNotFound {
		return nil, errorx.ExternalError(errors.New("lucker not found"))
	}
	//uuids去重
	inUuidList := utils.RemoveDuplicatesForString(in.LuckerUuids)

	//已领取的奖品
	var selectedLuckerList []*model2.Lucker   //已领取的奖品
	var noSelectedLuckerList []*model2.Lucker //未领取的奖品
	var mdUuidList []string
	var selectedUuidList []string
	for _, lucker := range luckerList {
		if lucker.Status == 1 {
			selectedLuckerList = append(selectedLuckerList, lucker)
			mdUuidList = append(mdUuidList, lucker.Uuid)
			selectedUuidList = append(selectedUuidList, lucker.Uuid)
		} else if lucker.Status == 0 {
			noSelectedLuckerList = append(noSelectedLuckerList, lucker)
			mdUuidList = append(mdUuidList, lucker.Uuid)
		}
	}

	if len(selectedLuckerList) >= int(activity.MaxPrizes) {
		return nil, errorx.ExternalError(nil, errors.New("The user's prizes in the current activity have been collected"))
	}

	//校验是否有多余的uuid传进来
	if utils.SliceContainString(mdUuidList, inUuidList...) == false {
		return nil, errorx.ExternalError(errors.New("LuckerUuidList存在多余的uuid"))
	}

	//此时inUuidList内可能含有已经领取过的uuid，也有可能全部是已经领取过得uuid
	//合并
	inUuidList = append(inUuidList, selectedUuidList...)
	//去重
	inUuidList = utils.RemoveDuplicatesForString(inUuidList)
	//判断是否会超额领取
	if len(inUuidList) > int(activity.MaxPrizes) {
		return nil, errorx.ExternalError(nil, errors.New("The number of prizes received exceeds the activity limit"))
	}

	//过滤已经未领取的奖品
	uuidList := utils.SliceFilterString(inUuidList, selectedUuidList...)
	if len(uuidList) == 0 {
		return nil, errorx.ExternalError(nil, errors.New("There is no prize to claim"))
	}

	//更新领取状态
	err = l.svcCtx.LuckerModel.UpdateSomeStatusByUuid(l.ctx, uuidList, 1)
	if err != nil {
		return nil, errorx.InternalError(fmt.Errorf("SelectPrizes LuckerModel.UpdateSomeStatusByUuid err:%+v", err))
	}

	//如果领取的奖品已达到限额值，其它奖品则为弃领状态
	//取出中奖名单
	newestLuckerList, err := l.svcCtx.LuckerModel.FindByUserAndActivity(l.ctx, in.UserId, activity.Id)
	if err != nil && err != model2.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("SelectPrizes LuckerModel.FindByUserAndActivity err:%+v", err))
	}
	if err == model2.ErrNotFound {
		return nil, errorx.ExternalError(nil, errors.New("There is no prize to claim"))
	}

	var selectNum int64
	var noIds []int64
	for _, lucker := range newestLuckerList {
		if lucker.Status == 1 {
			selectNum++
		} else if lucker.Status == 0 {
			noIds = append(noIds, lucker.Id)
		}
	}
	//如果已经领取完毕，其它奖品就等于弃领了
	if selectNum >= activity.MaxPrizes {
		err = l.svcCtx.LuckerModel.UpdateSomeStatusById(l.ctx, noIds, -1)
		if err != nil {
			return nil, errorx.InternalError(fmt.Errorf("SelectPrizes LuckerModel.UpdateSomeStatusById err:%+v", err))
		}
	}

	//查询用户已领取的中奖产品
	mdPrizeList, err := l.svcCtx.PrizeModel.FindsByUserAndActivityAndStatus(l.ctx, in.UserId, activity.Id, 1)
	if err != nil && err != model2.ErrNotFound {
		return nil, errorx.InternalError(fmt.Errorf("SelectPrizes PrizeModel.FindsByUserAndActivityAndStatus err:%+v", err))
	}

	pbPrizes := convertPrizes(mdPrizeList)
	return &luckdraw_pb.SelectPrizesResp{
		Prizes: pbPrizes,
	}, nil
}
