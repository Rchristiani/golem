package pokemon

import (
	"net/http"
	"encoding/json"
)

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Base_experience int `json:"base_experience"`
	Height int `json:"height"`
	Is_default bool `json:"is_default"`
	Order int `json:"order"`
	Weight int `json:"weight"`
}

func GetPokemon(id string) (pokemon Pokemon) {
	const pokeApi string = "http://pokeapi.co/api/v2/pokemon/"

	pokeRes, pokeErr := http.Get(pokeApi + id)

	if pokeErr != nil {
		panic("No Pokemon!")
	}

	defer pokeRes.Body.Close()

	pokemon = new(Pokemon)

	json.NewDecoder(pokeRes.Body).Decode(&pokemon)

	return pokemon
}