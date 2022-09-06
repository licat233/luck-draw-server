package logic

import (
	"errors"
	"luck_draw-server/service/v1/admin/rpc/admin_pb"
	"luck_draw-server/service/v1/model"
)

func newStatusResponse() *admin_pb.StatusResp {
	return &admin_pb.StatusResp{}
}

func mdActivityToPbActivity(mdActivity *model.Activity) *admin_pb.Activity {
	if mdActivity == nil {
		return nil
	}
	return &admin_pb.Activity{
		Id:        mdActivity.Id,
		Uuid:      mdActivity.Uuid,
		Name:      mdActivity.Name,
		ProjectId: mdActivity.ProjectId,
		Status:    mdActivity.Status,
		Pattern:   mdActivity.Pattern,
		MaxCount:  mdActivity.MaxCount,
		MaxPrizes: mdActivity.MaxPrizes,
		StartTime: mdActivity.StartTime.Unix(),
		EndTime:   mdActivity.EndTime.Unix(),
	}
}

func mdActivityListToPbActivityList(mdActivityList []*model.Activity) []*admin_pb.Activity {
	pbActivityList := make([]*admin_pb.Activity, 0, len(mdActivityList))
	for _, mdActivity := range mdActivityList {
		pbActivityList = append(pbActivityList, mdActivityToPbActivity(mdActivity))
	}
	return pbActivityList
}

func mdAdminerToPbAdminer(mdAdminer *model.Adminer) *admin_pb.Adminer {
	if mdAdminer == nil {
		return nil
	}
	return &admin_pb.Adminer{
		Id:           mdAdminer.Id,
		Uuid:         mdAdminer.Uuid,
		Name:         mdAdminer.Name,
		Avatar:       mdAdminer.Avatar,
		Passport:     mdAdminer.Passport,
		Password:     mdAdminer.Password,
		Email:        mdAdminer.Email,
		Status:       mdAdminer.Status,
		IsSuperAdmin: mdAdminer.IsSuperAdmin,
		LoginCount:   mdAdminer.LoginCount,
		LastLogin:    mdAdminer.LastLogin.Unix(),
	}
}

func mdAdminerListToPbAdminerList(mdAdminerList []*model.Adminer) []*admin_pb.Adminer {
	pbAdminerList := make([]*admin_pb.Adminer, 0, len(mdAdminerList))
	for _, mdAdminer := range mdAdminerList {
		pbAdminerList = append(pbAdminerList, mdAdminerToPbAdminer(mdAdminer))
	}
	return pbAdminerList
}

func mdLuckerToPbLucker(mdLucker *model.Lucker) *admin_pb.Lucker {
	if mdLucker == nil {
		return nil
	}
	return &admin_pb.Lucker{
		Id:           mdLucker.Id,
		Uuid:         mdLucker.Uuid,
		UserId:       mdLucker.UserId,
		PrizesPoolId: mdLucker.PrizesPoolId,
		Status:       mdLucker.Status,
		LuckDate:     mdLucker.LuckDate.Unix(),
		Ip:           mdLucker.Ip,
	}
}

func mdLuckerListToPbLuckerList(mdLuckerList []*model.Lucker) []*admin_pb.Lucker {
	pbLuckerList := make([]*admin_pb.Lucker, 0, len(mdLuckerList))
	for _, mdLucker := range mdLuckerList {
		pbLuckerList = append(pbLuckerList, mdLuckerToPbLucker(mdLucker))
	}
	return pbLuckerList
}

func mdPrizeToPbPrize(mdPrize *model.Prize) *admin_pb.Prize {
	if mdPrize == nil {
		return nil
	}
	return &admin_pb.Prize{
		Id:     mdPrize.Id,
		Uuid:   mdPrize.Uuid,
		Name:   mdPrize.Name,
		Image:  mdPrize.Image,
		Price:  mdPrize.Price,
		Status: mdPrize.Status,
	}
}

func mdPrizeListToPbPrizeList(mdPrizeList []*model.Prize) []*admin_pb.Prize {
	pbPrizeList := make([]*admin_pb.Prize, 0, len(mdPrizeList))
	for _, mdPrize := range mdPrizeList {
		pbPrizeList = append(pbPrizeList, mdPrizeToPbPrize(mdPrize))
	}
	return pbPrizeList
}

func mdPrizeGradeToPbPrizeGrade(mdPrizeGrade *model.PrizeGrade) *admin_pb.PrizeGrade {
	if mdPrizeGrade == nil {
		return nil
	}
	return &admin_pb.PrizeGrade{
		Id:    mdPrizeGrade.Id,
		Name:  mdPrizeGrade.Name,
		IsWin: 0,
	}
}

func mdPrizeGradeListToPbPrizeGradeList(mdPrizeGradeList []*model.PrizeGrade) []*admin_pb.PrizeGrade {
	pbPrizeGradeList := make([]*admin_pb.PrizeGrade, 0, len(mdPrizeGradeList))
	for _, mdPrizeGrade := range mdPrizeGradeList {
		pbPrizeGradeList = append(pbPrizeGradeList, mdPrizeGradeToPbPrizeGrade(mdPrizeGrade))
	}
	return pbPrizeGradeList
}

func mdPrizesPoolToPbPrizesPool(mdPrizesPool *model.PrizesPool) *admin_pb.PrizesPool {
	if mdPrizesPool == nil {
		return nil
	}
	return &admin_pb.PrizesPool{
		Id:         mdPrizesPool.Id,
		Uuid:       mdPrizesPool.Uuid,
		ActivityId: mdPrizesPool.ActivityId,
		PrizeId:    mdPrizesPool.PrizeId,
		GradeId:    mdPrizesPool.GradeId,
		LuckValue:  mdPrizesPool.LuckValue,
		Status:     mdPrizesPool.Status,
		Number:     mdPrizesPool.Number,
		Remainder:  mdPrizesPool.Remainder,
	}
}

func mdPrizesPoolListToPbPrizesPoolList(mdPrizesPoolList []*model.PrizesPool) []*admin_pb.PrizesPool {
	pbPrizesPoolList := make([]*admin_pb.PrizesPool, 0, len(mdPrizesPoolList))
	for _, mdPrizesPool := range mdPrizesPoolList {
		pbPrizesPoolList = append(pbPrizesPoolList, mdPrizesPoolToPbPrizesPool(mdPrizesPool))
	}
	return pbPrizesPoolList
}

func mdProjectToPbProject(mdProject *model.Project) *admin_pb.Project {
	if mdProject == nil {
		return nil
	}
	return &admin_pb.Project{
		Id:     mdProject.Id,
		Uuid:   mdProject.Uuid,
		Name:   mdProject.Name,
		Status: mdProject.Status,
	}
}

func mdProjectListToPbProjectList(mdProjectList []*model.Project) []*admin_pb.Project {
	pbProjectList := make([]*admin_pb.Project, 0, len(mdProjectList))
	for _, mdProject := range mdProjectList {
		pbProjectList = append(pbProjectList, mdProjectToPbProject(mdProject))
	}
	return pbProjectList
}

func mdUserToPbUser(mdUser *model.User) *admin_pb.User {
	if mdUser == nil {
		return nil
	}
	return &admin_pb.User{
		Id:       mdUser.Id,
		Uuid:     mdUser.Uuid,
		Passport: mdUser.Passport,
		Email:    mdUser.Email,
		Phone:    mdUser.Phone,
		Name:     mdUser.Name,
		Address:  mdUser.Address,
		Score:    mdUser.Status,
		Status:   mdUser.Status,
	}
}

func mdUserListToPbUserList(mdUserList []*model.User) []*admin_pb.User {
	pbUserList := make([]*admin_pb.User, 0, len(mdUserList))
	for _, mdUser := range mdUserList {
		pbUserList = append(pbUserList, mdUserToPbUser(mdUser))
	}
	return pbUserList
}

func mdUsersPoolToPbUsersPool(mdUsersPool *model.UsersPool) *admin_pb.UsersPool {
	if mdUsersPool == nil {
		return nil
	}
	return &admin_pb.UsersPool{
		Id:         mdUsersPool.Id,
		Uuid:       mdUsersPool.Uuid,
		ActivityId: mdUsersPool.ActivityId,
		UserId:     mdUsersPool.UserId,
		Status:     mdUsersPool.Status,
		Score:      mdUsersPool.Score,
	}
}

func mdUsersPoolListToPbUsersPoolList(mdUsersPoolList []*model.UsersPool) []*admin_pb.UsersPool {
	pbUsersPoolList := make([]*admin_pb.UsersPool, 0, len(mdUsersPoolList))
	for _, mdUsersPool := range mdUsersPoolList {
		pbUsersPoolList = append(pbUsersPoolList, mdUsersPoolToPbUsersPool(mdUsersPool))
	}
	return pbUsersPoolList
}

func verifyAdminerStatus(adminer *model.Adminer) error {
	if adminer.Status == 0 {
		return errors.New("this account has been disabled")
	}
	return nil
}

func convertLuckerDetailsList(in []*model.LuckerDetailsInfo) []*admin_pb.LuckerDetail {
	var res []*admin_pb.LuckerDetail
	for _, item := range in {
		res = append(res, &admin_pb.LuckerDetail{
			Lucker:     mdLuckerToPbLucker(item.Lucker),
			User:       mdUserToPbUser(item.User),
			PrizesPool: mdPrizesPoolToPbPrizesPool(item.PrizesPool),
			Activity:   mdActivityToPbActivity(item.Activity),
			Prize:      mdPrizeToPbPrize(item.Prize),
			PrizeGrade: mdPrizeGradeToPbPrizeGrade(item.PrizeGrade),
			Project:    mdProjectToPbProject(item.Project),
		})
	}
	return res
}
