package main

import (
	"log"
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
	"github.com/wetopi/docker-volume-rbd/lib"
)

const socketAddress = "/run/docker/plugins/rbd.sock"




func main() {
	debug := os.Getenv("DEBUG")
	if ok, _ := strconv.ParseBool(debug); ok {
		logrus.SetLevel(logrus.DebugLevel)
	}

	rbdDriver, err := dockerVolumeRbd.NewDriver()
	if err != nil {
		log.Fatal(err)
	}

	h := volume.NewHandler(rbdDriver)
	logrus.Infof("listening on %s", socketAddress)
	logrus.Error(h.ServeUnix(socketAddress, 0))
}