package clitools

import (
	"fmt"
	"os/exec"
)

func RpiCamJPEG(path string, width int, height int) error {
	//rpicam-jpeg -o test.jpg --width 640 --height 480
	cmd := exec.Command("rpicam-jpeg", "-o", path, "--width", fmt.Sprint(width), "--height", fmt.Sprint(height))
	return cmd.Run()
}

func RpiCamVID(path string, fileName string, width int, height int, duration int) (error, error) {
	return rpiCamVIDWithMKV(path, fileName, width, height, duration),
		mkvMerge(path, fileName)
}

func rpiCamVIDWithMKV(path string, fileName string, width int, height int, duration int) error {
	//rpicam-vid -o test.h264 -t 10000 --save-pts timestamps.txt
	cmd := exec.Command("rpicam-vid",
		"-o", path+fileName+".h264",
		"-t", fmt.Sprint(duration),
		"--width", fmt.Sprint(width),
		"--height", fmt.Sprint(height),
		"--save-pts", path+fileName+"-timestamps.txt")
	return cmd.Run()
}

func mkvMerge(path string, fileName string) error {
	// mkvmerge -o test.mkv --timecodes 0:timestamps.txt test.h264
	cmd := exec.Command("mkvmerge",
		"-o", path+fileName+".mkv",
		"--timecodes", "0:"+path+fileName+"-timestamps.txt",
		path+fileName+".h264")
	return cmd.Run()
}
