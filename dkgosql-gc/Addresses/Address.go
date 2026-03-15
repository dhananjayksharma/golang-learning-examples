package addresses

import "encoding/json"

type Addresser interface {
	Add(lineone string) Addresser
	GetLineOne() string
	GetZip() string
}

type Address struct {
	lineOne string
	zip     string
}

// constructor
func NewAddress(zip string) Addresser {
	return &Address{
		zip: zip,
	}
}

// add line
func (a *Address) Add(lineone string) Addresser {
	a.lineOne = lineone
	return a
}

func (a *Address) GetLineOne() string {
	return a.lineOne
}

func (a *Address) GetZip() string {
	return a.zip
}

// custom JSON marshal
func (a *Address) MarshalJSON() ([]byte, error) {
	type Alias struct {
		LineOne string `json:"line_one"`
		Zip     string `json:"zip"`
	}

	return json.Marshal(Alias{
		LineOne: a.lineOne,
		Zip:     a.zip,
	})
}
