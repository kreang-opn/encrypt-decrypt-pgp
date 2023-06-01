package main

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("./encrypt.go data.file public.asc")
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("error read file", err)
	}

	publicKey, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal("error read file", err)
	}

	armor, err := helper.EncryptBinaryMessageArmored(string(publicKey), data)
	if err != nil {
		log.Fatal("error encrypt", err)
	}

	err = ioutil.WriteFile(os.Args[1]+".pgp", []byte(armor), os.ModePerm)
	if err != nil {
		log.Fatal("error write file", err)
	}
	log.Println("Done", os.Args[1]+".pgp")
}
