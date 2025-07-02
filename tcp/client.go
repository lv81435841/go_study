package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	client.conn = conn

	return client
}

var serverIp string
var serverPort int

func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "server ip")
	flag.IntVar(&serverPort, "port", 8888, "server port")
}

func (client *Client) menu() bool {
	var flag int
	fmt.Println("1,all")
	fmt.Println("2.private")
	fmt.Println("3.change name")
	fmt.Println("0.exit")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("err input")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println("update name")
	fmt.Scanln(&client.Name)
	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		switch client.flag {
		case 1:
			//fmt.Println("all")
			client.PublicChat()
			break
		case 2:
			//fmt.Println("private")
			client.PrivateChat()
			break
		case 3:
			//fmt.Println("change name")
			client.UpdateName()
			break
		}
	}
}

func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn)
}

func (client *Client) PublicChat() {
	var charMsg string
	fmt.Println("public chat pls input")
	fmt.Scanln(&charMsg)
	for charMsg != "exit" {
		if len(charMsg) != 0 {
			sengMsg := charMsg + "\n"
			_, err := client.conn.Write([]byte(sengMsg))
			if err != nil {
				fmt.Println(err)
				break
			}
		} else {
			fmt.Println("input blank")
		}
		charMsg = ""
		fmt.Println("public chat pls input")
		fmt.Scanln(&charMsg)
	}
}

func (client *Client) SelectUsers() {
	sengMsg := "who\n"
	_, err := client.conn.Write([]byte(sengMsg))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (client *Client) PrivateChat() {
	var chatMsg string
	var remoteName string
	client.SelectUsers()
	fmt.Println("private chat pls input name")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
		fmt.Println("pls input")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sengMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sengMsg))
				if err != nil {
					fmt.Println(err)
					break
				}
			}
			chatMsg = ""
			fmt.Println("pls input")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		fmt.Println("private chat pls input")
		fmt.Scanln(&remoteName)
	}
}

func main() {
	flag.Parse()
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("err")
		return
	}
	go client.DealResponse()
	fmt.Println("success conn")
	client.Run()
}
