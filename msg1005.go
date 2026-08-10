package rtcm

type Msg1005 struct {
	MessageType               int
	StationID                 int
	ITRFRealizationYear       int
	GPSIndicator              bool
	GLONASSIndicator          bool
	GalileoIndicator          bool
	ReferenceStationIndicator bool
	ECEFXM                    float64
	OscillatorIndicator       bool
	ECEFYM                    float64
	QuarterCycleIndicator     int
	ECEFZM                    float64
}
