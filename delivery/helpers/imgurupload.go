package helpers

import (
	"Restobook/delivery/common"
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func ImgurUpload(filebytes []byte) []byte {

	url := "https://api.imgur.com/3/image"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("image", string(filebytes))
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	req.Header.Add("Authorization", fmt.Sprintf("Client-ID %v", common.IMGUR_CLIENTID))

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return body
}
