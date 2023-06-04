package service

import (
	"GinHelloWorld/testProject/dao"
	"GinHelloWorld/testProject/model"
	"GinHelloWorld/testProject/param"
	"GinHelloWorld/testProject/tool"
	"fmt"
	"math/rand"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) SmsLogin(param param.SmsLoginParam) *model.Member {
	md := &dao.MemberDao{}
	sms := md.ValidateSmsCode(param.Phone, param.Code)
	if sms.Id == 0 {
		return nil
	}
	return nil
}

func (ms *MemberService) SendCode(phone string) bool {

	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	smsCode := &model.SmsCode{Phone: phone, Code: code, CreatTime: time.Now().Unix()}
	memberDao := &dao.MemberDao{tool.DbEngine}
	result := memberDao.InsertCode(smsCode)

	//result := (&dao.MemberDao{tool.DbEngine}).InsertCode(smsCode)

	return result > 0
}
