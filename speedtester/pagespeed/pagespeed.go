package pagespeed

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"theo.dev/hello-universe/speedtester/model/PerformanceResult"
)

const pageSpeedUrl = "https://www.googleapis.com/pagespeedonline/v5/runPagespeed"

type Body struct {
	LoadingExperience struct {
		OverallCategory string `json:"overall_category"`
	} `json:"loadingExperience"`
}

func TestSpeedOfAUrl(siteUrl string) (PerformanceResult.PerformanceResult, error) {

	url := fmt.Sprintf("%s?url=%s", pageSpeedUrl, siteUrl)

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile("response.json", bytes, 0777)

	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		errorString := fmt.Sprintf("Status code=%d\n", resp.StatusCode)
		return "", errors.New(errorString)
	}

	var body Body

	err = json.Unmarshal(bytes, &body)

	if err != nil {
		return "", err
	}

	return PerformanceResult.PerformanceResult(body.LoadingExperience.OverallCategory), nil
}
