package model

type Member struct {
	UUIDColumn
	Tasks  []Task `gorm:"many2many:MemberTask"`
	UserID string `gorm:"column:userId;type:varchar(36);"`
	User   User   `gorm:"foreignKey:UserID"`
}

func (m *Member) TableName() string {
	return "Member"
}
