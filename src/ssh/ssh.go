package ssh

import (
	"fmt"
	"io"
	"os"
	"sync"

	"gihub.com/elmahdiben119/copyfiles/src/fs"
	"golang.org/x/crypto/ssh"
)

var config = &ssh.ClientConfig{
	User: fs.SSH_USER,

	Auth: []ssh.AuthMethod{
		ssh.Password(fs.SSH_PASS),
	},
	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
}

func CopyRemoteServer(originalFile string, remoteFile string) error {
	// connect to ssh
	config := &ssh.ClientConfig{
		User: fs.SSH_USER,

		Auth: []ssh.AuthMethod{
			ssh.Password(fs.SSH_PASS),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Create a session
	client, err := ssh.Dial("tcp", fs.SSH_HOST+":"+fs.SSH_PORT, config)
	if err != nil {
		fmt.Println("Failed to dial: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()

	if err != nil {
		fmt.Println("Failed to create session: ", err)
	}
	defer session.Close()

	// open file
	file, err := os.Open(originalFile)
	if err != nil {
		return err
	}
	defer file.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	// copy file to remote server
	go func() {
		hostIn, err := session.StdinPipe()
		if err != nil {
			fmt.Println("Failed to create stdin pipe: ", err)
		}
		defer hostIn.Close()

		io.Copy(hostIn, file)
		wg.Done()
	}()

	// run command on remote server to copy file to remote server and close stdin pipe when done
	session.Run("/usr/bin/scp -t " + remoteFile)
	wg.Wait()

	return nil
}
