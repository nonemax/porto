package parser

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/nonemax/porto-entity"
	t "github.com/nonemax/porto-transport"
)

// Config is for parser configuration
type Config struct {
	filename string
	buffer   int
	c        t.TransportClient
	Sender   Sender
}

// New creates new parser
func New(c t.TransportClient, buffer int, filename string, s Sender) Config {
	return Config{
		c:        c,
		buffer:   buffer,
		filename: filename,
		Sender:   s,
	}
}

// Start begining parse file
func (c *Config) Start() error {
	file, err := os.Open(c.filename)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	lines := []byte{}
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, c.buffer*1024*1024)
	for scanner.Scan() {
		if scanner.Text() == "{" || scanner.Text() == "}" {
			continue
		}
		if strings.Contains(scanner.Text(), "},") {
			lines = append(lines, []byte("  }")...)
			lines = []byte("{" + string(lines) + "}")
			err = c.ScanPort(lines)
			if err != nil {
				return err
			}
			lines = []byte{}
			continue
		}
		lines = append(lines, scanner.Bytes()...)
	}
	return nil
}

// ScanPort marshal and send port to the server
func (c *Config) ScanPort(lines []byte) error {
	port := map[string]entity.Port{}
	err := json.Unmarshal(lines, &port)
	if err != nil {
		return err
	}
	for _, v := range port {
		bytePort, err := json.Marshal(v)
		err = c.Sender.SendPort(bytePort)
		if err != nil {
			return err
		}
		break
	}
	return nil
}
