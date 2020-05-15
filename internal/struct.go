package internal

type Aconfig struct {
	Network Network `json:"network"`
	Global  Global  `json:"global"`
	Sensors Sensors `json:"sensors"`
	Time    Time    `json:"time"`
	Pump    Pump    `json:"pump"`
	Data    Data    `json:"data"`
}

type Network struct {
	DHCP    bool   `json:"dhcp"`
	IP      string `json:"ip"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
	DNS     string `json:"dns"`
}
type Global struct {
	LcdBacklightDuration int `json:"lcdbacklightduration"`
}

type Temp struct {
	Enabled bool    `json:"enabled"`
	Init    bool    `json:"init"`
	Addr    []int8  `json:"addr"`
	Val     float32 `json:"val"`
}
type AnalogSensor struct {
	Enabled   bool    `json:"enabled"`
	Val       float32 `json:"val"`
	Threshold float32 `json:"threshold"`
}

type Sensors struct {
	Twin              Temp         `json:"twin"`
	Twout             Temp         `json:"twout"`
	Tamb              Temp         `json:"tamb"`
	WaitForConversion bool         `json:"waitforconversion"`
	TempResolution    int          `json:"tempresolution"`
	Ph                AnalogSensor `json:"ph"`
	Ch                AnalogSensor `json:"ch"`
}

type Time struct {
	Initialized bool   `json:"initialized"`
	DayLight    bool   `json:"daylight"`
	NTPServer   string `json:"ntpserver"`
	TimeZone    int    `json:"timezone"`
}
type Pump struct {
	ForceFilter bool `json:"forcefilter"`
	ForcePH     bool `json:"forceph"`
	ForceCH     bool `json:"forcech"`
}

type Alarms struct {
	Filter  bool    `json:"filter"`
	PH      bool    `json:"ph"`
	PHVal   float32 `json:"phval"`
	RTC     bool    `json:"rtc"`
	Storage bool    `json:"storage"`
}

type Data struct {
	CurTempWater   float32 `json:"curtempwater"`
	SavedTempWater float32 `json:"savedtempwater"`
	Startup        bool    `json:"startup"`
	FilterOn       bool    `json:"filteron"`
	PHOn           bool    `json:"phon"`
	Hour           int     `json:"hour"`
	Alarms         Alarms  `json:"alarms"`
}
