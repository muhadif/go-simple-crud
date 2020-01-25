package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func getStudents(w http.ResponseWriter, r *http.Request) {
	var student Student
	var arr_student []Student
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select id,first_name,last_name from student")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&student.Id, &student.FirstName, &student.LastName); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_student = append(arr_student, student)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_student

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func addStudent(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
		return
	}

	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	_, err = db.Exec("INSERT INTO student (first_name, last_name) values (?,?)",
		first_name,
		last_name,
	)

	if err != nil {
		log.Print(err)
		return
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var response Response
	var result sql.Result

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("user_id")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	result, err = db.Exec("UPDATE student set first_name = ?, last_name = ? where id = ?",
		first_name,
		last_name,
		id,
	)

	if err != nil {
		log.Print(err)
		return
	}

	row, _ := result.RowsAffected()

	if row > 0 {
		response.Status = 1
		response.Message = "Success Update Data"
		log.Print("Update data to database")
	} else {
		response.Status = 404
		response.Message = "Data Not Found"
		log.Print("Data not found")
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var response Response


	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("user_id")

	result , err := db.Exec("DELETE from student where id = ?",
		id,
	)


	if err != nil {
		log.Print(err)
		return
	}

	row, _ := result.RowsAffected()

	if row == 1 {
		response.Status = 1
		response.Message = "Success Delete Data"
		log.Print("Delete data to database")
	} else {
		response.Status = 404
		response.Message = "Data Not Found"
		log.Print("Data not found")
	}



	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
