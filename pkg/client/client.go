package client

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/notEpsilon/lucy/pkg/constants"
)

func Send(filePath string, bytesPerIteration int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	conn, err := net.Dial("tcp", fmt.Sprintf("0.0.0.0:%d", constants.DefaultPort))
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("[lucy]: successfully connected to the receiver on %s\n", conn.RemoteAddr().String())

	buf := make([]byte, bytesPerIteration)

	for {
		n, err := file.Read(buf)
		if n == 0 {
			log.Println("[lucy]: successfully transefered the file")
			break
		}
		if err != nil {
			log.Printf("[lucy]: error reading from the file, sent file might get corrupted: %s\n", err.Error())
			continue
		}

		data := buf[:n]

		_, err = conn.Write(data)
		if err != nil {
			log.Printf("[lucy]: error writing to the server, sent file might get corrupted: %s\n", err.Error())
			continue
		}
	}

	return nil
}
