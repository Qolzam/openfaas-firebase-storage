package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/red-gold/telar-core/utils"
)

// GetFileHandle a function invocation
func GetFileHandle() func(http.ResponseWriter, *http.Request, httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		fmt.Println("File Upload Endpoint Hit")
		// params from /storage/:dir/:name
		dirName := ps.ByName("dir")
		if dirName == "" {
			errorMessage := fmt.Sprintf("Directory name is required!")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.MarshalError("dirNameRequired", errorMessage))

		}
		fmt.Printf("\n Directory name: %s", dirName)

		// params from /storage/:dir
		fileName := ps.ByName("name")
		if fileName == "" {
			errorMessage := fmt.Sprintf("File name is required!")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(utils.MarshalError("fileNameRequired", errorMessage))

		}
		fmt.Printf("\n File name: %s", fileName)

		userId := "userId"

		objectName := fmt.Sprintf("%s/%s/%s", userId, dirName, fileName)

		// Generate download URL
		downloadURL, urlErr := generateV4GetObjectSignedURL(os.Getenv("bucket_name"), objectName, "./serviceAccountKey.json")
		if urlErr != nil {
			fmt.Println(urlErr.Error())
		}

		code := 302 // Permanent redirect, request with GET method
		if r.Method != http.MethodGet {
			// Temporary redirect, request with same method
			// As of Go 1.3, Go does not support status code 308.
			code = 307
		}

		http.Redirect(w, r, "http://127.0.0.1:31112/function/storage-test"+downloadURL, code)

	}

}
