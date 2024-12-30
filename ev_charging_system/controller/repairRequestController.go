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
	"time"
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
		pages, count, err := dao.DaoService.RequestReDAO.FindByPage(page, req.PageSize)
		if err != nil {
			log.Error(err)
			response.RespondDefaultErr(g)
			return
		}

		pageResult := vo.Page[*model.RepairRequest]{
			Data:  pages,
			Count: count,
		}

		response.RespondWithData(g, pageResult)
		return
	}

	pages, count, err := dao.DaoService.RequestReDAO.Where(dao.DaoService.Query.RepairRequest.Status.Eq(req.UserType)).FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	pageResult := vo.Page[*model.RepairRequest]{
		Data:  pages,
		Count: count,
	}

	response.RespondWithData(g, pageResult)
}

// 新增充电站报修信息
func (s repairRequestController) AddMeRepairRequest(g *gin.Context) {
	req := model.RepairRequest{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	user, exists := g.Get("user")
	if !exists {
		response.RespondWithErrCode(g, 401, "not login")
		return
	}
	userInfo := user.(tool.User)

	staioninfo, err := dao.DaoService.StationDao.Where(dao.DaoService.Query.Station.RepairmanID.Eq(userInfo.RepairmanId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	req.StationID = staioninfo.StationID
	req.RequestTime = time.Now().Unix()
	req.RepairID = tool.GenerateUUIDWithoutDashes()
	err = dao.DaoService.RequestReDAO.Create(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondOK(g)
}

// 新增充电站
//func (s repairRequestController) AddRepairRequest(g *gin.Context) {
//
//	response.RespondOK(g)
//}
