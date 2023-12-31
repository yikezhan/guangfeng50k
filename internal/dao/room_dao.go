package dao

import (
	"guangfeng/internal/model"
	"time"
)

func (d *Dao) CreateRoom(info *model.RoomTab) bool {
	info.Number = 1
	info.CreateTime = time.Now().Unix()
	tx := d.db.Table(info.TableName()).Create(&info)
	return tx != nil && tx.RowsAffected == 1
}
func (d *Dao) UpdateRoom(info *model.RoomTab) bool {
	info.UpdateTime = time.Now().Unix()
	tx := d.db.Table(info.TableName()).Where("id", info.ID).Updates(info)
	return tx != nil && tx.RowsAffected == 1
}

func (d *Dao) UpdateNumber(id int64, number int64) bool {
	info := &model.RoomTab{
		ID:         id,
		Number:     number,
		UpdateTime: time.Now().Unix(),
	}
	tx := d.db.Table(info.TableName()).Where("id", info.ID).Select("number", "update_time").Updates(info)
	return tx != nil && tx.RowsAffected == 1
}
func (d *Dao) QueryRoom(roomName string) *model.RoomTab {
	info := &model.RoomTab{RoomName: roomName}
	var res *model.RoomTab
	var count int64
	tx := d.db.Table(info.TableName()).Where("room_name", info.RoomName).Count(&count)
	if count > 0 {
		tx.First(&res)
	}
	return res
}

func (d *Dao) QueryRoomById(roomId int64) *model.RoomTab {
	info := &model.RoomTab{ID: roomId}
	var res *model.RoomTab
	var count int64
	tx := d.db.Table(info.TableName()).Where("id", info.ID).Count(&count)
	if count > 0 {
		tx.First(&res)
	}
	return res
}
