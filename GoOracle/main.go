package main

import (
	crud "GoOracle/CRUD"
	db "GoOracle/DBConnection"
	"fmt"
)

func main() {
	fmt.Print("Test Database/sql: \n")
	db.DatabaseSQL_Select()

	fmt.Print("Insert Record using Oracle Driver\n")
	db.OraSes_Insert()

	fmt.Print("Get All Records  using Oracle Driver\n")
	db.OraSes_PrepAndQry()

	/*
			Notes:
			db.Query() and db.QueryRow() can only be used with SQL that returns data/rows
			Use Exec() with a prepared statement for statement that doesn’t return rows.

			INSERT, UPDATE, DELETE : Exec (SQL Command)
			SELECT : Query (SQL Query)

			Hence
			1) DatabaseSQL_Query(), DatabaseSQL_Prepare_Query(), DatabaseSQL_QueryRow() DatabaseSQL_Prepare_QueryRow()
			does not exist for INSERT and UPDATE
		 	2) DatabaseSQL_Prepare_Exec(), Ora_Prep_Exe, Ora_PrepAndExe() does not exist for SELECT

	*/

	// Call Different types of insert
	crud.Insert_using_DatabaseSQL_Prepare_Exec()
	crud.Insert_using_Ora_Ins()
	crud.Insert_using_Ora_Prep_Exe()
	crud.Insert_using_Ora_PrepAndExe()

	// Call Different types of Update
	crud.Update_using_DatabaseSQL_Prepare_Exec()
	crud.Update_using_Ora_Upd()
	crud.Update_using_Ora_Prep_Exe()
	crud.Update_using_Ora_PrepAndExe()

	// Call Different type of Select
	crud.Select_using_DatabaseSQL_Query()
	crud.Select_using_DatabaseSQL_Prepare_Query()
	crud.Select_using_DatabaseSQL_QueryRow()
	crud.Select_using_DatabaseSQL_Prepare_QueryRow()
	crud.Selete_using_Ora_Sel()
	crud.Select_using_Ora_Prep_Qry()
	crud.Select_using_Ora_PrepAndQry()

	// Call Different type of Delete
	crud.Delete_using_DatabaseSQL_Prepare_Exec()
	crud.Delete_using_Ora_Prep_Exe()
	crud.Delete_using_Ora_PrepAndExe()
}
