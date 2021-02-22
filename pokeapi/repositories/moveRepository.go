package repositories

import (
	"encoding/json"
	"errors"

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

func (self *moveRepository) FindMove(move *models.Move) (*models.Move, error) {
	criteria := move.Name
	// if &move.Name == nil {
	// 	criteria = fmt.Sprint(move.ID)
	// }
	resp, err := self.Client.R().Get("/move/" + criteria)
	if err != nil {
		return nil, err
	}
	var result models.Move

	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, errors.New("Could not retrieve move, " + err.Error())
	}

	return &result, nil

}
