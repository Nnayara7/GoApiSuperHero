package models

import (
	"strconv"
	"github.com/lib/pq"
	"api-go/types"
	"database/sql"
	"fmt"
  "strings"
)
type SuperRepository interface {
	Saving(idConection []string, super *types.AllSupers) bool
	DeleteSuper(name string) bool 
	GetSuperById(id string) ([]types.Get) 
	GetSuperName(name string) ([]types.Get) 
	GetSuperByAlignment(alignment string) ([]types.Get) 
}

type superRepositoryImpl struct {
	db *sql.DB 
}


func NewSuperRepository(db *sql.DB) SuperRepository {
	return &superRepositoryImpl{db}
}

func (repo *superRepositoryImpl) GetSuperById(id string) []types.Get {
  sql := "select name , intelligence ,power from super where uuid = $1"
  rs, err := repo.db.Query(sql, id)
  if err != nil {
    fmt.Println(err.Error())
        return []types.Get{}  
  }
   fmt.Println("entrou aqui ")
  var super types.Get
  super.Id = id
  for rs.Next() {
    err := rs.Scan(&super.Name, &super.Intelligence, &super.Power)
    if err != nil {
      return []types.Get{}
    }
    fmt.Println("Name : ",super.Name)

    sql = "select full_name from biography  where uuid = $1"
    rs, err = repo.db.Query(sql, id)
    _ = rs.Next()  
        _ = rs.Scan(&super.Full_Name)
     fmt.Println("Full Name : ",super.Full_Name)

    sql = "select occupation from work  where uuid = $1"
    rs, err = repo.db.Query(sql, id)
    _ = rs.Next()  
        _ = rs.Scan(&super.Occupation)
     fmt.Println("Occupation : ",super.Occupation)

    sql = "select url from image where uuid = $1"
    rs, err = repo.db.Query(sql, id)
    _ = rs.Next()  
        _ = rs.Scan(&super.Url)
     fmt.Println("Image : ",super.Url)

    sql = "select group_affiliation, relatives  from connections where uuid = $1"
    rs, err = repo.db.Query(sql, id)
    _ = rs.Next()  
        _ = rs.Scan(&super.Group_Affiliation, &super.Relatives)
     fmt.Println("Group_Affiliation: ",super.Group_Affiliation)
    numberRelatives := len(strings.Split(super.Relatives,","))-1 
    fmt.Println("Number Of Relatives : ", numberRelatives)     
    dataSuper := make([]types.Get,0)
        dataSuper = append(dataSuper, super)  
        return dataSuper 
  }
  fmt.Println("Nonexistent id")
  return []types.Get{}
}

func (repo *superRepositoryImpl) GetSuperName(name string) []types.Get {
  sql := "select uuid from super  where name ilike $1"
  rs, err := repo.db.Query(sql, name)
  if err != nil {
    fmt.Println(err.Error())
        return []types.Get{}
  }
  var (uuid   int)
  dataSearch := make([]types.Get,0)
  for rs.Next() {
    err := rs.Scan(&uuid)
    fmt.Println(err)
    if err != nil {
      fmt.Println("err.Error()")
            return dataSearch  
    }
    search := repo.GetSuperById(strconv.Itoa(uuid))
        if len(search)>0{
           dataSearch = append(dataSearch,search[0])   
            
        }
        continue
  }
  return dataSearch
}

func (repo *superRepositoryImpl) GetSuperByAlignment(alignment string) []types.Get {
  sql := "select uuid from biography where alignment ilike $1"
  rs, err := repo.db.Query(sql, alignment)
  if err != nil {
    fmt.Println(err.Error())
        return []types.Get{}
  }
  var id int
  dataSearch := make([]types.Get,0)
  for rs.Next() {
    err := rs.Scan(&id)
    if err != nil {
      fmt.Println(err.Error())
            return dataSearch  
    }
    search := repo.GetSuperById(strconv.Itoa(id))
        if len(search)>0{
           dataSearch = append(dataSearch,search[0])   
            
        }
        continue
  }
  return dataSearch
}


func (repo *superRepositoryImpl) erasingTheSuper(id int ){
	fmt.Println("aqui")
    _,_ = repo.db.Exec(`DELETE FROM appearance WHERE uuid = $1`, id)
    _,_ = repo.db.Exec(`DELETE FROM biography WHERE uuid = $1`, id)
    _,_ = repo.db.Exec(`DELETE FROM connections WHERE uuid = $1`, id)
    _,_ = repo.db.Exec(`DELETE FROM image WHERE uuid = $1`, id)
    _,_ = repo.db.Exec(`DELETE FROM work WHERE uuid = $1`, id)
    _,_ = repo.db.Exec(`DELETE FROM super WHERE uuid = $1`, id)
}

func (repo *superRepositoryImpl) DeleteSuper(name string) bool{
	sql := `SELECT  uuid FROM super WHERE name =$1`
	rs, err := repo.db.Query(sql,name)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	
	var id int 
	for rs.Next(){
        erroScan:=rs.Scan(&id)
        if erroScan != nil{
          fmt.Println(erroScan.Error())
          return false
        }
        fmt.Print(".")
        repo.erasingTheSuper(id)
    }
    return true
}

func (repo *superRepositoryImpl) RemoveDataBase(){

    _,err := repo.db.Exec("drop schema public cascade")
    if err != nil{
        panic(err.Error())
    } 
    fmt.Print(".")
    _,err = repo.db.Exec("create schema public")
    if err != nil{
        panic(err.Error())
    } 
    fmt.Println(".")
}

// a partir daqui meu

func (repo *superRepositoryImpl) Saving(idConection []string, super *types.AllSupers) bool {
  
  if len(idConection) == 0 {
    err := fmt.Errorf("there is no id's")
    panic(err)
  }
  fmt.Print(".")

  records, err := repo.db.Query("SELECT uuid FROM super")
  if err != nil {
    panic(err)
  }
  if !records.Next(){
      fmt.Println("successfully tiled")
      repo.NewSuper(super,0)
  }
  if err != nil {
    panic(err)
  }
  
  records, err = repo.db.Query("SELECT uuid FROM super")
  var id types.IdApi
  arrayId := make([]types.IdApi,0)
  //salva todos os ids que estão no database
  for records.Next(){
    erroScan := records.Scan(&id.Id)
    if erroScan != nil {
      fmt.Println(".")
      fmt.Println("id verification error")
      fmt.Println(erroScan.Error())
      continue
    }
      arrayId = append(arrayId,id)
    }
    //checa se o id do Nome do Super passado 
    flag :=  checkValidityId(idConection,arrayId)
    if flag{
    	for i:=1 ; i<len(idConection) ; i++{
        	repo.NewSuper(super,i)
        }
        fmt.Println(".")
        fmt.Println("Registrado com sucesso")
        return true
    }
     fmt.Println("Super Já está Registrado")
     return false
}

func checkValidityId(idConection []string, arrayId []types.IdApi) (bool) {
    flag := false
    for i:=0 ; i<len(idConection) ; i++{
      for j:=0 ; j<len(arrayId);j++{
          if idConection[i] == strconv.Itoa(arrayId[j].Id){
            flag = false
            break
          }
          flag = true
          continue
      }
      if flag{
        return true
      }
    }
    return false
} 

func (repo *superRepositoryImpl) NewSuper(super *types.AllSupers, posResponse int)  {

	convertIntId, _ := strconv.Atoi(super.Results[posResponse].Id)
	
	_, err := repo.db.Exec(`INSERT INTO super (uuid, name, intelligence, strength, speed, durability, power, combat) VALUES ($1,$2, $3,$4,$5,$6,$7, $8)`, convertIntId, super.Results[posResponse].Name,  super.Results[posResponse].Powerstats.Intelligence,  super.Results[posResponse].Powerstats.Strength, super.Results[posResponse].Powerstats.Speed, super.Results[posResponse].Powerstats.Durability, super.Results[posResponse].Powerstats.Power,  super.Results[posResponse].Powerstats.Combat)
	 if err != nil {
    	panic(err)
  	 }
	fmt.Println("primeiroinserção")

	_, err = repo.db.Exec(`INSERT INTO biography (uuid, full_name, alter_ego, aliases, place_birth, first_appearance, publisher, alignment) VALUES ($1,$2, $3,$4,$5,$6,$7, $8)`,	convertIntId, super.Results[posResponse].Biography.Full_Name, super.Results[posResponse].Biography.Alter_Ego,pq.Array(super.Results[posResponse].Biography.Aliases),super.Results[posResponse].Biography.Place_Birth, super.Results[posResponse].Biography.First_Appearance,super.Results[posResponse].Biography.Publisher, super.Results[posResponse].Biography.Alignment)
	if err != nil {
    	panic(err)
  	 }
	
	fmt.Println("insertBiography")

	_, err = repo.db.Exec(`insert into appearance (uuid, gender, race, height, weight, eye_color, hair_color) values ($1,$2, $3,$4,$5,$6,$7)`, convertIntId, super.Results[posResponse].Appearance.Gender, super.Results[posResponse].Appearance.Race, pq.Array(super.Results[posResponse].Appearance.Height), pq.Array(super.Results[posResponse].Appearance.Weight), super.Results[posResponse].Appearance.Eye_color, super.Results[posResponse].Appearance.Hair_color)
	if err != nil {
    	panic(err)
  	 }
	
	fmt.Println("insertAppearance")

	_, err = repo.db.Exec(`insert into work (uuid, occupation, base) values ($1,$2, $3)`, convertIntId, super.Results[posResponse].Work.Occupation, super.Results[posResponse].Work.Base)
	if err != nil {
    	panic(err)
  	 }
	
	fmt.Println("insertWork")

	_, err = repo.db.Exec(`insert into connections (uuid, group_affiliation, relatives) values ($1,$2, $3)`, convertIntId, super.Results[posResponse].Connections.Group_affiliation, super.Results[posResponse].Connections.Relatives)
	if err != nil {
    	panic(err)
  	 }
	
	fmt.Println("insertConnections")

	_, err = repo.db.Exec(`insert into image (uuid, url) values ($1,$2)`, convertIntId, super.Results[posResponse].Image.Url)
	if err != nil {
    	panic(err)
  	 }
	
	fmt.Println("insertImage")
}