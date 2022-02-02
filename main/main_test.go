package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

type Teacher struct {
	IdTeacher    int
	Snp          string
	Post         string
	DateOfHiring time.Time
}

func Test_inputTeachers(t *testing.T) {
	client := http.Client{}
	te := Teacher{IdTeacher: 9, Snp: "9", Post: "9", DateOfHiring: time.Date(2000, 12, 12, 0, 0, 0, 0, time.Local)}
	data, _ := json.Marshal(te)

	r := bytes.NewReader([]byte(data))
	req, err := http.NewRequest("POST", "http://localhost:8081/dbtable", r)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("CODE: %d BODY: %s", resp.StatusCode, string(body))

}

func Test_deleteTeachers(t *testing.T) {
	client := http.Client{}
	te := Teacher{IdTeacher: 9}
	data, _ := json.Marshal(te)
	r := bytes.NewReader([]byte(data))
	req, err := http.NewRequest("DELETE", "http://localhost:8081/dbtable", r)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("CODE: %d BODY: %s", resp.StatusCode, string(body))
}

func Test_updateTeachers(t *testing.T) {
	client := http.Client{}
	te := Teacher{IdTeacher: 8, Snp: "9", Post: "9", DateOfHiring: time.Date(2012, 12, 12, 0, 0, 0, 0, time.Local)}
	data, _ := json.Marshal(te)
	r := bytes.NewReader([]byte(data))
	req, err := http.NewRequest("PATCH", "http://localhost:8081/dbtable", r)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("CODE: %d BODY: %s", resp.StatusCode, string(body))
}
