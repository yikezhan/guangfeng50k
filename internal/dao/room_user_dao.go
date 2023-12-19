package dao

import (
	"guangfeng/internal/common"
	"guangfeng/internal/model"
	"time"
)

func (d *Dao) CreateRoomUser(info *model.RoomUserTab) bool {
	info.UpdateTime = time.Now().Unix()
	info.CreateTime = time.Now().Unix()
	info.IsDelete = common.Valid
	tx := d.db.Table(info.TableName()).Create(&info)
	return tx != nil && tx.RowsAffected == 1
}
func (d *Dao) QueryRoomUsers(roomId int64) []*model.RoomUserTab {
	info := &model.RoomUserTab{RoomID: roomId}
	var res []*model.RoomUserTab
	d.db.Table(info.TableName()).Where("room_id", info.RoomID).Scan(&res)
	return res
}
