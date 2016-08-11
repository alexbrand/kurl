# kurl: curl for Kubernetes
[![Build Status](https://travis-ci.org/alexbrand/kurl.svg?branch=master)](https://travis-ci.org/alexbrand/kurl)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexbrand/kurl)](https://goreportcard.com/report/github.com/alexbrand/kurl)

## Usage
kurl the first pod that matches the argument:
```
[root@localhost ~]# kubectl get pods
NAME                     READY     STATUS    RESTARTS   AGE
nginx-3137573019-jqvxa   1/1       Running   0          42m

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