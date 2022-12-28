package libhoggit

import "time"

type HoggitState struct {
	Players     int        `json:"players,omitempty"`
	UpdateTime  time.Time  `json:"updateTime,omitempty"`
	Theater     string     `json:"theater,omitempty"`
	Metar       string     `json:"metar,omitempty"`
	Objects     []Objects  `json:"objects,omitempty"`
	Airports    []Airports `json:"airports,omitempty"`
	Uptime      float64    `json:"uptime,omitempty"`
	Realtime    float64    `json:"realtime,omitempty"`
	ServerName  string     `json:"serverName,omitempty"`
	StartTime   int        `json:"start_time,omitempty"`
	MissionName string     `json:"missionName,omitempty"`
	MaxPlayers  int        `json:"maxPlayers,omitempty"`
	Wx          Wx         `json:"wx,omitempty"`
}
type LatLongAlt struct {
	Long float64 `json:"Long,omitempty"`
	Lat  float64 `json:"Lat,omitempty"`
	Alt  float64 `json:"Alt,omitempty"`
}
type Objects struct {
	ID         string     `json:"Id,omitempty"`
	LatLongAlt LatLongAlt `json:"LatLongAlt,omitempty"`
	Type       string     `json:"Type,omitempty"`
	Platform   string     `json:"Platform,omitempty"`
	Pilot      string     `json:"Pilot,omitempty"`
	Group      string     `json:"Group,omitempty"`
	Coalition  string     `json:"Coalition,omitempty"`
	Country    string     `json:"Country,omitempty"`
	LastSeen   int        `json:"LastSeen,omitempty"`
	Heading    float64    `json:"Heading,omitempty"`
}
type Type struct {
	Level1 int `json:"level1,omitempty"`
	Level2 int `json:"level2,omitempty"`
	Level3 int `json:"level3,omitempty"`
	Level4 int `json:"level4,omitempty"`
}
type Airports struct {
	LatLongAlt  LatLongAlt `json:"LatLongAlt,omitempty"`
	Coalition   string     `json:"Coalition,omitempty"`
	ID          int        `json:"Id,omitempty"`
	Name        string     `json:"Name,omitempty"`
	Type        Type       `json:"Type,omitempty"`
	Country     int        `json:"Country,omitempty"`
	CoalitionID int        `json:"CoalitionID,omitempty"`
}
type At8000 struct {
	Dir   int     `json:"dir,omitempty"`
	Speed float64 `json:"speed,omitempty"`
}
type At2000 struct {
	Dir   int     `json:"dir,omitempty"`
	Speed float64 `json:"speed,omitempty"`
}
type AtGround struct {
	Dir   int     `json:"dir,omitempty"`
	Speed float64 `json:"speed,omitempty"`
}
type Wind struct {
	At8000   At8000   `json:"at8000,omitempty"`
	At2000   At2000   `json:"at2000,omitempty"`
	AtGround AtGround `json:"atGround,omitempty"`
}
type Season struct {
	Temperature int `json:"temperature,omitempty"`
}
type Fog struct {
	Thickness  int `json:"thickness,omitempty"`
	Visibility int `json:"visibility,omitempty"`
}
type Visibility struct {
	Distance int `json:"distance,omitempty"`
}
type Clouds struct {
	Density   int     `json:"density,omitempty"`
	Thickness float64 `json:"thickness,omitempty"`
	Preset    string  `json:"preset,omitempty"`
	Base      float64 `json:"base,omitempty"`
	Iprecptns int     `json:"iprecptns,omitempty"`
}
type Wx struct {
	AtmosphereType   int        `json:"atmosphere_type,omitempty"`
	GroundTurbulence int        `json:"groundTurbulence,omitempty"`
	EnableFog        bool       `json:"enable_fog,omitempty"`
	Wind             Wind       `json:"wind,omitempty"`
	Season           Season     `json:"season,omitempty"`
	TypeWeather      int        `json:"type_weather,omitempty"`
	Qnh              float64    `json:"qnh,omitempty"`
	Name             string     `json:"name,omitempty"`
	Fog              Fog        `json:"fog,omitempty"`
	EnableDust       bool       `json:"enable_dust,omitempty"`
	DustDensity      int        `json:"dust_density,omitempty"`
	Visibility       Visibility `json:"visibility,omitempty"`
	Clouds           Clouds     `json:"clouds,omitempty"`
}
