package collection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


type CollectionEntity struct {
	EndPoint string
}

func (c *CollectionEntity) Init() {
	endpoint := "v1/database/collections"
	url := fmt.Sprintf("%s:%s/%s", os.Getenv("BASEPATH"), os.Getenv("PORT"), endpoint)
	c.EndPoint = url
}


func (c *CollectionEntity) GetCollections() (r CollectionResponse) {
	
	client := &http.Client{}
	req, reqErr := http.NewRequest("GET", c.EndPoint, nil)
	if reqErr != nil {
		panic(reqErr.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Appwrite-Project", os.Getenv("PROJECT_ID"))
	req.Header.Add("X-Appwrite-key", os.Getenv("TOKEN"))
	
	resp, respErr := client.Do(req)
	if respErr != nil {
		panic(respErr.Error())
	}
	defer resp.Body.Close()
	
	response := CollectionResponse{}
	parseErr := json.NewDecoder(resp.Body).Decode(&response)
	if parseErr != nil {
		panic(parseErr.Error())
	}
	return response
}


func (c *CollectionEntity) CreateCollection(input CollectionCreateInput) int {
	client := &http.Client{}
	data, dataErr := json.Marshal(input)
	if dataErr != nil {
		panic(dataErr)
	}
	req, reqErr := http.NewRequest("POST", c.EndPoint, bytes.NewBuffer(data))
	if reqErr != nil {
		panic(reqErr)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Appwrite-Project", os.Getenv("PROJECT_ID"))
	req.Header.Add("X-Appwrite-key", os.Getenv("TOKEN"))

	resp, respErr := client.Do(req)
	if respErr != nil {
		panic(respErr)
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
