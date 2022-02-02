package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"time"
)

type group struct {
	id_group int
	group    string
}

type discipline struct {
	id_discipline     int
	hours_of_practice int
	hours_of_lecture  int
	disciplines       string
}

type teacher struct {
	id_teacher     int
	snp            string
	post           string
	date_of_hiring time.Time
}

type load struct {
	id_load       int
	id_discipline int
	id_group      int
	id_teacher    int
	year          int
}

type table struct {
	table string
}

var ingN string

var ingN1 string
var ingN2 string
var ingN3 string
var ingN4 string

// подключение к бд
func main() {
	connStr := "host=localhost port=5432 user=postgres password=1 dbname=praktika1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	progN := true
	for progN != false {
		rows, err := db.Query("select tablename from pg_tables where schemaname = 'public'") // вывод таблиц в бд
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		tables := []table{}

		for rows.Next() {
			t := table{}
			err := rows.Scan(&t.table)
			if err != nil {
				fmt.Println(err)
				continue
			}
			tables = append(tables, t)
		}
		fmt.Println("| Номер таблицы | таблица |")
		i := 0
		for _, d := range tables {
			i += 1
			fmt.Println("|", i, "|", d.table, "|")
		}
		fmt.Println("| Выберите таблицу | 0 - выйти |")
		fmt.Print("-> ")
		fmt.Scanf("%s\n", &ingN)
		tableN, err := strconv.Atoi(ingN)
		if err != nil {
			log.Fatal(err)
		}
		if tableN == 1 {
			strN := true
			for strN != false { // вывод таблицы учителей
				rows, err := db.Query("select * from teachers")
				if err != nil {
					log.Fatal(err.Error())
					panic(err)
				}
				defer rows.Close()
				teachers := []teacher{}

				for rows.Next() {
					t := teacher{}
					err := rows.Scan(&t.id_teacher, &t.snp, &t.post, &t.date_of_hiring)
					if err != nil {
						fmt.Println(err)
						continue
					}
					teachers = append(teachers, t)
				}
				fmt.Println("| Номер учителя | ФИО | должность | дата наема |")
				for _, t := range teachers {
					fmt.Println("|", t.id_teacher, "|", t.snp, "|", t.post, "|", t.date_of_hiring.Format("02.01.2006"), "|")
				}

				fmt.Println("| 1 - добавить | 2 - удалить | 3 - изменить | 0 - выйти |")
				fmt.Print("-> ")
				fmt.Scanf("%s\n", &ingN1)
				funcN, err := strconv.Atoi(ingN1)
				if err != nil {
					log.Fatal(err)
				}

				if funcN == 1 { // добавление
					t := teacher{}
					fmt.Print("Номер:")
					fmt.Scan(&t.id_teacher)
					fmt.Print("Имя:")
					fmt.Scan(&t.snp)
					fmt.Print("Должность:")
					fmt.Scan(&t.post)
					fmt.Print("Введите дату: ")
					var dateString = ""
					fmt.Scan(&dateString)
					result, err := db.Exec("insert into teachers (id_teacher, snp, post, date_of_hiring) values ($1, $2, $3, $4::date) ", t.id_teacher, t.snp, t.post, dateString)
					if err != nil {
						panic(err)

					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 2 { // удаление
					t := teacher{}
					fmt.Print("Номер учителя: ")
					fmt.Scan(&t.id_teacher)
					result, err := db.Exec("delete from teachers where id_teacher = $1", t.id_teacher)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 3 { // обновление
					t := teacher{}
					fmt.Print("Выберите номер: ")
					fmt.Scan(&t.id_teacher)
					fmt.Print("Имя: ")
					fmt.Scan(&t.snp)
					fmt.Print("Должность: ")
					fmt.Scan(&t.post)
					fmt.Print("Дата принятия: ")
					var dateString = ""
					fmt.Scan(&dateString)
					result, err := db.Exec("update teachers set snp = $2, post = $3::date, date_of_hiring = $4 where id_teacher = $1", t.id_teacher, t.snp, t.post, dateString)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 0 {
					strN = false
				}
			}
		}
		if tableN == 2 {
			strN := true
			for strN != false {
				// вывод таблицы нагрузки
				rows, err := db.Query("select * from load")
				if err != nil {
					log.Fatal(err.Error())
					panic(err)
				}
				defer rows.Close()
				loads := []load{}

				for rows.Next() {
					l := load{}
					err := rows.Scan(&l.id_load, &l.id_discipline, &l.id_group, &l.id_teacher, &l.year)
					if err != nil {
						fmt.Println(err)
						continue
					}
					loads = append(loads, l)
				}
				fmt.Println("| Номер нагрузки | номер дисциплины | номер группы | номер учителя | год |")
				for _, l := range loads {
					fmt.Println("|", l.id_load, "|", l.id_discipline, "|", l.id_group, "|", l.id_teacher, "|", l.year, "|")
				}

				fmt.Println("| 1 - добавить | 2 - удалить | 3 - изменить | 0 - выйти |")
				fmt.Print("-> ")
				fmt.Scanf("%s\n", &ingN2)
				funcN, err := strconv.Atoi(ingN2)
				if err != nil {
					log.Fatal(err)
				}

				if funcN == 1 { // добавление
					l := load{}
					print("Номер нагрузки: ")
					fmt.Scan(&l.id_load)
					print("Номер дисциплины: ")
					fmt.Scan(&l.id_discipline)
					print("Номер группы: ")
					fmt.Scan(&l.id_group)
					print("Номер учителя")
					fmt.Scan(&l.id_teacher)
					print("Год: ")
					fmt.Scan(&l.year)
					result, err := db.Exec("insert into load (id_load, id_discipline, id_group, id_teacher, year) values ($1, $2, $3, $4, $5) ", l.id_load, l.id_discipline, l.id_group, l.id_teacher, l.year)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 2 { // удаление
					l := load{}
					fmt.Print("Номер нагрузки: ")
					fmt.Scan(&l.id_load)
					result, err := db.Exec("delete from load where id_load = $1", l.id_load)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 3 { // обновление
					l := load{}
					print("Номер нагрузки: ")
					fmt.Scan(&l.id_load)
					print("Номер дисциплины: ")
					fmt.Scan(&l.id_discipline)
					print("Номер группы: ")
					fmt.Scan(&l.id_group)
					print("Номер учителя")
					fmt.Scan(&l.id_teacher)
					print("Год: ")
					fmt.Scan(&l.year)
					result, err := db.Exec("update load set id_load = $2, id_discipline = $3, id_group = $4, id_teacher=$5 where id_load = $1", l.id_load, l.id_discipline, l.id_group, l.id_teacher, l.year)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 0 {
					strN = false
				}
			}
		}
		if tableN == 3 {
			strN := true
			for strN != false { //вывод таблицы групп
				rows, err := db.Query("select * from grupi")
				if err != nil {
					log.Fatal(err.Error())
					panic(err)
				}
				defer rows.Close()
				groups := []group{}

				for rows.Next() {
					g := group{}
					err := rows.Scan(&g.id_group, &g.group)
					if err != nil {
						fmt.Println(err)
						continue
					}
					groups = append(groups, g)
				}
				fmt.Println("| Номер группы | группа |")
				for _, g := range groups {
					fmt.Println("|", g.id_group, "|", g.group, "|")
				}

				fmt.Println("| 1 - добавить | 2 - удалить | 3 - изменить | 0 - выйти |")
				fmt.Print("-> ")
				fmt.Scanf("%s\n", &ingN3)
				funcN, err := strconv.Atoi(ingN3)
				if err != nil {
					log.Fatal(err)
				}

				if funcN == 1 { // добавление
					g := group{}
					fmt.Print("Номер группы: ")
					fmt.Scan(&g.id_group)
					fmt.Print("Группа: ")
					fmt.Scan(&g.group)
					result, err := db.Exec("insert into grupi (id_group, group) values ($1, $2) ", g.id_group, g.group)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 2 { // удаление
					g := group{}
					fmt.Print("Номер группы: ")
					fmt.Scan(&g.id_group)
					result, err := db.Exec("delete from grupi where id_group = $1", g.id_group)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 3 { // обновление
					g := group{}
					fmt.Print("Номер группы: ")
					fmt.Scan(&g.id_group)
					fmt.Print("Группа: ")
					fmt.Scan(&g.group)
					result, err := db.Exec("update grupi set id_group = $2 where id_group = $1", g.id_group, g.group)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 0 {
					strN = false
				}
			}
		}
		if tableN == 4 {
			strN := true
			for strN != false { // вывод таблицы дисциплин
				rows, err := db.Query("select * from disciplines")
				if err != nil {
					log.Fatal(err.Error())
					panic(err)
				}
				defer rows.Close()
				disciplines := []discipline{}

				for rows.Next() {
					d := discipline{}
					err := rows.Scan(&d.id_discipline)
					if err != nil {
						fmt.Println(err)
						continue
					}
					disciplines = append(disciplines, d)
				}
				fmt.Println("Номер дисциплины | часов практики | часов лекций | дисциплина |")
				for _, d := range disciplines {
					fmt.Println("|", d.id_discipline, "|", d.hours_of_practice, "|", d.hours_of_lecture, "|", d.disciplines, "|")
				}

				fmt.Println("| 1 - добавить | 2 - удалить | 3 - изменить | 0 - выйти |")

				fmt.Print("-> ")
				fmt.Scanf("%s\n", &ingN4)
				funcN, err := strconv.Atoi(ingN4)
				if err != nil {
					log.Fatal(err)
				}

				if funcN == 1 { // добавление
					d := discipline{}
					fmt.Print("Номер дисциплины: ")
					fmt.Scan(&d.id_discipline)
					fmt.Print("Часов практики: ")
					fmt.Scan(&d.hours_of_practice)
					fmt.Print("Часов лекций: ")
					fmt.Scan(&d.hours_of_lecture)
					fmt.Print("Дисциплина: ")
					fmt.Scan(&d.disciplines)
					result, err := db.Exec("insert into disciplines (id_discipline, hours_of_practice, hours_of_lecture, disciplines) values ($1, $2, $3, $4) ", d.id_discipline, d.hours_of_practice, d.hours_of_lecture, d.disciplines)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 2 { // удаление
					d := discipline{}
					fmt.Print("Номер дисциплины: ")
					fmt.Scan(&d.id_discipline)
					result, err := db.Exec("delete from disciplines where id_discipline = $1", d.id_discipline)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 3 { // обновление
					d := discipline{}
					fmt.Print("Номер дисциплины: ")
					fmt.Scan(&d.id_discipline)
					fmt.Print("Часов практики: ")
					fmt.Scan(&d.hours_of_practice)
					fmt.Print("Часов лекций: ")
					fmt.Scan(&d.hours_of_lecture)
					fmt.Print("Дисциплина: ")
					fmt.Scan(&d.disciplines)
					result, err := db.Exec("update teachers set hours_of_practice = $2, hours_of_lecture = $3, disciplines = $4 where id_discipline = $1", d.id_discipline, d.hours_of_practice, d.hours_of_lecture, d.disciplines)
					if err != nil {
						panic(err)
					}
					fmt.Println(result.RowsAffected())
					strN = false
				}
				if funcN == 0 {
					strN = false
				}
			}
		}
		if tableN == 0 {
			progN = false
		}
	}
}
