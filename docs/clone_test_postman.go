//--------------
//[1] Change pass private key					-> runScripts()
//[2] Change path private key					-> runScripts()
//[3] Change path of scripts clone template   	-> scripPath -> clone_template()
//[4] Change path of scripts rename template	-> scripPath -> rename_template()

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

// http://networkbit.ch/golang-ssh-client/

func input(s string) string {
	var input string
	fmt.Print("\n", s)
	fmt.Scan(&input)
	return input
}

func menu() int {
	fmt.Println("\n-------------Menu-------------")
	fmt.Println("\n[1] Clone template")
	fmt.Println("\n[2] Rename template")
	fmt.Println("\n[3] Clone all template")
	fmt.Println("\n[4] Exit program")
	var choice int
	fmt.Print("\n[*]Your choice: ")
	fmt.Scan(&choice)
	return choice
}

func runScripts(scriptPath, str1, str2, str3 string) {
	host := "ip_host"
	port := "22"
	user := "root"
	cmd := scriptPath + " " + str1 + " " + str2 + " " + str3

	//##pass privatekey
	keypass := "pass_auth"

	//pass root if not use privatekey
	//pass := "ohCwgZc4fWdAWbhgioWt"

	//##path of privatekey
	pathOfPriKey := "docs/id_rsa"

	key, err := ioutil.ReadFile(pathOfPriKey)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	priKey, err := ssh.ParsePrivateKeyWithPassphrase(key, []byte(keypass))
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	// ssh client config
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			//sử dụng private key = hàm publickey hoặc dùng pass root = hàm password, 1 trong 2
			//ssh.Password(pass),   //-> use password root
			ssh.PublicKeys(priKey), //-> use privatekey
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

func main() {
	src_id := "srcid_template"
	id := "id_template"
	template := "node_template"
	runScripts("bash /root/tan/script/go_clone_all_template.sh", src_id, id, template)
}
