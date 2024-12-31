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
	reqJson, err := json.Marshal(req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	reqHash := tool.CalculateSHA256Hash(tool.EncodeToString(reqJson))
	err = fabric.RegisterGun(req.PileID, req.GunID, reqHash)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	response.RespondOK(g)
}

func (s guneController) GetGunById(g *gin.Context) {
	gunId := g.Param("gunId")
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

func (s guneController) GetGunHistory(g *gin.Context) {
	gunId := g.Param("gunId")
	if len(gunId) == 0 {
		response.RespondDefaultErr(g)
		return
	}
	data, err := dao.DaoService.GunDao.Where(dao.DaoService.Query.Gun.GunID.Eq(gunId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "Gun not exist")
		return
	}
	log.Info(data.PileID, data.GunID)
	guns, err := fabric.QueryGunHistory(data.PileID, data.GunID)
	if err != nil {
		response.RespondDefaultErr(g)
		return
	}
	response.RespondWithData(g, guns)
}

func (s guneController) GetGunHistoryByPileId(g *gin.Context) {

	pileId := g.Param("pileId")
	if len(pileId) == 0 {
		response.RespondDefaultErr(g)
		return
	}
	guns, err := fabric.QueryGunByPile(pileId)
	if err != nil {
		response.RespondDefaultErr(g)
		return
	}
	response.RespondWithData(g, guns)

}
func (s guneController) DeleteGun(g *gin.Context) {
	gunId := g.Param("gunId")
	if len(gunId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	info, err := dao.DaoService.GunDao.Where(dao.DaoService.Query.Gun.GunID.Eq(gunId)).Delete()
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
	} else if update.RowsAffected == 0 {
		response.RespondErr(g, "affected 0 rows")
		return
	}

	GunData, err := dao.DaoService.GunDao.Where(dao.DaoService.Query.Gun.GunID.Eq(req.GunID)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "Gun not exist")
		return
	}
	gunJson, err := json.Marshal(GunData)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	gunHash := tool.CalculateSHA256Hash(string(gunJson))
	err = fabric.UpdateGun(GunData.PileID, GunData.GunID, gunHash)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

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
