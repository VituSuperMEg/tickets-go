package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VituSuperMEg/tickets-go/application/usecase"
	"github.com/VituSuperMEg/tickets-go/domain/model"
	"github.com/VituSuperMEg/tickets-go/infra/db"
	"github.com/VituSuperMEg/tickets-go/infra/repository"
	"github.com/gorilla/mux"
)

var filmUseCase *usecase.FilmUseCast

func main() {
	router := mux.NewRouter()

	db := db.InitDB()
	filmRepository := &repository.FilmRepositoryDB{DB: db}
	filmUseCase = &usecase.FilmUseCast{FilmRepository: filmRepository}

	router.HandleFunc("/films", registerFilmHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":1010", router))
}

func registerFilmHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var filmData model.Film
	err := json.NewDecoder(r.Body).Decode(&filmData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newFilm, err := filmUseCase.Register(filmData.Film_name, filmData.Film_count, filmData.Film_time)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newFilm)
}
