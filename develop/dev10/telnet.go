/*
Реализовать простейший telnet-клиент.

Примеры вызовов:

go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Требования:

1. Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP. После подключения
	STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT

2. Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)

3. При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера,
	программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться
	через timeout
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Args ...
type Args struct {
	timeout time.Duration
	host    string
	port    string
}

func main() {
	args, err := getArgs()
	if err != nil {
		log.Fatalln(err)
	}

	runTelnetClient(args)
}

func getArgs() (*Args, error) {
	timeout := flag.Duration("timeout", time.Duration(10)*time.Second, "timeout")
	flag.Parse()
	hp := flag.Args()[0:]

	if len(hp) != 2 {
		return nil, fmt.Errorf("host or port is empty")
	}

	host := hp[0]
	port := hp[1]

	return &Args{
		timeout: *timeout,
		host:    host,
		port:    port,
	}, nil
}

func runTelnetClient(args *Args) {
	conn, err := connectToHost(args)
	if err != nil {
		handleError(err)
		return
	}
	defer conn.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go readFromSocket(conn)
	go writeToSocket(conn)

	select {
	case <-quit:
		fmt.Println("closing connection")
	}
}

func getHost(args *Args) string {
	return args.host + ":" + args.port
}

func connectToHost(args *Args) (net.Conn, error) {
	host := getHost(args)
	conn, err := net.DialTimeout("tcp", host, args.timeout)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func readFromSocket(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		handleError(err)
	}
}

func writeToSocket(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Fprintf(conn, "%s\n", input)
	}

	if err := scanner.Err(); err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
