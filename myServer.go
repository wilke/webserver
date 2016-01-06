package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/wilke/webserver/Frame"
	"github.com/wilke/webserver/MICCoM"
	"log"
	"net/http"
	"strconv"
)

type Item Frame.Item
type Collection Frame.Collection

var url string = ""
var baseURL string = ""

func init() {
	url = "http://localhost:8000"
	baseURL = url
	// i = new(Item)
	// 	c = new(Frame.Collection)
	// 	fmt.Printf("%+v\n", c)
	// 	fmt.Printf("%s\n", "Test")
}

func main() {

	fmt.Printf("%s\n", "Starting Server")

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.

	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/experiment", ExperimentHandler)
	r.HandleFunc("/experiment/{id:[a-z]*}", ExperimentHandler).Name("experiment")
	r.HandleFunc("/search", SearchHandler)
	r.HandleFunc("/search/{path:.+}", SearchHandler)
	r.HandleFunc("/register", RegisterHandler)
	r.HandleFunc("/register/{path:[a-z+]+}", RegisterHandler)
	r.HandleFunc("/upload", UploadHandler)
	r.HandleFunc("/download", DownloadHandler)
	r.HandleFunc("/transfer", TransferHandler)
	r.HandleFunc("/transfer/{id}", SearchHandler)

	// Bind to a port and pass our router in
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Printf("%s\n", "Started Server at port 8000")
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("MICCoM!\n"))
	//fmt.Printf("%+v\n", c)
	fmt.Printf("%s\n", "Test")
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	path := vars["path"]

	//w.Write([]byte(r.UserAgent))
	//println(r.UserAgent())
	fmt.Printf("%+v\n", r.UserAgent())
	println(path)
	w.Write([]byte(r.UserAgent()))
	//c.version = 1
	//fmt.Printf("%+v\n", c)
	fmt.Printf("%s\n", "Test")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	path := vars["path"]

	href := url + "register" + "/" + path

	item := Frame.Item{
		Href:  href,
		Data:  []MICCoM.Experiment{}, //[]string{"Hallo", "Du"},
		Links: "no links",
		//Queries:  nil,
		//Template: nil,
	}

	// if path != "" {
	// 		alist := item.Data.([]string)
	// 		//alist.([]string)
	// 		//append(alist, path)
	// 		item.Data = append(alist, path)
	// 	}

	experiment := MICCoM.Experiment{}
	experiment.Data.ID = "A1001"

	if path != "" {
		alist := item.Data.([]MICCoM.Experiment)
		//alist.([]string)
		//append(alist, path)
		item.Data = append(alist, experiment)
		//item.AddData(experiment)
	}

	var list []MICCoM.Experiment
	col := Frame.Collection{
		Version: "1",
		Href:    baseURL + r.URL.String(),
		Items:   list,
	}

	nr := experiment.AddItem(&col)

	//list = col.Items.([]Frame.Item)
	//col.Items = append(list, item)

	fmt.Printf("%+v\n", col.Items)
	fmt.Printf("%+v\n", "Items: "+strconv.Itoa(nr))

	register := Frame.Frame{
		Collection: col,
		ID:         1,
	}

	println("Collection", register.Collection.Version)

	jb, err := json.Marshal(register)
	if err != nil {
		println(jb)
		w.Write([]byte(err.Error()))
		http.Error(w, err.Error(), 500)
	} else {
		fmt.Printf("%s\n", jb)
		fmt.Printf("%+v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jb))

	}
}

func ExperimentHandler(w http.ResponseWriter, r *http.Request) {

	// Retrieve path and query parameters for experiment resource
	vars := mux.Vars(r)
	id := vars["id"]

	// initialize r.Form for query parameters
	r.ParseForm()

	// fmt.Printf("%s\n", "Length: "+strconv.Itoa(len(r.Form)))
	//
	// 	if id != "" {
	// 		w.Write([]byte(id))
	// 	}
	//
	// 	for key, value := range r.Form {
	// 		fmt.Printf("%s -> %s\n", key, value)
	// 		w.Write([]byte(key + "\n"))
	//
	// 	}
	//
	// 	for key, value := range vars {
	// 		fmt.Printf("%s :: %s\n", key, value)
	// 		w.Write([]byte(key + "\n"))
	//
	// 	}

	// initialize experiment
	experiment, err := MICCoM.NewExperiment(id)

	if err != nil {
		Frame.SendError(w, err)
		return
		panic(err)
	}

	var list []MICCoM.Experiment
	col := Frame.Collection{
		Version: "1",
		Href:    baseURL + r.URL.String(),
		Items:   list,
	}

	nr := experiment.AddItem(&col)

	//list = col.Items.([]Frame.Item)
	//col.Items = append(list, item)

	fmt.Printf("%+v\n", col.Items)
	fmt.Printf("%+v\n", "Items: "+strconv.Itoa(nr))

	frame := Frame.Frame{
		Collection: col,
		ID:         1,
	}

	jb, err := json.Marshal(frame)
	if err != nil {
		println(jb)
		w.Write([]byte(err.Error()))
		http.Error(w, err.Error(), 500)
	} else {
		fmt.Printf("%s\n", jb)
		fmt.Printf("%+v\n", err)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jb))

	}

}

func UploadHandler(w http.ResponseWriter, r *http.Request)   {}
func DownloadHandler(w http.ResponseWriter, r *http.Request) {}
func TransferHandler(w http.ResponseWriter, r *http.Request) {}
