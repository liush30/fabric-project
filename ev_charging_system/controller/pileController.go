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

var PileController = &pileController{}

type pileController struct {
}

// 新增充电站
func (s pileController) AddPile(g *gin.Context) {
	req := model.Pile{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	req.PileID = tool.GenerateUUIDWithoutDashes()
	err := dao.DaoService.PileDao.Create(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondOK(g)
}

func (s pileController) GetPileById(g *gin.Context) {
	pileId := g.Param("pileId")
	if len(pileId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	PileData, err := dao.DaoService.PileDao.Where(dao.DaoService.Query.Pile.PileID.Eq(pileId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "Pile not exist")
		return
	}
	response.RespondWithData(g, PileData)
}

func (s pileController) UpdatePile(g *gin.Context) {
	req := model.Pile{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	update, err := dao.DaoService.PileDao.Where(dao.DaoService.Query.Pile.PileID.Eq(req.PileID)).Updates(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	log.Info(update)
	response.RespondOK(g)
}

// 查询充电站分页信息
func (s pileController) PileByPage(g *gin.Context) {
	req := dto.UserPageDto{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	page := (req.PageNum - 1) * req.PageSize

	if req.UserType == 3 {
		pages, _, err := dao.DaoService.PileDao.FindByPage(page, req.PageSize)
		if err != nil {
			log.Error(err)
			response.RespondDefaultErr(g)
			return
		}

		response.RespondWithData(g, pages)
		return
	}

	pages, _, err := dao.DaoService.PileDao.Where(dao.DaoService.Query.Pile.Status.Eq(req.UserType)).FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondWithData(g, pages)
}

// 新增充电站
//func (s pileController) AddPile(g *gin.Context) {
//
//	response.RespondOK(g)
//}
