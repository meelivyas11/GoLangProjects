package CRUD

import (
	"database/sql"
	"fmt"
	"log"

	ora "gopkg.in/rana/ora.v3"
)

// db.Query() actually prepares, executes, and closes a prepared statement.
//That’s three round-trips to the database.
func Select_using_DatabaseSQL_Query() {
	fmt.Println("Get First and Last name of all employees using database/sql Prepare Query method")
	dsn := "testDB/testDB@localhost:32769/xe"
	db, err := sql.Open("ora", dsn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT First, Last FROM EMP_TABLE")
	if err != nil {
		log.Fatal("Could not execute query due to " + err.Error())
	}
	// Explicitly calling in case the loop exist in-between
	// rows.Close() is a harmless no-op, so you can call it multiple times
	defer rows.Close()

	for rows.Next() {
		var First, Last string
		err = rows.Scan(&First, &Last)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("First: " + First + " Last: " + Last)
	}
	// When the last row is read, rows.Next() return an EOF error which calls rows.Close
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

// Use Prepare queries to be used multiple times.
// The result of preparing the query is a prepared statement, which can have placeholders
//for parameters that you’ll provide when you execute the statement.
func Select_using_DatabaseSQL_Prepare_Query() {
	fmt.Println("Get First and Last name of all employees with ID =100 (not unique) using database/sql Prepare Query method")
	dsn := "testDB/testDB@localhost:32769/xe"
	db, err := sql.Open("ora", dsn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	//In MySQL, the parameter placeholder is ?
	// In PostgreSQL it is $N, where N is a number.
	//SQLite accepts either of these.
	//In Oracle placeholders begin with a colon and are named, like :param1.
	stmt, err := db.Prepare("SELECT First,Last FROM EMP_TABLE where ID = (:1)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(100)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var First, Last string
		err = rows.Scan(&First, &Last)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("First: " + First + " Last: " + Last)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func Select_using_DatabaseSQL_QueryRow() {
	fmt.Println("Get First and Last name of all employees with ID =101 (unique) using database/sql QueryRow methods")
	dsn := "testDB/testDB@localhost:32769/xe"
	db, err := sql.Open("ora", dsn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	var first, last string
	err = db.QueryRow("SELECT First, Last FROM EMP_TABLE where ID = (:1)", 101).Scan(&first, &last)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Selected employee is: " + first + " " + last)
}

func Select_using_DatabaseSQL_Prepare_QueryRow() {
	fmt.Println("Get First and Last name of all employees with ID =101 (unique) using database/sql Prepare and QueryRow methods")
	dsn := "testDB/testDB@localhost:32769/xe"
	db, err := sql.Open("ora", dsn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("SELECT First, Last FROM EMP_TABLE where ID = (:1)")
	if err != nil {
		log.Fatal(err)
	}
	var first, last string
	err = stmt.QueryRow(101).Scan(&first, &last)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Selected employee is: " + first + " " + last)
}

func Selete_using_Ora_Sel() {
	fmt.Println("Select using oracle driver Ses.Sel method")
	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}

	result, err := ses.Sel("EMP_TABLE",
		"ID", ora.I32, // integer32
		"First", ora.S, // string
		"Last", ora.S,
		"PHONE", ora.I32,
		"EMAIL", ora.S)
	if err != nil {
		fmt.Printf("Unable to insert: %s", err)
		panic(err)
	}
	fmt.Println("Selected Data is:")

	for result.Next() {
		fmt.Println(result.Index, result.Row[0], result.Row[1], result.Row[2], result.Row[3], result.Row[4])
	}
}

func Select_using_Ora_PrepAndQry() {
	fmt.Println("Select using oracle driver PrepAndQry method")

	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}

	result, err := ses.PrepAndQry("SELECT FIRST FROM EMP_TABLE")

	if err != nil {
		fmt.Printf("Unable to select: %s", err)
		panic(err)
	}
	for result.Next() {
		fmt.Println(result.Row[0])
	}
	if result.Err != nil {
		panic(result.Err)
	}
}

func Select_using_Ora_Prep_Qry() {
	fmt.Println("Select using oracle driver Prep and Qry method")
	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}
	stmt, err := ses.Prep("SELECT First, Last, Phone FROM EMP_TABLE")
	result, err := stmt.Qry()

	if err != nil {
		fmt.Printf("Unable to select: %s", err)
		panic(err)
	}
	for result.Next() {
		fmt.Println(result.Index, result.Row[0], result.Row[1], result.Row[2])
	}
	if result.Err != nil {
		panic(result.Err)
	}
}
