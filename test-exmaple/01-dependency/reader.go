package reader

import (
	"fmt"
)

type InputServicer interface {
	Data(url string) ([]byte, error)
}

type Reader struct {
	Servicer InputServicer
}

func (r Reader) Fetch(url string) ([]byte, error) {
	data, err := r.Servicer.Data(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to fetch the data from url : %v", err)
	}
	return data, nil
}
