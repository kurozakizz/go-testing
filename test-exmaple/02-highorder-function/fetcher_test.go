package fetcher

import (
	"testing"
)

var data = func(url string) (string, error) {
	return url, nil
}

func TestFetch_Test_ShouldReturn_Test(t *testing.T) {
	actualResult, err :=fetchData(data, "Test")
	if err != nil {
		t.Fatal("Expected no error")
	}

	expectedResult := "Test"
	if actualResult != expectedResult {
		t.Fatalf("Expected %s but actual is %s",expectedResult, actualResult)
	}
}
