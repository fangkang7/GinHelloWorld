package model

type Member struct {
	Id           int64  `xorm:"pk autoincr" json:"id"`
	UserName     string `xorm:"varchar(20)" json:"user_name"`
	Mobile       string `xorm:"varchar(20)" json:"mobile"`
	Password     string `xorm:"varchar(255)" json:"password"`
	RegisterTime int64  `xorm:"bigint" json:"register_time"`
}
