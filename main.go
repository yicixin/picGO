package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/yicixin/pigo/config"
	uploader2 "github.com/yicixin/pigo/uploader"
	"github.com/yicixin/pigo/uploader/ali"
)

var (
	configPath = flag.String("f", "", "Path to the configuration file")
	filePath   = flag.String("p", "", "Path to the file to upload")
)

func main() {
	flag.Parse()

	if configPath == nil || len(*configPath) == 0 {
		fmt.Println("No configuration file provided")
		return
	}

	if filePath == nil || len(*filePath) == 0 {
		fmt.Println("No file path provided")
		return
	}

	// Load the configuration file
	content, err := os.ReadFile(*configPath)
	if err != nil {
		fmt.Println("Error reading configuration file")
		return
	}
	cfg, err := config.Load(content)
	if err != nil {
		fmt.Printf("Error loading configuration file:%s\n", err)
		return
	}

	var uploader uploader2.Uploader
	switch cfg.Type {
	case config.ConfigTypeAliOSS:
		var c *config.AliOSSConfig
		c, err = config.LoadAliOSSConfig(content)
		if err != nil {
			fmt.Printf("Error loading AliOSS configuration:%s\n", err)
			return
		}

		uploader, err = ali.NewAliUploader(c)
		if err != nil {
			fmt.Printf("Error creating AliUploader:%s\n", err)
			return
		}
	default:
		fmt.Println("Unsupported configuration type")
	}

	fd, err := os.Open(*filePath)
	if err != nil {
		fmt.Printf("Error opening file:%s\n", err)
		return
	}
	defer func() {
		_ = fd.Close()
	}()

	filename := fmt.Sprintf("%s-%s%s", time.Now().Format("2006-01-02"), uuid.New().String(), path.Ext(fd.Name()))
	url, err := uploader.Upload(filename, fd)
	if err != nil {
		fmt.Printf("Upload failed:%s\n", err)
		return
	}
	fmt.Printf("%s\n", url)
}
