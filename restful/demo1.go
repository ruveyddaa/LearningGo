package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"` // bunu bir restful yapı gibi tanımladık int den sonra yazdığımız kısım apı de neye karşılık geldiğini gösteriyo
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {

	// get yapma yolu bunun için http paketini kullınıcaz ve get methodunu uyhukuyucaz get methoduna paramtre polarak neyi çağırcağımızı yazıcaz
	response, err := http.Get("https://jsonplaceholder.typicode.com/todo/1")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close() // burda defer ile mutlaka response değikeninin kapanmasını garanti ettik

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todo Todo // struct objsi oluşturduk

	json.Unmarshal(bodyBytes, &todo) // json tipini todoya unmarshall ile çeviricez
	fmt.Println(todo)

}

func Demo2() {

	todo := Todo{1, 2, "alışveriş yapılacak", false}
	jsonTodo, err := json.Marshal(todo) // todo yu jsaon tipinie dönüştürücez marshall ile

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer((jsonTodo))) // 1. post fonmkiyonu bu adresebir kayıt ekler 2. datanın formatını belirtir buradki utf 8 formatı 3. de gödereceğeimiz data
	if err != nil {

		fmt.Println(err)
	}
	defer response.Body.Close()

}
