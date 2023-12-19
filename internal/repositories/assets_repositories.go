package repositories

import (
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/domain"
	"gorm.io/gorm"
)

type AssetsRepository interface {
	Insert(asset *domain.Assets) (*domain.Assets, error)
	FindByID(id string) (*domain.Assets, error)
	FindAll() ([]domain.Assets, error)
	Update(id string, assets *domain.Assets) (*domain.Assets, error)
	Delete(id string) (*domain.Assets, error)
	AddDamageReport(damageReport *domain.DamageReport) (*domain.DamageReport, error)
	AssetExists(assetID string) (bool, error)
}

type assetsRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) *assetsRepository {
	return &assetsRepository{db}
}

func (r *assetsRepository) Insert(asset *domain.Assets) (*domain.Assets, error) {
	err := r.db.Create(&asset).Error

	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (r *assetsRepository) FindByID(id string) (*domain.Assets, error) {
	var asset domain.Assets

	err := r.db.Where("id = ?", id).First(&asset).Error

	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetsRepository) FindAll() ([]domain.Assets, error) {
	var assets []domain.Assets

	err := r.db.Find(&assets).Error

	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (r *assetsRepository) Update(id string, asset *domain.Assets) (*domain.Assets, error) {
	err := r.db.Where("id = ?", id).Updates(&asset).Error

	if err != nil {
		return nil, err
	}

	return asset, nil
}

func (r *assetsRepository) Delete(id string) (*domain.Assets, error) {
	var asset domain.Assets

	err := r.db.Where("id = ?", id).Delete(&asset).Error

	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (r *assetsRepository) AddDamageReport(damageReport *domain.DamageReport) (*domain.DamageReport, error) {
	err := r.db.Create(&damageReport).Error

	if err != nil {
		return nil, err
	}

	return damageReport, nil
}

func (r *assetsRepository) AssetExists(assetID string) (bool, error) {
	var asset domain.Assets

	err := r.db.Where("id = ?", assetID).First(&asset).Error

	if err != nil {
		return false, err
	}

	return true, nil
}