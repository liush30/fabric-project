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

var GunController = &guneController{}

type guneController struct {
}

// 新增充电站
func (s guneController) AddGun(g *gin.Context) {
	req := model.Gun{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	req.GunID = tool.GenerateUUIDWithoutDashes()
	err := dao.DaoService.GunDao.Create(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondOK(g)
}

func (s guneController) GetGunById(g *gin.Context) {
	gunId := g.Param("GunId")
	if len(gunId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	GunData, err := dao.DaoService.GunDao.Where(dao.DaoService.Query.Gun.GunID.Eq(gunId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "Gun not exist")
		return
	}
	response.RespondWithData(g, GunData)
}

func (s guneController) UpdateGun(g *gin.Context) {
	req := model.Gun{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	update, err := dao.DaoService.GunDao.Where(dao.DaoService.Query.Gun.GunID.Eq(req.GunID)).Updates(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	log.Info(update)
	response.RespondOK(g)
}

// 查询充电站分页信息
func (s guneController) GunByPage(g *gin.Context) {
	req := dto.UserPageDto{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	page := (req.PageNum - 1) * req.PageSize

	//if req.UserType == 3 {
	pages, count, err := dao.DaoService.GunDao.FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	pageResult := vo.Page[*model.Gun]{
		Data:  pages,
		Count: count,
	}

	response.RespondWithData(g, pageResult)
	return
	//}

	//pages, _, err := dao.DaoService.GunDao.Where(dao.DaoService.Query.Gun.Status.Eq(req.UserType)).FindByPage(page, req.PageSize)
	//if err != nil {
	//	log.Error(err)
	//	response.RespondDefaultErr(g)
	//	return
	//}
	//
	//response.RespondWithData(g, pages)
}

func (s guneController) GetGunListByPileId(g *gin.Context) {
	pileId := g.Param("pileId")
	if len(pileId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	GunData, err := dao.DaoService.GunDao.Where(dao.DaoService.Query.Gun.PileID.Eq(pileId)).Find()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "Gun not exist")
		return
	}
	response.RespondWithData(g, GunData)
}

// 新增充电站
//func (s guneController) AddGun(g *gin.Context) {
//
//	response.RespondOK(g)
//}
