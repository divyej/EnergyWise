package models

type WeatherForecast struct {
	Timelines Timelines
	Location  Location
}

type RealtimeWeather struct {
	Data     Minutely
	Location Location
}

// Returns Timelines with Minutely nil
type WeatherHistory struct {
	Timelines Timelines
	Location  Location
}

// Timelines (Minutely, Hourly, Daily) for Weather related data. 
// Each data has its own fields, sometimes common, sometimes unique.
// Minutely is nil for WeatherHistory
type Timelines struct {
	Minutely []Minutely
	Hourly   []Hourly
	Daily    []Daily
}

type Minutely struct {
	Time   string
	Values MinutelyData
}

type MinutelyData struct {
	CloudBase                float64
	CloudCeiling             float64
	CloudCover               float64
	DewPoint                 float64
	FreezingRainIntensity    float64
	Humidity                 float64
	PrecipitationProbability float64
	PressureSurfaceLevel     float64
	RainIntensity            float64
	SleetIntensity           float64
	SnowIntensity            float64
	Temperature              float64
	TemperatureApparent      float64
	UvHealthConcern          float64
	UvIndex                  float64
	Visibility               float64
	WeatherCode              float64
	WindDirection            float64
	WindGust                 float64
	WindSpeed                float64
}

type Hourly struct {
	Time   string
	Values HourlyData
}

type HourlyData struct {
	CloudBase                float64
	CloudCeiling             float64
	CloudCover               float64
	DewPoint                 float64
	Evapotranspiration       float64
	FreezingRainIntensity    float64
	Humidity                 float64
	IceAccumulation          float64
	IceAccumulationLwe       float64
	PrecipitationProbability float64
	PressureSurfaceLevel     float64
	RainAccumulation         float64
	RainAccumulationLwe      float64
	RainIntensity            float64
	SleetAccumulation        float64
	SleetAccumulationLwe     float64
	SleetIntensity           float64
	SnowAccumulation         float64
	SnowAccumulationLwe      float64
	SnowDepth                float64
	SnowIntensity            float64
	Temperature              float64
	TemperatureApparent      float64
	UvHealthConcern          float64
	UvIndex                  float64
	Visibility               float64
	WeatherCode              float64
	WindDirection            float64
	WindGust                 float64
	WindSpeed                float64
}

type Daily struct {
	Time   string
	Values DailyData
}

type DailyData struct {
	CloudBaseAvg                float64
	CloudBaseMax                float64
	CloudBaseMim                float64
	CloudCeilingAvg             float64
	CloudCeilingMax             float64
	CloudCeilingMin             float64
	CloudCoverAvg               float64
	CloudCoverMax               float64
	CloudCoverMin               float64
	DewPointAvg                 float64
	DewPointMax                 float64
	DewPointMin                 float64
	EvapotranspirationAvg       float64
	EvapotranspirationMax       float64
	EvapotranspirationMin       float64
	EvapotranspirationSum       float64
	FreezingRainIntensityAvg    float64
	FreezingRainIntensityMax    float64
	FreezingRainIntensityMin    float64
	HumidityAvg                 float64
	HumidityMax                 float64
	HumidityMin                 float64
	IceAccumulationAvg          float64
	IceAccumulationLweAvg       float64
	IceAccumulationLweMax       float64
	IceAccumulationLweMin       float64
	IceAccumulationLweSum       float64
	IceAccumulationMax          float64
	IceAccumulationMin          float64
	IceAccumulationSum          float64
	MoonriseTime                string
	MoonsetTime                 string
	PrecipitationProbabilityAvg float64
	PrecipitationProbabilityMax float64
	PrecipitationProbabilityMin float64
	PressureSurfaceLevelAvg     float64
	PressureSurfaceLevelMax     float64
	PressureSurfaceLevelMin     float64
	RainAccumulationAvg         float64
	RainAccumulationLweAvg      float64
	RainAccumulationLweMax      float64
	RainAccumulationLweMin      float64
	RainAccumulationMax         float64
	RainAccumulationMin         float64
	RainAccumulationSum         float64
	RainIntensityAvg            float64
	RainIntensityMax            float64
	RainIntensityMin            float64
	SleetAccumulationAvg        float64
	SleetAccumulationLweAvg     float64
	SleetAccumulationLweMax     float64
	SleetAccumulationLweMin     float64
	SleetAccumulationLweSum     float64
	SleetAccumulationMax        float64
	SleetAccumulationMin        float64
	SleetIntensityAvg           float64
	SleetIntensityMax           float64
	SleetIntensityMin           float64
	SnowAccumulationAvg         float64
	SnowAccumulationLweAvg      float64
	SnowAccumulationLweMax      float64
	SnowAccumulationLweMin      float64
	SnowAccumulationLweSum      float64
	SnowAccumulationMax         float64
	SnowAccumulationMin         float64
	SnowAccumulationSum         float64
	SnowDepthAvg                float64
	SnowDepthMax                float64
	SnowDepthMin                float64
	SnowDepthSum                float64
	SnowIntensityAvg            float64
	SnowIntensityMax            float64
	SnowIntensityMin            float64
	SunriseTime                 string
	SunsetTime                  string
	TemperatureApparentAvg      float64
	TemperatureApparentMax      float64
	TemperatureApparentMin      float64
	TemperatureAvg              float64
	TemperatureMax              float64
	TemperatureMin              float64
	UvHealthConcernAvg          float64
	UvHealthConcernMax          float64
	UvHealthConcernMin          float64
	UuvIndexAvg                 float64
	UvIndexMax                  float64
	UvIndexMin                  float64
	VisibilityAvg               float64
	VisibilityMax               float64
	VisibilityMin               float64
	WeatherCodeMax              float64
	WeatherCodeMin              float64
	WindDirectionAvg            float64
	WindGustAvg                 float64
	WindGustMax                 float64
	WindGustMin                 float64
	WindSpeedAvg                float64
	WindSpeedMax                float64
	WindSpeedMin                float64
}
