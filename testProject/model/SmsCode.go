package model

type SmsCode struct {
	Id        int64  `xorm:"pk autoincr" json:"id"`
	Phone     string `xorm:"varchar(11)" json:"phone"`
	Code      string `xorm:"varchar(30)" json:"code"`
	CreatTime int64  `xorm:"bigint" json:"creat_time"`
}
