package handler

import (
	"errors"
	"net/http"
	"strconv"

	"avaliacao-ii/internal/domain"
	"avaliacao-ii/internal/patient"
	"avaliacao-ii/pkg/web"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

func (h *patientHandler) ReadById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		patient, err := h.s.ReadById(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, patient)
	}
}

func (h *patientHandler) ReadByRg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		patient, err := h.s.ReadByRg(ctx.Param("rg"))
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, patient)
	}
}

func (h *patientHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patient domain.Patient
		err := ctx.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyValuesPatient(&patient)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		createdPatient, err := h.s.Create(patient)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdPatient)
	}
}

func (h *patientHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		var patient domain.Patient
		err = ctx.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptyValuesPatient(&patient)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		createdPatient, err := h.s.Update(id, patient)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdPatient)
	}
}

func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name         string `json:"name,omitempty"`
		LastName     string `json:"last_name,omitempty"`
		Rg           string `json:"rg,omitempty"`
		RegisterDate string `json:"register_date,omitempty"`
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

		update := domain.Patient{
			Name:         request.Name,
			LastName:     request.LastName,
			Rg:           request.Rg,
			RegisterDate: request.RegisterDate,
		}

		updatedPatient, err := h.s.Patch(id, update)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}

		web.Success(ctx, http.StatusOK, updatedPatient)
	}
}

func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}

		web.Success(ctx, http.StatusNoContent, nil)
	}
}

func validateEmptyValuesPatient(patient *domain.Patient) (bool, error) {
	switch {
	case patient.Name == "" || patient.LastName == "" || patient.Rg == "" || patient.RegisterDate == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}
