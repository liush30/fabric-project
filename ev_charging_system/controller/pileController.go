package controller

import (
	"encoding/json"
	"ev_charging_system/dao"
	"ev_charging_system/fabric"
	"ev_charging_system/log"
	"ev_charging_system/model"
	"ev_charging_system/model/dto"
	"ev_charging_system/model/vo"
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

	pileData, err := json.Marshal(req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	pileHash := tool.CalculateSHA256Hash(tool.EncodeToString(pileData))

	err = fabric.RegisterPile(req.PileID, pileHash)
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
		pages, count, err := dao.DaoService.PileDao.FindByPage(page, req.PageSize)
		if err != nil {
			log.Error(err)
			response.RespondDefaultErr(g)
			return
		}
		pageResult := vo.Page[*model.Pile]{
			Data:  pages,
			Count: count,
		}
		response.RespondWithData(g, pageResult)
		return
	}

	pages, count, err := dao.DaoService.PileDao.Where(dao.DaoService.Query.Pile.Status.Eq(req.UserType)).FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	pageResult := vo.Page[*model.Pile]{
		Data:  pages,
		Count: count,
	}

	response.RespondWithData(g, pageResult)
}

func (s pileController) GetMePilePage(g *gin.Context) {
	req := dto.PilePageDto{}
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

	page := (req.PageNum - 1) * req.PageSize

	if req.Status == 3 || req.Type == 3 {
		pages, count, err := dao.DaoService.PileDao.Where(dao.DaoService.Query.Pile.StationID.Eq(staioninfo.StationID)).FindByPage(page, req.PageSize)
		if err != nil {
			log.Error(err)
			response.RespondDefaultErr(g)
			return
		}
		pageResult := vo.Page[*model.Pile]{
			Data:  pages,
			Count: count,
		}

		response.RespondWithData(g, pageResult)
		return
	}

	pages, count, err := dao.DaoService.PileDao.Where(dao.DaoService.Query.Pile.StationID.Eq(staioninfo.StationID), dao.DaoService.Query.Pile.Status.Eq(req.Status), dao.DaoService.Query.Pile.Type.Eq(req.Type)).FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	pageResult := vo.Page[*model.Pile]{
		Data:  pages, // 这里直接使用原切片
		Count: count,
	}
	response.RespondWithData(g, pageResult)
}

// 新增充电桩
func (s pileController) AddMePile(g *gin.Context) {
	req := model.Pile{}
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
	req.PileCode = tool.GenerateUUIDWithoutDashes()
	req.StationID = staioninfo.StationID
	req.PileID = tool.GenerateUUIDWithoutDashes()
	err = dao.DaoService.PileDao.Create(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	pileData, err := json.Marshal(req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	pileHash := tool.CalculateSHA256Hash(tool.EncodeToString(pileData))

	err = fabric.RegisterPile(req.PileID, pileHash)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondOK(g)
}

// 新增充电站
//func (s pileController) AddPile(g *gin.Context) {
//
//	response.RespondOK(g)
//}
