package model

type WeatherForecastItem struct {
	Date          string `json:"date"`
	Time          string `json:"time"`
	Status        string `json:"status"`
	Temperature   string `json:"temperature"`
	Humidity      string `json:"humidity"`
	WindDirection string `json:"wind_direction"`
	WindVelocity  string `json:"wind_velocity"`
	DamUpstream   string `json:"dam_upstream"`
}

type BMKGLocation struct {
	Provinsi  string `json:"provinsi"`
	KotaKab   string `json:"kotkab"`
	Kecamatan string `json:"kecamatan"`
	Desa      string `json:"desa"`
}

type BMKGResponse struct {
	BMKGLocation BMKGLocation `json:"lokasi"`
	Data         []struct {
		Cuaca [][]struct {
			LocalDatetime string  `json:"local_datetime"`
			T             float64 `json:"t"`
			Hu            int     `json:"hu"`
			WeatherDesc   string  `json:"weather_desc"`
			Wd            string  `json:"wd"`
			Ws            float64 `json:"ws"`
		} `json:"cuaca"`
	} `json:"data"`
}
