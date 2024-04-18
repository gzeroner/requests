package requests

import (
	"fmt"
	"testing"
)

const url = "https://uapis.cn/api/mcinfo?username=T8k_"

func TestGet(t *testing.T) {
	res, err := Get(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.StatusCode)

	result := make(map[string]any)
	err = res.Unmarshal(&result)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(result)
}

func TestHead(t *testing.T) {
	res, err := Head(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.StatusCode)
}

func TestOptions(t *testing.T) {
	res, err := Options(url)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.StatusCode)
}
