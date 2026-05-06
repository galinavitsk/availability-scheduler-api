package handlers

import (
	"errors"
	"net/http"

	"github.com/galinavitsk/availability-scheduler-api/internal/models"
	"github.com/galinavitsk/availability-scheduler-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type availabilityHandler struct {
	service *services.AvailabilityService
}

// CreateAvailability godoc
// @Summary      Create a availability
// @Tags         availability
// @Accept       json
// @Produce      json
// @Param        availability  body      models.CreateAvailabilityRequest  true  "Availability to create"
// @Success      201      {object}  models.Availability
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /availability [post]
func (handler *availabilityHandler) CreateAvailability(c *gin.Context) {
	var req models.CreateAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	availability, err := handler.service.CreateAvailability(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created the availability", "data": availability})
}

// GetAllAvailabilitiesForSlug godoc
// @Summary      Get a all names and icons by Slug
// @Tags         availability
// @Produce      json
// @Param        slug   path      string  true  "Availability Slug"
// @Success      200  {object}  models.Availability
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /availability/{slug} [get]
func (handler *availabilityHandler) GetAllAvailabilitiesForSlug(c *gin.Context) {
	s, err := handler.service.GetAllAvailabilitiesForSlug(c.Request.Context(), c.Param("slug"))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "availability not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully retrieved the availability", "data": s})
}

// UpdateAvailability gdoc
// @Summary      Update a availability
// @Tags         availability
// @Accept       json
// @Produce      json
// @Param        id       path      string                       true  "Availability ID"
// @Param        availability  body      models.UpdateAvailabilityRequest  true  "Fields to update"
// @Success      200      {object}  models.Availability
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /availability/{id} [put]
func (handler *availabilityHandler) UpdateAvailability(c *gin.Context) {
	var req models.UpdateAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := handler.service.UpdateAvailability(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "availability not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated the availability", "data": s})
}

/*
// Update godoc
// @Summary      Update a availability
// @Tags         availabilitys
// @Accept       json
// @Produce      json
// @Param        id       path      string                       true  "Availability ID"
// @Param        availability  body      models.UpdateAvailabilityRequest  true  "Fields to update"
// @Success      200      {object}  models.Availability
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /availabilitys/{id} [put]
func (handler *availabilityHandler) Update(c *gin.Context) {
	var req models.UpdateAvailabilityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := handler.service.Update(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "availability not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s)
}

// Delete godoc
// @Summary      Delete a availability
// @Tags         availabilitys
// @Param        id   path  string  true  "Availability ID"
// @Success      204
// @Failure      500  {object}  map[string]string
// @Router       /availabilitys/{id} [delete]
func (handler *availabilityHandler) Delete(c *gin.Context) {
	if err := handler.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
*/
