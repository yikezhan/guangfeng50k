package dao

import (
	"entryTask/internal/common"
	"entryTask/internal/model"
	"time"
)

func (d *Dao) AllGameResult(roomID int64, roomUser string) []*model.GameResultTab {
	query := &model.GameResultTab{
		RoomID:   roomID,
		RoomUser: roomUser,
	}
	var res []*model.GameResultTab
	d.db.Table(query.TableName()).Where("room_id = ?", query.RoomID, "room_user = ?", query.RoomUser).
		Order("id desc").Limit(500).Scan(&res)
	return res
}
func (d *Dao) LatestGameResult(roomID int64, roomUser string) []*model.GameResultTab {
	query := &model.GameResultTab{
		RoomID:   roomID,
		RoomUser: roomUser,
	}
	var res []*model.GameResultTab
	d.db.Table(query.TableName()).Where("room_id = ? and room_user = ?", query.RoomID, query.RoomUser).
		Order("id desc").Limit(1).Scan(&res)
	return res
}
func (d *Dao) QueryGameResult(roomID int64, number int64) []model.GameResultTab {
	query := &model.GameResultTab{
		RoomID: roomID,
		Number: number,
	}
	var res []model.GameResultTab
	d.db.Table(query.TableName()).Where("room_id = ? and number = ?", query.RoomID, query.Number).
		Limit(4).Scan(&res)
	return res
}
func (d *Dao) CreateGameResult(info *model.GameResultTab) bool {
	info.CreateTime = time.Now().Unix()
	info.UpdateTime = time.Now().Unix()
	info.IsDelete = common.Valid
	info.Status = common.Draft
	info.Amount = 0
	tx := d.db.Table(info.TableName()).Create(&info)
	return tx != nil && tx.RowsAffected == 1
}
func (d *Dao) UpdateGameResult(info *model.GameResultTab) bool {
	info.UpdateTime = time.Now().Unix()
	tx := d.db.Table(info.TableName()).Where("id", info.ID).Updates(info)
	return tx != nil && tx.RowsAffected == 1
}
func (d *Dao) UpdateAmount(ID int64, amount int64) bool {
	info := &model.GameResultTab{
		ID:         ID,
		Amount:     amount,
		UpdateTime: time.Now().Unix(),
	}
	tx := d.db.Table(info.TableName()).Where("id", info.ID).Select("amount", "update_time").Updates(info)
	return tx != nil && tx.RowsAffected == 1
}
func (d *Dao) UpdateStatus(resultID int64) bool {
	info := &model.GameResultTab{
		ID:         resultID,
		Status:     common.Confirm,
		UpdateTime: time.Now().Unix(),
	}
	tx := d.db.Table(info.TableName()).Where("id", info.ID).Select("status", "update_time").Updates(info)
	return tx != nil && tx.RowsAffected == 1
}
