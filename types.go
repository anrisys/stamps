package main

import "time"

type WeatherResponse struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
				Dt   int `json:"dt"`
				Main struct {
					Temp float64 `json:"temp"`
				} `json:"main"`
				DtTxt string `json:"dt_txt"`
			} `json:"list"`
	City struct {
				Name    string `json:"name"`
				Country string `json:"country"`
			} `json:"city"`
}

type Weather struct {
	Date time.Time
	Temp float64
}
