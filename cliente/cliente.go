package main

import(
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
	"bytes"
)

const dir = "127.0.0.1:8000"
const tamBufer = 256
const finLinea = 10

var apodo string
var in *bufio.Reader

func main() {
	in = bufio.NewReader(os.Stdin)
	for apodo == ""{
		fmt.Printf("Escribe tu apodo: ")
		buf, _, _ := in.ReadLine()
		apodo = string(buf)
	}
	var conn net.Conn
	var err error
	for{
		fmt.Printf("Conectando a %s...", dir)
		conn, err = net.Dial("tcp", dir)
		if err == nil{
			break
		}
	}
	defer conn.Close()

	go reciveMsj(conn)
	handleConexion(conn)
}

func handleConexion(conn net.Conn){
	for{
		buf, _, _ := in.ReadLine()
		if len(buf) > 0 {
			conn.Write( append([]byte(apodo+" dice: "), append(buf, finLinea)...))
		}
	}
}

func reciveMsj(conn net.Conn){
	var msj []byte
	bufer := make([]byte, tamBufer)
	for{
		for{
			n, err := conn.Read(bufer)
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			bufer = bytes.Trim(bufer[:n], "\x00")
			msj = append(msj, bufer...)
			if msj[len(msj)-1] == finLinea{
				break
			}
		}
		fmt.Printf("%s\n", msj[:len(msj)-1])
		msj = make([]byte, 0)
	}
}

