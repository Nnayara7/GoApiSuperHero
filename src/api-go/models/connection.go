package models

import (
	"errors"
	"api-go/types"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	// "os"
)

type SuperHeroApi struct {
	ApiUrl string
	Token  string
}

type superHeroServiceImpl struct {
	api  *SuperHeroApi
	repo SuperRepository
}

func NewSuperHeroService(api *SuperHeroApi, repo SuperRepository) types.SuperHeroService {
	return &superHeroServiceImpl{api, repo}
} 

func NewSuperHeroApi(apiUrl string, token string) *SuperHeroApi {
	return &SuperHeroApi{apiUrl, token}
}

func (api *SuperHeroApi) buildRequestPath(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", api.ApiUrl, api.Token, endpoint)
}

func (api *SuperHeroApi) FindByName(name string) (*types.AllSupers, error) {
	tgt := api.buildRequestPath("search/" + name)
	resp, err := http.Get(tgt)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result types.AllSupers
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *superHeroServiceImpl) AddSuper(name string) error {
	if name == "" {
		return errors.New("Super name is required")
	}

	response, err := s.api.FindByName(name)

	if err != nil {
		panic(err.Error())
		return nil
	}

	fmt.Print("Buscando...")
	ids := make([]string,0)
	for i:=0; i<len(response.Results);i++{
		ids = append(ids,response.Results[i].Id)
	}
	s.repo.Saving(ids, response)

	return nil
}

func (s *superHeroServiceImpl) ListSuper(superType string) ([]types.Get) {
	r := s.repo.GetSuperByAlignment(superType)
	if len(r) > 0 {
 		return r
	}
	return nil
}

func (s *superHeroServiceImpl) FindSuperByName(name string) ([]types.Get) {
	r :=  s.repo.GetSuperName(name)
	if len(r) > 0 {
 		return r
	}
	return nil
}

func (s *superHeroServiceImpl) FindSuperById(id string) ([]types.Get) {
	r :=  s.repo.GetSuperById(id)
	if len(r) > 0 {
 		return r
	}
	return nil
}

func (s *superHeroServiceImpl) RemoveSuper(nameSuper string) error {
	fmt.Print("Loading .")
	if s.repo.DeleteSuper(nameSuper){
		fmt.Println("complete removal")
		return nil
	}
	fmt.Println("there was an error, maybe the super does not exist")
	return nil
}
