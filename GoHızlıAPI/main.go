package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	ID    int    `json:"id"`
	FName string `json:"firstname"`
	LName string `json:"lastname"`
	Info  *Info  `json:"info"`
}

type Info struct {
	City string `json:"city"`
	Job  string `json:"job"`
}

var people []Person

func Hata(err error) {
	log.Fatalln("Hata var hocam dikkat et bak", err.Error())
}

// Herkesi getirir
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// Seçili kişiyi getirir
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	Hata(err)
	for _, item := range people {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// Yeni birisini oluşturur
func PostPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	id, err := strconv.Atoi(params["id"])
	Hata(err)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&person)
	Hata(err)
	person.ID = id
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// Seçili Kişiyi Günceller
// // func PutPerson(w http.ResponseWriter, r *http.Request) {
// // 	params := mux.Vars(r)
// // 	id, err := strconv.Atoi(params["id"])
// // 	Hata(err)
// // 	for index, item := range people {
// // 		if item.ID == id {
// // 			people = append(people[:index], people[index+1]...)
// // 			var person Person
// // 			json.NewDecoder(r.Body).Decode(&person)
// // 			person.ID = id
// // 			people = append(people, person)
// // 			json.NewEncoder(w).Encode(people)
// // 			return
// // 		}
// // 		json.NewEncoder(w).Encode(people)
// // 	}
// // }

func PutPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	Hata(err)
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			var person Person
			json.NewDecoder(r.Body).Decode(&person)
			person.ID = id
			people = append(people, person)
			json.NewEncoder(w).Encode(people)
			return
		}
		json.NewEncoder(w).Encode(people)
	}
}

// Seçili Kişiyi Siler
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	Hata(err)
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {

	r := mux.NewRouter()

	people = append(people, Person{ID: 1, FName: "Yasin", LName: "Yumrutaş", Info: &Info{City: "Bursa", Job: "Jr.Developer"}})
	people = append(people, Person{ID: 2, FName: "Enver", LName: "Yumrutaş", Info: &Info{City: "İstanbul", Job: "Sn.Developer"}})
	people = append(people, Person{ID: 3, FName: "Tyler", LName: "Durdan", Info: &Info{City: "New York", Job: "Club Manangement"}})

	r.HandleFunc("/people", GetPeople).Methods("GET")
	r.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	r.HandleFunc("/people/{id}", PostPerson).Methods("POST")
	r.HandleFunc("/people/{id}", PutPerson).Methods("PUT")
	r.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	err := http.ListenAndServe(":1000", r)
	Hata(err)
}
