package main

import (
	"fmt"
	"os"
	"github.com/go-fsnotify/fsnotify"
)

func main() {

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

			case err := <-watcher.Errors:
				fmt.Println("File System Error ! ", err)
			}
		}
	}()

	filePath := os.Args[1]

	if err := watcher.Add(filePath); err != nil {
		fmt.Println("ERROR", err)
	}

	<-done
}
