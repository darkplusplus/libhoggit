package libhoggit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	Campaign_GeorgiaAtWar     = "gaw"
	Campaign_PersianGulfAtWar = "pgaw"
)

type HoggitClient struct {
	Campaign string
}

func Client(Campaign string) HoggitClient {
	c := HoggitClient{
		Campaign: Campaign,
	}

	return c
}

func (c *HoggitClient) GetState() (*HoggitState, error) {
	url := fmt.Sprintf("https://statecache.hoggitworld.com/%s?", c.Campaign)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var state HoggitState
	err = json.NewDecoder(res.Body).Decode(&state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}
