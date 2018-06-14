package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"os"

	// avoid the error: "image: unknown format"
	_ "image/gif"
	"image/jpeg"
	"image/png"
)

const (
	//IMG_INPUT_PATH = "game.png"
	IMG_INPUT_PATH = "sun.jpg"

	IMG_OUTPUT_HTML_PATH = "output.html"
	IMG_OUTPUT_FILE_PATH = "output"
)

func main() {
	// read image
	imgType, err := readImg(IMG_INPUT_PATH)
	if err != nil {
		log.Printf("Failed readImg: %+v", err)
		return
	}
	fmt.Println("Image type: ", imgType)

	// encode into base64 string
	imgBase64Str, err := imgToBase64(IMG_INPUT_PATH)
	if err != nil {
		log.Printf("Failed imgToBase64: %+v", err)
		return
	}
	fmt.Println()
	fmt.Println("image to base64 string:")
	fmt.Printf("%s ... %s \n", imgBase64Str[:10], imgBase64Str[len(imgBase64Str)-10:])

	// decode base64 string into image file
	err = base64ToFile(imgBase64Str, imgType, generagePath(imgType))
	if err != nil {
		log.Printf("Failed base64ToFilePng: %+v", err)
		return
	}
	fmt.Println()
	fmt.Println("base64 string to image file -> ok")

	// decode base64 string into html
	err = base64ToHTMLPng(imgBase64Str, IMG_OUTPUT_HTML_PATH)
	if err != nil {
		log.Printf("Failed base64ToImg: %+v", err)
		return
	}
	fmt.Println()
	fmt.Println("base64 string to html -> ok")
}

func generagePath(imgType string) string {
	return fmt.Sprintf("%s.%s", IMG_OUTPUT_FILE_PATH, imgType)
}

func readImg(path string) (string, error) {
	imgFile, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("Failed open %s: %+v", path, err)
	}
	defer imgFile.Close()

	_, imageType, err := image.Decode(imgFile)
	if err != nil {
		return "", fmt.Errorf("Failed Decode: %+v", err)
	}

	return imageType, nil
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
}

func base64ToFile(base64Str, imgType, outputPath string) error {
	imgByte, err := base64.StdEncoding.DecodeString(string(base64Str))
	if err != nil {
		return fmt.Errorf("Failed DecodeString: %+v", err)
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("Failed Create %s: %+v", outputPath, err)
	}
	defer f.Close()

	switch imgType {
	case "png":
		pngImg, err := png.Decode(bytes.NewReader(imgByte))
		if err != nil {
			return fmt.Errorf("Failed png Decode: %+v", err)
		}

		err = png.Encode(f, pngImg)
		if err != nil {
			return fmt.Errorf("Failed png Encode: %+v", err)
		}
	case "jpeg":
		jpegImg, err := jpeg.Decode(bytes.NewReader(imgByte))
		if err != nil {
			return fmt.Errorf("Failed jpeg Decode: %+v", err)
		}

		err = jpeg.Encode(f, jpegImg, nil)
		if err != nil {
			return fmt.Errorf("Failed jpeg Encode: %+v", err)
		}
	default:
		return fmt.Errorf("Unknown image type %s", imgType)
	}

	return nil
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
