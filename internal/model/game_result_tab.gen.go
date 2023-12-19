// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGameResultTab = "game_result_tab"

// GameResultTab game_result_tab
type GameResultTab struct {
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoomID     int64  `gorm:"column:room_id;not null;comment:room_id" json:"room_id"`                         // room_id
	WxID       string `gorm:"column:wx_id;not null;comment:wx_id" json:"wx_id"`                               // wx_id
	Number     int64  `gorm:"column:number;not null;comment:number" json:"number"`                            // number
	ResultJSON string `gorm:"column:result_json;comment:result_json" json:"result_json"`                      // result_json
	Amount     int64  `gorm:"column:amount;not null;comment:amount" json:"amount"`                            // amount
	Status     int32  `gorm:"column:status;not null;comment:0 draft,1confirm" json:"status"`                  // 0 draft,1confirm
	CreateTime int64  `gorm:"column:create_time;not null;comment:create time" json:"create_time"`             // create time
	UpdateTime int64  `gorm:"column:update_time;not null;comment:update time" json:"update_time"`             // update time
	IsDelete   int32  `gorm:"column:is_delete;not null;comment:0:valid,1:logically deleted" json:"is_delete"` // 0:valid,1:logically deleted
}

// TableName GameResultTab's table name
func (*GameResultTab) TableName() string {
	return TableNameGameResultTab
}
