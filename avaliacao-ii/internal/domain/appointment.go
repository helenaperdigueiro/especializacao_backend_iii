package domain

type Appointment struct {
	Id          int     `json:"id"`
	Patient     Patient `json:"name"`
	Dentist     Dentist `json:"last_name"`
	Date        string  `json:"date" binding:"required"`
	Description string  `json:"description" binding:"required"`
}
