package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	runScripts("echo 'Hello world!'")
	hello_world()
}

func runScripts(cmd string) {
	host := "103.200.21.199"
	port := "22"
	user := "root"

	//pass privatekey
	//keypass := "namluu"
	pass := "vietnix@2016"

	//path of privatekey
	//pathOfPriKey := "/Users/koongnam/.ssh/auth.vietnix.vn"

	// key, err := ioutil.ReadFile(pathOfPriKey)
	// if err != nil {
	// 	log.Fatalf("unable to read private key: %v", err)
	// }

	// Create the Signer for this private key.
	// priKey, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(keypass))
	// if err != nil {
	// 	log.Fatalf("unable to parse private key: %v", err)
	// }

	// ssh client config
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			//sử dụng private key = hàm publickey hoặc dùng pass root = hàm password, 1 trong 2
			ssh.Password(pass), //-> use password root
			//ssh.PublicKeys(priKey), //-> use privatekey
		},
		// allow any host key to be used (non-prod)
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// connect
	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// start session
	sess, err := client.NewSession()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// setup standard out and error
	// uses writer interface
	sess.Stdout = os.Stdout
	sess.Stderr = os.Stderr

	// run single command
	err = sess.Run(cmd)
	if err != nil {
		log.Fatal(err)
	}
}

func hello_world() {
	fmt.Print("Hello world!")
	fmt.Print("Hello world - feature!!!!")
	fmt.Print("Hello world - setup golang!!	!!!")
}
