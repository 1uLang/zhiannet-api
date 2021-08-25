package request

import (
	"fmt"
	"testing"
)

func TestRequest_GetToken(t *testing.T) {

	err := InitServerUrl("https://localhost:55000/")
	if err != nil {
		t.Fatal(err)
	}
	err = InitToken("wazuh","wazuh")
	if err != nil {
		t.Fatal(err)
	}
	req,err := NewRequest()
	if err != nil {
		t.Fatal(err)
	}
	token,err := req.GetToken()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}
