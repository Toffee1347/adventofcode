package utils

import (
	"embed"
	"fmt"
	"os"
)

//go:embed assets/*
var assets embed.FS

func GetAsset(file string) string {
	rawData, err := assets.ReadFile("assets/" + file)
	if err != nil {
		fmt.Printf("Failed to open asset "+file+": %s", err)
		os.Exit(1)
	}
	return string(rawData)
}
