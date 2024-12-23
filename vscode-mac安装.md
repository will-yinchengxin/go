### 配置 GO 环境变量
````
nano ~/.zshrc
````
按住 ` ^V` 到最后一行, 添加内容
```
export GO_HOME=/Users/<youer_base_path>/go/go1.23.2/bin
export GO_PATH=/Users/<youer_base_path>/go/bin
export SQLToDSL=/Users/<youer_base_path>/tools
export PATH=$GO_HOME:$GO_PATH:$SQLToDSL:$PATH
```
保存文件：
- 按下 Ctrl + O（即按住 Ctrl 键并按 O 键）。
- 然后会提示您确认文件名（通常是 ~/.zshrc）。直接按 Enter 键即可保存。

退出 nano：
- 按下 Ctrl + X 退出 nano 编辑器。
- 如果没有保存，nano 会询问您是否保存更改，您可以按 Y（Yes）确认保存，然后按 Enter 完成退出。

### 配置 vscode 
下载 GO 拓展, 然后在 GO 拓展设置中选择 「在 setting.json 中编辑」, 添加 gopath 选项(前提开启了 gopath)
````
.....

"go.gopath": "/Users/<youer_base_path>/go/bin"
````
保存退出

Ctrl + shift + P 调出快捷命令, 选择 「Go:Install/Update Tools」, 将所有内容全部下载下来, 重启 vscode 即可
