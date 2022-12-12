package handler

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"avaliacao-ii/internal/dentist"
	"avaliacao-ii/internal/domain"
	"avaliacao-ii/pkg/web"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

func (h *dentistHandler) ReadById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		dentist, err := h.s.ReadById(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, dentist)
	}
}

func (h *dentistHandler) ReadByEnrollment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentist, err := h.s.ReadByEnrollment(ctx.Param("enrollment"))
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, dentist)
	}
}

func (h *dentistHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentist domain.Dentist
		err := ctx.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyValues(&dentist)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		createdDentist, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdDentist)
	}
}

func (h *dentistHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		var dentist domain.Dentist
		err = ctx.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyValues(&dentist)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		createdDentist, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdDentist)
	}
}

func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name       string `json:"name,omitempty"`
		LastName   string `json:"last_name,omitempty"`
		Enrollment string `json:"enrollment,omitempty"`
	}
	return func(ctx *gin.Context) {
		var request Request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}

		update := domain.Dentist{
			Name:       request.Name,
			LastName:   request.LastName,
			Enrollment: request.Enrollment,
		}

		updatedDentist, err := h.s.Patch(id, update)
		if err.Error() == "dentist not found" {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}

		web.Success(ctx, http.StatusOK, updatedDentist)
	}
}

func (h *dentistHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err.Error() == "dentist not found" {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}

		web.Success(ctx, http.StatusNoContent, nil)
	}
}

func validateEmptyValues(dentist *domain.Dentist) (bool, error) {
	switch {
	case dentist.Name == "" || dentist.LastName == "" || dentist.Enrollment == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
