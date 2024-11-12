package openweathermap

import (
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/rclancey/apiclient"
	"github.com/rclancey/cache/fs"
)

type OpenWeatherMapClient struct {
	client *apiclient.APIClient
}

type ClientAuth struct {
	apiKey string
}

func NewClientAuth(apiKey string) (*ClientAuth, error) {
	return &ClientAuth{apiKey: apiKey}, nil
}

func (a *ClientAuth) AuthenticateRequest(req *http.Request) error {
	query := req.URL.Query()
	query.Set("appid", a.apiKey)
	req.URL.RawQuery = query.Encode()
	return nil
}

func NewOpenWeatherMapClient(apiKey, cacheDir string, cacheTime time.Duration) (*OpenWeatherMapClient, error) {
	auth, err := NewClientAuth(apiKey)
	if err != nil {
		return nil, err
	}
	opts := apiclient.APIClientOptions{
		BaseURL: "https://api.openweathermap.org/data/",
		RequestTimeout: 0,
		CacheStore: fscache.NewFSCacheStore(cacheDir),
		MaxCacheTime: cacheTime,
		MaxRequestsPerSecond: 1.0,
		Auth: auth,
	}
	api, err := apiclient.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}
	return &OpenWeatherMapClient{client: api}, nil
}

func (c *OpenWeatherMapClient) Get(rsrc string, args url.Values) (*http.Response, error) {
	return c.client.Get(rsrc, args)
}

func (c *OpenWeatherMapClient) CurrentConditions(args url.Values) (*Observation, error) {
	obs := &Observation{}
	rsrc := "2.5/weather"
	err := c.client.GetObj(rsrc, args, obs)
	if err != nil {
		return nil, err
	}
	return obs, nil
}

func (c *OpenWeatherMapClient) Forecast(lat, lon float64) (*Forecast, error) {
	fcst := &Forecast{}
	rsrc := "3.0/onecall"
	args := url.Values{}
	args.Set("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	args.Set("lon", strconv.FormatFloat(lon, 'f', -1, 64))
	err := c.client.GetObj(rsrc, args, fcst)
	if err != nil {
		return nil, err
	}
	return fcst, nil
}
