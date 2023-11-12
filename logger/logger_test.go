package logger

import (
	"reflect"
	"testing"
	"time"

	"go.uber.org/zap/zapcore"
)

func Test_getZapLevel(t *testing.T) {
	type args struct {
		level string
	}
	tests := []struct {
		name string
		args args
		want zapcore.Level
	}{
		{
			name: "givenInfo_returnIntoLevel_success",
			args: args{
				level: "info",
			},
			want: zapcore.InfoLevel,
		},
		{
			name: "givenWarn_returnWarnLevel_success",
			args: args{
				level: "warn",
			},
			want: zapcore.WarnLevel,
		},
		{
			name: "givenDebug_returnDebugLevel_success",
			args: args{
				level: "debug",
			},
			want: zapcore.DebugLevel,
		},
		{
			name: "givenError_returnErrorLevel_success",
			args: args{
				level: "error",
			},
			want: zapcore.ErrorLevel,
		},
		{
			name: "givenFatal_returnFatalLevel_success",
			args: args{
				level: "fatal",
			},
			want: zapcore.FatalLevel,
		},
		{
			name: "givenPanic_returnPanicLevel_success",
			args: args{
				level: "panic",
			},
			want: zapcore.PanicLevel,
		},
		{
			name: "givenDPanic_returnDPanicLevel_success",
			args: args{
				level: "dpanic",
			},
			want: zapcore.DPanicLevel,
		},
		{
			name: "givenEmpty_returnInfoLevel_success",
			args: args{
				level: "",
			},
			want: zapcore.InfoLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getZapLevel(tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getZapLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customTimeEncoder(t *testing.T) {
	type args struct {
		t   time.Time
		enc zapcore.PrimitiveArrayEncoder
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customTimeEncoder(tt.args.t, tt.args.enc)
		})
	}
}

func TestInit(t *testing.T) {
	type args struct {
		lvl        string
		timeFormat string
		msName     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "givenArguments_returnNothing_success",
			args: args{
				lvl:        "DEBUG",
				timeFormat: time.Now().Format(time.RFC3339),
				msName:     "myMicroservice",
			},
		},
		{
			name: "givenArgumentsWithoutTime_returnNothing_failure",
			args: args{
				lvl:        "DEBUG",
				timeFormat: "",
				msName:     "myMicroservice",
			},
			wantErr: false,
		},

		{
			name: "givenArgument_returnSomething_success",
			args: args{
				lvl:        "testlvl",
				timeFormat: "testtimeFormat",
				msName:     "testmsName",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init(tt.args.lvl, tt.args.timeFormat, tt.args.msName); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAppLogger(t *testing.T) {
	type args struct {
		level       string
		description string
		time        int64
		items       []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "givenInfoArguments_returnNothing_success",
			args: args{
				level:       "info",
				description: "description",
				time:        1593332975979,
				items:       []string{"description | info"},
			},
		},
		{
			name: "givenDebugArguments_returnNothing_success",
			args: args{
				level:       "debug",
				description: "description",
				time:        1593332975979,
				items:       []string{"description | debug"},
			},
		},
		{
			name: "givenErrorArguments_returnNothing_success",
			args: args{
				level:       "error",
				description: "description",
				time:        1593332975979,
				items:       []string{"description | error"},
			},
		},
		{
			name: "givenWarnArguments_returnNothing_success",
			args: args{
				level:       "warn",
				description: "description",
				time:        1593332975979,
				items:       []string{"description | warn"},
			},
		},
		{
			name: "givenPipeArguments_returnNothing_Error",
			args: args{
				level:       "debug",
				description: "description",
				time:        1593332975979,
				items:       []string{"description fatal"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AppLogger(tt.args.level, tt.args.description, tt.args.time, tt.args.items...)
		})
	}
}
