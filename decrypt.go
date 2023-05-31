package main

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("./decrypt.go file.pgp private.key passphrase")
	}
	fileArg := os.Args[1]
	privateKeyArg := os.Args[2]
	passphraseArg := os.Args[3]

	file, err := ioutil.ReadFile(fileArg)
	if err != nil {
		log.Fatal("error read file", err)
	}

	privateKey, err := ioutil.ReadFile(privateKeyArg)
	if err != nil {
		log.Fatal("error read file", err)
	}

	decrypted, err := helper.DecryptBinaryMessageArmored(string(privateKey), []byte(passphraseArg), string(file))
	if err != nil {
		log.Fatal("error decrypt", err)
	}

	fileWithoutPGP := strings.TrimSuffix(fileArg, ".pgp")
	err = ioutil.WriteFile(fileWithoutPGP, decrypted, os.ModePerm)
	if err != nil {
		log.Fatal("error write file", err)
	}
	log.Println("Done", fileWithoutPGP)
}
