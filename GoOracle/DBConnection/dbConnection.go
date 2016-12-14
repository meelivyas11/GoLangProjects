package DBConnection

import (
	"database/sql"
	"fmt"
	"log"

	ora "gopkg.in/rana/ora.v3"
)

// Can be used with different Drivers
func DatabaseSQL_Select() {
	fmt.Println("Database/Sql connection code with Select from SYSDATE Query")
	dsn := "testDB/testDB@localhost:32769/xe"
	db, err := sql.Open("ora", dsn)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT SYSDATE FROM DUAL")
	if err != nil {
		log.Fatal("Could not execute query due to " + err.Error())
	}
	for rows.Next() {
		var SYSDATE string
		rows.Scan(&SYSDATE)
		fmt.Println("Results: " + SYSDATE)
	}
}

// specific to Oracle driver
func OraSes_Insert() {
	fmt.Println("Oracle drive connection code with Insert Query")
	env, err := ora.OpenEnv(nil)
	defer env.Close()
	if err != nil {
		panic(err)
	}

	srvCfg := ora.SrvCfg{Dblink: "localhost:32769/xe"}
	srv, err := env.OpenSrv(&srvCfg)
	defer srv.Close()
	if err != nil {
		panic(err)
	}
	sesCfg := ora.SesCfg{
		Username: "testDB",
		Password: "testDB",
	}
	ses, err := srv.OpenSes(&sesCfg)
	err = ses.Ins("EMP_TABLE",
		"First", "Meeli",
		"Last", "Vyas",
		"PHONE", "1111111111",
		"EMAIL", "m@test.com",
		"ID", "100",
		"ID", "100")
	if err != nil {
		fmt.Printf("Unable to insert: %s", err)
		panic(err)
	}
	fmt.Println("Sample row inserted")
}

// specific to oracle driver
func OraSes_PrepAndQry() {
	fmt.Println("Oracle drive connection code with Select Query")
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
