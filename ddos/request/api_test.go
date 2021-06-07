package request

import (
	"fmt"
	"testing"
)

func Test_ApiLogin(t *testing.T) {
	par := APIKeys{"cdadmin", "A16pBIzVJOwHSC%23Q"}
	cook, err := Login(par)
	fmt.Println(cook, err)
}
