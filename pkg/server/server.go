package server

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/notEpsilon/lucy/pkg/constants"
)

func Start(outputFilePath string, bytesPerIteration int) error {
	file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	ln, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", constants.DefaultPort))
	if err != nil {
		return err
	}
	defer ln.Close()
	log.Printf("[lucy]: Listening on %s for incoming files...\n", ln.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("[lucy]: a device attempted to connect but failed: %s\n", err.Error())
			continue
		}
		defer conn.Close()

		buf := make([]byte, bytesPerIteration)
		done := false

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				log.Println("[lucy]: successfully received the file")
				done = true
				break
			}
			if err != nil {
				log.Printf("[lucy]: failed to read from the sender, transefered file might get corrupted: %s\n", err.Error())
				continue
			}
			data := buf[:n]

			_, err = file.Write(data)
			if err != nil {
				log.Printf("[lucy]: failed to write to the output file, transefered file might get corrupted: %s\n", err.Error())
				continue
			}
		}

		if done {
			break
		}
	}

	return nil
}
