package dir

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func GetWD() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

func LookPath() {
	fmt.Println(os.Args[0])
	binary, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(binary)
	root := filepath.Dir(filepath.Dir(path))
	fmt.Println(root)
}