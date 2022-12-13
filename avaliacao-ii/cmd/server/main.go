package main

import (
	"avaliacao-ii/cmd/server/handler"
	// "avaliacao-ii/internal/patient"
	"avaliacao-ii/internal/dentist"
	// "avaliacao-ii/internal/appointment"
	"avaliacao-ii/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	sqlStore := store.NewSQLStore();

	// sqlStoragePatient := store.NewSQLStorePatient(sqlStore)
	// repoPatient := patient.NewRepository(sqlStoragePatient)
	// servicePatient := patient.NewService(repoPatient)
	// patientHandler := handler.NewPatientHandler(servicePatient)

	// r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// patients := r.Group("/patients")
	// {
	// 	patients.GET(":id", patientHandler.ReadById())
	// 	patients.POST("", patientHandler.Create())
	// 	patients.PUT("", patientHandler.Update())
	// 	patients.PATCH("", patientHandler.Patch())
	// 	patients.DELETE(":id", patientHandler.Delete())
	// }

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

	// sqlStorageAppointment := store.NewSQLStoreAppointment(sqlStore)
	// repoAppointment := appointment.NewRepository(sqlStorageAppointment)
	// serviceAppointment := appointment.NewService(repoAppointment)
	// appointmentHandler := handler.NewAppointmentHandler(serviceAppointment)

	// r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// appointments := r.Group("/appointments")
	// {
	// 	appointments.GET(":id", appointmentHandler.ReadById())
	// 	appointments.POST("/id", appointmentHandler.CreateById())
	// 	appointments.POST("/rg-enrollment", appointmentHandler.CreateByRgEnrollment())
	// 	appointments.PUT("", appointmentHandler.Update())
	// 	appointments.PATCH("", appointmentHandler.Patch())
	// 	appointments.DELETE(":id", appointmentHandler.Delete())
	// }

	r.Run(":8080")
}
