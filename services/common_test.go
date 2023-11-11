package commands

import (
	"TDBackend/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func Test_readFile(t *testing.T) {
	logger.Init("debug", time.RFC3339, "TDBackend")

	os.Chdir("../")

	type args struct {
		endPoint string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "getKit_Success",
			args: args{
				endPoint: "Kit",
			},
			want: `[
				{
					"id": "1",
					"url": "https://images.squarespace-cdn.com/content/v1/5679680d2399a3acd195c64e/1623501020384-GJYROM88JEGSSNYTU6DS/ke17ZwdGBToddI8pDm48kIY9XMijbNp5oh3Amrd02-EUqsxRUqqbr1mOJYKfIPR7LoDQ9mXPOjoJoqy81S2I8PaoYXhp6HxIwZIk7-Mi3Tsic-L2IOPH3Dwrhl-Ne3Z2hmbgTVhMEoAcv2ODUGwXknK0vl1tv8aASOUMAN9pVXEKMshLAGzx4R3EDFOm1kBS/image001.png?format=1000w"
				}
			]`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.endPoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("readFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.EqualFold(jsonMarshaller(got), jsonMarshaller(tt.want)) {
				t.Errorf("readFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getImage(t *testing.T) {
	logger.Init("debug", time.RFC3339, "TDBackend")

	r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/v1/img/FFMC.jpg", nil)
	r = mux.SetURLVars(r, map[string]string{"name": "FFMC.jpg"})

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "getFFMC_Success",
			args: args{
				w: httptest.NewRecorder(),
				r: r,
			},
			want:    381170,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getImage(tt.args.w, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("getImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_loadFileToMemory(t *testing.T) {
	logger.Init("debug", time.RFC3339, "TDBackend")

	tests := []struct {
		name string
	}{
		{
			name: "getKit_Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loadFileToMemory()
		})
	}
}

func jsonMarshaller(str string) string {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, []byte(str)); err != nil {
		fmt.Println(err)
	}

	return buffer.String()
}
