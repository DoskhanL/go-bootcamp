package helper

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/google/uuid"
)

// OperatePoints parsing data from csv
func OperatePoints() {

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	static := path.Join(wd, "static")
	operateCsvFile(static)
}

func operateCsvFile(path string) {
	f, err := os.Open(path + "/atmsst.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	var wholeBody []byte
	wholeBody = append(wholeBody, []byte(`{"elements":[`)...)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		point := NewPoint(row)
		if point == nil {
			continue
		}

		// fmt.Printf("%#+v\n", point)

		body, err := json.Marshal(point)
		if err != nil {
			fmt.Printf("err while parse %s\n", err.Error())
		}
		fmt.Printf("%s\n", body)
		wholeBody = append(wholeBody, body...)
		delimiter := []byte(",")
		wholeBody = append(wholeBody, delimiter...)
		fmt.Println("-------------------------------------------------------------------------------------------------------------")
	}

	wholeBody = append(wholeBody, []byte("]}")...)
	uID := uuid.New()
	fileName := strings.Replace(uID.String(), "-", "", -1) + ".json"
	err = ioutil.WriteFile(fileName, wholeBody, os.ModePerm)
	if err != nil {
		fmt.Printf("err while saving to file %s\n", err.Error())
	}
}

// JSONMarshal encodes
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
