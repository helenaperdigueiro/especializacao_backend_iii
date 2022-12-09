package main

import (
	"avaliacao-ii/cmd/server/handler"
	"avaliacao-ii/internal/patient"
	"avaliacao-ii/internal/dentist"
	"avaliacao-ii/internal/appointment"
	"avaliacao-ii/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(".Env can't be load")
	}

	sqlStore := store.NewSQLStore();

	sqlStoragePatient := store.NewSQLStorePatient(sqlStore)
	repoPatient := patient.NewRepository(sqlStoragePatient)
	servicePatient := patient.NewService(repoPatient)
	patientHandler := handler.NewPatientHandler(servicePatient)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	patients := r.Group("/patients")
	{
		patients.GET(":id", patientHandler.ReadById())
		patients.POST("", patientHandler.Create())
		patients.PUT("", patientHandler.Update())
		patients.PATCH("", patientHandler.Patch())
		patients.DELETE(":id", patientHandler.Delete())
	}

	r.Run(":8080")
}
