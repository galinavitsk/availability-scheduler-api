package handlers

import (
	"errors"
	"net/http"

	"github.com/galinavitsk/availability-scheduler-api/internal/models"
	"github.com/galinavitsk/availability-scheduler-api/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type sessionHandler struct {
	service *services.SessionService
}

// Create godoc
// @Summary      Create a session
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Param        session  body      models.CreateSessionRequest  true  "Session to create"
// @Success      201      {object}  models.Session
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /sessions [post]
func (handler *sessionHandler) Create(c *gin.Context) {
	var req models.CreateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := handler.service.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Successfully created the session", "data": session})
}

// GetBySlug godoc
// @Summary      Get a session by Slug
// @Tags         sessions
// @Produce      json
// @Param        slug   path      string  true  "Session Slug"
// @Success      200  {object}  models.Session
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /sessions/{slug} [get]
func (handler *sessionHandler) GetBySlug(c *gin.Context) {
	s, err := handler.service.GetBySlug(c.Request.Context(), c.Param("slug"))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully retrieved the session", "data": s})
}

// Update godoc
// @Summary      Update a session
// @Tags         sessions
// @Accept       json
// @Produce      json
// @Param        id       path      string                       true  "Session ID"
// @Param        session  body      models.UpdateSessionRequest  true  "Fields to update"
// @Success      200      {object}  models.Session
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /sessions/{id} [put]
func (handler *sessionHandler) Update(c *gin.Context) {
	var req models.UpdateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s, err := handler.service.Update(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, s)
}

// Delete godoc
// @Summary      Delete a session
// @Tags         sessions
// @Param        id   path  string  true  "Session ID"
// @Success      204
// @Failure      500  {object}  map[string]string
// @Router       /sessions/{id} [delete]
func (handler *sessionHandler) Delete(c *gin.Context) {
	if err := handler.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
