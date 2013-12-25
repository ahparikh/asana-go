package asana

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"encoding/json"
)

type asanaclient struct {
	api_key  string
	client   *http.Client
}

// Public Functions
func NewAsanaClient(api_key string) *asanaclient {
	return &asanaclient{
		api_key,
		&http.Client{},
	}
}

func (this *asanaclient) GetResponse(path string) (body []byte) {
	req := this.GetRequest(path)
	resp, _ := this.client.Do(req)
	body, _ = ioutil.ReadAll(resp.Body)
	var dbg interface{}
	_ = json.Unmarshal(body, &dbg)
	fmt.Printf("%v\n", dbg)
	return
}

func (this *asanaclient) GetRequest(path string) (req *http.Request) {
	req, _ = http.NewRequest("GET", path, nil)
	this.setBasicAuth(req)
	return
}

// Private Functions
func (this *asanaclient) setBasicAuth(req *http.Request) {
	req.SetBasicAuth(this.api_key, "")
}