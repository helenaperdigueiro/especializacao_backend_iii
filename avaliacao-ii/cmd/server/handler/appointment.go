package handler

import (
	"errors"
	"net/http"

	// "os"
	"strconv"
	// "strings"

	"avaliacao-ii/internal/appointment"
	"avaliacao-ii/internal/domain"
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

func (h *appointmentHandler) ReadByRg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rg := ctx.Param("rg")
		appointments, err := h.s.ReadByRg(rg)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, appointments)
	}
}

func (h *appointmentHandler) CreateById() gin.HandlerFunc {
	type Request struct {
		Date        string `json:"date" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var request Request
		idPatient, err := strconv.Atoi(ctx.Param("id-patient"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid patient id"))
			return
		}
		idDentist, err := strconv.Atoi(ctx.Param("id-dentist"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid dentist id"))
			return
		}

		err = ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		appointment := domain.Appointment{
			Date:        request.Date,
			Description: request.Description,
		}
		valid, err := validateEmptyValuesAppointment(&appointment)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		createdAppointment, err := h.s.CreateById(appointment, idPatient, idDentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdAppointment)
	}
}

func (h *appointmentHandler) CreateByRgEnrollment() gin.HandlerFunc {
	type Request struct {
		Date        string `json:"date" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var request Request
		rgPatient := ctx.Param("rg-patient")
		enrollmentDentist := ctx.Param("enrollment-dentist")

		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		appointment := domain.Appointment{
			Date:        request.Date,
			Description: request.Description,
		}
		valid, err := validateEmptyValuesAppointment(&appointment)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		createdAppointment, err := h.s.CreateByRgEnrollment(appointment, rgPatient, enrollmentDentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdAppointment)
	}
}

func (h *appointmentHandler) Update() gin.HandlerFunc {
	type Request struct {
		PatientId   int `json:"patient_id"`
		DentistId   int `json:"dentist_id"`
		Date        string `json:"date"`
		Description string `json:"description"`
	}
	return func(ctx *gin.Context) {
		var request Request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		if request.PatientId == 0 || request.DentistId == 0 || request.Date == "" || request.Description == "" {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		updateRequestAppointment := domain.Appointment{
			Patient: domain.Patient{
				Id: request.PatientId,
			},
			Dentist: domain.Dentist{
				Id: request.DentistId,
			},
			Date:        request.Date,
			Description: request.Description,
		}
		updatedAppointment, err := h.s.Update(id, updateRequestAppointment)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, updatedAppointment)
	}
}

func (h *appointmentHandler) Patch() gin.HandlerFunc {
	type Request struct {
		PatientId   int `json:"patient_id"`
		DentistId   int `json:"dentist_id"`
		Date        string `json:"date"`
		Description string `json:"description"`
	}
	return func(ctx *gin.Context) {
		var request Request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		updateRequestAppointment := domain.Appointment{
			Patient: domain.Patient{
				Id: request.PatientId,
			},
			Dentist: domain.Dentist{
				Id: request.DentistId,
			},
			Date:        request.Date,
			Description: request.Description,
		}
		updatedAppointment, err := h.s.Patch(id, updateRequestAppointment)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, updatedAppointment)
	}
}

func (h *appointmentHandler) Delete() gin.HandlerFunc {
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

func validateEmptyValuesAppointment(appointment *domain.Appointment) (bool, error) {
	switch {
	case appointment.Date == "" || appointment.Description == "":
		return false, errors.New("date and description can't be empty")
	}
	return true, nil
}
