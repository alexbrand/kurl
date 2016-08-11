// Copyright 2016 Alexander Brand
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/client/restclient"
	client "k8s.io/kubernetes/pkg/client/unversioned"
)

const defaultServer string = "http://127.0.0.1:8080"

var server = flag.String("s", defaultServer, "address of the K8s API Server")

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		usage()
	}
	podName := args[0]

	config := &restclient.Config{
		Host: *server,
	}
	c, err := client.New(config)
	if err != nil {
		log.Fatalf("error building client: %v\n", err)
	}

	res := c.Get().Resource("pods").Timeout(5 * time.Second).Do()
	obj, err := res.Get()
	if err != nil {
		log.Fatalf("error getting object from response: %v\n", err)
	}
	list, ok := obj.(*api.PodList)
	if !ok {
		log.Fatal("got something other than a pod list")
	}

	var podIP string
	for _, p := range list.Items {
		if strings.HasPrefix(p.Name, podName) {
			podIP = p.Status.PodIP
		}
	}
	if podIP == "" {
		log.Fatal("Pod not found")
	}

	resp, err := http.Get("http://" + podIP)
	if err != nil {
		log.Fatalf("error connecting to pod: %v\n", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatalf("error writing response to STDOUT: %v\n", err)
	}
}

func usage() {
	fmt.Println("Usage: kurl POD_NAME")
	os.Exit(1)
}
