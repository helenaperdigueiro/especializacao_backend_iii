package handler

import (
	"errors"
	"net/http"
	// "os"
	"strconv"
	// "strings"

	"avaliacao-ii/internal/appointment"
	"avaliacao-ii/pkg/web"
	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	s appointment.Service
}

func NewAppointmentHandler(s appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		s: s,
	}
}

func (h *appointmentHandler) ReadById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		appointment, err := h.s.ReadById(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, appointment)
	}
}

// func (h *appointmentHandler) CreateById() gin.HandlerFunc {
	
// }

// func (h *appointmentHandler) CreateByRgEnrollment() gin.HandlerFunc {
	
// }

// func (h *appointmentHandler) Update() gin.HandlerFunc {
	
// }

// func (h *appointmentHandler) Patch() gin.HandlerFunc {
	
// }

// func (h *appointmentHandler) Delete() gin.HandlerFunc {
	
// }
