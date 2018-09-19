package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	maxUploadSize = 10 * 1024 * 1024 // 10 MB
)

func main() {
	http.HandleFunc("/upload1", handlerUpload1)
	http.HandleFunc("/upload2", handlerUpload2)
	http.HandleFunc("/index", handlerIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, uploadHTML)
}

func handlerUpload1(w http.ResponseWriter, r *http.Request) {
	// check file-size
	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		log.Printf("Failed ParseMultipartForm: %+v", err)
		return
	}

	for i := 1; i <= 2; i++ {
		file, _, err := r.FormFile(fmt.Sprintf("uploadfile%d", i))
		if err != nil {
			log.Printf("Failed FormFile for file %d: %+v", i, err)
			return
		}
		defer file.Close()

		// copy to local file
		f, err := os.Create(fmt.Sprintf("tem_local%d", i))
		if err != nil {
			log.Printf("Failed Create file: %+v", err)
			return
		}
		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			log.Printf("Failed Copy file: %+v", err)
			return
		}

		log.Printf("upload file %d ok", i)
	}

	w.Write([]byte("SUCCESS"))
}

func handlerUpload2(w http.ResponseWriter, r *http.Request) {
	// check file-size
	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		log.Printf("Failed ParseMultipartForm: %+v", err)
		return
	}

	formdata := r.MultipartForm
	for inputName, files := range formdata.File {
		for i := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Printf("Failed open file %s-%d: %+v", inputName, i, err)
				return
			}

			// display the content
			b, err := ioutil.ReadAll(file)
			if err != nil {
				log.Printf("Failed ReadAll %s-%d: %+v", inputName, i, err)
				return
			}

			log.Printf("Content of file %s: %s", inputName, b)
		}
	}

	w.Write([]byte("SUCCESS"))
}

const (
	uploadHTML = `
	<html>
<head>
       <title>Upload file</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://localhost:8080/upload1" method="post">
	<input type="file" name="uploadfile1" /> <br />
	<input type="file" name="uploadfile2" /> <br />
    <input type="submit" value="upload" />
</form>
</body>
</html>
`
)
