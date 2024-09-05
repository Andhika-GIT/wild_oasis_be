package main

import (
	"net/http"

	api "github.com/Andhika-GIT/wild_oasis_be/internal/api/handlers"
	"github.com/Andhika-GIT/wild_oasis_be/internal/infrastructure/config"
)


func main(){
	v := config.NewViper()
	l := config.NewLogger()
	config.NewDatabase(v, &l)
	r:= api.NewRouter()

	http.ListenAndServe(":3000", r)
}