package reader

import (
	"fmt"
	"reflect"
	"testing"
)

type StubInputServicer struct {
}

func (s StubInputServicer) Data(url string) ([]byte, error) {
	if url != "" {
		return []byte(url), nil
	}
	return nil, fmt.Errorf("Error")
}

func TestFetch_String_ShoulReturn_ByteArray(t *testing.T) {
	reader := Reader{
		Servicer: StubInputServicer{},
	}

	expectedResult := []byte("getInput")
	actualResult, err := reader.Fetch("getInput")
	if err != nil {
		t.Fatal("Expected no error")
	}

	if !reflect.DeepEqual(expectedResult, actualResult) {
		t.Fatalf("Expected %v but got %v", expectedResult, actualResult)
	}
}

func TestFetch_EmptyString_ShoulReturn_Error(t *testing.T) {
	reader := Reader{
		Servicer: StubInputServicer{},
	}

	expectedResult := "Unable to fetch the data from url : Error"
	actualResult, err := reader.Fetch("")

	if actualResult != nil {
		t.Fatal("actualResult must be nil")
	}

	if err != nil && err.Error() != expectedResult {
		t.Fatal("Error message is not expected")
	}
}
