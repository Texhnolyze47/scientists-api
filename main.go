package main

import (
	"cientificos/internal/database"
	"database/sql"
	"encoding/json"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

type Scientist struct {
	Dni      string `json:"dni"`
	Nomapels string `json:"nomApels"`
}

type Project struct {
	Id     string `json:"id"`
	Nombre string `json:"nombre"`
	Horas  int32  `json:"horas"`
}

type AsignadoA struct {
	Dni        string `json:"dni"`
	Id         string `json:"id"`
	Dedicacion int32  `json:"dedicacion"`
}

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUrl := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbUrl)

	queries := database.New(conn)

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	router := http.NewServeMux()

	router.HandleFunc("POST /scientists", func(writer http.ResponseWriter, request *http.Request) {

		var scientist Scientist

		err := json.NewDecoder(request.Body).Decode(&scientist)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		newScientist, err := queries.CreateScientist(request.Context(), database.CreateScientistParams{
			Dni:      scientist.Dni,
			Nomapels: scientist.Nomapels,
		})

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(newScientist)
	})

	router.HandleFunc("POST /projects", func(writer http.ResponseWriter, request *http.Request) {

		var project Project

		err := json.NewDecoder(request.Body).Decode(&project)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		newProject, err := queries.CreateProject(request.Context(), database.CreateProjectParams{
			ID:     project.Id,
			Nombre: project.Nombre,
			Horas:  project.Horas,
		})

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(newProject)

	})

	router.HandleFunc("POST /asginado_a", func(writer http.ResponseWriter, request *http.Request) {

		var asignadoA AsignadoA

		err := json.NewDecoder(request.Body).Decode(&asignadoA)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		newAsignadoA, err := queries.CreateProyectoAsignado(request.Context(), database.CreateProyectoAsignadoParams{
			Cientifico: asignadoA.Dni,
			Proyecto:   asignadoA.Id,
			Dedicacion: sql.NullInt32{Int32: asignadoA.Dedicacion, Valid: true},
		})

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(newAsignadoA)

	})

	router.HandleFunc("GET /consulta/1", func(writer http.ResponseWriter, request *http.Request) {

		consulta, err := queries.InfoProjectsScientists(request.Context())

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(consulta)

	})

	router.HandleFunc("GET /consulta/2", func(writer http.ResponseWriter, request *http.Request) {

		consulta, err := queries.NumberProjectsScientist(request.Context())

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(consulta)
	})

	router.HandleFunc("GET /consulta/3", func(writer http.ResponseWriter, request *http.Request) {

		consulta, err := queries.NumberScientistsProject(request.Context())

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(consulta)
	})

	router.HandleFunc("GET /consulta/4", func(writer http.ResponseWriter, request *http.Request) {

		consulta, err := queries.NumberHourScientificProject(request.Context())

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(consulta)
	})

	router.HandleFunc("GET /consulta/5", func(writer http.ResponseWriter, request *http.Request) {

		consulta, err := queries.MoreEightyHoursProject(request.Context())

		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(writer).Encode(consulta)
	})

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))

}
