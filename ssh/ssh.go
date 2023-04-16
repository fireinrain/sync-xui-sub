package ssh

import (
	"fmt"
	"github.com/melbahja/goph"
	"log"
)

func UpdateSubconverter(vmessStrs string) {
	// Start new ssh connection with private key.
	auth, err := goph.Key("/home/mohamed/.ssh/id_rsa", "")
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root", "192.1.1.3", auth)
	if err != nil {
		log.Fatal(err)
	}

	// Defer closing the network connection.
	defer client.Close()

	// Execute your command.
	out, err := client.Run("ls /tmp/")

	if err != nil {
		log.Fatal(err)
	}

	// Get your output as []byte.
	fmt.Println(string(out))
}
