package test

import (
	"iot-platform/api"
	"testing"
)

func TestCreateAuthUser(t *testing.T) {
	user := api.CreateAuthUserRequest{
		UserId:   "testclientid",
		Password: "123456789",
	}
	err := api.CreateAuthUser(&user)
	if err != nil {
		t.Fatal(err)
	}

}
