package CRUD

import (
	"database/sql"
	"fmt"
	"log"

	ora "gopkg.in/rana/ora.v3"
)

func Update_using_DatabaseSQL_Prepare_Exec() {
	fmt.Println("Update using database/sql using Prepare and Exec")
	dsn := "testDB/testDB@localhost:32769/xe"
	db, err := sql.Open("ora", dsn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("UPDATE EMP_TABLE SET First = (:1), Last =(:2) WHERE Id = (:3)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Meeli_upd", "Vyas_upd", "101")
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of rows updated = %d\n", rowCnt)
}

func Update_using_Ora_Upd() {
	fmt.Println("Update using oracle driver Ses.Upd method")

	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}
	// Update First and Last name of the record where ID = 102
	err = ses.Upd("EMP_TABLE",
		"First", "Hello",
		"Last", "World",
		"ID", "102")
	if err != nil {
		fmt.Printf("Unable to Update: %s", err)
		panic(err)
	}

	fmt.Println("Successfully updated the Record")
}

func Update_using_Ora_Prep_Exe() {
	fmt.Println("Update using oracle driver Prep and Exe method")
	tableName := "EMP_TABLE"

	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}
	stmtIns, err := ses.Prep(fmt.Sprintf("UPDATE %v SET First = (:1), Last =(:2)  WHERE ID = (:3)", tableName))
	defer stmtIns.Close()
	rowsAffected, err := stmtIns.Exe("Meeli_upd", "Vyas_upd", "103")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of rows Updated: %v\n", rowsAffected)
}

func Update_using_Ora_PrepAndExe() {
	fmt.Println("Update using oracle driver PrepAndExe method")
	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := ses.PrepAndExe("UPDATE EMP_TABLE SET First = (:1), Last =(:2) WHERE ID = (:3) RETURNING First INTO (:2)",
		"MY_upd", "Vyas_upd", "104")
	fmt.Printf("Number of rows updated: %v\n", rowsAffected)
}
