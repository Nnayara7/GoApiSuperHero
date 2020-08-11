package types

import (
	"net/http"
)

type SuperHeroService interface {
	AddSuper(name string) error
	ListSuper(superType string) ([]Get)
	FindSuperById(id string) ([]Get)
	FindSuperByName(name string) ([]Get)
	RemoveSuper(name string) error
}

type SuperHeroCtrl interface {
	AddSuper(w http.ResponseWriter, r *http.Request)
	RemoveSuper(w http.ResponseWriter, r *http.Request)
	ListSuper(w http.ResponseWriter, r *http.Request)
	FindSuperById(w http.ResponseWriter, r *http.Request)
	FindSuperByName(w http.ResponseWriter, r *http.Request)
}

type AllSupers struct {
	Response    string 	`json:"response"`
	Result_for  string	`json:"results-for"`
	Results 	[]*Results 	`json:"results"`
}

type Results struct{
	Id 			string  	`json:"id"`
	Name 		string 		`json:"name"`
	Powerstats  Powerstats 	`json:"powerstats"`
	Biography   Biography 	`json:"biography"`
	Appearance  Appearance  `json:"appearance"`
	Work 		Work		`json:"Work"`
	Connections Connections `json:"connections"`
	Image 		Image 		`json:"image"`
}

type IdApi struct {
	Id 		 int  	    `json:"id"`
}


type UUID struct {
	UUID 		 int  	    `json:"uuid"`
}


type Powerstats struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

type Appearance struct {
	IdSuper 	int      `json:"id"`
	UID         string   `json:"uuid"`
	Gender      string   `json:"gender"`
	Race    	string   `json:"race"`
	Height		[]string 	 `json:"height"`
	Weight		[]string 	 `json:"weight'`
	Eye_color	string 	 `json:"eye-color'`
	Hair_color	string 	 `json:"hair-color'`
}

type Biography struct {
	IdSuper 		 int  	    `json:"id"`
	UID              string     `json:"uuid"`
	Full_Name        string   	`json:"full-name"`
	Aliases    		 []string    `json:"aliases"`
	Alter_Ego         string    `json:"alter-ego"`
	Place_Birth		 string 	`json:"place-of-birth"`
	First_Appearance string 	`json:"first-appearance'`
	Publisher		string 		`json:publisher`
	Alignment		string 		`json:alignment`
}


type Connections struct {
	IdSuper 		 	int  	   `json:"id"`
	UID              	string     `json:"uuid"`
	Group_affiliation   string     `json:"group-affiliation"`
	Relatives    		string     `json:"relatives"`
}

type Image struct {
	IdSuper 		 int  	    `json:"id"`
	UID              string     `json:"uuid"`
	Url        		string   	`json:"url"`
}

type Work struct {
	IdSuper 		 int  	    `json:"id"`
	UID              string     `json:"uuid"`
	Occupation        string   	`json:"occupation"`
	Base    		 string     `json:"base"`
}

type Get struct{
	Id 			 	 string
	Name 			 string `json:"name"`
	Full_Name 	 	 string
	Intelligence 	 string
	Power 		 	 string
	Occupation 	 	 string
	Url 		 	 string
	Group_Affiliation string
	Relatives 	 	 string
}


