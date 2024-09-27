package web

import "github.com/Andhika-GIT/wild_oasis_be/internal/domain/entities"

type CabinResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MaxCapacity  int    `json:"max_capacity"`
	RegulerPrice int    `json:"regular_price"`
	Discount     int    `json:"discount,omitempty"`
	Description  string `json:"description"`
	Image        string `json:"image"`
}

func ToCabinResponse(cabin entities.Cabin) CabinResponse {
	return CabinResponse{
		ID:           cabin.ID,
		Name:         cabin.Name,
		MaxCapacity:  cabin.MaxCapacity,
		RegulerPrice: cabin.RegulerPrice,
		Discount:     cabin.Discount,
		Description:  cabin.Description,
		Image:        cabin.Image,
	}
}

func ToCabinResponses(cabins []entities.Cabin) []CabinResponse {
	var responses []CabinResponse

	for _, cabin := range cabins {
		responses = append(responses, ToCabinResponse(cabin))
	}

	return responses
}
