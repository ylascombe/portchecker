package services

import (
	"bytes"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
)

func SendResultToApiserver(apiserverUrl string, mode string, res []byte) {

	postRes, err := http.Post(apiserverUrl, "application/json", bytes.NewBuffer(res))
	fmt.Fprintf(os.Stdout, "POST Result \n%v. Err %v", postRes, err)

	fmt.Println("")
	fmt.Println("")
	fmt.Fprintf(os.Stdout, "JSON RESULT \n%v", string(res))

	err = ioutil.WriteFile(fmt.Sprintf("/tmp/%v-result.json", mode), res, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot write to /tmp/%v-result.json", mode)
	}
}
