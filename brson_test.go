package golang_brson

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestReadBrson(t *testing.T) {
	data, err := os.ReadFile("./857a7be3908d78496f8577f35b611c339707904e4d0bacfa826288acea5aaa33")
	if err != nil {
		log.Fatal(err)
		return
	}
	res, err := DecodeBrsonJson(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf(string(res) + "\n")
	brson, err := EncodeJsonBrson(res)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("RESULT -> %v\n", brson)
	file, err := os.OpenFile("output.brson", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	_, err = file.Write(brson)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func TestEncodeJsonBrson(t *testing.T) {
	data, err := os.ReadFile("./input.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	brson, err := EncodeJsonBrson(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	file, err := os.OpenFile("output.brson", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	_, err = file.Write(brson)
	if err != nil {
		log.Fatal(err)
		return
	}
}
