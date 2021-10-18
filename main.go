package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var projectName string

func main() {
	flag.StringVar(&projectName, "p", "", "project name")
	flag.Parse()

	if projectName == "" {
		flag.PrintDefaults()
		return
	}

	createDir("/cmd")
	createFile("/cmd/readme.md", "cmd dir")
	createDir("/pkg")
	createFile("/pkg/readme.md", "pkg dir")

	createDir("/internal")
	createFile("/internal/readme.md", "internal dir")

	createDir("/api")
	createFile("/api/readme.md", "api dir")

	createDir("/configs")
	createFile("/configs/readme.md", "configs dir")

	createDir("/scripts")
	createFile("/scripts/readme.md", "configs dir")

	createDir("/build")
	createFile("/build/readme.md", "build dir")

	createDir("/depoyments")
	createFile("/depoyments/readme.md", "build dir")

	createDir("/docs")
	createFile("/docs/readme.md", "build dir")

	createDir("/examples")
	createFile("/examples/readme.md", "build dir")

	createDir("/third_party")
	createFile("/third_party/readme.md", "build dir")

	fmt.Printf("init project %s successfully\n", projectName)
}

func createDir(path string) error {
	path = fmt.Sprintf("%s%s", projectName, path)
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Printf("create dir failed, path=%s, error=%v\n", path, err)
	}
	return nil
}

func createFile(path string, content string) error {
	path = fmt.Sprintf("%s%s", projectName, path)
	err := ioutil.WriteFile(path, []byte(content), 0755)
	if err != nil {
		fmt.Printf("create file failed, path=%s, error=%v\n", path, err)
	}
	return nil
}
