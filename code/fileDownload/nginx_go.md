# golang + nginx 配置文件下载

## 配置安装 nginx

```shell
cd /usr/local/ && mkdir nginx

cd nginx/

yum -y install gcc zlib zlib-devel pcre-devel openssh openssh-devel

wget http://nginx.org/download/nginx-1.24.0.tar.gz

tar  -zxvf nginx-1.24.0.tar.gz 

cd nginx-1.24.0/

./configure 

make && make install

cd ../
````

修改配置文件

```shell
vim /usr/local/nginx/conf/nginx.conf
```

```
# 打开注释
pid        logs/nginx.pid;


events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    #tcp_nopush     on;

    #keepalive_timeout  0;
    keepalive_timeout  65;

    #gzip  on;

    # 添加的代码
    server {
        listen      8900;
        server_name _;
        location /
        {
                root  /home/share;
                autoindex on;
                autoindex_localtime on;
                autoindex_exact_size off;
        }
    }
    server {
        listen       80;
        server_name  localhost;

    

        location / {
            root   html;
            index  index.html index.htm;
        }

    
        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }

````
启动 nginx

```shell
[root@master ~]# /usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf 
````

访问浏览器

````
http://172.16.27.95:8900/
````

或者是如下配置

````
server {
    listen       8900;
    server_name  _;

    location /cdnfiles/ {
        alias /home/share/ ;
        autoindex on;
        autoindex_exact_size off;
        autoindex_localtime on;
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
````
访问浏览器

````
http://172.16.27.95:8900/cdnfiles/
````


## golang 程序

```go
package file

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"testing"
)

func TestRemoteDownloadFile(t *testing.T) {
	http.HandleFunc("/download", downloadAndPassToFrontend)
	http.ListenAndServe(":8080", nil)
}

func downloadAndPassToFrontend(w http.ResponseWriter, r *http.Request) {
	// 远程服务器上的.tar.gz文件URL
	remoteFileURL := "http://172.16.27.95:8900/download/www.bb.com/2.tar.gz"

	client := &http.Client{}

	resp, err := client.Get(remoteFileURL)
	if err != nil {
		log.Println("Error fetching file:", err)
		http.Error(w, "Error fetching file", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Failed to fetch file, status code:", resp.StatusCode)
		http.Error(w, "Failed to fetch file", resp.StatusCode)
		return
	}

	var buffer bytes.Buffer

	_, err = io.Copy(&buffer, resp.Body)
	if err != nil {
		log.Println("Error copying file content:", err)
		http.Error(w, "Error copying file content", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+filepath.Base(remoteFileURL))
	w.Header().Set("Content-Type", "application/octet-stream")

	_, err = io.Copy(w, &buffer)
	if err != nil {
		log.Println("Error writing file to response:", err)
		http.Error(w, "Error writing file to response", http.StatusInternalServerError)
		return
	}
}
````

访问服务

```shell
curl localhost:8080/download
````
