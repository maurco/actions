package toolkit

import (
	"fmt"
	"os"
)

func ChdirFromEnv(key string) {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		err := os.Chdir(val)
		if err != nil {
			panic(err)
		}
	}
}

func AddFlagFromEnv(flags *[]string, flagFormat int, name, key string) {
	var format string
	switch flagFormat {
	case 1:
		format = "--%s=%s"
	case 2:
		format = "--%s %s"
	case 3:
		format = "-%s=%s"
	case 4:
		format = "-%s %s"
	default:
		panic(fmt.Sprintf("AddFlagByEnvVar{flagFormat} must be an int between 1-4, got %d", flagFormat))
	}

	if val, ok := os.LookupEnv(key); ok && val != "" {
		*flags = append(*flags, fmt.Sprintf(format, name, val))
	}
}
