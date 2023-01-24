package main

import (
	"strings"
	"testing"
)

func TestGetContents(t *testing.T) {
	resp := GetContents()

	if !strings.HasPrefix(resp, "{") {
		t.Fatal("Prefix is not valid json:", resp)
	}
	print(resp)
}
