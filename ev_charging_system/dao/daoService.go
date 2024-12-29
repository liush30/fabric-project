package dao

import (
	"context"
	"ev_charging_system/model"
)

var DaoService *daoService

type daoService struct {
	FeeRuleDao   IFeeRuleDo
	GunDao       IGunDo
	ParameterDao IParameterDo
	PileDao      IPileDo
	RequestReDAO IRepairRequestDo
	RepairmanDAO IRepairmanDo
	StationDao   IStationDo
	Query        *Query
}

func init() {
	temp := new(daoService)
	var bk context.Context
	bk = context.Background()
	daoEntity := Use(model.DB)
	temp.GunDao = daoEntity.Gun.WithContext(bk)
	temp.FeeRuleDao = daoEntity.FeeRule.WithContext(bk)
	temp.ParameterDao = daoEntity.Parameter.WithContext(bk)
	temp.PileDao = daoEntity.Pile.WithContext(bk)
	temp.RequestReDAO = daoEntity.RepairRequest.WithContext(bk)
	temp.RepairmanDAO = daoEntity.Repairman.WithContext(bk)
	temp.StationDao = daoEntity.Station.WithContext(bk)
	temp.Query = daoEntity
	DaoService = temp
}
