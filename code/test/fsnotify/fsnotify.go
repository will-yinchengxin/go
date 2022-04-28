package fsnotify

import (
	"log"
	"github.com/fsnotify/fsnotify"
)

func Fsnotify() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <- watcher.Events:
				if !ok {
					return
				}
				// 打印监听事件
				log.Println("event:", event)
			case _, ok := <- watcher.Errors:
				if !ok {
					return
				}
			}
		}
	}()

	// 监听当前目录
	err = watcher.Add("./")
	if err != nil {
		log.Fatal(err)
	}
	<- done
}

/*
	操作系统提供了三个接口来支撑:
		inotify_init1
		inotify_add_watch
		inotify_rm_watch

	内核要上报这些文件 api 事件必然要采集这些事件。在哪一个内核层次采集的呢？
		系统调用 -> vfs -> 具体文件系统（ ext4 ）-> 块层 -> scsi 层
	答案是 vfs 曾, vfs 是所有"文件"操作的入口

	fsnotify的调用栈:
	fsnotify
		-> send_to_group
			-> inotify_handle_event
				-> fsnotify_add_event
					-> wake_up （唤醒等待队列，也就是 epoll）
	把事件通知给关注的 fsnotify_group结构体,也就是通知给 inotify fd

	inotify 也支持 epoll机制
		inotify fd 支持 epoll 机制。最明显的两个特征：
			1) notify fd 的 inotify_fops 实现了 .poll 接口；
			2) inotify fd 相关的某个结构体一定有个 wait 队列的表头

	1) Go 的 fsnotify 其实操作系统能力的浅层封装，Linux 本质就是对 inotify 机制；
	2) inotify 也是一个特殊句柄，属于匿名句柄之一，这个句柄用于文件的事件监控；
*/