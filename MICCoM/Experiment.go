package MICCoM

import (
	//"fmt"
	"errors"
	"fmt"
	"github.com/wilke/webserver/Frame"
	"time"
)

type Data struct {
	Name     string
	ID       string
	Version  string
	Date     string
	Duration string
	Files    []string
	Samples  []string
}

type Experiment struct {
	Frame.Item
	Data Data `json:"data"`
}

func NewExperiment(id string) (Experiment, error) {

	var e Experiment

	if id != "" {

		t := time.Now()
		fmt.Print(t)

		e = Experiment{}
		e.Data.ID = id
		e.Data.Version = "1"
		e.Data.Date = time.Now().Format(time.ANSIC)
	} else {
		return e, errors.New("Can't initialize experiment, no ID given")
	}

	return e, nil
}

func (e Experiment) GetItem() Experiment {

	var err error

	e, err = NewExperiment(e.Data.ID)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	return e
}

func (e Experiment) AddData(c interface{}) {
	//var alist []ExperimentStruct
	t := c.(Frame.Collection)
	alist := t.Items.([]Experiment)
	t.Items = append(alist, e)

}

// Add experiment to items list in collection
func (e Experiment) AddItem(c *Frame.Collection) int {
	//var alist []ExperimentStruct

	alist := c.Items.([]Experiment)
	c.Items = append(alist, e)

	return len(c.Items.([]Experiment))
}
