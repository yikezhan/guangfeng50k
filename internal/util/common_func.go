package util

import (
	"guangfeng/internal/common"
	"reflect"
	"time"
)

func ReflectBuildDefaultTimeAndValid(str interface{}) {
	v := reflect.ValueOf(str).Elem()
	curr := time.Now().Unix()
	createTime := v.FieldByName("CreateTime")
	createTime.SetInt(curr)
	updateTime := v.FieldByName("UpdateTime")
	updateTime.SetInt(curr)
	isDelete := v.FieldByName("IsDelete")
	isDelete.Set(reflect.ValueOf(common.Valid))
}
