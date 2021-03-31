package main 

import (
	"testing"
	"api-go/models"
	"database/sql"
	"os"
	"log"
	"fmt"
)

const (
	vapiURL = "https://superheroapi.com/api"
	token  = "3426078140759658"
)

const (
	USER = "postgres"
	PASS = "123456"
	DBNAME = "postgres"
	SSLMODE = "disable"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func Connect() *sql.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASS, DBNAME, SSLMODE)
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}

func TestConnection(t *testing.T) {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return 
	}
	fmt.Println("Database connected")
}


func TestFindByNameApi(t *testing.T) {
	api := models.NewSuperHeroApi(vapiURL, token)
	rsp, err := api.FindByName("batman")
	if err != nil {
		t.Errorf("Could not make request to API:\n%s", err)
	}

	// Need to have something on the response
	if err != nil || len(rsp.Results) == 0 {
		t.Error("Failed to fetch values from API")
	}
}

func TestFindByName(t *testing.T) {

	con := Connect()
	defer con.Close()

	super := models.NewSuperRepository(con)
	nameSuper := "Batman"
	
	consultSuper := super.GetSuperName(nameSuper)
	if len(consultSuper) > 0 {
		fmt.Println("==========Que legal seu super foi encontrado =================")
	} else {
		t.Error("==========Não foi Possivel buscar esse super em banco local =================")
	}
}

func TestSuperDeleteLocal(t *testing.T) {
	nameSuper := "Batman"
		con := Connect()
	defer con.Close()

	super := models.NewSuperRepository(con)
	
	resposta := super.DeleteSuper(nameSuper)
	if resposta {
		fmt.Println("==========Super deletado  =================")
	} else {
			t.Error("==========Não foi Possivel deletar esse super=================")
		}
}

func TestFoundAllVillanLocal(t *testing.T) {
	alignment := "bad"
	con := Connect()
	defer con.Close()

	super := models.NewSuperRepository(con)
	
	consultSuper := super.GetSuperByAlignment(alignment)
	if len(consultSuper) > 0 {
		fmt.Println("==========Todos os vilões foram listados =================")
	} else {
		t.Error("==========Não foi possível realizar a busca =================")
	}
}

func TestFoundAllSupersLocal(t *testing.T) {
	alignment := "good"
	con := Connect()
	defer con.Close()

	super := models.NewSuperRepository(con)
	
	consultSuper := super.GetSuperByAlignment(alignment)
	if len(consultSuper) > 0 {
		fmt.Println("==========Todos os Supers foram listados =================")
	} else {
		t.Error("==========Não foi possível realizar a busca =================")
	}
}
