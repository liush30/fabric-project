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

var StationController = &stationController{}

type stationController struct {
}

// 新增充电站
func (s stationController) AddStation(g *gin.Context) {
	req := model.Station{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}
	id := tool.GenerateUUIDWithoutDashes()
	man := model.Repairman{
		RepairmanID: id,
		UserName:    req.StationName,
		Password:    req.LoginPwd,
		UserType:    1,
	}
	err := dao.DaoService.RepairmanDAO.Create(&man)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	req.StationID = tool.GenerateUUIDWithoutDashes()
	req.RepairmanID = id
	err = dao.DaoService.StationDao.Create(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	response.RespondOK(g)
}

func (s stationController) GetStationById(g *gin.Context) {
	stationId := g.Param("stationId")
	if len(stationId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	stationData, err := dao.DaoService.StationDao.Where(dao.DaoService.Query.Station.StationID.Eq(stationId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "Station not exist")
		return
	}
	response.RespondWithData(g, stationData)
}

func (s stationController) UpdateStation(g *gin.Context) {
	req := model.Station{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	update, err := dao.DaoService.StationDao.Where(dao.DaoService.Query.Station.StationID.Eq(req.StationID)).Updates(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	log.Info(update)
	response.RespondOK(g)
}

// 查询充电站分页信息
func (s stationController) StationByPage(g *gin.Context) {
	req := dto.UserPageDto{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	page := (req.PageNum - 1) * req.PageSize

	if req.UserType == 3 {
		pages, count, err := dao.DaoService.StationDao.FindByPage(page, req.PageSize)
		if err != nil {
			log.Error(err)
			response.RespondDefaultErr(g)
			return
		}
		pageResult := vo.Page[*model.Station]{
			Data:  pages,
			Count: count,
		}
		response.RespondWithData(g, pageResult)
		return
	}

	pages, count, err := dao.DaoService.StationDao.Where(dao.DaoService.Query.Station.Status.Eq(req.UserType)).FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	pageResult := vo.Page[*model.Station]{
		Data:  pages,
		Count: count,
	}

	response.RespondWithData(g, pageResult)
}

func (s stationController) GetMeStationInfo(g *gin.Context) {
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

	response.RespondWithData(g, staioninfo)
}

// 新增充电站
// func (s stationController) AddStation(g *gin.Context) {
//
//		response.RespondOK(g)
//	}
func (s stationController) DeleteStation(g *gin.Context) {
	stationId := g.Param("stationId")
	if len(stationId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	info, err := dao.DaoService.StationDao.Where(dao.DaoService.Query.Station.StationID.Eq(stationId)).Delete()
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
