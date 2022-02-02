package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"
)

var db *sql.DB

type Teacher struct {
	IdTeacher    int
	Snp          string
	Post         string
	DateOfHiring time.Time
}

type Discipline struct {
	IdDiscipline    int
	HoursOfPractice int
	HoursOfLecture  int
	Disciplines     string
}

type Group struct {
	IdGroup int
	Group   string
}

type Load struct {
	IdLoad       int
	IdDiscipline int
	IdGroup      int
	IdTeacher    int
	Year         int
}

func inputLoads(l Load) error {
	_, err := db.Exec("insert into load (id_load, id_discipline, id_group, id_teacher, year) values ($1, $2, $3, $4, $5) ", l.IdLoad, l.IdDiscipline, l.IdGroup, l.IdTeacher, l.Year)
	return err
}

func deleteLoads(l Load) error {
	_, err := db.Exec("delete from load where id_load = $1", l.IdLoad)
	return err
}

func updateLoads(l Load) error {
	_, err := db.Exec("update load set id_load = $2, id_discipline = $3, id_group = $4, id_teacher=$5 where id_load = $1", l.IdLoad, l.IdDiscipline, l.IdGroup, l.IdTeacher, l.Year)
	return err
}

func outputLoads() ([]Load, error) {
	rows, err := db.Query("select * from load")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var loads []Load

	for rows.Next() {
		l := Load{}
		err := rows.Scan(&l.IdLoad, &l.IdDiscipline, &l.IdGroup, &l.IdTeacher, &l.Year)
		if err != nil {
			fmt.Println(err)
			continue
		}
		loads = append(loads, l)
	}
	return loads, err
}

func inputGroups(g Group) error {
	_, err := db.Exec("insert into grupi (id_group, group) values ($1, $2) ", g.IdGroup, g.Group)
	return err
}

func deleteGroups(g Group) error {
	_, err := db.Exec("delete from grupi where id_group = $1", g.IdGroup)
	return err
}

func updateGroups(g Group) error {
	_, err := db.Exec("update grupi set id_group = $2 where id_group = $1", g.IdGroup, g.Group)
	return err
}

func outputGroups() ([]Group, error) {
	rows, err := db.Query("select * from grupi")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var groups []Group

	for rows.Next() {
		g := Group{}
		err := rows.Scan(&g.IdGroup, &g.Group)
		if err != nil {
			fmt.Println(err)
			continue
		}
		groups = append(groups, g)
	}

	return groups, err
}

func inputDisciplines(d Discipline) error {
	_, err := db.Exec("insert into disciplines (id_discipline, hours_of_practice, hours_of_lecture, disciplines) values ($1, $2, $3, $4) ", d.IdDiscipline, d.HoursOfPractice, d.HoursOfLecture, d.Disciplines)
	return err
}

func deleteDisciplines(d Discipline) error {
	_, err := db.Exec("delete from disciplines where id_discipline = $1", d.IdDiscipline)
	return err
}

func updateDisciplines(d Discipline) error {
	_, err := db.Exec("update teachers set hours_of_practice = $2, hours_of_lecture = $3, disciplines = $4 where id_discipline = $1", d.IdDiscipline, d.HoursOfPractice, d.HoursOfLecture, d.Disciplines)
	return err
}

func outputDisciplines() ([]Discipline, error) { // вывод таблицы учителей
	rows, err := db.Query("select * from disciplines")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var disciplines []Discipline

	for rows.Next() {
		d := Discipline{}
		err := rows.Scan(&d.IdDiscipline, &d.HoursOfPractice, &d.HoursOfLecture, &d.Disciplines)
		if err != nil {
			fmt.Println(err)
			continue
		}
		disciplines = append(disciplines, d)
	}

	return disciplines, err
}

func inputTeachers(t Teacher) error {
	_, err := db.Exec("insert into teachers (id_teacher, snp, post, date_of_hiring) values ($1, $2, $3, $4) ", t.IdTeacher, t.Snp, t.Post, t.DateOfHiring)
	return err
}

func deleteTeachers(t Teacher) error {
	_, err := db.Exec("delete from teachers where id_teacher = $1", t.IdTeacher)
	return err
}

func updateTeachers(t Teacher) error {
	_, err := db.Exec("update teachers set snp = $2, post = $3, date_of_hiring = $4 where id_teacher = $1", t.IdTeacher, t.Snp, t.Post, t.DateOfHiring)
	return err
}

func outputTeachers() ([]Teacher, error) { // вывод таблицы учителей
	rows, err := db.Query("select * from teachers")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	defer rows.Close()
	var teachers []Teacher

	for rows.Next() {
		t := Teacher{}
		err := rows.Scan(&t.IdTeacher, &t.Snp, &t.Post, &t.DateOfHiring)
		if err != nil {
			fmt.Println(err)
			continue
		}
		teachers = append(teachers, t)
	}

	return teachers, err
}

func Teachers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method)
	switch r.Method {
	case "GET":
		data, err := outputTeachers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO: Проверка типа
		w.Header().Set("Content-Type", "text/html")
		t, err := template.ParseFiles("./templates/teacher.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var tec Teacher
		err = json.Unmarshal(body, &tec)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = inputTeachers(tec)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var tec Teacher
		err = json.Unmarshal(body, &tec)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = deleteTeachers(tec)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var tec Teacher
		err = json.Unmarshal(body, &tec)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = updateTeachers(tec)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func Disciplines(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method)
	switch r.Method {
	case "GET":
		data, err := outputDisciplines()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		d, err := template.ParseFiles("./templates/discipline.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = d.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var dis Discipline
		err = json.Unmarshal(body, &dis)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = inputDisciplines(dis)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var dis Discipline
		err = json.Unmarshal(body, &dis)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = deleteDisciplines(dis)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var dis Discipline
		err = json.Unmarshal(body, &dis)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = updateDisciplines(dis)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func Groups(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method)
	switch r.Method {
	case "GET":
		data, err := outputGroups()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		g, err := template.ParseFiles("./templates/group.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = g.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var gro Group
		err = json.Unmarshal(body, &gro)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = inputGroups(gro)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var gro Group
		err = json.Unmarshal(body, &gro)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = deleteGroups(gro)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var gro Group
		err = json.Unmarshal(body, &gro)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = updateGroups(gro)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func Loads(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Получен запрос: " + r.Method)
	switch r.Method {
	case "GET":
		data, err := outputLoads()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		l, err := template.ParseFiles("./templates/load.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = l.Execute(w, data)

		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var loa Load
		err = json.Unmarshal(body, &loa)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = inputLoads(loa)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "DELETE":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var loa Load
		err = json.Unmarshal(body, &loa)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = deleteLoads(loa)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	case "PATCH":
		body, err := ioutil.ReadAll(r.Body)
		fmt.Println(string(body))
		defer r.Body.Close()
		var loa Load
		err = json.Unmarshal(body, &loa)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
		}
		err = updateLoads(loa)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func main() {
	var err error
	connStr := "host=localhost port=5433 user=postgres password=1 dbname=praktika1 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/teachers", Teachers)
	http.HandleFunc("/disciplines", Disciplines)
	http.HandleFunc("/groups", Groups)
	http.HandleFunc("/loads", Loads)

	http.ListenAndServe("0.0.0.0:8081", nil)
}
