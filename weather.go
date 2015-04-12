package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type WunderJson struct {
	Forecast struct {
		Simpleforecast struct {
			Forecastday []struct {
				Avehumidity float64 `json:"avehumidity"`
				Avewind     struct {
					Degrees float64 `json:"degrees"`
					Dir     string  `json:"dir"`
					Kph     float64 `json:"kph"`
					Mph     float64 `json:"mph"`
				} `json:"avewind"`
				Conditions string `json:"conditions"`
				Date       struct {
					Ampm           string  `json:"ampm"`
					Day            float64 `json:"day"`
					Epoch          string  `json:"epoch"`
					Hour           float64 `json:"hour"`
					Isdst          string  `json:"isdst"`
					Min            string  `json:"min"`
					Month          float64 `json:"month"`
					Monthname      string  `json:"monthname"`
					MonthnameShort string  `json:"monthname_short"`
					Pretty         string  `json:"pretty"`
					Sec            float64 `json:"sec"`
					TzLong         string  `json:"tz_long"`
					TzShort        string  `json:"tz_short"`
					Weekday        string  `json:"weekday"`
					WeekdayShort   string  `json:"weekday_short"`
					Yday           float64 `json:"yday"`
					Year           float64 `json:"year"`
				} `json:"date"`
				High struct {
					Celsius    string `json:"celsius"`
					Fahrenheit string `json:"fahrenheit"`
				} `json:"high"`
				Icon    string `json:"icon"`
				IconURL string `json:"icon_url"`
				Low     struct {
					Celsius    string `json:"celsius"`
					Fahrenheit string `json:"fahrenheit"`
				} `json:"low"`
				Maxhumidity float64 `json:"maxhumidity"`
				Maxwind     struct {
					Degrees float64 `json:"degrees"`
					Dir     string  `json:"dir"`
					Kph     float64 `json:"kph"`
					Mph     float64 `json:"mph"`
				} `json:"maxwind"`
				Minhumidity float64 `json:"minhumidity"`
				Period      float64 `json:"period"`
				Pop         float64 `json:"pop"`
				QpfAllday   struct {
					In float64 `json:"in"`
					Mm float64 `json:"mm"`
				} `json:"qpf_allday"`
				QpfDay struct {
					In float64 `json:"in"`
					Mm float64 `json:"mm"`
				} `json:"qpf_day"`
				QpfNight struct {
					In float64 `json:"in"`
					Mm float64 `json:"mm"`
				} `json:"qpf_night"`
				Skyicon    string `json:"skyicon"`
				SnowAllday struct {
					Cm float64 `json:"cm"`
					In float64 `json:"in"`
				} `json:"snow_allday"`
				SnowDay struct {
					Cm float64 `json:"cm"`
					In float64 `json:"in"`
				} `json:"snow_day"`
				SnowNight struct {
					Cm float64 `json:"cm"`
					In float64 `json:"in"`
				} `json:"snow_night"`
			} `json:"forecastday"`
		} `json:"simpleforecast"`
		TxtForecast struct {
			Date        string `json:"date"`
			Forecastday []struct {
				Fcttext       string  `json:"fcttext"`
				FcttextMetric string  `json:"fcttext_metric"`
				Icon          string  `json:"icon"`
				IconURL       string  `json:"icon_url"`
				Period        float64 `json:"period"`
				Pop           string  `json:"pop"`
				Title         string  `json:"title"`
			} `json:"forecastday"`
		} `json:"txt_forecast"`
	} `json:"forecast"`
	MoonPhase struct {
		AgeOfMoon   string `json:"ageOfMoon"`
		CurrentTime struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"current_time"`
		Hemisphere         string `json:"hemisphere"`
		PercentIlluminated string
		PhaseofMoon        string `json:"phaseofMoon"`
		Sunrise            struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunrise"`
		Sunset struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunset"`
	} `json:"moon_phase"`
	Response struct {
		Features struct {
			Astronomy float64 `json:"astronomy"`
			Forecast  float64 `json:"forecast"`
		} `json:"features"`
		TermsofService string `json:"termsofService"`
		Version        string `json:"version"`
	} `json:"response"`
	SunPhase struct {
		Sunrise struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunrise"`
		Sunset struct {
			Hour   string `json:"hour"`
			Minute string `json:"minute"`
		} `json:"sunset"`
	} `json:"sun_phase"`
}

func main() {
	weatherUndergroundKey := ""
	zipCode := ""
	weatherUrl := fmt.Sprintf("http://api.wunderground.com/api/%s/forecast/astronomy/q/%s.json", weatherUndergroundKey, zipCode)
	response, err := http.Get(weatherUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	byteContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var weatherJson WunderJson
	if err := json.Unmarshal(byteContent, &weatherJson); err != nil {
		panic(err)
	}
	_, month, day := time.Now().Date()
	fmt.Printf("Data Generated somewhere around %s %d at %s\n\n", month.String(), day, weatherJson.Forecast.TxtForecast.Date)
	for _, day := range weatherJson.Forecast.TxtForecast.Forecastday {
		fmt.Printf("# %s:\n", day.Title)
		fmt.Printf("  +%s\n\n", day.Fcttext)
	}
	fmt.Printf("The Moon is %s%% illuminated tonight.\n", weatherJson.MoonPhase.PercentIlluminated)
}
