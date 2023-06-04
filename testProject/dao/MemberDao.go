package dao

import (
	"GinHelloWorld/testProject/model"
	"GinHelloWorld/testProject/tool"
	"fmt"
)

type MemberDao struct {
	*tool.Orm
}

func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}

	return &sms
}

func (md *MemberDao) InsertCode(sms *model.SmsCode) int64 {
	result, err := md.InsertOne(sms)
	if err != nil {
		fmt.Println(err.Error())
	}

	return result
}
