package openweathermap

type ForecastBase struct {
	DateEpochS int64 `json:"dt"`
	PressureHPa float64 `json:"pressure"` // Atmospheric pressure on the sea level, hPa
	HumidityPct float64 `json:"humidity"` // Humidity, %
	DewPointK float64 `json:"dew_point"` // Atmospheric temperature (varying according to pressure and humidity) below which water droplets begin to condense and dew can form. Units – default: kelvin, metric: Celsius, imperial: Fahrenheit.
	CloudPct float64 `json:"clouds"` // Cloudiness, %
	UVIndex float64 `json:"uvi"` // Current UV index
	VisibilityM float64 `json:"visibility"` // Average visibility, metres
	WindSpeedMPS float64 `json:"wind_speed"` // Wind speed. Wind speed. Units – default: metre/sec, metric: metre/sec, imperial: miles/hour. How to change units used
	WindGustMPS float64 `json:"wind_gust"` // (where available) Wind gust. Units – default: metre/sec, metric: metre/sec, imperial: miles/hour. How to change units used
	WindDeg float64 `json:"wind_deg"` // Wind direction, degrees (meteorological)
	Weather ObsWeather `json:"weather"`
}

type Current struct {
	ForecastBase
	SunriseEpochS int64 `json:"sunrise"` // Sunrise time, Unix, UTC
	SunsetEpochS int64 `json:"sunset"` // Sunset time, Unix, UTC
	TempK float64 `json:"temp"` // Temperature. Units - default: kelvin, metric: Celsius, imperial: Fahrenheit. How to change units used
	FeelsLikeK float64 `json:"feels_like"` // Temperature. This temperature parameter accounts for the human perception of weather. Units – default: kelvin, metric: Celsius, imperial: Fahrenheit.
	Rain ObsPrecip `json:"rain"`
	Snow ObsPrecip `json:"snow"`
}

type MinuteForecast struct {
	DateEpochS int64 `json:"dt"`
	PrecipitationMM float64 `json:"precipitation"`
}

type DailyTemp struct {
	MorningK float64 `json:"morn"`
	DayK float64 `json:"day"`
	EveningK float64 `json:"eve"`
	NightK float64 `json:"night"`
	MinK float64 `json:"min"`
	MaxK float64 `json:"max"`
}

type DailyForecast struct {
	ForecastBase
	SunriseEpochS int64 `json:"sunrise"` // Sunrise time, Unix, UTC
	SunsetEpochS int64 `json:"sunset"` // Sunset time, Unix, UTC
	MoonriseEpochS int64 `json:"moonrise"` // The time of when the moon rises for this day, Unix, UTC
	MoonsetEpochS int64 `json:"moonset"` // The time of when the moon sets for this day, Unix, UTC
	MoonPhase float64 `json:"moon_phase"` // Moon phase. 0 and 1 are 'new moon', 0.25 is 'first quarter moon', 0.5 is 'full moon' and 0.75 is 'last quarter moon'. The periods in between are called 'waxing crescent', 'waxing gibous', 'waning gibous', and 'waning crescent', respectively.
	Temp DailyTemp `json:"temp"` // Units – default: kelvin, metric: Celsius, imperial: Fahrenheit. How to change units used
	FeelsLike DailyTemp `json:"feels_like"`
	PrecipitationPct float64 `json:"pop"` // Probability of precipitation
	RainMM float64 `json:"rain"` // (where available) Precipitation volume, mm
	SnowMM float64 `json:"snow"` // (where available) Snow volume, mm
}

type Alert struct {
	SenderName string `json:"sender_name"` // Name of the alert source. Please read here the full list of alert sources
	Event string `json:"event"` // Alert event name
	StartEpochS int64 `json:"start"` // Date and time of the start of the alert, Unix, UTC
	EndEpochS int64 `json:"end"` // Date and time of the end of the alert, Unix, UTC
	Description string `json:"description"` // Description of the alert
	Tags []string `json:"tags"` // Type of severe weather
}

type Forecast struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	TZName string `json:"timezone"`
	TZOffsetS int64 `json:"timezone_offset"`
	Current Current `json:"current"`
	Minutely []MinuteForecast `json:"minutely"`
	Hourly []Current `json:"hourly"`
	Daily []DailyForecast `json:"daily"`
	Alerts []Alert `json:"alerts"`
}
