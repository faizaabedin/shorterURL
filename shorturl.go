package main

// 1st read long urls
// do the has ids
// and then create the short ids and store then back from the database
// check each time if the url is there if not create a hash -
// redirect shorturl to longurl

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

type Url struct {
	hashID   string `bson:"hashID"`
	longURL  string `bson:"longURL"`
	shortURL string `bson:"shortURL"`
}

// func readLongURL() []string {
//
// }
// func ExpandEndpoint(w http.ResponseWriter, req *http.Request) {
//     var n1qlParams []interface{}
//     query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "` WHERE shortUrl = $1")
//     params := req.URL.Query()
//     n1qlParams = append(n1qlParams, params.Get("shortUrl"))
//     rows, _ := bucket.ExecuteN1qlQuery(query, n1qlParams)
//     var row MyUrl
//     rows.One(&row)
//     json.NewEncoder(w).Encode(row)
// }
//
// func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
//     var url MyUrl
//     _ = json.NewDecoder(req.Body).Decode(&url)
//     var n1qlParams []interface{}
//     n1qlParams = append(n1qlParams, url.LongUrl)
//     query := gocb.NewN1qlQuery("SELECT `" + bucketName + "`.* FROM `" + bucketName + "` WHERE longUrl = $1")
//     rows, err := bucket.ExecuteN1qlQuery(query, n1qlParams)
//     if err != nil {
//         w.WriteHeader(401)
//         w.Write([]byte(err.Error()))
//         return
//     }
//     var row MyUrl
//     rows.One(&row)
//     if row == (MyUrl{}) {
//         hd := hashids.NewData()
//         h := hashids.NewWithData(hd)
//         now := time.Now()
//         url.ID, _ = h.Encode([]int{int(now.Unix())})
//         url.ShortUrl = "http://localhost:12345/" + url.ID
//         bucket.Insert(url.ID, url, 0)
//     } else {
//         url = row
//     }
//     json.NewEncoder(w).Encode(url)
// }

func RootEndpoint(w http.ResponseWriter, req *http.Request) {
	url := req.URL.String() //
	fmt.Println(strings.TrimPrefix(url, "/"))
	session, _ := mgo.Dial("localhost")
	URLcollection := session.DB("local").C("urls")
	// var urls []Url
	// _ = URLcollection.Find(bson.M{}).All(&urls)
	// fmt.Println(urls)
	var urls []Url

	err := URLcollection.Find(bson.M{}).All(&urls)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("Results All: ")
		fmt.Println(urls)
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{url}", RootEndpoint).Methods("GET")
	// router.HandleFunc("/expand/", ExpandEndpoint).Methods("GET")
	// router.HandleFunc("/create", CreateEndpoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
