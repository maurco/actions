package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/maurerlabs/actions/toolkit"
)

func main() {
	flags := []string{"--gc"}

	toolkit.ChdirFromEnv("GITHUB_WORKSPACE")
	toolkit.ChdirFromEnv("INPUT_BASE_DIR")
	toolkit.AddFlagFromEnv(&flags, 1, "baseURL", "INPUT_BASE_URL")
	toolkit.AddFlagFromEnv(&flags, 1, "buildDrafts", "INPUT_BUILD_DRAFTS")
	toolkit.AddFlagFromEnv(&flags, 1, "config", "INPUT_CONFIG")
	toolkit.AddFlagFromEnv(&flags, 1, "configDir", "INPUT_CONFIG_DIR")
	toolkit.AddFlagFromEnv(&flags, 1, "destination", "INPUT_DESTINATION")
	toolkit.AddFlagFromEnv(&flags, 1, "environment", "INPUT_ENVIRONMENT")
	toolkit.AddFlagFromEnv(&flags, 1, "minify", "INPUT_MINIFY")
	toolkit.AddFlagFromEnv(&flags, 1, "path-warnings", "INPUT_PATH_WARNINGS")

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
