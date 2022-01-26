package entity

import (
	"time"

	"gorm.io/gorm"
)

type PremiumMember struct {
	gorm.Model
	PremiumMemberID string    
	CreateAt        time.Time 
	Point           int    

	MemberID *uint
	Member   Member `gorm:"references:ID"`

	MemberClassID *uint
	MemberClass   MemberClass `gorm:"references:ID"`

	PremiumMemberPeriodID *uint
	PremiumMemberPeriod   PremiumMemberPeriod `gorm:"references:ID"`
}

type MemberClass struct {
	gorm.Model
	Name   string
	Detail string

	PremiumMember []PremiumMember `gorm:"foreignKey:MemberClassID"`
}

type PremiumMemberPeriod struct {
	gorm.Model
	Period string

	PremiumMember []PremiumMember `gorm:"foreignKey:PremiumMemberPeriodID"`
}
