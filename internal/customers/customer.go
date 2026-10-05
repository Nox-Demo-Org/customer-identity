package customers

import "time"

// Customer is the profile returned by GET /v1/customers/{id}. Table: customers.
type Customer struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Postcode      string    `json:"postcode"`
	CustomerSince time.Time `json:"customer_since"`
	TenureYears   int       `json:"tenure_years"`
}

// Tenure counts whole years since CustomerSince. Gaps under 90 days between policies do not reset it
// (ADR-0004); a longer gap starts the count again.
func Tenure(since, now time.Time) int {
	years := now.Year() - since.Year()
	if now.YearDay() < since.YearDay() {
		years--
	}
	if years < 0 {
		return 0
	}
	return years
}
