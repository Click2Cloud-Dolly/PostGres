package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type employee struct {
	gorm.Model
	ID       int
	Username string
	MobNo    int
	Location string
	Position string
}

var db *gorm.DB

var err error
var (
	emp = []employee{
		{ID: 1, Username: "Jimmy", MobNo: 4551, Location: "XYZ", Position: "ABC"},
		{ID: 2, Username: "John", MobNo: 4551, Location: "rys", Position: "def"},
		{ID: 3, Username: "jasmin", MobNo: 4551, Location: "Xye", Position: "hik"},
		{ID: 4, Username: "Jeet", MobNo: 4551, Location: "Ysw", Position: "kuy"},
	}
)

func main() {
	router := mux.NewRouter()
	db, err = gorm.Open("postgres", "host=192.168.3.141 port=5432 user=postgres dbname=postgres sslmode=disable password=somePassword")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&employee{})

	for index := range emp {

		db.Create(&emp[index])

	}

	router.HandleFunc("/Allemps", AllEmps).Methods("GET")

	router.HandleFunc("/Allemps/{id}", GetEmp).Methods("GET")

	router.HandleFunc("/Allemps/add", AddEmps).Methods("POST")

	router.HandleFunc("/Allemps/{id}", DeleteEmp).Methods("DELETE")

	router.HandleFunc("/Allemps/update", UpdateEmps).Methods("PUT")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))

}

func AllEmps(w http.ResponseWriter, r *http.Request) {

	var emp []employee

	db.Find(&emp)

	json.NewEncoder(w).Encode(&emp)

}

func GetEmp(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	var emp employee

	db.First(&emp, params["id"])

	json.NewEncoder(w).Encode(&emp)

}

func AddEmps(w http.ResponseWriter, r *http.Request) {

	var emp employee
	json.NewDecoder(r.Body).Decode(&emp)

	NewEmp := db.Create(&emp)
	err = NewEmp.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&emp)
	}

}

func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var emp employee

	db.First(&emp, params["id"])

	db.Delete(&emp)

	var emps []employee

	db.Find(&emps)

	json.NewEncoder(w).Encode(&emps)

}

func UpdateEmps(w http.ResponseWriter, r *http.Request) {

	var emp employee
	db.Model(&emp).Where("Username = ?", "Jimmy").Update("Location", "Nagpur")

	json.NewEncoder(w).Encode(&emp)

}
