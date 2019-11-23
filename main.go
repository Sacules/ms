package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"

	ms "gitlab.com/sacules/ms/schedule"
)

var lockfile = filepath.Join(ms.DataDir, "ms.lock")

func main() {
	exitsig := make(chan os.Signal, 1)
	signal.Notify(exitsig, os.Interrupt)
	locked := false

	go func() {
		<-exitsig
		if !locked {
			return
		}

		err := os.Remove(lockfile)
		if err != nil {
			fmt.Printf("couldn't remove the lock: %v", err)
			return
		}
	}()

	_, err := os.Stat(lockfile)
	if err == nil {
		fmt.Println("ms: another instance already running")
		return
	}

	if os.IsNotExist(err) {
		file, err := os.Create(lockfile)
		if err != nil {
			fmt.Printf("couldn't create the lock: %v", err)
			return
		}
		defer os.Remove(lockfile)

		file.Close()
		locked = true
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "newblock":
			newblock()

		case "status":
			err = status()
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return
	}

	tui := newTui()

	tui.init()
	tui.run()
}
