package Frame

import (
	"encoding/json"
	"fmt"
	//"github.com/gorilla/mux"
	"net/http"
)

type Itemer interface {
	AddData(d interface{})
	AddItem(c *Collection) int
	GetItem(i interface{}) interface{}
}

type url string

type Item struct {
	Href     string      `json:"href"`
	Data     interface{} `json:"data"`
	Links    string      `json:"links"`
	Queries  string      `json:"queries"`
	Template string      `json:"template"`
}

type Collection struct {
	Version string      `json:"version"`
	Href    string      `json:"href"`
	Items   interface{} `json:"items"`
}

type Frame struct {
	Collection Collection `json:"collection"`
	ID         int
	Previous   url
	Next       url
	Offset     int
	Count      int
	Limit      int
}

func (i Item) AddData(d Itemer) {

	alist := i.Data.([]interface{})
	i.Data = append(alist, d)

}

func SendError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), 500)
}

func (c Collection) ToJson() ([]byte, error) {

	jb, err := json.Marshal(c)
	if err != nil {

	} else {
		fmt.Printf("%s\n", jb)
		fmt.Printf("%+v\n", err)

	}

	return jb, err
}
