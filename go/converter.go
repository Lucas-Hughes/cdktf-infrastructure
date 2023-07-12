package main

import (
	"encoding/json"
	"flag"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := flag.String("root", "", "The root directory")
	outputDir := flag.String("output", "", "The output directory")
	flag.Parse()

	err := filepath.Walk(*rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), "cdk.tf.json") {
			err = processFile(path, *outputDir, *rootDir)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func processFile(path string, outputDir string, rootDir string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var data map[string]interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return err
	}

	delete(data, "data")

	relativePath, _ := filepath.Rel(rootDir, path)
	newPath := filepath.Join(outputDir, relativePath)
	err = os.MkdirAll(filepath.Dir(newPath), os.ModePerm)
	if err != nil {
		return err
	}

	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, strings.NewReader(serialize(data)))
	return err
}

func serialize(data interface{}) string {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
