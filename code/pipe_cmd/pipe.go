package pipe

import (
	"fmt"
	"os/exec"
)

/*
func main() {
	// simple 模式
	//pipe.SimpleCmd()
	pipe.Pipe_One()
	//pipe.Pipe_Two()
}
*/

func SimpleCmd() {
	cmd := exec.Command("ls", "-l")

	// 执行命令并等待完成
	// cmd.Run()已经启动了命令并等待它完成，因此再次尝试使用cmd.Output()会导致"exec: already started"错误。
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("Error executing command:", err)
	//	return
	//}

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting command output:", err)
		return
	}

	fmt.Println("Command Output:")
	fmt.Println(string(output))
}

func Pipe_One() {
	cmd1 := exec.Command("echo", "Hello")
	cmd2 := exec.Command("grep", "Hello")

	cmd2.Stdin, _ = cmd1.StdoutPipe()
	err := cmd1.Start()
	if err != nil {
		fmt.Println("Error Start command 1:", err)
		return
	}
	output, err := cmd2.CombinedOutput()
	if err != nil {
		fmt.Println("Error running command 2:", err)
		fmt.Println("Command 2 error output:")
		fmt.Println(string(output))
		return
	}

	err = cmd1.Wait()
	if err != nil {
		fmt.Println("Error Wait command 1:", err)
		return
	}

	fmt.Println("Final Output: " + string(output))
}

func Pipe_Two() {
	cmd := exec.Command("sh", "-c", "echo Hello | grep Hello")

	// 获取命令的标准输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// 打印输出结果
	fmt.Println("Command Output:")
	fmt.Println(string(output))
}
