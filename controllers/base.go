package controllers

import (
	"bufio"
	"context"
	"database/sql"
	"net"
	"os"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rapando/tictactoe/utils"
)

type Base struct {
	Ctx   context.Context
	Cache *redis.Client
	DB    *sql.DB
}

func (base *Base) Init() error {
	var err error
	base.Ctx = context.Background()
	if base.Cache, err = utils.RedisConnect(base.Ctx); err != nil {
		return err
	}

	if base.DB, err = utils.DbConnect(); err != nil {
		return err
	}

	return nil
}

func (base *Base) RunServer() {
	port := ":" + os.Getenv("PORT")
	host := os.Getenv("HOST")
	utils.Log("INFO", "app", "starting socket server on port %v", port)
	listener, err := net.Listen("tcp", host+port)
	if err != nil {
		utils.Log("ERROR", "app", "unable to start listener because %v", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			utils.Log("ERROR", "app", "unable to connect because %v", err)
			return
		}
		utils.Log("INFO", "app", "client connected from : %v", conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	if err != nil {
		utils.Log("INFO", "app", "client left")
		conn.Close()
		return
	}

	utils.Log("INFO", "app", "client message : %v", string(buffer))

	response := "This is a response to : " + string(buffer)
	conn.Write([]byte(response))
	handleConnection(conn)
}
