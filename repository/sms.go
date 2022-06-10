package repository

import (
	"super_admin/library/env"
	"super_admin/providers"
	"time"
)

func SetSMSEntry(phone, verifyCode string) (ok bool, err error) {
	//TODO: 这个接口有安全风险，需要限流（网关那里做）
	/*
		ip限制daily调用次数100次。 防止恶意刷接口
	*/
	// 1分钟内一个手机只能发一次。
	smsTimeMinute := env.GetIntVal("SMS_TIME_DURARTION_MINUTES")
	cmd := providers.RedisClient.SetNX(phone, verifyCode, (time.Duration(smsTimeMinute) * time.Minute))
	ok, err = cmd.Result()
	return
}

func GetSMSEntry(phone string) (res string, err error) {
	cmd := providers.RedisClient.Get(phone)
	res, err = cmd.Result()
	return
}
