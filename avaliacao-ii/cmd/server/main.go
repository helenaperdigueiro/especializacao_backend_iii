package main

import (
	"avaliacao-ii/cmd/server/handler"
	"avaliacao-ii/internal/patient"
	"avaliacao-ii/internal/dentist"
	"avaliacao-ii/internal/appointment"
	"avaliacao-ii/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	sqlStore := store.NewSQLStore();

	sqlStoragePatient := store.NewSQLStorePatient(sqlStore)
	repoPatient := patient.NewRepository(sqlStoragePatient)
	servicePatient := patient.NewService(repoPatient)
	patientHandler := handler.NewPatientHandler(servicePatient)

	patients := r.Group("/patients")
	{
		patients.GET("/id/:id", patientHandler.ReadById())
		patients.GET("/rg/:rg", patientHandler.ReadByRg())
		patients.POST("", patientHandler.Create())
		patients.PUT(":id", patientHandler.Update())
		patients.PATCH(":id", patientHandler.Patch())
		patients.DELETE(":id", patientHandler.Delete())
	}

	sqlStorageDentist := store.NewSQLStoreDentist(sqlStore)
	repoDentist := dentist.NewRepository(sqlStorageDentist)
	serviceDentist := dentist.NewService(repoDentist)
	dentistHandler := handler.NewDentistHandler(serviceDentist)

	dentists := r.Group("/dentists")
	{
		dentists.GET("/id/:id", dentistHandler.ReadById())
		dentists.GET("/enrollment/:enrollment", dentistHandler.ReadByEnrollment())
		dentists.POST("", dentistHandler.Create())
		dentists.PUT(":id", dentistHandler.Update())
		dentists.PATCH(":id", dentistHandler.Patch())
		dentists.DELETE(":id", dentistHandler.Delete())
	}

	sqlStorageAppointment := store.NewSQLStoreAppointment(sqlStore)
	repoAppointment := appointment.NewRepository(sqlStorageAppointment)
	serviceAppointment := appointment.NewService(repoAppointment)
	appointmentHandler := handler.NewAppointmentHandler(serviceAppointment)

	appointments := r.Group("/appointments")
	{
		appointments.GET("/id/:id", appointmentHandler.ReadById())
		appointments.GET("/rg/:rg", appointmentHandler.ReadByRg())
		appointments.POST("/id/:id-patient/:id-dentist", appointmentHandler.CreateById())
		appointments.POST("/rg-enrollment/:rg-patient/:enrollment-dentist", appointmentHandler.CreateByRgEnrollment())
		appointments.PUT(":id", appointmentHandler.Update())
		appointments.PATCH(":id", appointmentHandler.Patch())
		appointments.DELETE(":id", appointmentHandler.Delete())
	}

	r.Run(":8080")
}
