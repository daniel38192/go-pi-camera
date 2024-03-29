package rpicam

import (
	clitools "github.com/daniel38192/go-pi-camera/cli-tools"
)

type RpiCamera struct {
	outputPath string //expected directory path, ex: "/home/user/Images/"
	fileName   string // "test1", without any extension
	width      int    //640
	height     int    //480
}

type RpiVideo struct {
	outputPath string //expected directory path, ex: "/home/user/Images/"
	fileName   string // "test1", without any extension
	width      int    // 640
	height     int    // 480
	duration   int    // duration of the recording in miliseconds
}

func NewRpiCamera(outputPath string, fileName string, width int, height int) RpiCamera {
	rpiCamera := RpiCamera{outputPath: outputPath, fileName: fileName, width: width, height: height}
	return rpiCamera
}

func NewRpiVideo(outputPath string, fileName string, width int, height int, duration int) RpiVideo {
	rpiVideo := RpiVideo{outputPath: outputPath, fileName: fileName, width: width, height: height, duration: duration}
	return rpiVideo
}

func (rpicam RpiCamera) TakeSnap() error {
	return clitools.RpiCamJPEG(rpicam.outputPath+rpicam.fileName+".jpeg", rpicam.width, rpicam.height)
}

func (rpicam RpiVideo) Record() (error, error) {
	return clitools.RpiCamVID(rpicam.outputPath, rpicam.fileName, rpicam.width, rpicam.height, rpicam.duration)
}
