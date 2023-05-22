package helpers

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Read User Id from Environment
func ReadUserId() (value string, ok bool){
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		log.Fatal(err.Error())
	} 
	
	value, ok = os.LookupEnv("USERID")

	if !ok {
		log.Fatal("USERID not set in environment")
	}
	return value, ok
}

// Read Join Code from Environment 
func ReadJoinCode() (value string, ok bool){

	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		log.Fatal(err.Error())
	} 


	value, ok = os.LookupEnv("JOINCODE")

	if !ok {
		log.Fatal("JOINCODE not set in environment")
	}

	return value, ok
}

// Read API End Point from Environment 
func ReadEndPoint() (value string, ok bool){

	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		log.Fatal(err.Error())
	} 


	value, ok = os.LookupEnv("BASEURL")

	if !ok {
		log.Fatal("BASEURL not set in environment")
	}
	return value, ok
}

// Read Scheme from Environment 
func ReadScheme() (value string, ok bool){

	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		log.Fatal(err.Error())
	} 


	value, ok = os.LookupEnv("SCHEME")

	if !ok {
		log.Fatal("SCHEME not set in environment")
	}
	return value, ok
}