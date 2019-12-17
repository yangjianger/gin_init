package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

//三元运算
// common.If(uid > 0, uid, 0).(int64)
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

//md5 加密
func Md5(key string) string{
	h := md5.New()
	h.Write([]byte(key)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr) // 输出加密结果
}

//时间戳转日期格式
func TimeTDate(times int64) (res string){
	if times <= 0{
		return
	}
	return time.Unix(times, 0).Format("2006-01-02 15:04:05")
}

//日期格式转时间戳
func DateToTime(dates string) (times int64){
	res, err := time.Parse("2006-01-02 15:04:05", dates)
	if err != nil{
		panic(err)
	}
	times = res.Unix()
	return
}
