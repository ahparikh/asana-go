package asana

import (
	"encoding/json"
	"log"
	"net/http"
)

type asanaclient struct {
	api_key string
	client  *http.Client
}

type ResponseData struct {
	Data interface{} `json:"data"`
}

// Public Functions
func NewAsanaClient(api_key string) *asanaclient {
	return &asanaclient{
		api_key,
		&http.Client{},
	}
}

func (this *asanaclient) GetResponse(path string, out interface{}) {
	req := this.GetRequest(path)

	resp, _ := this.client.Do(req)
	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber()

	var rd ResponseData
	err := decoder.Decode(&rd)
	if err != nil {
		log.Print(err)
	}

	json_resp, _ := json.Marshal(rd.Data)
	err = json.Unmarshal(json_resp, out)
	if err != nil {
		log.Print(err)
	}
}

func (this *asanaclient) GetRequest(path string) (req *http.Request) {
	req, _ = http.NewRequest("GET", path, nil)
	this.setBasicAuth(req)
	return req
}

// Private Functions
func (this *asanaclient) setBasicAuth(req *http.Request) {
	req.SetBasicAuth(this.api_key, "")
}
