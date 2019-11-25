package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func addFlag(flags []string, key, envKey string) []string {
	if val, ok := os.LookupEnv(envKey); ok && val != "" {
		return append(flags, fmt.Sprintf("--%s=%s", key, val))
	}

	return flags
}

func main() {
	flags := []string{"--gc"}
	flags = addFlag(flags, "baseURL", "INPUT_BASE_URL")
	flags = addFlag(flags, "buildDrafts", "INPUT_BUILD_DRAFTS")
	flags = addFlag(flags, "config", "INPUT_CONFIG")
	flags = addFlag(flags, "configDir", "INPUT_CONFIG_DIR")
	flags = addFlag(flags, "destination", "INPUT_DESTINATION")
	flags = addFlag(flags, "environment", "INPUT_ENVIRONMENT")
	flags = addFlag(flags, "minify", "INPUT_MINIFY")
	flags = addFlag(flags, "path-warnings", "INPUT_PATH_WARNINGS")

	if val, ok := os.LookupEnv("GITHUB_WORKSPACE"); ok && val != "" {
		err := os.Chdir(val)
		if err != nil {
			log.Fatal(err)
		}
	}

	if val, ok := os.LookupEnv("INPUT_BASE_DIR"); ok && val != "" {
		err := os.Chdir(val)
		if err != nil {
			log.Fatal(err)
		}
	}

	cmd := exec.Command("hugo", flags...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
