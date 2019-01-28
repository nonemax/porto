package tester

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nonemax/porto/integtests/datafactory"
	"github.com/nonemax/porto/integtests/structs"
)

// Config describes tester configuration
type Config struct {
	Address string
}

// New return new tester
func New(addres string) Config {
	return Config{
		Address: addres,
	}
}

// Test start testing
func (t *Config) Test() error {
	positiveData := datafactory.GetPositiveDubai()
	err := t.check(positiveData)
	if err != nil {
		return err
	}
	negativeData := datafactory.GetNegativeDubai()
	err = t.check(negativeData)
	if err != nil {
		return err
	}
	return nil
}

func (t *Config) check(testData datafactory.TestData) error {
	respBytes, respStatus, err := t.Request(testData.Request.Enpoint + testData.Request.PortUnloc)
	if err != nil {
		return err
	}
	if respStatus != testData.Response.Status {
		return fmt.Errorf("Wrong status")
	}
	var testPort structs.Port
	err = json.Unmarshal(respBytes, &testPort)
	if respStatus == "400 Bad Request" && err.Error() == "unexpected end of JSON input" {
		return nil
	}
	if err != nil {
		return err
	}
	if testPort.Code != testData.Response.Port.Code {
		return fmt.Errorf("Wrong Code field. Want %s, but have %s", testData.Response.Port.Code, testPort.Code)
	}
	if testPort.Name != testData.Response.Port.Name {
		return fmt.Errorf("Wrong Name field. Want %s, but have %s", testData.Response.Port.Name, testPort.Name)
	}
	if testPort.City != testData.Response.Port.City {
		return fmt.Errorf("Wrong City field. Want %s, but have %s", testData.Response.Port.City, testPort.City)
	}
	if testPort.Country != testData.Response.Port.Country {
		return fmt.Errorf("Wrong Country field. Want %s, but have %s", testData.Response.Port.Country, testPort.Country)
	}
	if testPort.TimeZone != testData.Response.Port.TimeZone {
		return fmt.Errorf("Wrong TimeZone field. Want %s, but have %s", testData.Response.Port.TimeZone, testPort.TimeZone)
	}
	return nil
}

// Request send test request
func (t *Config) Request(endpoint string) ([]byte, string, error) {
	resp, err := http.Get(t.Address + endpoint)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	return body, resp.Status, nil
}
