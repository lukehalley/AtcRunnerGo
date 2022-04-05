package main

import (
	"atc-runner/internal/env"
	"fmt"
)

func main() {

	DBName := env.GetEnvFromFile("DB_NAME")

	fmt.Println(DBName)

}