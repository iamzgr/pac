package blackwhite

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func InitWhiteList() {
	for _, v := range china_cidr {
		_, in, err := net.ParseCIDR(v)
		if err != nil {
			continue
		}
		chinaNet = append(chinaNet, in)
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	res, err := client.Get("https://brook.txthinking.com/white.list")
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	ss := strings.Split(string(bytes.TrimSpace(data)), "\n")
	for i := 0; i < len(ss); i++ {
		if _, ok := white_list[ss[i]]; ok {
			continue
		}
		white_list[ss[i]] = 0
	}
}
