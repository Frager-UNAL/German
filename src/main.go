package main

import (
	"os"

	"./dbconn"
	router "./router"

	"github.com/joho/godotenv"
)

type Env struct {
	DATABASE_URL, PORT string
}

func getEnv() Env {
	godotenv.Load(".env")
	myEnv := Env{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		PORT:         os.Getenv("PORT"),
	}

	return myEnv

}

func main() {

	println("Corriendo servicio FRAGER-GO")

	myEnv := getEnv()

	err := dbconn.Connect(myEnv.DATABASE_URL)

	if err != nil {
		println("Error on DB Connection")
		println(err.Error())
		print("db url: ")
		println(myEnv.DATABASE_URL)
		return
	}

	err = router.StartServe(":" + myEnv.PORT)
	if err != nil {
		println("Error on start server")
		println(err)
		return
	}

}
