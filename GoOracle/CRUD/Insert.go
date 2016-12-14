package CRUD

import (
	"database/sql"
	"fmt"
	"log"

	ora "gopkg.in/rana/ora.v3"
)

func Insert_using_DatabaseSQL_Prepare_Exec() {
	fmt.Println("Insert using database/sql using Prepare and Exec")
	dsn := "testDB/testDB@localhost:32769/xe"
	db, err := sql.Open("ora", dsn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO EMP_TABLE (ID, First, Last) VALUES (:1, :2, :3)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("101", "Meeli", "Vyas")
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of rows affected = %d\n", rowCnt)
}

func Insert_using_Ora_Ins() {
	fmt.Println("Insert using oracle driver Ses.Ins method")

	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}

	err = ses.Ins("EMP_TABLE",
		"First", "Meeli",
		"Last", "Vyas",
		"PHONE", "1111111111",
		"EMAIL", "m@test.com",
		"ID", "102",
		"ID", "102")
	if err != nil {
		fmt.Printf("Unable to insert: %s", err)
		panic(err)
	}

	fmt.Println("Data inserted successfully")
}

func Insert_using_Ora_Prep_Exe() {
	fmt.Println("Insert using oracle driver Prep and Exe method")
	tableName := "EMP_TABLE"

	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}
	stmtIns, err := ses.Prep(fmt.Sprintf("INSERT INTO %v (ID, First, Last) VALUES (:1, :2, :3) RETURNING First INTO (:2)", tableName))
	defer stmtIns.Close()
	rowsAffected, err := stmtIns.Exe("103", "Meeli", "Vyas")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of rows inserted: %v\n", rowsAffected)
}

func Insert_using_Ora_PrepAndExe() {
	fmt.Println("Insert using oracle driver PrepAndExe method")
	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := ses.PrepAndExe("INSERT INTO EMP_TABLE (ID, First, Last) VALUES (:1, :2, :3) RETURNING First INTO (:2)",
		"104", "MY", "Vyas")
	fmt.Printf("Number of rows inserted: %v\n", rowsAffected)
}
