package main

import (
	"io/ioutil"
	"os"
)

var configdir string = os.Getenv("HOME") + "/.config/git/"

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "help" {
			println("Usage: gitprofile profilename")
		} else {
			// Changes config file
			_, err := os.Lstat(configdir + "config")
			if err == nil {
				os.Remove(configdir + "config")
			}
			os.Symlink(configdir+os.Args[1]+".gitprofile", configdir+"config")
		}
	} else {
		// Prints current config file
		configfile, err := ioutil.ReadFile(configdir + "config")
		if err != nil {
			println(err)
		}
		println(string(configfile))
	}
}
