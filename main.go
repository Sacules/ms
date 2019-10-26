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
	lockstate := false

	go func() {
		<-exitsig
		if lockstate {
			err := os.Remove(lockfile)
			if err != nil {
				fmt.Printf("couldn't remove the lock: %v", err)
				return
			}
		}
		os.Exit(0)
	}()

	if _, err := os.Stat(lockfile); err == nil {
		fmt.Printf("Cannot run program : Another Instance already running")
		return

	} else if os.IsNotExist(err) {
		var file, err = os.Create(lockfile)
		if err != nil {
			fmt.Printf("couldn't create the lock: %v", err)
			return
		}
		file.Close()
		defer os.Remove(lockfile)
		lockstate = true
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "newblock" {
			newblock()
		}

		return
	}

	tui := newTui()

	tui.init()
	tui.run()
}
