package main

import (
	"fmt"
	"os/exec"
	"os"
	"github.com/go-fsnotify/fsnotify"
)
func configMapUpdate(path string, configMapName string ) {
	out, err := exec.Command("bash", "-c", "kubectl create configmap"+configMapName+"--from-file=./"+path+" -o yaml --dry-run | kubectl replace -f -").Output()
	if err != nil {
			fmt.Println("Error For Kubernetes ConfigMap Update ", err)
	}else{
		fmt.Println(out)
	}
}

func main() {
	filePath := os.Args[1]
	configMapName := os.Args[2]

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()


	done := make(chan bool)


	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Printf("File System Changing !  %#v\n", event)
				configMapUpdate(filePath, configMapName)


			case err := <-watcher.Errors:
				fmt.Println("File System Error ! ", err)
			}
		}
	}()


	if err := watcher.Add(filePath); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
}
