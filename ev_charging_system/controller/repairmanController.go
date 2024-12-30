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

var RepairmanController = &repairmanController{}

type repairmanController struct {
}

func (r repairmanController) Login(g *gin.Context) {
	req := dto.UserDto{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}
	userInfo, err := dao.DaoService.RepairmanDAO.Where(dao.DaoService.Query.Repairman.UserName.Eq(req.UserName), dao.DaoService.Query.Repairman.Password.Eq(req.Password)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "username or passwd error")
		return
	}

	token, err := tool.GenerateJWT(tool.User{
		RepairmanId: userInfo.RepairmanID,
		UserId:      userInfo.UserName,
		UserType:    int(userInfo.UserType),
	})

	if err != nil {
		log.Error(err)
		response.RespondServerError(g)
		return
	}

	response.RespondWithData(g, token)
}

func (r repairmanController) Info(g *gin.Context) {
	user, exists := g.Get("user")
	if !exists {
		response.RespondWithErrCode(g, 401, "not login")
		return
	}
	userInfo := user.(tool.User)
	userData, err := dao.DaoService.RepairmanDAO.Where(dao.DaoService.Query.Repairman.RepairmanID.Eq(userInfo.RepairmanId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "username or passwd error")
		return
	}
	response.RespondWithData(g, userData)
}

// 查询所有用户分页信息
func (r repairmanController) ListAndPage(g *gin.Context) {
	req := dto.UserPageDto{}
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
	if userInfo.UserType != 2 {
		response.RespondErr(g, "not admin")
		return
	}
	page := (req.PageNum - 1) * req.PageSize

	if req.UserType == 3 {
		pages, count, err := dao.DaoService.RepairmanDAO.FindByPage(page, req.PageSize)
		if err != nil {
			log.Error(err)
			response.RespondDefaultErr(g)
			return
		}

		pageResult := vo.Page[*model.Repairman]{
			Data:  pages,
			Count: count,
		}
		response.RespondWithData(g, pageResult)
		return
	}

	pages, count, err := dao.DaoService.RepairmanDAO.Where(dao.DaoService.Query.Repairman.UserType.Eq(req.UserType)).FindByPage(page, req.PageSize)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}

	pageResult := vo.Page[*model.Repairman]{
		Data:  pages,
		Count: count,
	}
	response.RespondWithData(g, pageResult)
}

func (r repairmanController) GetUserById(g *gin.Context) {
	userId := g.Param("userId")
	print("userid:", userId)
	if len(userId) == 0 {
		response.RespondDefaultErr(g)
		return
	}

	userData, err := dao.DaoService.RepairmanDAO.Where(dao.DaoService.Query.Repairman.RepairmanID.Eq(userId)).Take()
	if err != nil {
		log.Error(err)
		response.RespondErr(g, "user not exist")
		return
	}
	response.RespondWithData(g, userData)
}

func (r repairmanController) AddUser(g *gin.Context) {
	req := model.Repairman{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}
	req.RepairmanID = tool.GenerateUUIDWithoutDashes()
	err := dao.DaoService.RepairmanDAO.Create(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	response.RespondOK(g)
}

func (r repairmanController) UserUpdateUser(g *gin.Context) {
	req := model.Repairman{}
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
	if userInfo.UserType != 2 {
		response.RespondErr(g, "not admin")
		return
	}

	update, err := dao.DaoService.RepairmanDAO.Where(dao.DaoService.Query.Repairman.RepairmanID.Eq(userInfo.RepairmanId)).Updates(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	log.Info(update)
	response.RespondOK(g)
}

func (r repairmanController) UpdateUser(g *gin.Context) {
	req := model.Repairman{}
	if err := g.Bind(&req); err != nil {
		log.Error(err)
		response.RespondInvalidArgsErr(g)
		return
	}

	update, err := dao.DaoService.RepairmanDAO.Where(dao.DaoService.Query.Repairman.RepairmanID.Eq(req.RepairmanID)).Updates(&req)
	if err != nil {
		log.Error(err)
		response.RespondDefaultErr(g)
		return
	}
	log.Info(update)
	response.RespondOK(g)
}
