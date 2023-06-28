package dao

import (
	"GinHelloWorld/testProject/model"
	"GinHelloWorld/testProject/tool"
	"fmt"
)

type MemberDao struct {
	*tool.Orm
}

// 用户使用账号密码进行登录
func (md *MemberDao) Query(name string, password string) *model.Member {
	var member model.Member

	_, err := md.Where(" user_name = ? and password = ? ", name, password).Get(&member)
	if err != nil {
		return nil
	}

	return &member
}

// 验证手机号和验证码是否存在
func (md *MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	if _, err := md.Where("phone = ? and code = ?", phone, code).Get(&sms); err != nil {
		fmt.Println(err.Error())
	}

	return &sms
}

// 根据手机号查询数据
func (md *MemberDao) QueryByPhone(phone string) (*model.Member, error) {
	var member model.Member
	if _, err := md.Where("mobile = ?", phone).Get(&member); err != nil {
		return nil, err
	}

	return &member, nil
}

// 插入验证码
func (md *MemberDao) InsertCode(sms *model.SmsCode) int64 {
	result, err := md.InsertOne(sms)
	if err != nil {
		fmt.Println(err.Error())
	}

	return result
}

// 新用户的数据插入
func (md *MemberDao) InsertMember(member model.Member) int64 {
	result, err := md.InsertOne(&member)
	if err != nil {
		return 0
	}

	return result
}
