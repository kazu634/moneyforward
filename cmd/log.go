package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/logutils"
)

func logInit() *logutils.LevelFilter {
	level, notEmpty := os.LookupEnv("EnvLog")
	if !notEmpty {
		level = "INFO"
	}

	return &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel(level),
		Writer:   os.Stderr,
	}
}

func logDebug(msg string) {
	target := fmt.Sprintf("[DEBUG] %s", msg)
	log.Println(target)
}

func logInfo(msg string) {
	target := fmt.Sprintf("[INFO] %s", msg)
	log.Println(target)
}

func logWarn(msg string) {
	target := fmt.Sprintf("[WARN] %s", msg)
	log.Println(target)
}

func logError(msg string) {
	target := fmt.Sprintf("[ERROR] %s", msg)
	log.Println(target)
}
