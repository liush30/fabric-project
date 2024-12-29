package controller

import (
	"ev_charging_system/dao"
	"ev_charging_system/log"
	"ev_charging_system/model"
	"ev_charging_system/model/dto"
	"ev_charging_system/response"
	"ev_charging_system/tool"
	"github.com/gin-gonic/gin"
)

var FeeController = &feeController{}

type feeController struct {
}

// 新增充电站
func (s feeController) AddFee(g *gin.Context) {
	req := model.FeeRule{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

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
	req := dto.UserPageDto{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	page := (req.PageNum - 1) * req.PageSize

	//if req.UserType == 3 {
	pages, _, err := dao.DaoService.FeeRuleDao.FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondWithData(g, pages)
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
