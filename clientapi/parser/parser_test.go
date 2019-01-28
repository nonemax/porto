package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type TestSender struct{}

func (t *TestSender) SendPort(port []byte) error {
	return nil
}

func TestStart(t *testing.T) {
	filename := "test.json"
	err := ioutil.WriteFile(filename, []byte(`{
		"AEAJM": {
			"name": "Ajman",
			"city": "Ajman",
			"country": "United Arab Emirates",
			"alias": [],
			"regions": [],
			"coordinates": [
			  55.5136433,
			  25.4052165
			],
			"province": "Ajman",
			"timezone": "Asia/Dubai",
			"unlocs": [
			  "AEAJM"
			],
			"code": "52000"
		},
	}`), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	sender := TestSender{}
	p := Config{
		buffer:   10,
		filename: filename,
		Sender:   &sender,
	}
	err = p.Start()
	if err != nil {
		t.Error(err)
	}

	err = os.Remove(filename)
}
