package main

import (
	"net/http"
	"strconv"

	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi/repositories"
	"github.com/flickyiyo/pokemon-api/pokeapi/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// if os.Getenv("PORT") != "" {
	// 	port = os.Getenv("PORT")
	// }
	baseURL := "https://pokeapi.co/api/v2/"
	// moveRepo := repositories.NewMoveRepository(baseURL)
	typeRepo := repositories.NewTypeRepository(baseURL)
	pokeRepo := repositories.NewPokemonRepository(baseURL)

	matchService := services.NewMatchService(pokeRepo, typeRepo)
	// moveService := services.NewMovesService(pokeRepo, moveRepo)
	// typeService := services.NewTypeService(&typeRepo)

	// matchHandler := rest.NewMatchHandler(matchService, typeService)

	router := gin.Default()
	router.GET("/match/:pokemonA/:pokemonB", func(c *gin.Context) {
		idA := c.Params.ByName("pokemonA")
		idB := c.Params.ByName("pokemonB")
		var pokemonA models.Pokemon
		var pokemonB models.Pokemon
		if val, err := strconv.Atoi(idA); err != nil {
			pokemonA.Name = idA
		} else {
			pokemonA.ID = val
		}

		if val, err := strconv.Atoi(idB); err != nil {
			pokemonB.Name = idB
		} else {
			pokemonB.ID = val
		}

		pokeMatch := models.PokemonMatch{Pokemons: []models.Pokemon{
			pokemonA, pokemonB,
		}}

		response, err := matchService.MatchPokemons(&pokeMatch)
		if err != nil {
			c.String(http.StatusInternalServerError, "Invalid match response %s", err.Error())
		}
		c.JSON(200, response)

	})

	router.Run()
}
