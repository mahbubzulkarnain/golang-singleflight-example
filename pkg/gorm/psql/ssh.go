package psql

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type SSH struct {
	Host string
	User string
	Pass string
	Port string
}

func NewSSHConfig() *SSH {
	return &SSH{
		Host: os.Getenv("SSH_HOST"),
		User: os.Getenv("SSH_USER"),
		Pass: os.Getenv("SSH_PASS"),
		Port: os.Getenv("SSH_PORT"),
	}
}

func (c SSH) NewConnection() (Net net.Conn, SSH *ssh.Client, err error) {
	var agentClient agent.Agent

	// Establish a connection to the local ssh-agent
	if Net, err = net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(Net)
	}

	// The client configuration with configuration option to use the ssh-agent
	sshConfig := &ssh.ClientConfig{
		User:            c.User,
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	// When there's a non empty password add the password AuthMethod
	sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
		return c.Pass, nil
	}))

	// Connect to the SSH Server
	if SSH, err = ssh.Dial("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port), sshConfig); err == nil {
		// Now we register the ViaSSHDialer with the ssh connection as a parameter
		sql.Register("postgres+ssh", &ViaSSHDialer{Client: SSH})
	}
	return
}
