package dto

import "time"

type AssetRequest struct {
	AssetName   string    `json:"asset_name"`
	AssetType   string    `json:"asset_type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AssetUpdateRequest struct {
	AssetName   string    `json:"asset_name"`
	AssetType   string    `json:"asset_type"`
	Description string    `json:"description"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AssetResponse struct {
	ID          string    `json:"id"`
	AssetName   string    `json:"asset_name"`
	AssetType   string    `json:"asset_type"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type DamageReportRequest struct {
	TypeOfDamage         string    `json:"type_of_damage"`
	HasBeenDamagedBefore bool      `json:"has_been_damaged_before"`
	EstimatedTimeToFix   int       `json:"estimated_time_to_fix"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type DamageReportResponse struct {
	ID                   string    `json:"id"`
	AssetID              string    `json:"asset_id"`
	TypeOfDamage         string    `json:"type_of_damage"`
	HasBeenDamagedBefore bool      `json:"has_been_damaged_before"`
	EstimatedTimeToFix   int       `json:"estimated_time_to_fix"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
