package datafactory

import "github.com/nonemax/porto/integtests/structs"

// TestData is struct for testing
type TestData struct {
	Request  Request
	Response Response
}

// Request represents data for test request
type Request struct {
	Enpoint   string
	PortUnloc string
}

// Response represents data of test request
type Response struct {
	Status string
	Port   structs.Port
}

// GetPositiveDubai return data for positive test
func GetPositiveDubai() TestData {
	return TestData{
		Request: Request{
			Enpoint:   "/port",
			PortUnloc: "/AEDXB",
		},
		Response: Response{
			Status: "200 OK",
			Port: structs.Port{
				Name:     "Dubai",
				City:     "Dubai",
				Country:  "United Arab Emirates",
				TimeZone: "Asia/Dubai",
				Code:     "52005",
			},
		},
	}
}

// GetNegativeDubai return data for negative test
func GetNegativeDubai() TestData {
	return TestData{
		Request: Request{
			Enpoint:   "/port",
			PortUnloc: "/AEDXB1",
		},
		Response: Response{
			Status: "400 Bad Request",
			Port: structs.Port{
				Name:     "",
				City:     "",
				Country:  "",
				TimeZone: "",
				Code:     "",
			},
		},
	}
}
