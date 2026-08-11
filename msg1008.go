package rtcm

type Msg1008 struct {
	MessageType       int
	StationID         int
	AntennaDescriptor string
	SetupID           int
	AntennaSerial     string
}
