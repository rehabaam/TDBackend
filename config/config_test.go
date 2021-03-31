package config

import (
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name             string
		configPathFormat string
		cfg              interface{}
		wantPanic        bool
	}{
		// TODO: Add test cases.
		{
			name:             "givenExistPath_loadFile_success",
			configPathFormat: "prod.yml",
			cfg:              &Config,
			wantPanic:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("Load(), wantPanic = %v", tt.wantPanic)
				}
			}()
			configPath = tt.configPathFormat
			Load()
		})
	}
}

func Test_readFile(t *testing.T) {
	type args struct {
		cfg  interface{}
		path string
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "givenExistPath_success",
			args: args{
				cfg:  &Config,
				path: "prod.yml",
			},
			wantPanic: false,
		},
		{
			name: "givenNotExistPath_getPanic_failure",
			args: args{
				cfg:  nil,
				path: "dummy-file.yml",
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("readFile(), wantPanic = %v", tt.wantPanic)
				}
			}()
			readFile(tt.args.cfg, tt.args.path)
		})
	}
}

func Test_readEnv(t *testing.T) {
	type args struct {
		cfg interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		// TODO: Add test cases.
		{
			name: "givenExistPath_readEnv_success",
			args: args{
				cfg: &Config,
			},
			wantPanic: false,
		},
		{
			name: "givenNilAsStruct_readEnv_failure",
			args: args{
				cfg: nil,
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.wantPanic {
					t.Errorf("readEnv(), wantPanic = %v", tt.wantPanic)
				}
			}()
			readEnv(tt.args.cfg)
		})
	}
}
