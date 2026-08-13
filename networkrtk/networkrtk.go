package networkrtk

import "github.com/Monster0506/rtcm-go/core"

type NetworkRTKHeader struct {
	MessageType                 int
	NetworkID                   int
	SubnetworkID                int
	GPSEpochTimeS               float64
	MultipleMessageIndicator    bool
	MasterReferenceStationID    int
	AuxiliaryReferenceStationID int
	NumGPSSats                  int
}

func decodeNetworkRTKHeader(r *core.BitReader) NetworkRTKHeader {
	var h NetworkRTKHeader
	h.MessageType = int(r.ReadUint(12))
	h.NetworkID = int(r.ReadUint(8))
	h.SubnetworkID = int(r.ReadUint(4))
	h.GPSEpochTimeS = float64(r.ReadUint(23)) * 0.1
	h.MultipleMessageIndicator = r.ReadUint(1) != 0
	h.MasterReferenceStationID = int(r.ReadUint(12))
	h.AuxiliaryReferenceStationID = int(r.ReadUint(12))
	h.NumGPSSats = int(r.ReadUint(4))
	return h
}
