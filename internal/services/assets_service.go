package services

import (
	"net/http"

	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/config"
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/domain"
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/repositories"
)

type AssetsServices interface {
	Add(asset *domain.Assets) (*domain.Assets, error)
	CreateDamageReport(damageReport *domain.DamageReport) (*domain.DamageReport, error)
	Show(id string) (*domain.Assets, error)
	ShowAll() ([]domain.Assets, error)
	Update(id string, asset *domain.Assets) (*domain.Assets, error)
	Delete(id string) (*domain.Assets, error)
}

type assetsServices struct {
	assetsRepository repositories.AssetsRepository
}

func NewAssetsServices(assetsRepository repositories.AssetsRepository) *assetsServices {
	return &assetsServices{assetsRepository}
}

func (s *assetsServices) Add(asset *domain.Assets) (*domain.Assets, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAssetRepository(conn)

	// insert asset to database
	asset, err = repo.Insert(asset)

	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to insert asset to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return asset, nil
}

func (s *assetsServices) CreateDamageReport(damageReport *domain.DamageReport) (*domain.DamageReport, error) {
	conn, err := config.Connect()
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAssetRepository(conn)

	// Check if the Asset exists
	exists, err := repo.AssetExists(damageReport.AssetID)
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to check asset existence",
			StatusCode: http.StatusInternalServerError,
		}
	}
	if !exists {
		return nil, &ErrorMessages{
			Message: "Asset not found",
			StatusCode: http.StatusNotFound,
		}
	}

	// Insert damage report to database
	damageReport, err = repo.AddDamageReport(damageReport)
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to insert damage report to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return damageReport, nil
}

func (s *assetsServices) Show(id string) (*domain.Assets, error) {
	conn, err := config.Connect()
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAssetRepository(conn)

	// Check if the Asset exists
	exists, err := repo.AssetExists(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to check asset existence",
			StatusCode: http.StatusInternalServerError,
		}
	}
	if !exists {
		return nil, &ErrorMessages{
			Message: "Asset not found",
			StatusCode: http.StatusNotFound,
		}
	}

	// Find asset by id
	asset, err := repo.FindByID(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to find asset by id",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return asset, nil
}

func (s *assetsServices) ShowAll() ([]domain.Assets, error) {
	conn, err := config.Connect()
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAssetRepository(conn)

	// Find all assets
	assets, err := repo.FindAll()
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to find all assets",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return assets, nil
}

func (s *assetsServices) Update(id string, asset *domain.Assets) (*domain.Assets, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAssetRepository(conn)

	// Check if the Asset exists
	exists, err := repo.AssetExists(id)

	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to check asset existence",
			StatusCode: http.StatusInternalServerError,
		}
	}

	if !exists {
		return nil, &ErrorMessages{
			Message: "Asset not found",
			StatusCode: http.StatusNotFound,
		}
	}

	// Update asset by id
	asset, err = repo.Update(id, asset)
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to update asset by id",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return asset, nil
}

func (s *assetsServices) Delete(id string) (*domain.Assets, error) {
	conn, err := config.Connect()
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to connect to database",
			StatusCode: http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAssetRepository(conn)

	// Check if the Asset exists
	exists, err := repo.AssetExists(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to check asset existence",
			StatusCode: http.StatusInternalServerError,
		}
	}
	if !exists {
		return nil, &ErrorMessages{
			Message: "Asset not found",
			StatusCode: http.StatusNotFound,
		}
	}

	// Delete asset by id
	asset, err := repo.Delete(id)
	if err != nil {
		return nil, &ErrorMessages{
			Message: "Failed to delete asset by id",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return asset, nil
}

