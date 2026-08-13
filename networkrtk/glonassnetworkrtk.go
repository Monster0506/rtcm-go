package networkrtk

import "github.com/Monster0506/rtcm-go/core"

type GLONASSNetworkRTKHeader struct {
	MessageType                 int
	NetworkID                   int
	SubnetworkID                int
	GLONASSEpochTimeS           float64
	MultipleMessageIndicator    bool
	MasterReferenceStationID    int
	AuxiliaryReferenceStationID int
	NumGLONASSDataEntries       int
}

func decodeGLONASSNetworkRTKHeader(r *core.BitReader) GLONASSNetworkRTKHeader {
	var h GLONASSNetworkRTKHeader
	h.MessageType = int(r.ReadUint(12))
	h.NetworkID = int(r.ReadUint(8))
	h.SubnetworkID = int(r.ReadUint(4))
	h.GLONASSEpochTimeS = float64(r.ReadUint(20)) * 0.1
	h.MultipleMessageIndicator = r.ReadUint(1) != 0
	h.MasterReferenceStationID = int(r.ReadUint(12))
	h.AuxiliaryReferenceStationID = int(r.ReadUint(12))
	h.NumGLONASSDataEntries = int(r.ReadUint(4))
	return h
}
