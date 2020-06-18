package models

import (
	"time"
)

type Message struct {
	Id               int64     `xorm:"pk autoincr"`                      // 主键
	Status           string    `xorm:"varchar(20) notnull default('')"`  // 状态
	Creater          string    `xorm:"varchar(100) default('')"`         // 创建者
	Editor           string    `xorm:"varchar(100) default('')"`         // 编辑者
	MessageId        string    `xorm:"varchar(50) notnull default('')"`  // 消息ID
	MessageBody      string    `xorm:"varchar(255) notnull default('')"` // 消息体
	Remark           string    `xorm:"varchar(200) default('')"`         // 备注
	MessageSendTimes int       `xorm:"int notnull defalut('0')"`         // 消息重发次数
	MessageDataType  string    `xorm:"varchar(50)  null default('')"`    // 消息数据类型
	ConsumerQueue    string    `xorm:"varchar(100) notnull default('')"` // 消费队列
	AreadlyDead      string    `xorm:"varchar(20) notnull default('')"`  // 是否死亡
	Field1           string    `xorm:"varchar(200) notnull default('')"` // 扩展字段1
	Field2           string    `xorm:"varchar(200) notnull default('')"` // 扩展字段2
	Field3           string    `xorm:"varchar(200) notnull default('')"` // 扩展字段3
	Version          int       `xorm:"version"`                          //版本
	CreatedTime      time.Time `xorm:"created"`                          //创建时间
	DeletedTime      time.Time `xorm:"created"`                          //删除时间
	UpdatedTime      time.Time `xorm:"updated"`                          //修改后自动更新时间
}
