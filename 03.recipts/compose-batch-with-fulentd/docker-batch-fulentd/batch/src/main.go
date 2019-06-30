package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"io/ioutil"
	"time"
	
	_ "github.com/lib/pq"
	"github.com/urfave/cli"
)

func main() {
	run(os.Args)
}

func run(args []string) error {
	app := cli.NewApp()
	
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "SqlFile",
			Value: "",
		},
	}
	app.Action = func(c *cli.Context) error {
		InfoLog("Batch started")
		
		sqlfile := c.String("SqlFile")
		if sqlfile == "" {
			return errors.New("Parameter 'SqlFile' is required.")
		}
		
		db, err := sql.Open("postgres", GetConnectionString())
		if err != nil {
			return err
		}
		defer db.Close()
		
		file, err := os.Open(sqlfile)
		if err != nil {
			return err
		}
		defer file.Close()
		bytes, _ := ioutil.ReadAll(file)
		query := string(bytes)
		
		InfoLog("Read SQL from `%s` and run query...", sqlfile)
		InfoLog("Query: `%s`", query)
		
		result, err := db.Exec(query)
		if err != nil {
			fmt.Printf(err.Error())
			return nil
		}
		rowsAffected, _ := result.RowsAffected()
		InfoLog("%d rows affected", rowsAffected)
		InfoLog("Batch finished")
		
		return nil
	}

	return app.Run(args)
}

func InfoLog(message string, args ...interface{}) {
	nowTime := time.Now()
	const timefmt = "2006/01/02 15:04:05"
	
	msg := fmt.Sprintf(message, args...)
	
	fmt.Printf("[INFO] %s: %s\n", nowTime.Format(timefmt), msg)
}


func GetConnectionString() (string) {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
}
