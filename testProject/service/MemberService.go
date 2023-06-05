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
	// 验证手机号 + 验证码是否正确
	md := &dao.MemberDao{tool.DbEngine}
	sms := md.ValidateSmsCode(param.Phone, param.Code)
	if sms.Id == 0 {
		return nil
	}

	// 根据手机号member表中查询记录
	member, _ := md.QueryByPhone(param.Phone)
	if member.Id != 0 {
		return member
	}

	// 新创建一个member记录，并保存
	user := model.Member{
		UserName:     param.Phone,
		Mobile:       param.Phone,
		RegisterTime: time.Now().Unix(),
	}

	user.Id = md.InsertMember(user)

	return &user
}

func (ms *MemberService) SendCode(phone string) bool {

	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	smsCode := &model.SmsCode{Phone: phone, Code: code, CreatTime: time.Now().Unix()}
	memberDao := &dao.MemberDao{tool.DbEngine}
	result := memberDao.InsertCode(smsCode)

	//result := (&dao.MemberDao{tool.DbEngine}).InsertCode(smsCode)

	return result > 0
}
