package mysql

import (
	"gift2grow_backend/loaders/mysql/model"

	"gorm.io/gorm"
)

var CampaignImagesModel *gorm.DB
var CampaignsModel *gorm.DB
var DonateHistoriesModel *gorm.DB
var EvidenceCampaignImagesModel *gorm.DB
var NotiObjectsModel *gorm.DB
var RanksModel *gorm.DB
var UserNotisModel *gorm.DB
var UsersModel *gorm.DB
var WantListsModel *gorm.DB
var UserTokensModel *gorm.DB

func assignModel() {
	CampaignImagesModel = Gorm.Model(new(model.CampaignImage))
	CampaignsModel = Gorm.Model(new(model.Campaign))
	DonateHistoriesModel = Gorm.Model(new(model.DonateHistory))
	EvidenceCampaignImagesModel = Gorm.Model(new(model.EvidenceCampaignImage))
	NotiObjectsModel = Gorm.Model(new(model.NotiObject))
	UserNotisModel = Gorm.Model(new(model.UserNoti))
	UsersModel = Gorm.Model(new(model.User))
	WantListsModel = Gorm.Model(new(model.WantList))
	UserTokensModel = Gorm.Model(new(model.UserToken))
}
