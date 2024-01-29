package tools

import "work7z-go/core/log"

func IsDockerMode() bool {
	return false
}

var IsDevMode = false

func IsOnlineMode() bool {
	return true
}

func ShouldNoErr(e error, label string) {
	if e != nil {
		log.Ref().Panic("FATAL_ERROR:" + label + " -> " + e.Error())
	}
}

func ShouldShowWarning(e error, label string) {
	if e != nil {
		log.Ref().Warn("FATAL_WARNING:" + label + " -> " + e.Error())
	}
}
