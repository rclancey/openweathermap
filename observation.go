package openweathermap

import (
)

type ObsClouds struct {
	All float64 `json:"all"`
}

type ObsCoord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type ObsMain struct {
	TemperatureK float64 `json:"temp"`
	TemperatureMaxK float64 `json:"temp_max"`
	TemperatureMinK float64 `json:"temp_min"`
	FeelsLikeK float64 `json:"feels_like"`
	HumidityPct float64 `json:"humidity"`
	PressureHPa float64 `json:"pressure"`
	GroundLevelHPa float64 `json:"grnd_level"`
	SeaLevelHPa float64 `json:"sea_level"`
}

type ObsSys struct {
	ID int64 `json:"id"`
	Country string `json:"country"`
	SunriseEpochS int64 `json:"sunrise"`
	SunsetEpochS int64 `json:"sunset"`
}

type ObsWeather struct {
	ID int64 `json:"id"`
	Name string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type ObsWind struct {
	SpeedMPS float64 `json:"speed"`
	DirectionDeg float64 `json:"deg"`
	GustMPS float64 `json:"gust"`
}

type ObsPrecip struct {
	LastHourMm float64 `json:"1h"`
	Last3HoursMm float64 `json:"3h"`
}

type Observation struct {
	Coord ObsCoord `json:"coord"`
	Weather []ObsWeather `json:"weather"`
	Main ObsMain `json:"main"`
	VisibilityM float64 `json:"visibility"`
	Wind ObsWind `json:"wind"`
	Clouds ObsClouds `json:"clouds"`
	Rain ObsPrecip `json:"rain"`
	Snow ObsPrecip `json:"snow"`
	DateEpochS int64 `json:"dt"`
	Sys ObsSys `json:"sys"`
	TZOffsetS int64 `json:"timezone"`
	Id int64 `json:"id"`
	Name string `json:"name"`
}
