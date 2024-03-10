package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Conditions struct {
	Summary     string
	Temperature float64
	//Temperature Temperature
}

//Temperature commented out as openWeather api supports units flag in endpoint
//type Temperature float64

/*func (t Temperature) Celsius() float64 {
	return float64(t) - 273.15
}*/

type OWMResponse struct {
	Weather []struct {
		Main string
	}
	Main struct {
		Temp float64
	}
}

func ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %w", data, err)
	}
	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf("invalid API response %s: want at least one Weather element", data)
	}
	conditions := Conditions{
		Summary: resp.Weather[0].Main,
		//Temperature: Temperature(resp.Main.Temp),
		Temperature: resp.Main.Temp,
	}
	return conditions, nil
}

type Client struct {
	APIKey     string
	BaseURL    string
	Units      string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		BaseURL:    "https://api.openweathermap.org",
		Units:      "metric",
		HTTPClient: http.DefaultClient,
	}
}

func (c Client) FormatURL(zip string) string {
	return fmt.Sprintf("%s/data/2.5/weather?zip=%s&appid=%s&units=%s", c.BaseURL, zip, c.APIKey, c.Units)
}

func (c Client) GetWeather(zip string) (Conditions, error) {
	URL := c.FormatURL(zip)
	resp, err := c.HTTPClient.Get(URL)
	if err != nil {
		return Conditions{}, fmt.Errorf("unexpected response status %q", resp.Status)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Conditions{}, err
	}
	conditions, err := ParseResponse(data)
	if err != nil {
		return Conditions{}, fmt.Errorf("invalid API response %s: %w", data, err)
	}
	return conditions, nil

}

// Get : handle internal paperwork
func Get(location, key string) (Conditions, error) {
	c := NewClient(key)
	conditions, err := c.GetWeather(location)
	if err != nil {
		return Conditions{}, err
	}
	return conditions, nil
}

const Usage = `Usage: weather ZIP,CountryCode
Example: weather 800010,IN`

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		return 0
	}
	//export OPENWEATHERMAP_API_KEY in zshrc ,
	//avoid entring sensitive information in shell
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		_, err := fmt.Fprintln(os.Stderr, "Please set the environment variable OPENWEATHERMAP_API_KEY.")
		if err != nil {
			panic(err)
		}
		return 1
	}
	//fmt.Printf("%v", os.Args[2])
	// os.rgs[1] is "weather" no use for now , later it can be used in printing
	// details like humidity , wind ...
	location := os.Args[2]
	conditions, err := Get(location, key)
	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)
		if err != nil {
			panic(err)
		}
		return 1
	}
	fmt.Printf(" %s %.1fºC\n", conditions.Summary, conditions.Temperature)
	return 0
}
