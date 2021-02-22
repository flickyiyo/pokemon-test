package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/flickyiyo/pokemon-api/pokeapi/repositories"
	"github.com/flickyiyo/pokemon-api/pokeapi/services"
	"github.com/flickyiyo/pokemon-api/rest"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	baseUrl := "https://pokeapi.co/api/v2/"
	moveRepo := repositories.NewMoveRepository(baseUrl)
	typeRepo := repositories.NewTypeRepository(baseUrl)
	pokeRepo := repositories.NewPokemonRepository(baseUrl)

	matchService := services.NewMatchService(pokeRepo, typeRepo)
	// moveService := services.NewMovesService(pokeRepo, moveRepo)
	typeService := services.NewTypeService(&typeRepo)

	matchHandler := rest.NewMatchHandler(matchService, typeService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)

	router.Get("/match/{pokemonId}/{pokemonId}", matchHandler.Get)
	router.Post("/moves/common/")

	errs := make(chan error, 2)

	go func() {
		fmt.Println("Listening on port " + string(port))
		errs <- http.ListenAndServe(port, router)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println("Terminated %s", <-errs)
}
