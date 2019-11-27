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

func command(val string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(val, args...))
}

func ExportVariable(key, val string) {
	key = strings.TrimSpace(key)

	os.Setenv(key, val)
	command("::set-env name=%s::%v", key, val)
}

func SetSecret(val string) {
	command("::add-mask::%s", val)
}

func AddPath(val string) {
	os.Setenv("PATH", fmt.Sprintf("%s%c%s", strings.TrimSpace(val), os.PathListSeparator, os.Getenv("PATH")))
	command("::add-path::%s", val)
}

func GetInput(key string, options ...*InputOptions) string {
	key = strings.ReplaceAll(strings.TrimSpace(key), " ", "_")
	envVar := "INPUT_" + strings.ToUpper(key)

	if val, ok := os.LookupEnv(envVar); ok {
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
	command("::set-output name=%s::%v", strings.TrimSpace(key), val)
}

func SetFailed(msg string) {
	Error(msg)
	os.Exit(1)
}

func Info(msg string) {
	command("%s", msg)
}

func Debug(msg string) {
	if _, file, line, ok := runtime.Caller(1); ok {
		command("::debug file=%s,line=%d::%s", file, line, msg)
	} else {
		command("::debug ::%s", msg)
	}
}

func Warning(msg string) {
	if _, file, line, ok := runtime.Caller(1); ok {
		command("::warning file=%s,line=%d::%s", file, line, msg)
	} else {
		command("::warning ::%s", msg)
	}
}

func Error(msg string) {
	if _, file, line, ok := runtime.Caller(1); ok {
		command("::error file=%s,line=%d::%s", file, line, msg)
	} else {
		command("::error ::%s", msg)
	}
}

func Pause() func() {
	id := time.Now().UnixNano()
	command("::stop-commands::%d", id)

	return func() {
		command("::%d::", id)
	}
}

func StartGroup(key string) {
	command("::group ::%s", strings.TrimSpace(key))
}

func EndGroup() {
	command("::endgroup")
}

func SaveState(key string, val interface{}) {
	command("::save-state name=%s::%v", strings.TrimSpace(key), val)
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
