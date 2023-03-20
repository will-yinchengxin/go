# MUSL-GCC

## 关于 gcc、glibc

**gcc(gnu collect compiler)是一组编译工具的总称**。它主要完成的工作任务是**“预处理”和“编译”**，以及提供了与编译器紧密相关的运行库的支持，如libgcc_s.so、libstdc++.so等。

**glibc是gnu发布的libc库，也即c运行库**。glibc是linux系统中最底层的api(应用程序开发接口)，几乎其它任何的运行库都会倚赖于glibc。glibc除了封装linux操作系统所提供的系统服务外，**glibc最主要的功能就是对系统调用的封装**，你想想看，你怎么能在C代码中直接用fopen函数就能打开文件？ 打开文件最终还是要触发系统中的 sys_open 系统调用，而这中间的处理过程都是glibc来完成的, 它本身也提供了许多其它一些必要功能服务的实现，主要的如下：

- string，字符串处理
- signal，信号处理
- dlfcn，管理共享库的动态加载
- direct，文件目录操作
- elf，共享库的动态加载器，也即interpreter
- iconv，不同字符集的编码转换
- inet，socket接口的实现
- ......

## 那么怎么查看并安装 gcc 和 glibc 呢?

查看 gcc 的版本号

```shell
[root@localhost ~]# gcc --version
gcc (GCC) 4.8.5 20150623 (Red Hat 4.8.5-44)
Copyright (C) 2015 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.


[root@localhost ~]# /lib/gcc/x86_64-redhat-linux
4.8.2/ 4.8.5/
```

查看 glibc 软件版本号, 以下两种方式都可 (千万不要在[CentOS7](https://so.csdn.net/so/search?q=Centos7&spm=1001.2101.3001.7020)上升级GLIBC,)

```shell
[root@localhost ~]# rpm -qa | grep glibc
glibc-2.17-326.el7_9.x86_64
glibc-headers-2.17-326.el7_9.x86_64
glibc-common-2.17-326.el7_9.x86_64
glibc-devel-2.17-326.el7_9.x86_64


[root@localhost ~]# ldd --version
ldd (GNU libc) 2.17
Copyright (C) 2012 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
Written by Roland McGrath and Ulrich Drepper.
```

**centos7 默认的 gcc 版本是4.8.5，无法编译高版本的 glibc 2.28, 需要升级到 gcc 8.2 版本**

## CentOS7 上 GLibC 版本低，常规编译版本不能使用

在使用 Go 编程语言时，需要考虑所使用的 Go 版本与所使用的 gcc 和 glibc 版本之间的兼容性

CentOS 7默认提供的 gcc 版本为4.8.5，glibc 版本为 2.17。这些版本的 gcc 和 glibc 都已经足够支持 Golang 1.x版本。但是，如果您要编译 Golang 的源代码，则需要使用至少 gcc 4.9 版本或更高版本。[建议使用官方推荐的 gcc 版本进行编译](https://golang.org/doc/install/source#introduction)。

官方建议使用以下版本的GCC进行编译：

- Go 1.16: GCC 9.3
- Go 1.15: GCC 8.3
- Go 1.14: GCC 8.3
- Go 1.13: GCC 8.3

我们的项目使用的是 go1.19 的版本, CentOS7上 glibc 版本低，常规编译版本不能使用。需要自行源码编译,或使用使用 musl 编译版本

musl是一个轻量级的C标准库，相对于传统的glibc而言，它更小、更快、更安全，而且兼容性也很好。而且，与glibc相比，musl使用的是静态链接，这意味着Go程序使用musl编译版本时，可以不依赖于操作系统本身的glibc库，而是将所有需要的依赖项静态链接到二进制文件中，从而使得程序在运行时更加稳定和可移植。

此外，musl还可以显著减小程序的二进制大小，并提供更好的性能和更小的内存占用。因此，一些需要运行在嵌入式设备或者容器等资源受限环境中的Go程序，可能会选择使用musl编译版本。

```shell
- sudo yum install -y make automake gcc patch linux-headers glibc-static

- wget https://copr.fedorainfracloud.org/coprs/ngompa/musl-libc/repo/epel-7/ngompa-musl-libc-epel-7.repo -O /etc/yum.repos.d/ngompa-musl-libc-epel-7.repo 

- yum install -y musl-libc-static --no-check-certificate

- wget https://musl.libc.org/releases/musl-1.2.3.tar.gz

- tar -zxvf musl-1.2.3.tar.gz 

- mv musl-1.2.3 musl
 
- ./configure
 
- make & make install
 
- 安装位置在: /usr/local/musl
 

添加环境变量: vim /etc/profile
export PATH=$PATH:/usr/local/musl/bin


- source /etc/profile


# --------------------------------------------------------------------------------------------------------
sudo yum install -y make automake gcc patch linux-headers glibc-static

wget https://www.musl-libc.org/releases/musl-1.2.2.tar.gz

tar -xvzf musl-1.2.2.tar.gz

cd musl-1.2.2

./configure --prefix=/usr/local/musl && make && sudo make install

export LD_LIBRARY_PATH=/usr/local/musl/lib:$LD_LIBRARY_PATH

```

安装完成后程序的编译命令更正为

```shell
CGO_ENABLED=1 GOOS=linux  GOARCH=amd64  CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++ go build main.go -o grafana-server
```

