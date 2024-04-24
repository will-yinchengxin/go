# GPU

驱动程序一般指的是设备驱动程序（Device Driver），是一种可以使计算机和设备进行相互通信的特殊程序。相当于硬件的接口，操作系统只有通过这个接口，才能控制硬件设备的工作，假如某设备的驱动程序未能正确安装，便不能正常工作。因此，驱动程序被比作“硬件的灵魂”、“硬件的主宰”、“硬件和系统之间的桥梁”等。

简单来说驱动程序就是驱动硬件动起来的程序。

- 显卡驱动
- 声卡驱动
- 网卡驱动
- 外设驱动
- 主板驱动

关于I卡、N卡和A卡，这些通常是指显卡的品牌和型号。

- **N卡**：NVIDIA公司生产的显卡，以高性能和良好的驱动支持著称，在游戏和专业领域有广泛应用。
- **A卡**：AMD公司生产的显卡，以其价格性能比和某些特定领域（如高清解码）的优势受到市场欢迎。
- **I卡**：Intel公司生产的集成显卡，通常集成在CPU中，适用于日常使用和一些不太要求图形处理能力的应用。

---

Intel 提供了 `xpu-smi` 工具

```shell
xpu-smi dump -d -1 -m 0,24,25
```

Nvidia 提供了 `nvidia-smi` 工具

```shell
# 查看 device 数量
nvidia-smi -L
# 产看 gpuUtil,encUtil,decUtil 等信息
nvidia-smi stats -d gpuUtil,encUtil,decUtil
```

Netint 提供了 `ni_rsrc_mon` 工具

```shell
ni_rsrc_mon -n1 -R0 -o json1 -d
```

----

完整的 golang 代码

```go
package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	Intel             = "intel"
	Nvidia            = "nvidia"
	Netint            = "netint"
	InterlTool        = "xpu-smi"
	NvidialTool       = "nvidia-smi"
	NetIntTool        = "ni_rsrc_mon"
	InterlArgs        = []string{"dump", "-d", "-1", "-m", "0,24,25"}
	NvidiaArgs        = []string{"stats", "-d", "gpuUtil,encUtil,decUtil"}
	NvidiaMechineArgs = []string{"-L"}
	NetIntArgs        = []string{"-n0", "-R0", "-otext"}

	IntelResErr         = errors.New("Get Wrong Result By Tool " + InterlTool)
	NvidiaMechineResErr = errors.New("Get Nvidia Mechine Num Fail")
	NetIntMechineResErr = errors.New("Get Nvidia Mechine Num Fail")
)

func main() {
	// get GPU type
	gpuType := "netint"
	//gpuType := "nvdia"
	if gpuType == Intel {
		fmt.Println(getIntelInfo())
	}
	if gpuType == Nvidia {
		fmt.Println(getNvidiaInfo())
	}
	if gpuType == Netint {
		fmt.Println(getNetIntInfo())
	}
}

func getNetIntInfo() (interface{}, error) {
	cmd := exec.Command(NetIntTool, NetIntArgs...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe Error: ", err)
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("Error starting command: ", err)
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)
	var (
		mechineNum  = 0
		currentLine = 0
		decoder     bool

		res = make(map[string]interface{})
	)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Num scalers:") {
			break
		}
		if strings.Contains(line, "*************") ||
			strings.Contains(line, "from current pool at start up") ||
			regexp.MustCompile("(\\d{2}:\\d{2}:\\d{2})").MatchString(line) ||
			strings.Contains(line, "Num encoders:") ||
			strings.Contains(line, "INDEX LOAD MODEL_LOAD") {
			continue
		}
		if strings.Contains(line, "Num decoders:") {
			split := strings.Split(line, ":")
			if len(split) <= 1 {
				return nil, NetIntMechineResErr
			}
			mechineNum, err = strconv.Atoi(strings.Trim(split[1], " "))
			continue
		}
		split := strings.Split(line, " ")
		newSplit := make([]string, 0)
		for key, val := range split {
			if val == "" {
				continue
			}
			newSplit = append(newSplit, strings.Trim(split[key], " "))
		}
		//fmt.Println("newSplit", newSplit)
		//fmt.Println("mechineNum", mechineNum)
		if currentLine == mechineNum && !decoder {
			currentLine = 0
			decoder = true
		}
		currentLine++
		if decoder {
			res["Encoder_Engine_LOAD_"+newSplit[0]] = newSplit[1]
			res["Encoder_Engine_MEM_"+newSplit[0]] = newSplit[4]
			continue
		}
		res["Decoder_Engine_LOAD_"+newSplit[0]] = newSplit[1]
		res["Decoder_Engine_MEM_"+newSplit[0]] = newSplit[4]
	}

	return res, nil
}

func getIntelInfo() (interface{}, error) {
	cmd := exec.Command(InterlTool, InterlArgs...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe Error: ", err)
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("Error starting command: ", err)
		return nil, err
	}
	scanner := bufio.NewScanner(stdout)

	var (
		i        int
		splitStr = []string{}
	)
	for scanner.Scan() {
		if i >= 2 {
			break
		}
		i++
		if i == 1 {
			continue
		}
		line := scanner.Text()
		//fmt.Println("line", line)
		splitStr = strings.Split(line, ",")
		if len(splitStr) <= 3 || (len(splitStr)-3)%2 > 0 {
			return nil, IntelResErr
		}
	}

	return dealIntelInfo(splitStr), nil
}

func getNvidiaInfo() (interface{}, error) {
	mechineNum, err := getNvidiaMechineNum()
	if err != nil {
		fmt.Println("getNvidiaMechineNum Err: ", err)
		return NvidiaMechineResErr, err
	}
	mechineNum -= 1

	cmd := exec.Command(NvidialTool, NvidiaArgs...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe Error: ", err)
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("Error starting command: ", err)
		return nil, err
	}
	var (
		res = make(map[string]interface{})
	)
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		splitStr := strings.Split(line, ",")
		for key, _ := range splitStr {
			splitStr[key] = strings.Trim(splitStr[key], " ")
		}
		//fmt.Println("line", line, "splitStr", splitStr)
		_, gpuOk := res["GPU_Utilization"]
		if !gpuOk {
			res["GPU_Utilization"] = splitStr[3]
		}
		_, deOk := res["Decoder_Engine_"+splitStr[0]]
		if !deOk && splitStr[1] == "decUtil" {
			res["Decoder_Engine_"+splitStr[0]] = splitStr[3]
		}
		_, enOk := res["Encoder_Engine_"+splitStr[0]]
		if !enOk && splitStr[1] == "encUtil" {
			res["Encoder_Engine_"+splitStr[0]] = splitStr[3]
		}
		if enOk && splitStr[1] == "decUtil" && splitStr[0] == fmt.Sprintf("%d", mechineNum) {
			break
		}
	}

	return res, nil
}

func getNvidiaMechineNum() (num int, err error) {
	cmd := exec.Command(NvidialTool, NvidiaMechineArgs...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("cmd.StdoutPipe Error: ", err)
		return 0, err
	}
	if err = cmd.Start(); err != nil {
		fmt.Println("Error starting command: ", err)
		return num, err
	}
	scanner := bufio.NewScanner(stdout)
	var (
		i int
	)
	for scanner.Scan() {
		i++
	}
	num = i
	return num, nil
}

func dealIntelInfo(i []string) map[string]interface{} {
	var (
		res = make(map[string]interface{})
	)
	res["GPU_Utilization"] = i[2]
	var (
		numEngine = (len(i) - 3) / 2
	)
	decoder := i[3 : 3+numEngine]
	encoder := i[3+numEngine:]
	//fmt.Println("numEngine", numEngine)
	//fmt.Println("decoder", decoder)
	//fmt.Println("encoder", encoder)
	for k, _ := range decoder {
		res["Decoder_Engine_"+fmt.Sprintf("%d", k+1)] = decoder[k]
	}
	for k, _ := range encoder {
		res["Encoder_Engine_"+fmt.Sprintf("%d", k+1)] = encoder[k]
	}
	return res
}
```

