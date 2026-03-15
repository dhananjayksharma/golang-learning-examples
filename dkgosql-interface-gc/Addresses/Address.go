package addresses

import "encoding/json"

type Addresser interface {
	Add(lineOne string) Address
}

type Address struct {
	LineOne string
	zip     string
}

func NewZip(zip string) Addresser {
	return &Address{zip: zip}
}

func (a *Address) Add(lineOne string) Address {
	a.LineOne = lineOne
	return *a
}

func (a Address) MarshalJSON() ([]byte, error) {
	type addressJSON struct {
		LineOne string `json:"LineOne"`
		Zip     string `json:"zip"`
	}

	return json.Marshal(addressJSON{
		LineOne: a.LineOne,
		Zip:     a.zip,
	})
}
