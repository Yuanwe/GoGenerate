package service

import (
	"GoGenerate/global"
	"GoGenerate/model"
)

// @title    GetAllCampaign
// @description   GetAllCampaign, GetAllCampaign
// @auth                     （2020/04/05  20:22）
// @return    err             error
func GetAllFields(tableName string) (err error, campaign []model.TableField) {
	sql := "select COLUMN_NAME,DATA_TYPE from information_schema.COLUMNS where table_name = '" + tableName + "';"
	err = global.GG_DB.Raw(sql).Scan(&campaign).Error
	return err, campaign
}
