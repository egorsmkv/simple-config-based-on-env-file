package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pleycpl/godotenv"
)

var env = envFileParams()

var (
	dbHost     = env["DB_HOST"]
	dbUser     = env["DB_USER"]
	dbPassword = env["DB_PASSWORD"]
	dbPort     = mustBeInt(env["DB_PORT"])
)

var (
	DbHost     = flag.String("db-host", dbHost, "Database hostname")
	DbUser     = flag.String("db-user", dbUser, "Database user")
	DbPassword = flag.String("db-password", dbPassword, "Database password")
	DbPort     = flag.Int("db-port", dbPort, "Database port")
)

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("Database settings")
	fmt.Println()
	fmt.Println("Host:", *DbHost)
	fmt.Println("User:", *DbUser)
	fmt.Println("Password:", *DbPassword)
	fmt.Println("Port:", *DbPort)
}

func envFileParams() map[string]string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return godotenv.Godotenv(dir + "/.env")
}

func mustBeInt(v string) int {
	result, err := strconv.Atoi(v)
	if err != nil {
		log.Fatal(err)
	}
	return result
}
