package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/domain"
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/dto"
	"github.com/rafidfadhil/Backend-Tubes-Mobile/internal/helper"
)

func (c *AssetController) Add(ctx *fiber.Ctx) error {
	var addRequest dto.AssetRequest

	if err := ctx.BodyParser(&addRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	asset := &domain.Assets{
		AssetName: addRequest.AssetName,
		AssetType: addRequest.AssetType,
		Description: addRequest.Description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	asset, err := c.assetService.Add(asset)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.AssetResponse{
		ID: asset.ID,
		AssetName: asset.AssetName,
		AssetType: asset.AssetType,
		Description: asset.Description,
		CreatedAt: asset.CreatedAt,
		UpdatedAt: asset.UpdatedAt,

	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Asset successfully added",
		"data":    response,
	})
}

func (c *AssetController) CreateDamageReport(ctx *fiber.Ctx) error {
	assetID := ctx.Params("id")
	var createDamageReportRequest dto.DamageReportRequest

	if err := ctx.BodyParser(&createDamageReportRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	estimatedTime := helper.CountEstimatedTimeToFix(createDamageReportRequest.TypeOfDamage, createDamageReportRequest.HasBeenDamagedBefore)

	damageReport := &domain.DamageReport{
		AssetID: assetID,
		TypeOfDamage: createDamageReportRequest.TypeOfDamage,
		HasBeenDamagedBefore: createDamageReportRequest.HasBeenDamagedBefore,
		EstimatedTimeToFix: estimatedTime,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	damageReport, err := c.assetService.CreateDamageReport(damageReport)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.DamageReportResponse{
		ID: damageReport.ID,
		AssetID: damageReport.AssetID,
		TypeOfDamage: damageReport.TypeOfDamage,
		HasBeenDamagedBefore: damageReport.HasBeenDamagedBefore,
		EstimatedTimeToFix: damageReport.EstimatedTimeToFix,
		CreatedAt: damageReport.CreatedAt,
		UpdatedAt: damageReport.UpdatedAt,
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Damage report successfully created",
		"data":    response,
	})
}

func (c *AssetController) Show(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	asset, err := c.assetService.Show(id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.AssetResponse{
		ID: asset.ID,
		AssetName: asset.AssetName,
		AssetType: asset.AssetType,
		Description: asset.Description,
		CreatedAt: asset.CreatedAt,
		UpdatedAt: asset.UpdatedAt,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Asset successfully retrieved",
		"data":    response,
	})
}

func (c *AssetController) ShowAll(ctx *fiber.Ctx) error {
	assets, err := c.assetService.ShowAll()

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var response []dto.AssetResponse

	for _, asset := range assets {
		response = append(response, dto.AssetResponse{
			ID: asset.ID,
			AssetName: asset.AssetName,
			AssetType: asset.AssetType,
			Description: asset.Description,
			CreatedAt: asset.CreatedAt,
			UpdatedAt: asset.UpdatedAt,
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Assets successfully retrieved",
		"data":    response,
	})
}

func (c *AssetController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var updateRequest dto.AssetUpdateRequest

	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	asset := &domain.Assets{
		AssetName: updateRequest.AssetName,
		AssetType: updateRequest.AssetType,
		Description: updateRequest.Description,
		UpdatedAt: time.Now(),
	}

	asset, err := c.assetService.Update(id, asset)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.AssetResponse{
		ID: asset.ID,
		AssetName: asset.AssetName,
		AssetType: asset.AssetType,
		Description: asset.Description,
		CreatedAt: asset.CreatedAt,
		UpdatedAt: asset.UpdatedAt,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Asset successfully updated",
		"data":    response,
	})
	
}

func (c *AssetController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	_, err := c.assetService.Delete(id)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Asset successfully deleted",
	})
}

