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

var RepairRequestController = &repairRequestController{}

type repairRequestController struct {
}

// 新增充电站
func (s repairRequestController) AddRepairRequest(g *gin.Context) {
	req := model.RepairRequest{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	req.RepairID = tool.GenerateUUIDWithoutDashes()
	err := dao.DaoService.RequestReDAO.Create(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondOK(g)
}

func (s repairRequestController) GetRepairRequestById(g *gin.Context) {
	RepairRequestId := g.Param("repairRequestId")
	if len(RepairRequestId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	RepairRequestData, err := dao.DaoService.RequestReDAO.Where(dao.DaoService.Query.RepairRequest.RepairmanID.Eq(RepairRequestId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "RepairRequest not exist")
		return
	}
	response.RespondWithData(g, RepairRequestData)
}

func (s repairRequestController) UpdateRepairRequest(g *gin.Context) {
	req := model.RepairRequest{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	update, err := dao.DaoService.RequestReDAO.Where(dao.DaoService.Query.RepairRequest.RepairID.Eq(req.RepairID)).Updates(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	log.Info(update)
	response.RespondOK(g)
}

// 查询充电站分页信息
func (s repairRequestController) RepairRequestByPage(g *gin.Context) {
	req := dto.UserPageDto{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	page := (req.PageNum - 1) * req.PageSize

	if req.UserType == 3 {
		pages, _, err := dao.DaoService.RequestReDAO.FindByPage(page, req.PageSize)
		if err != nil {
			log.Error(err)
			response.RespondDefaultErr(g)
			return
		}

		response.RespondWithData(g, pages)
		return
	}

	pages, _, err := dao.DaoService.RequestReDAO.Where(dao.DaoService.Query.RepairRequest.Status.Eq(req.UserType)).FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondWithData(g, pages)
}

// 新增充电站
//func (s repairRequestController) AddRepairRequest(g *gin.Context) {
//
//	response.RespondOK(g)
//}
