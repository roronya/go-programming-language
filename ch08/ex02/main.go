package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func port(data string) (*net.TCPAddr, error) {
	// ex. 127,0,0,1,250,113 => 127.0.0.1:64113
	s := strings.Split(data, ",")
	p1, _ := strconv.Atoi(s[4])
	p2, _ := strconv.Atoi(s[5])
	p := p1*256 + p2 // ポート番号は8bitずつ分けられて送られてくるので10進数に変換する
	return net.ResolveTCPAddr(
		"tcp",
		fmt.Sprintf("%s.%s.%s.%s:%d", s[0], s[1], s[2], s[3], p),
	)
}

func handleConn(conn net.Conn) {
	fmt.Fprintln(conn, "220 Welcome to gopl Ex8.2 FTP Server")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		var dtpUserAddr *net.TCPAddr
		dtpServerAddr, _ := net.ResolveTCPAddr("tcp", ":20") // ignore error

		data := scanner.Text()
		log.Println(data)

		split := strings.Split(data, " ")
		cmd, arg := strings.ToUpper(split[0]), split[1:]

		switch cmd {
		case "USER":
			fmt.Fprintf(conn, "331 User %s OK. Password required\n", arg)
		case "PASS":
			fmt.Fprintln(conn, "230 OK. Current directory is /")
		case "QUIT":
			fmt.Fprintln(conn, "221 Goodbye")
		case "PORT":
			if len(arg) != 1 {
				fmt.Fprintln(conn, "argument error")
			}
			dtpUserAddr, _ = port(arg[0])
			fmt.Fprintf(conn, "200 Using %s to transfer files\n", dtpUserAddr)
		case "TYPE":
			fmt.Fprintln(conn, "Using ascii mode to transfer files")
		case "MODE":
			fmt.Fprintln(conn, "We only support stream mode, sorry")
		case "STRU":
			fmt.Fprintln(conn, "We only support file structure, sorry.")
		case "RETR":
			if len(arg) != 1 {
				fmt.Fprintln(conn, "argument error")
			}
			dtpConn, err := net.DialTCP("tcp", dtpServerAddr, dtpUserAddr)
			if err != nil {
				fmt.Fprintln(conn, err) // TODO
				break
			}
			defer dtpConn.Close()
			f, err := os.Open(arg[0])
			defer f.Close()
			if err != nil {
				fmt.Fprintln(conn, err)
			}
			for {
				var b []byte
				read, err := f.Read(b)
				if err != nil {
					fmt.Fprintln(conn, err)
					break
				}
				if read == 0 {
					break
				}
				_, err = dtpConn.Write(b)
				if err != nil {
					fmt.Fprintln(conn, err)
					break
				}
			}
			dtpConn.Close()
			f.Close()
		case "STOR":
			// TODO
		case "NOOP":
			fmt.Fprintln(conn, "200")
		default:
			fmt.Fprintf(conn, "502 %s is not supported.\n", cmd)
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":21")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
