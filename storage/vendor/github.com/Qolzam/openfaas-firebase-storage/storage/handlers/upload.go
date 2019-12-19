package handlers

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/api/option"

	"github.com/red-gold/telar-core/utils"
)

// UploadeHandle a function invocation
func UploadeHandle() func(http.ResponseWriter, *http.Request, httprouter.Params) {
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		fmt.Println("File Upload Endpoint Hit")
		// params from /storage/:dir
		dirName := ps.ByName("dir")
		if dirName == "" {
			errorMessage := fmt.Sprintf("Directory name is required!")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.MarshalError("dirNameRequired", errorMessage))
		}
		fmt.Printf("\n Directory name: %s", dirName)

		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handlerFile, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)

		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handlerFile.Filename)
		fmt.Printf("File Size: %+v\n", handlerFile.Size)
		fmt.Printf("MIME Header: %+v\n", handlerFile.Header)

		extension := filepath.Ext(handlerFile.Filename)
		fileNameUUID, uuidErr := uuid.NewV4()
		if uuidErr != nil {
			errorMessage := fmt.Sprintf("File name from UUID error: %s", uuidErr.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.MarshalError("fileNameUUIDError", errorMessage))

		}

		fileName := fileNameUUID.String()
		fileNameWithExtension := fmt.Sprintf("%s%s", fileName, extension)

		objectName := fmt.Sprintf("%s/%s/%s", "userId", dirName, fileNameWithExtension)
		config := &firebase.Config{
			StorageBucket: os.Getenv("bucket_name"),
		}

		opt := option.WithCredentialsFile("./serviceAccountKey.json")
		app, err := firebase.NewApp(ctx, config, opt)
		if err != nil {
			log.Fatalln(err)
		}

		client, err := app.Storage(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		bucket, err := client.DefaultBucket()
		if err != nil {
			log.Fatalln(err)
		}

		wc := bucket.Object(objectName).NewWriter(ctx)
		if _, err = io.Copy(wc, file); err != nil {
			fmt.Println(err.Error())
		}
		if err := wc.Close(); err != nil {
			fmt.Println(err.Error())
		}

		downloadURL := fmt.Sprintf("%s/%s/%s", "http://127.0.0.1:31112/function/storage-test", dirName, fileNameWithExtension)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("{ \"success\": true, \"payload\": { \"url\": \"%s\"}}", downloadURL)))

	}

}
