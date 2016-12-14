package CRUD

import (
	"database/sql"
	"fmt"
	"log"

	ora "gopkg.in/rana/ora.v3"
)

func Delete_using_DatabaseSQL_Prepare_Exec() {
	fmt.Println("Delete using database/sql using Prepare and Exec")
	dsn := "testDB/testDB@localhost:32769/xe"
	db, err := sql.Open("ora", dsn)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("DELETE FROM EMP_TABLE WHERE ID = (:1)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("101")
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of rows Deleted = %d\n", rowCnt)
}

func Delete_using_Ora_Prep_Exe() {
	fmt.Println("Delete using oracle driver Prep and Exe method WHERE FirstName is like Meeli")
	tableName := "EMP_TABLE"

	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}
	stmtIns, err := ses.Prep(fmt.Sprintf("DELETE FROM %v WHERE ID = (:1) AND FIRST like (:2)", tableName))
	defer stmtIns.Close()
	rowsAffected, err := stmtIns.Exe("103", "%Meeli%")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number of rows Deleted: %v\n", rowsAffected)
}

func Delete_using_Ora_PrepAndExe() {
	fmt.Println("Delete using oracle driver PrepAndExe method")
	_, _, ses, err := ora.NewEnvSrvSes("testDB/testDB@localhost:32769/xe", nil)
	if err != nil {
		panic(err)
	}
	rowsAffected, err := ses.PrepAndExe("DELETE FROM EMP_TABLE WHERE ID =(:1)", "104")
	fmt.Printf("Number of rows Deleted: %v\n", rowsAffected)
}
