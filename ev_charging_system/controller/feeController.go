package controller

import (
	"ev_charging_system/dao"
	"ev_charging_system/log"
	"ev_charging_system/model"
	"ev_charging_system/model/dto"
	"ev_charging_system/model/vo"
	"ev_charging_system/response"
	"ev_charging_system/tool"
	"github.com/gin-gonic/gin"
)

var FeeController = &feeController{}

type feeController struct {
}

// 新增充电站
func (s feeController) AddFee(g *gin.Context) {
	user, exists := g.Get("user")
	if !exists {
		response.RespondWithErrCode(g, 401, "not login")
		return
	}
	userInfo := user.(tool.User)

	req := model.FeeRule{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}
	req.StationID = userInfo.RepairmanId

	req.RuleID = tool.GenerateUUIDWithoutDashes()
	err := dao.DaoService.FeeRuleDao.Create(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondOK(g)
}

func (s feeController) GetFeeById(g *gin.Context) {
	feeId := g.Param("FeeId")
	if len(feeId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	FeeData, err := dao.DaoService.FeeRuleDao.Where(dao.DaoService.Query.FeeRule.RuleID.Eq(feeId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "Fee not exist")
		return
	}
	response.RespondWithData(g, FeeData)
}

func (s feeController) DeleteFee(g *gin.Context) {
	feeId := g.Param("FeeId")
	if len(feeId) == 0 {
		response.RespondDefaultErr(g)
		return
	}
	info, err := dao.DaoService.FeeRuleDao.Where(dao.DaoService.Query.FeeRule.RuleID.Eq(feeId)).Delete()
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	} else if info.RowsAffected == 0 {
		response.RespondErr(g, "affected 0 rows")
		return
	}
	response.RespondOK(g)
}

func (s feeController) UpdateFee(g *gin.Context) {
	req := model.FeeRule{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	update, err := dao.DaoService.FeeRuleDao.Where(dao.DaoService.Query.FeeRule.RuleID.Eq(req.RuleID)).Updates(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	log.Info(update)
	response.RespondOK(g)
}

// 查询充电站分页信息
func (s feeController) FeeByPage(g *gin.Context) {
	user, exists := g.Get("user")
	if !exists {
		response.RespondWithErrCode(g, 401, "not login")
		return
	}
	userInfo := user.(tool.User)
	req := dto.UserPageDto{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}
	log.Info(req)

	page := (req.PageNum - 1) * req.PageSize
	var pages []*model.FeeRule
	var count int64
	var err error
	if req.IsUseState {
		pages, count, err = dao.DaoService.FeeRuleDao.Where(dao.DaoService.Query.FeeRule.StationID.Eq(userInfo.RepairmanId), dao.DaoService.Query.FeeRule.ChargingType.Eq(req.State)).FindByPage(page, req.PageSize)
	} else {
		pages, count, err = dao.DaoService.FeeRuleDao.Where(dao.DaoService.Query.FeeRule.StationID.Eq(userInfo.RepairmanId)).FindByPage(page, req.PageSize)
	}
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	pageResult := vo.Page[*model.FeeRule]{
		Data:  pages,
		Count: count,
	}

	response.RespondWithData(g, pageResult)
	return
	//}

	//pages, _, err := dao.DaoService.FeeRuleDao.Where(dao.DaoService.Query.FeeRule.Status.Eq(req.UserType)).FindByPage(page, req.PageSize)
	//if err != nil {
	//	log.Error(err)
	//	response.RespondDefaultErr(g)
	//	return
	//}

	//response.RespondWithData(g, pages)
}

// 新增充电站
//func (s feeController) AddFee(g *gin.Context) {
//
//	response.RespondOK(g)
//}
