package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type user struct {
	gorm.Model
	UserId   int
	Username string
	mobNo    int
	Location string
	Position string
}

var db *gorm.DB
var err error
var (
	emp = []user{
		{UserId: 1, Username: "JEET", mobNo: 1234, Location: "XXRX1", Position: "SDE"},
		{UserId: 2, Username: "Preet", mobNo: 1223, Location: "XEXRT1", Position: "SDE"},
		{UserId: 3, Username: "John", mobNo: 1230, Location: "XFF1TT", Position: "SDE"},
		{UserId: 4, Username: "Janny", mobNo: 1294, Location: "Xf1SS", Position: "SDE"},
		{UserId: 5, Username: "Meet", mobNo: 1834, Location: "XX31", Position: "SDE"},
		{UserId: 6, Username: "Geet", mobNo: 7234, Location: "X2X1", Position: "SDE"},
	}
)

func main() {
	router := mux.NewRouter()
	db, err = gorm.Open("postgres", "host=192.168.3.141 port=5432 user=postgres dbname=postgres sslmode=disable password=somePassword")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&user{})
	for index := range emp {
		db.Create(&emp[index])
	}
	router.HandleFunc("/Allemps", AllEmps).Methods("GET")
	router.HandleFunc("/Allemps/{id}", GetEmp).Methods("GET")
	router.HandleFunc("/Allemps/{id}", DeleteEmp).Methods("DELETE")
	router.HandleFunc("/Allemps/add", AddEmps).Methods("POST")
	router.HandleFunc("/Allemps/update", UpdateEmps1).Methods("PUT")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func AllEmps(w http.ResponseWriter, r *http.Request) {
	var emp []user
	db.Find(&emp)
	json.NewEncoder(w).Encode(&emp)
}
func GetEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var emp user
	db.First(&emp, params["id"])
	json.NewEncoder(w).Encode(&emp)
}
func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var emp user
	db.First(&emp, params["id"])
	db.Delete(&emp)
	var emps []user
	db.Find(&emps)
	json.NewEncoder(w).Encode(&emps)
}
func AddEmps(w http.ResponseWriter, r *http.Request) {
	var emp user
	json.NewDecoder(r.Body).Decode(&emp)
	NewEmp := db.Create(&emp)
	err = NewEmp.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&emp)
	}
}
func UpdateEmps1(w http.ResponseWriter, r *http.Request) {
	var emp user
	//db.First(&product, "Username = ?", "rkf") // find product with code D42
	db.Model(&emp).Where("Username = ?", "Preet").Update("Location", "Nagpur")
	json.NewEncoder(w).Encode(&emp)
}
