package model

type MemberTask struct {
	TaskID   string `gorm:"uuid;column:taskId;primaryKey"`
	MemberID string `gorm:"uuid;column:memberId;primaryKey"`
}

func (m *MemberTask) TableName() string {
	return "MemberTask"
}
