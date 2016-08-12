# kurl: curl for Kubernetes
[![Build Status](https://travis-ci.org/alexbrand/kurl.svg?branch=master)](https://travis-ci.org/alexbrand/kurl)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexbrand/kurl)](https://goreportcard.com/report/github.com/alexbrand/kurl)

kurl makes it easy to quickly issue a GET request to a pod running on your Kubernetes cluster. 
It also supports a more advanced "proxy mode", which allows you to talk to pods using the 
tool of your choice.

## Usage
```
[root@localhost ~]# ./kurl -h
kurl: curl for Kubernetes

Usage: kurl POD_NAME
      --proxy[=false]: start in proxy mode
      --proxy-port="9090": set the port when running in proxy mode
  -s, --server="http://127.0.0.1:8080": address of the K8s API Server
```

## Example
Assuming you have an nginx pod running on your Kubernetes cluster:
```
[root@localhost ~]# kubectl get pods
NAME                     READY     STATUS    RESTARTS   AGE
nginx-3137573019-jqvxa   1/1       Running   0          42m
```

### Basic (GET request)
```
[root@localhost ~]# kurl nginx
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>
``` 

### Advanced (Proxy mode)
```
[root@localhost ~] kurl --proxy &
[1] 24271
Starting proxy on port 9090

[root@localhost ~] export http_proxy=localhost:9090
[root@localhost ~]# curl --head nginx
HTTP/1.1 200 OK
Accept-Ranges: bytes
Content-Length: 612
Content-Type: text/html
Date: Fri, 12 Aug 2016 02:37:13 GMT
Etag: "574da256-264"
Last-Modified: Tue, 31 May 2016 14:40:22 GMT
Server: nginx/1.11.1
```