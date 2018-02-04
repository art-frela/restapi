package main

import (
	"database/sql"
	"flag"
	"fmt"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"golang.org/x/text/encoding/charmap"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", DNS["password"], "the database password")
	port     *int = flag.Int("port", Atoi1(DNS["port"]), "the database port")
	server        = flag.String("server", DNS["server"], "the database server")
	user          = flag.String("user", DNS["userid"], "the database user")
	database      = flag.String("database", DNS["dbname"], "the database name")
)

//конвертация строки utf-8 в cp1251
func EncodeWindows1251(ba string) string {
	enc := charmap.Windows1251.NewEncoder()
	out, _ := enc.String(ba)
	return out
}

//Atoi1 is function converting string to integer and return singl value
func Atoi1(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}

//конвертация строки cp1251 -> utf-8
func DecodeWindows1251(ba string) string {
	dec := charmap.Windows1251.NewDecoder()
	out, _ := dec.String(ba)
	return out
}

//Execution sql query "CMD"
func exec(db *sql.DB, cmd string) error {
	rows, err := db.Query(cmd)
	if err != nil {
		return err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return err
	}
	if cols == nil {
		return nil
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
		}
		fmt.Print(cols[i])
	}
	fmt.Println()
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println(err)
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
			}
			printValue(vals[i].(*interface{}))
		}
		fmt.Println()

	}
	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
}

//Print values
func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}

//slaGet - get one or list SLAs
func slaGet(cmd string) (Slas, error) {
	if *debug {
		fmt.Printf("DNS is %s\n", DNS)
		fmt.Printf("Query is %s\n", cmd)
	}

	dsn := "server=" + DNS["server"] + ";user id=" + DNS["userid"] + ";password=" + DNS["password"] + ";database=" + DNS["dbname"]
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(cmd)
	var result Slas

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	if cols == nil {
		return nil, nil
	}
	var vals SLA

	for rows.Next() {
		err = rows.Scan(&vals.ID, &vals.Name, &vals.Description, &vals.Value, &vals.Icon)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			if *debug {
				fmt.Printf("Cols: %v\t%v\n", cols, err)
				fmt.Printf("row(vals): %v\n", vals)
				fmt.Printf("result: %v\n", result)
			}
			result = append(result, vals)
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return result, nil
}

//
func ObjectGet(cmd string) (Objects, error) {
	if *debug {
		fmt.Printf("DNS is %s\n", DNS)
		fmt.Printf("Query is %s\n", cmd)
	}

	dsn := "server=" + DNS["server"] + ";user id=" + DNS["userid"] + ";password=" + DNS["password"] + ";database=" + DNS["dbname"]
	db, err := sql.Open("mssql", dsn)
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect: ", err.Error())
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(cmd)
	var result Objects

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	if cols == nil {
		return nil, nil
	}
	var vals OBJECT

	for rows.Next() {
		err = rows.Scan(
			&vals.ID,
			&vals.Name,
			&vals.ID_City,
			&vals.Address,
			&vals.ID_ObjectType,
			&vals.DT_Open,
			&vals.DT_Close,
			&vals.ID_Focus,
			&vals.ID_Datamanager,
			&vals.IsEnabled,
		)
		if err != nil {
			fmt.Println(err)
			continue
		} else {
			if *debug {
				fmt.Printf("Cols: %v\t%v\n", cols, err)
				fmt.Printf("row(vals): %v\n", vals)
				fmt.Printf("result: %v\n", result)
			}
			result = append(result, vals)
		}
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return result, nil
}
