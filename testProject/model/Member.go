package model

type Member struct {
	Id           int64  `xorm:"ok autoincr" json:"id"`
	UserName     string `xorm:"varchar(20)" json:"user_name"`
	Mobile       string `xorm:"varchar(20)" json:"mobile"`
	Password     string `xorm:"varchar(255)" json:"password"`
	RegisterTime string `xorm:"bigint" json:"register_time"`
}
