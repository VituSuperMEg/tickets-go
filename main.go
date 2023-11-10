package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/VituSuperMEg/tickets-go/application/usecase"
	"github.com/VituSuperMEg/tickets-go/auth"
	"github.com/VituSuperMEg/tickets-go/domain/model"
	"github.com/VituSuperMEg/tickets-go/infra/db"
	"github.com/VituSuperMEg/tickets-go/infra/repository"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var filmUseCase *usecase.FilmUseCast
var userUseCase *usecase.UserUseCast

func main() {
	router := mux.NewRouter()

	db := db.InitDB()
	// Inicializar o Case & Repository
	filmRepository := &repository.FilmRepositoryDB{DB: db}
	filmUseCase = &usecase.FilmUseCast{FilmRepository: filmRepository}
	userRepository := &repository.UserRepositoryDB{DB: db}
	userUseCase = &usecase.UserUseCast{UserRepository: userRepository}

	fs := http.FileServer(http.Dir("public"))
	router.Handle("/", http.StripPrefix("/", fs))

	// Films
	router.HandleFunc("/films", listFilmHandler).Methods("GET")
	router.HandleFunc("/films", registerFilmHandler).Methods("POST")
	router.HandleFunc("/films/{id}", findFilmHandler).Methods("GET")
	// UserRoutes
	router.HandleFunc("/user", registerUserHandler).Methods("POST")

	// Login
	router.HandleFunc("/login", loginHandler).Methods("POST")

	// Cors
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"}) // Altere "*" para o domínio específico que você deseja permitir

	log.Fatal(http.ListenAndServe(":1010", handlers.CORS(headers, methods, origins)(router)))
}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Método não encontrado"})
		return
	}

	var requestData map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erro ao decodificar a solicitação"})
		return
	}

	username, password := requestData["login"], requestData["password"]

	user, err := userUseCase.Find(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Usuário não encontrado"})
		return
	}

	if err = comparePasswords(user.Password, password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "Credenciais inválidas"})
		return
	}

	token, err := auth.GenerateJWT(username, password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Falha ao gerar token"})
		return
	}

	response := map[string]string{"token": token, "user": username}
	json.NewEncoder(w).Encode(response)
}
func comparePasswords(hashedPwd string, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}
func registerUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var userData model.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUser, err := userUseCase.Register(userData.Login, userData.Email, userData.Password, userData.Perfil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// Films
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

	newFilm, err := filmUseCase.Register(filmData.Film_name, filmData.Film_count, filmData.Film_time, filmData.Description, filmData.ImagePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newFilm)
}
func listFilmHandler(w http.ResponseWriter, r *http.Request) {
	films, err := filmUseCase.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(films)
}
func findFilmHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filmID := params["id"]

	film, err := filmUseCase.Find(filmID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(film)
}
