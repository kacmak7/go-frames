package main

import (
	"encoding/hex"
	"github.com/akamensky/argparse"
	"github.com/kacmak7/go-p2p-packets/sender"
	"github.com/kacmak7/go-p2p-packets/receiver"
	"github.com/mdlayher/raw"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Setup commands and arguments parser
	parser := argparse.NewParser("commands", "Frames commands")
	interfaceFlag := parser.String("i", "interface", &argparse.Options{Required: false, Help: "Interface to be used. If not specified \"FRAMES_INTERFACE\" variable will be used"})

	sendCmd := parser.NewCommand("send", "Sender mode - sends a message")
	msgFlag := sendCmd.String("m", "msg", &argparse.Options{Required: true, Help: "Message to be send"})
	destFlag := sendCmd.String("d", "destination", &argparse.Options{Required: true, Help: "MAC Address of the destination of the message (e.g. 09:d8:61:bd:52:1a)"})

	receiveCmd := parser.NewCommand("receive", "Receiver mode - receive messages")

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	// Check if interface is specified
	inter := *interfaceFlag
	if inter == "" {
		inter = os.Getenv("FRAMES_INTERFACE")
		if inter == "" {
			log.Fatal("Interface is not specified.")
		}
	}

	// Setup interface
	ifi, err := net.InterfaceByName(inter)
	if err != nil {
		log.Fatal("Interface setup error")
	}
	log.Printf("Using %v", inter)

	// Setup connection
	conn, err := raw.ListenPacket(ifi, 0x1234, nil) // TODO protocol.EtherType type
	if err != nil {
		log.Fatal("Connection setup error")
	}

	if sendCmd.Happened() {
		dest, err := hex.DecodeString(strings.Replace(*destFlag, ":", "", -1))
		if err != nil {
			log.Fatal("Destination MAC address formatting error")
		}
		sender.SendMsg(conn, dest, ifi.HardwareAddr, *msgFlag)
	} else if receiveCmd.Happened() {
		receiver.ReceiveAndProcess(conn)
	}
}
