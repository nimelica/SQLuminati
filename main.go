package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/olekukonko/tablewriter"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./draw <database-file-path> <query>")
		os.Exit(1)
	}

	dbFilename := os.Args[1]
	sqlQuery := os.Args[2]

	dbPath, err := filepath.Abs(dbFilename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query(sqlQuery)
	if err != nil {
		fmt.Println("Error executing query:", err)
		os.Exit(1)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Println("Error getting column names:", err)
		os.Exit(1)
	}

	data := make([][]string, 0)
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			os.Exit(1)
		}

		rowData := make([]string, len(columns))
		for i, v := range values {
			switch v := v.(type) {
			case nil:
				rowData[i] = "NULL"
			case []byte:
				rowData[i] = string(v)
			default:
				rowData[i] = fmt.Sprintf("%v", v)
			}
		}

		data = append(data, rowData)
	}

	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(columns)
		for _, rowData := range data {
			table.Append(rowData)
		}
		table.Render()
	} else {
		fmt.Println("No rows found")
	}
}

