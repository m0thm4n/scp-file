package main

import (
	"flag"
	"fmt"
	"github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	host := flag.String("ip", "IP Address: 192.168.1.1", "Enter the IP of the device you want to copy to")
	username := flag.String("user", "Username: TimBob", "Enter the user for the remote device")
	password := flag.String("pass", "Password: P@$$W0rD!", "Enter the password of the user")
	localFilePath := flag.String("local", "Local File Path: /path/to/file", "Enter the path for the local file")
	remoteFilePath := flag.String("remote", "Remote File Path: /remote/path/to/file", "Enter the path for the remote"+
		" location")
	flag.Parse()

	if len(os.Args) < 5 {
		useage(*host, *username, *password, *localFilePath, *remoteFilePath)
	} else {
		scpFileHandler(*host, *username,*password, *localFilePath, *remoteFilePath)
	}
}

func scpFileHandler(host, username, password, localFilePath, remoteFilePath string) {
	// Use SSH key authentication from the auth package
	// we ignore the host key in this example, please change this if you use this library
	clientConfig, _ := auth.PrivateKey("username", "/path/to/rsa/key", ssh.InsecureIgnoreHostKey())

	// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod

	// Create a new SCP client
	client := scp.NewClient("example.com:22", &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}

	// Open a file
	f, _ := os.Open("/path/to/local/file")

	// Close client connection after the file has been copied
	defer client.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	err = client.CopyFile(f, "/home/server/test.txt", "0655")

	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
}

func useage(host, username, password, localFilePath, remoteFilePath string) {
	fmt.Printf("Usage of %s\n", host)
	fmt.Printf("Usage of %s\n", username)
	fmt.Printf("Usage of %s\n", password)
	fmt.Printf("Usage of %s\n", localFilePath)
	fmt.Printf("Usage of %s\n", remoteFilePath)
}