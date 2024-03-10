package weather_test

import (
	"github.com/devashishRaj/goTools/weather"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestParseResponse_CorrectlyParsesJsonData(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}
	want := weather.Conditions{
		Summary:     "Clouds",
		Temperature: 284.1,
	}
	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponse_ReturnsErrorOnEmptyData(t *testing.T) {
	t.Parallel()
	_, err := weather.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty response, got nil")
	}

}

func TestParseResponse_ReturnsErrorOnInvalidResponse(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather_invalid.json")
	if err != nil {
		t.Fatal(err)
	}
	_, err = weather.ParseResponse(data)
	if err == nil {
		t.Fatal("want error on invalid json input, got nil")
	}
}

func TestFormatURL_ReturnsCorrectURLForGivenInput(t *testing.T) {
	t.Parallel()
	c := weather.NewClient("dummyAPIKey")
	c.Units = "metric"
	zip := "800010,IN"
	want := "https://api.openweathermap.org/data/2.5/weather?zip=800010,IN&appid=dummyAPIKey&units=metric"
	got := c.FormatURL(zip)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetWeather_ReturnsExpectedConditions(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "testdata/weather.json")
		}))
	defer ts.Close()
	c := weather.NewClient("dummyAPIKey")
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()
	want := weather.Conditions{
		Summary:     "Clouds",
		Temperature: 284.1,
	}
	got, err := c.GetWeather("Paris,FR")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetWeather_ReturnsExpectedResponse(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(
		http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {
				http.ServeFile(writer, request, "testdata/weather.json")
			}))
	defer ts.Close()
	c := weather.NewClient("dummyAPIKey")
	c.BaseURL = ts.URL
	c.HTTPClient = ts.Client()
	want := weather.Conditions{
		Summary:     "Clouds",
		Temperature: 284.1,
	}
	got, err := c.GetWeather("Paris,FR")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

/*func TestCelsiusCorrectlyConvertsFahrenheitToCelsius(t *testing.T) {
	t.Parallel()
	input := weather.Temperature(274.15)
	want := 1.0
	got := input.Celsius()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}*/
