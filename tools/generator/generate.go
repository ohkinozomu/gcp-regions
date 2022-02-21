//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sanity-io/litter"
)

type Region struct {
	Name      string  `json:"name"`
	Flag      string  `json:"flag"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func main() {
	regionsJsonURL := "https://raw.githubusercontent.com/GoogleCloudPlatform/region-picker/main/data/regions.json"
	resp, err := http.Get(regionsJsonURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	regionData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var m map[string]Region
	if err := json.Unmarshal(regionData, &m); err != nil {
		panic(err)
	}

	fmt.Println(strings.Replace(litter.Sdump(m), "main.Region", "Region", -1))
}
