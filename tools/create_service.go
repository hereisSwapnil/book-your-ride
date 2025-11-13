package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var serviceName = flag.String("name", "", "The name of the service to create")
	flag.Parse()

	if *serviceName == "" {
		log.Fatal("Service name is required")
		os.Exit(1)
	}

	var basePath = filepath.Join("services", *serviceName+"-service")
	dirs := []string{
		"cmd",
		"internal/domain",
		"internal/service",
		"internal/infrastructure/events",
		"internal/infrastructure/grpc",
		"internal/infrastructure/repository",
		"pkg/types",
	}

	for _, dir := range dirs {
		var fullPath = filepath.Join(basePath, dir)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			log.Fatal("Failed to create directory: ", err)
			os.Exit(1)
		}
	}

	var readmePath = filepath.Join(basePath, "README.md")
	readmeContent := fmt.Sprintf("# %s-service", *serviceName)
	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		log.Fatal("Failed to create README.md: ", err)
		os.Exit(1)
	}
	fmt.Println("Service created successfully")
}