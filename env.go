package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func download(url string, filename string) (string, error) {
	var filepath string
	home, err := os.UserHomeDir()

	if err != nil {
		return filepath, err
	}

	root := home + "/.gosplash/"

	_, err = os.Stat(root)

	if os.IsNotExist(err) {
		err = os.Mkdir(root, 0700)
		if err != nil {
			return filepath, err
		}
	}

	filepath = root + filename
	resp, err := http.Get(url)
	if err != nil {
		return filepath, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return filepath, err
	}
	err = os.WriteFile(filepath, b, 0700)
	return filepath, err
}

func setWallpaper(filepath string) error {
	cmd := exec.Command("/usr/lib/x86_64-linux-gnu/io.elementary.contract.set-wallpaper", filepath)
	b, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	if len(b) > 0 {
		log.Println(string(b))
	}
	return err
}
