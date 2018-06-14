package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

const (
	IMG_INPUT_PATH       = "game.png"
	IMG_OUTPUT_HTML_PATH = "output.html"
)

func main() {
	imgBase64Str, err := imgToBase64(IMG_INPUT_PATH)
	if err != nil {
		log.Printf("Failed imgToBase64: %+v", err)
		return
	}

	fmt.Println("image to base64 string:")
	fmt.Println(imgBase64Str)

	err = base64ToHTMLPng(imgBase64Str, IMG_OUTPUT_HTML_PATH)
	if err != nil {
		log.Printf("Failed base64ToImg: %+v", err)
		return
	}
	fmt.Println()
	fmt.Println("base64 string to html -> ok")
}

// write to into a file (display in html)
func base64ToHTMLPng(base64Str, outputPath string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("Failed Create %s: %+v", outputPath, err)
	}
	defer f.Close()

	imghtml := "<html><body><img src=\"data:image/png;base64," + base64Str + "\" /></body></html>"

	_, err = f.Write([]byte(imghtml))
	if err != nil {
		return fmt.Errorf("Failed Write: %+v", err)
	}

	return nil

	///////////////////////////////////////////////////
	////////////////// Write to a file   //////////////
	///////////////////////////////////////////////////
	/*
		pngImg, err := png.Decode(bytes.NewReader([]byte(base64Str)))
		if err != nil {
			return fmt.Errorf("Failed png Decode: %+v", err)
		}

		f, err := os.Create(outputPath)
		if err != nil {
			return fmt.Errorf("Failed Create %s: %+v", outputPath, err)
		}
		defer f.Close()

		err = png.Encode(f, pngImg)
		if err != nil {
			return fmt.Errorf("Failed png Encode: %+v", err)
		}

		return nil
	*/
	///////////////////////////////////////////////////
}

func imgToBase64(intputPath string) (string, error) {
	imgFile, err := os.Open(intputPath)
	if err != nil {
		return "", fmt.Errorf("Failed Open image (%s): %+v", intputPath, err)
	}
	defer imgFile.Close()

	// create a new buffer base on file size
	fInfo, err := imgFile.Stat()
	if err != nil {
		return "", fmt.Errorf("Failed fetch image Stat: %+v", err)
	}

	size := fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)

	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)

	return imgBase64Str, nil
}
