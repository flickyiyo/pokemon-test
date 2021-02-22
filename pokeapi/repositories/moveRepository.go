package repositories

import (
	"github.com/flickyiyo/pokemon-api/models"
	"github.com/flickyiyo/pokemon-api/pokeapi"
	"github.com/go-resty/resty/v2"
)

type moveRepository struct {
	Client resty.Client
}

func NewMoveRepository(baseUrl string) pokeapi.MoveRepository {
	client := resty.New().SetHostURL(baseUrl)
	return &moveRepository{Client: *client}
}

func (self *moveRepository) FindMove(request *models.CommonMovesRequest) (*models.Move, error) {
	// lang := request.LangID
	// pokemons := request.Pokemons
	// langName := request.LangName
	// numPage := request.NumPage

	// movesMap := make(map[string][]string)
	// for _, pokemon := range pokemons {

	// }
	return nil, nil

}
