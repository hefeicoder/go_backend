package main

import (
	"bytes"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
	"time"

	dsp "demo/proto/demo_service_proto"
)

const (
	content_type = "application/octet-stream"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	req := dsp.DemoRequest{
		Keys: []int32{1, 2, 3, 4, 5},
	}

	b, _ := proto.Marshal(&req)
	body := bytes.NewBuffer(b)

	resp, err := c.Post("http://localhost:8080/get_feature", content_type, body)
	if err != nil {
		fmt.Printf("Post request error %s", err)
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	respProto := dsp.DemoResponse{}
	err = proto.Unmarshal(respBody, &respProto)
	if err != nil {
		fmt.Printf("Unmarshal error %s", err)
		return
	}

	fmt.Printf("Response: %s", respProto.String())
}
