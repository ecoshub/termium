package demo

import (
	"fmt"
	"net"

	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/style"
)

func StartListen(addr string, infoPanel *panel.Stack) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		infoPanel.Push(err.Error())
		return
	}
	infoPanel.Push(fmt.Sprintf("listen started at %s. write 'connect %s' to connect", addr, addr), style.DefaultStyleInfo)

	for {
		conn, err := listener.Accept()
		if err != nil {
			infoPanel.Push(err.Error())
			continue
		}
		infoPanel.Push(fmt.Sprintf("connection accepted at %s", conn.RemoteAddr()), style.DefaultStyleInfo)
		go reader(conn, infoPanel)
	}
}

func reader(conn net.Conn, infoPanel *panel.Stack) {
	buff := make([]byte, 128)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			infoPanel.Push(err.Error(), style.DefaultStyleError)
			return
		}
		s := string(buff[:n])
		commandSwitch(conn, s)
	}
}

func commandSwitch(conn net.Conn, command string) {
	switch command {
	case "hello":
		conn.Write([]byte("Hi!"))
		return
	case "what is your name":
		conn.Write([]byte("My name is eco"))
		return
	case "test":
		conn.Write([]byte("this is a test line"))
		return
	}
	conn.Write([]byte(command))
}

func ReadClient(conn net.Conn, commPanel *panel.Stack, infoPanel *panel.Stack) {
	buff := make([]byte, 128)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			infoPanel.Push(err.Error(), style.DefaultStyleError)
			return
		}
		s := string(buff[:n])
		commPanel.Push(fmt.Sprintf(">> %s", s), style.DefaultStyleSuccess)
	}
}
