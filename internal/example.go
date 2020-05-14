package internal

var Example = Aconfig{
	Global: Global{
		LcdBacklightDuration: 3000,
	},
	NetConfig: Network{
		DHCP: true,
	},
	SensConfig: Sensors{
		Twout:             Temp{},
		Tamb:              Temp{},
		Twin:              Temp{},
		WaitForConversion: true,
		TempResolution:    12,
		Ph:                AnalogSensor{},
		Ch:                AnalogSensor{},
	},
	Data: Data{},
	Time: Time{},
	Pump: Pump{},
}
