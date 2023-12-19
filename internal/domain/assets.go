package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assets struct {
	ID          string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	AssetName   string    `json:"asset_name"`
	AssetType   string    `json:"asset_type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (a *Assets) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return
}

type DamageReport struct {
	ID                   string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	AssetID              string    `json:"asset_id"`
	Assets               Assets    `gorm:"foreignKey:AssetID"`
	TypeOfDamage         string    `json:"type_of_damage"`
	HasBeenDamagedBefore bool      `json:"has_been_damaged_before"`
	EstimatedTimeToFix   int       `json:"estimated_time_to_fix"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

func (d *DamageReport) BeforeCreate(tx *gorm.DB) (err error) {
	d.ID = uuid.NewString()
	return
}
