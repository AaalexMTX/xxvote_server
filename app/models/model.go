package models

import (
	"time"
)

type User struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Uid        int       `gorm:"column:uid;type:int(11);comment:可展示的uid" json:"uid"`
	Name       string    `gorm:"column:name;type:varchar(50);NOT NULL" json:"name"`
	Password   string    `gorm:"column:password;type:varchar(50)" json:"password"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (m *User) TableName() string {
	return "user"
}

type Vote struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	UserId     int       `gorm:"column:user_id;type:int(11)" json:"user_id"`
	Title      string    `gorm:"column:title;type:varchar(255)" json:"title"`
	Time       int       `gorm:"column:time;type:datetime;comment:有效时长" json:"time"`
	Status     int       `gorm:"column:status;type:tinyint(1);comment:是否有效" json:"status"`
	Type       int       `gorm:"column:type;type:tinyint(1);comment:0单选 1多选" json:"type"`
	Repeat     int       `gorm:"column:repeat;type:tinyint(1);comment:0不可重复 1可重复" json:"repeat"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (m *Vote) TableName() string {
	return "vote"
}

type VoteOption struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	VoteId     int       `gorm:"column:vote_id;type:int(11)" json:"vote_id"`
	Name       string    `gorm:"column:name;type:varchar(255)" json:"name"`
	Count      int       `gorm:"column:count;type:int(11)" json:"count"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (m *VoteOption) TableName() string {
	return "vote_option"
}

type VoteOptionUser struct {
	Id           int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	UserId       int       `gorm:"column:user_id;type:int(11)" json:"user_id"`
	VoteId       int       `gorm:"column:vote_id;type:int(11)" json:"vote_id"`
	VoteOptionId int       `gorm:"column:vote_option_id;type:int(11)" json:"vote_option_id"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`
}

func (m *VoteOptionUser) TableName() string {
	return "vote_option_user"
}

// VoteWithOptions 收集用户对应的投票信息
type VoteWithOptions struct {
	Vote   Vote         `json:"Vote"`
	Option []VoteOption `json:"Opt"`
}
