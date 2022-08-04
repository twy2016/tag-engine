package service

import (
	"tag-engine/model"
	"tag-engine/pkg/global"
)

func ListTag() (interface{}, error) {
	res := make([]*model.Tag, 0)
	return res, global.MysqlClient.Find(&res).Error
}

func GetTag(tag *model.Tag) (interface{}, error) {
	return tag, global.MysqlClient.First(&tag, &tag.Id).Error
}

func AddTag(tag *model.Tag) (interface{}, error) {
	return tag, global.MysqlClient.Save(&tag).Error
}

func UpdateTag(tag *model.Tag) (interface{}, error) {
	return tag, global.MysqlClient.Where("id=?", &tag.Id).Updates(&tag).Error
}

func DeleteTag(id string) (interface{}, error) {
	return id, global.MysqlClient.Where("id = ?", id).Delete(&model.Tag{}).Error
}
