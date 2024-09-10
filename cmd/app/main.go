package main

import (
	"net/http"

	"github.com/Andhika-GIT/wild_oasis_be/internal/infrastructure/config"
)

func main() {

	r := config.Bootstrap()

	http.ListenAndServe(":3000", r)
}
