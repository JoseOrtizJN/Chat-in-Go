package main

import(
	"log"
	"net"
	"os"
	"io"
	"bytes"
)

const dir = "127.0.0.1:8000"
const tamBufer = 256
const finLinea = 10

var clientes []net.Conn

func main(){
	clientes = make([]net.Conn, 0)
	oyente, err := net.Listen("tcp", dir)
	if err != nil{
		log.Fatal("No puedo entrar a " + dir)
		os.Exit(1)
	}
	for{
		conn, _ := oyente.Accept()
		clientes = append(clientes, conn)
		go handleConexion(conn)
	}
}

func handleConexion(conn net.Conn){
	defer conn.Close()
	var msj []byte
	bufer := make([]byte, tamBufer)
	for{
		for{
			n, err := conn.Read(bufer)
			if err != nil{
				if err == io.EOF{
					break
				}
			}
			bufer = bytes.Trim(bufer[:n], "\x00")
			msj = append(msj, bufer...)
			if msj[len(msj)-1] == finLinea{
				break
			}
		}
		enviarClientes(conn, msj)
		msj = make([]byte, 0)
	}
}

func enviarClientes(remitente net.Conn, msj []byte){
	for i := 0; i < len(clientes); i++ {
		if clientes[i] != remitente{
			clientes[i].Write(msj)
		}	
	}
}

