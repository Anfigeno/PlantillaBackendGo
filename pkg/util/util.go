package util

import (
	"os"
	"os/signal"
	"syscall"
)

func CongelarHastaRecibirSeñal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
