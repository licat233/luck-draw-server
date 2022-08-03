package logic

import (
	"context"
	"fmt"
	"luck-draw/common/errorx"
	"luck-draw/common/utils"
	"luck-draw/service/v1/luckdraw/rpc/luckdraw_pb"
	"luck-draw/service/v1/model"
)

/** SearchUser 搜索用户
 * 返回的error均为 errorx 错误
 * 先匹配passport，然后再去匹配 email ,再去匹配 phone
 */
func SearchUser(ctx context.Context, UserModel model.UserModel, passport string) (*model.User, error) {
	mdUser, err := UserModel.FindByPasspord(ctx, passport)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("SearchUser  UserModel.FindByPasspord error, err:%+v", err))
	}
	if err == nil {
		return mdUser, nil
	}

	mdUser, err = UserModel.FindByEmail(ctx, passport)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("SearchUser UserModel.FindByEmail error, err:%+v", err))
	}
	if err == nil {
		return mdUser, nil
	}

	mdUser, err = UserModel.FindByPhone(ctx, passport)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntStatusErr(fmt.Errorf("SearchUser UserModel.FindByPhone error, err:%+v", err))
	}
	if err == nil {
		return mdUser, nil
	}

	if err == model.ErrNotFound {
		return nil, nil
	}
	return mdUser, nil
}

func convertActivity(mdActivity *model.Activity) *luckdraw_pb.Activity {
	if mdActivity == nil {
		return nil
	}
	return &luckdraw_pb.Activity{
		Id:        mdActivity.Id,
		Uuid:      mdActivity.Uuid,
		Name:      mdActivity.Name,
		ProjectId: mdActivity.ProjectId,
		Status:    utils.Int64ToBool(mdActivity.Status),
		Pattern:   int32(mdActivity.Pattern),
		MaxCount:  mdActivity.MaxCount,
		MaxPrizes: mdActivity.MaxPrizes,
		StartTime: mdActivity.StartTime.Unix(),
		EndTime:   mdActivity.EndTime.Unix(),
	}
}

func convertPrize(mdPrize *model.Prize) *luckdraw_pb.Prize {
	if mdPrize == nil {
		return nil
	}
	return &luckdraw_pb.Prize{
		Id:     mdPrize.Id,
		Uuid:   mdPrize.Uuid,
		Name:   mdPrize.Name,
		Image:  mdPrize.Image,
		Price:  mdPrize.Price,
		Status: int32(mdPrize.Status),
	}
}

func convertPrizesPool(mdPrizesPool *model.PrizesPool) *luckdraw_pb.PrizesPool {
	if mdPrizesPool == nil {
		return nil
	}
	return &luckdraw_pb.PrizesPool{
		Id:         mdPrizesPool.Id,
		Uuid:       mdPrizesPool.Uuid,
		ActivityId: mdPrizesPool.ActivityId,
		PrizeId:    mdPrizesPool.PrizeId,
		LuckValue:  mdPrizesPool.LuckValue,
		Status:     int32(mdPrizesPool.Status),
		GradeName:  mdPrizesPool.GradeName,
		IsWin:      utils.Int64ToBool(mdPrizesPool.IsWin),
		Number:     mdPrizesPool.Number,
		Remain:     mdPrizesPool.Remain,
	}
}

func convertUser(mdUser *model.User) *luckdraw_pb.User {
	if mdUser == nil {
		return nil
	}
	return &luckdraw_pb.User{
		Id:       mdUser.Id,
		Uuid:     mdUser.Uuid,
		Passport: mdUser.Passport,
		Email:    mdUser.Email,
		Phone:    mdUser.Phone,
		Name:     mdUser.Name,
		Address:  mdUser.Address,
		Score:    mdUser.Score,
		Status:   int32(mdUser.Status),
	}
}

func convertLucker(mdLucker *model.Lucker) *luckdraw_pb.Lucker {
	if mdLucker == nil {
		return nil
	}
	return &luckdraw_pb.Lucker{
		Id:           mdLucker.Id,
		Uuid:         mdLucker.Uuid,
		UserId:       mdLucker.UserId,
		PrizesPoolId: mdLucker.PrizesPoolId,
		Status:       int32(mdLucker.Status),
		LuckDate:     mdLucker.LuckDate.Unix(),
		Ip:           mdLucker.Ip,
	}
}

func convertPrizes(mPrizes []*model.Prize) (pbPrizes []*luckdraw_pb.Prize) {
	if len(mPrizes) == 0 {
		return
	}
	for _, mPrize := range mPrizes {
		lprize := convertPrize(mPrize)
		if lprize == nil {
			continue
		}
		pbPrizes = append(pbPrizes, lprize)
	}
	return
}

func convertLuckerList(mdLuckerList []*model.Lucker) (pbLuckerlist []*luckdraw_pb.Lucker) {
	for _, mdLucker := range mdLuckerList {
		pbLucker := convertLucker(mdLucker)
		if pbLucker == nil {
			continue
		}
		pbLuckerlist = append(pbLuckerlist, pbLucker)
	}
	return
}

func convertPrizesPoolList(mdPrizesPoolList []*model.PrizesPool) (pbPrizesPoolList []*luckdraw_pb.PrizesPool) {
	for _, mdPrizesPool := range mdPrizesPoolList {
		pbPrizesPool := convertPrizesPool(mdPrizesPool)
		if pbPrizesPool == nil {
			continue
		}
		pbPrizesPoolList = append(pbPrizesPoolList, pbPrizesPool)
	}
	return
}

func convertLuckerDetail(mdLuckerDetail *model.LuckerDetail) *luckdraw_pb.LuckerDetail {
	if mdLuckerDetail == nil {
		return nil
	}
	return &luckdraw_pb.LuckerDetail{
		Lucker: &luckdraw_pb.Lucker{
			Id:           mdLuckerDetail.LuckerId,
			Uuid:         mdLuckerDetail.LuckerUuid,
			UserId:       mdLuckerDetail.UserId,
			PrizesPoolId: mdLuckerDetail.PrizesPoolId,
			Status:       int32(mdLuckerDetail.LuckerStatus),
			LuckDate:     mdLuckerDetail.LuckDate.Unix(),
		},
		PrizesPool: &luckdraw_pb.PrizesPool{
			Id:         mdLuckerDetail.PrizesPoolId,
			Uuid:       mdLuckerDetail.PrizesPoolUuid,
			ActivityId: mdLuckerDetail.ActivityId,
			PrizeId:    mdLuckerDetail.PrizeId,
			LuckValue:  mdLuckerDetail.PrizesPoolLuckValue,
			Status:     int32(mdLuckerDetail.PrizesPoolStatus),
			GradeName:  mdLuckerDetail.PrizesPoolGradeName,
			IsWin:      utils.Int64ToBool(mdLuckerDetail.PrizesPoolIsWin),
			Number:     mdLuckerDetail.PrizesPoolNumber,
			Remain:     mdLuckerDetail.PrizesPoolRemain,
		},
		Activity: &luckdraw_pb.Activity{
			Id:        mdLuckerDetail.ActivityId,
			Uuid:      mdLuckerDetail.ActivityUuid,
			Name:      mdLuckerDetail.ActivityName,
			ProjectId: mdLuckerDetail.ActivityProjectId,
			Status:    utils.Int64ToBool(mdLuckerDetail.ActivityStatus),
			Pattern:   int32(mdLuckerDetail.ActivityPattern),
			MaxCount:  mdLuckerDetail.ActivityMaxCount,
			MaxPrizes: mdLuckerDetail.ActivityMaxPrizes,
			StartTime: mdLuckerDetail.ActivityStartTime.Unix(),
			EndTime:   mdLuckerDetail.ActivityEndTime.Unix(),
		},
		Prize: &luckdraw_pb.Prize{
			Id:     mdLuckerDetail.PrizeId,
			Uuid:   mdLuckerDetail.PrizeUuid,
			Name:   mdLuckerDetail.PrizeName,
			Image:  mdLuckerDetail.PrizeImage,
			Price:  mdLuckerDetail.PrizePrice,
			Status: int32(mdLuckerDetail.PrizeStatus),
		},
	}
}

func convertLuckerDetailsList(mdLuckerDetailsList []*model.LuckerDetail) (res []*luckdraw_pb.LuckerDetail) {
	for _, mdluckerDetail := range mdLuckerDetailsList {
		pbLuckerDetail := convertLuckerDetail(mdluckerDetail)
		if pbLuckerDetail == nil {
			continue
		}
		res = append(res, pbLuckerDetail)
	}
	return
}
