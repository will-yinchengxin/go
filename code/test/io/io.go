package io

/*
	计算机和信息领域里 I/O这个术语表示输入 / 输出

	1) Go语言里使用io.Reader和io.Writer两个 interface 来抽象I/O
		io.Reader/Writer 常用的几种实现
		net.Conn: 表示网络连接。
		os.Stdin, os.Stdout, os.Stderr: 标准输入、输出和错误。
		os.File: 网络,标准输入输出,文件的流读取。
		strings.Reader: 字符串抽象成 io.Reader 的实现。
		bytes.Reader: []byte抽象成 io.Reader 的实现。
		bytes.Buffer: []byte抽象成 io.Reader 和 io.Writer 的实现。
		bufio.Reader/Writer: 带缓冲的流读取和写入（比如按行读写）


*/

func IO()  {
	//net.Conn()
}
