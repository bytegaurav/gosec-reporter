package messenger

import (
	"Reporter/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendMessage(url string, chat models.Chat) (models.ChatResponse, error) {
	requestbody, err := json.Marshal(chat)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestbody))
	req.Header.Set("content-type", "application/json")
	if err != nil {
		return models.ChatResponse{}, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		errObj := models.ErrorResponse{}
		json.Unmarshal(body, &errObj)
		return models.ChatResponse{}, fmt.Errorf(errObj.Error.Message)
	}
	chatResponse := models.ChatResponse{}
	json.Unmarshal(body, &chatResponse)
	return chatResponse, nil

}
