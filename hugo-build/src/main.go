package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/maurerlabs/github-actions/toolkit"
)

func main() {
	flags := []string{"--gc"}

	toolkit.ChangeDirByEnvVar("GITHUB_WORKSPACE")
	toolkit.ChangeDirByEnvVar("INPUT_BASE_DIR")
	toolkit.AddFlagByEnvVar(&flags, 1, "baseURL", "INPUT_BASE_URL")
	toolkit.AddFlagByEnvVar(&flags, 1, "buildDrafts", "INPUT_BUILD_DRAFTS")
	toolkit.AddFlagByEnvVar(&flags, 1, "config", "INPUT_CONFIG")
	toolkit.AddFlagByEnvVar(&flags, 1, "configDir", "INPUT_CONFIG_DIR")
	toolkit.AddFlagByEnvVar(&flags, 1, "destination", "INPUT_DESTINATION")
	toolkit.AddFlagByEnvVar(&flags, 1, "environment", "INPUT_ENVIRONMENT")
	toolkit.AddFlagByEnvVar(&flags, 1, "minify", "INPUT_MINIFY")
	toolkit.AddFlagByEnvVar(&flags, 1, "path-warnings", "INPUT_PATH_WARNINGS")

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
