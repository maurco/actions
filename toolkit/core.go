package toolkit

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type InputOptions struct {
	Required bool
	Fallback string
}

func ExportVar(key, val string) {
	key = strings.TrimSpace(key)

	os.Setenv(key, val)
	Info("::set-env name=%s::%s", key, val)
}

func SetSecret(val string) {
	Info("::add-mask::%s", val)
}

func AddPath(val string) {
	os.Setenv("PATH", fmt.Sprintf("%s%c%s", strings.TrimSpace(val), os.PathListSeparator, os.Getenv("PATH")))
	Info("::add-path::%s", val)
}

func GetInput(key string, options ...*InputOptions) string {
	key = strings.ReplaceAll(strings.TrimSpace(key), " ", "_")
	envVar := "INPUT_" + strings.ToUpper(key)

	if val, ok := os.LookupEnv(envVar); ok && val != "" {
		return val
	}

	if len(options) > 0 && options[0].Required {
		panic(fmt.Sprintf("Missing input value: %s", key))
	}

	if len(options) > 0 && options[0].Fallback != "" {
		return options[0].Fallback
	}

	return ""
}

func SetOutput(key string, val interface{}) {
	Info("::set-output name=%s::%v", strings.TrimSpace(key), val)
}

func SetFailed(msg string) {
	Error(msg)
	os.Exit(1)
}

func Info(msg string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(msg, args...))
}

func Debug(msg string) {
	if _, file, line, ok := runtime.Caller(1); ok {
		Info("::debug file=%s,line=%d::%s", file, line, msg)
	} else {
		Info("::debug::%s", msg)
	}
}

func Warning(msg string) {
	if _, file, line, ok := runtime.Caller(1); ok {
		Info("::warning file=%s,line=%d::%s", file, line, msg)
	} else {
		Info("::warning::%s", msg)
	}
}

func Error(msg string) {
	if _, file, line, ok := runtime.Caller(1); ok {
		Info("::error file=%s,line=%d::%s", file, line, msg)
	} else {
		Info("::error::%s", msg)
	}
}

func Pause() func() {
	id := time.Now().UnixNano()
	Info("::stop-commands::%d", id)

	return func() {
		Info("::%d::", id)
	}
}

func StartGroup(key string) {
	Info("::group::%s", strings.TrimSpace(key))
}

func EndGroup() {
	Info("::endgroup::")
}

func SaveState(key, val string) {
	Info("::save-state name=%s::%s", strings.TrimSpace(key), val)
}

func GetState(key string, options ...*InputOptions) string {
	envVar := "STATE_" + strings.TrimSpace(key)

	if val, ok := os.LookupEnv(envVar); ok {
		return val
	}

	if len(options) > 0 && options[0].Required {
		panic(fmt.Sprintf("Missing state value: %s", key))
	}

	if len(options) > 0 && options[0].Fallback != "" {
		return options[0].Fallback
	}

	return ""
}
