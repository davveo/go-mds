package entity

import "time"

type BaseEntity struct {
	id         int32     //主键id
	version    int       // 版本号默认为0
	status     string    // 状态 PublicStatusEnum
	creater    string    // 创建人
	editor     string    // 修改人
	remark     string    // 描述
	createTime time.Time // 创建时间
	editTime   time.Time // 修改时间
}

func (b *BaseEntity) Id() int32 {
	return b.id
}

func (b *BaseEntity) SetId(id int32) {
	b.id = id
}

func (b *BaseEntity) EditTime() time.Time {
	return b.editTime
}

func (b *BaseEntity) SetEditTime(editTime time.Time) {
	b.editTime = editTime
}

func (b *BaseEntity) CreateTime() time.Time {
	return b.createTime
}

func (b *BaseEntity) SetCreateTime(createTime time.Time) {
	b.createTime = createTime
}

func (b *BaseEntity) Remark() string {
	return b.remark
}

func (b *BaseEntity) SetRemark(remark string) {
	b.remark = remark
}

func (b *BaseEntity) Editor() string {
	return b.editor
}

func (b *BaseEntity) SetEditor(editor string) {
	b.editor = editor
}

func (b *BaseEntity) Creater() string {
	return b.creater
}

func (b *BaseEntity) SetCreater(creater string) {
	b.creater = creater
}

func (b *BaseEntity) Status() string {
	return b.status
}

func (b *BaseEntity) SetStatus(status string) {
	b.status = status
}

func (b *BaseEntity) Version() int {
	return b.version
}

func (b *BaseEntity) SetVersion(version int) {
	b.version = version
}
