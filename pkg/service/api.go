package service

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

func GetAgeAPI(baseURL, param string) int {
	params := url.Values{}
	params.Add("name", param)

	u, _ := url.ParseRequestURI(baseURL)
	u.RawQuery = params.Encode()
	uStr := fmt.Sprintf("%v", u)

	r, err := http.Get(uStr)
	if err != nil {
		logrus.Errorln("Error while performing GET request:", err)
		return -1
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)

	var age struct {
		Age int `json:"age"`
	}

	json.Unmarshal(b, &age)
	//fmt.Println(age.Age)
	return age.Age
}

func GetGenderAPI(baseURL, param string) string {
	params := url.Values{}
	params.Add("name", param)

	u, _ := url.ParseRequestURI(baseURL)
	u.RawQuery = params.Encode()
	uStr := fmt.Sprintf("%v", u)

	r, err := http.Get(uStr)
	if err != nil {
		logrus.Errorln("Error while performing GET request:", err)
		return ""
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)

	var gender struct {
		Gender string `json:"gender"`
	}

	json.Unmarshal(b, &gender)
	//fmt.Println(age.Age)
	return gender.Gender
}

func GetNationalityAPI(baseURL, param string) string {
	params := url.Values{}
	params.Add("name", param)

	u, _ := url.ParseRequestURI(baseURL)
	u.RawQuery = params.Encode()
	uStr := fmt.Sprintf("%v", u)

	r, err := http.Get(uStr)
	if err != nil {
		logrus.Errorln("Error while performing GET request:", err)
		return ""
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)

	var countries struct {
		Countries []map[string]any `json:"country"`
	}

	err = json.Unmarshal(b, &countries)

	var nlty string
	var maxProb float64 = 0

	for _, dict := range countries.Countries {
		probFloat64, _ := dict["probability"].(float64)
		if probFloat64 > maxProb {
			maxProb = probFloat64
			nlty, _ = dict["country_id"].(string)
		}
	}

	return nlty
}
