package method

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostJSON(requestString string, url string) (string, error) {
	requestByte := []byte(requestString)

	request, err := http.NewRequest("POST", url, bytes.NewReader(requestByte))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(responseBytes), nil
}
