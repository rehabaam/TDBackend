package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_startServer(t *testing.T) {
	tests := []struct {
		name string
		want error
	}{
		// {
		// 	name: "startServer_success",
		// 	want: nil,
		// },
		// {
		// 	name: "startServer_fail",
		// 	want: nil,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := startServer()
			assert.Equal(t, tt.want, got, "GetReqID returned value as expected")

		})
	}
}
