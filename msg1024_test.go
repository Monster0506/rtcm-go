package rtcm

import "testing"



func TestDecodeMsg1024(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x4a, 0x40, 0x00, 0x3b, 0x67, 0x69, 0x80, 0x3d, 0x09, 0x00,
		0x7d, 0x05, 0xdc, 0x06, 0x5e, 0x22, 0xee, 0x03, 0x27, 0x38, 0x14, 0x65,
		0xe6, 0x82, 0xcc, 0xdc, 0xc0, 0x61, 0x9f, 0x96, 0x0d, 0x34, 0x72, 0x81,
		0xc6, 0x9e, 0x48, 0x3c, 0xd5, 0xc8, 0x08, 0x1a, 0xf8, 0xe1, 0x13, 0x67,
		0x18, 0x24, 0x6d, 0xe2, 0x84, 0xcd, 0xdc, 0x40, 0xa1, 0xbf, 0x86, 0x15,
		0x38, 0x70, 0x82, 0xc7, 0x1e, 0x08, 0x5c, 0xe5, 0xc0, 0x0c, 0x1c, 0xf7,
		0xe1, 0x96, 0xbb, 0x97, 0x84, 0xbe, 0x9a, 0xad,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1024(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1024: unexpected error: %v", err)
	}
	if m.MessageType != 1024 || m.SystemIdentificationNumber != 3 ||
		m.HorizontalShiftIndicator != true || m.VerticalShiftIndicator != false ||
		m.HorizontalInterpolationMethod != 1 ||
		m.VerticalInterpolationMethod != 2 ||
		m.HorizontalGridQualityIndicator != 5 ||
		m.VerticalGridQualityIndicator != 6 || m.MJDNumber != 58849 {
		t.Fatalf("mismatch: %+v", m)
	}
	checkFloat(t, "NorthingOfOriginM", m.NorthingOfOriginM, -50000000, 1e-9)
	checkFloat(t, "EastingOfOriginM", m.EastingOfOriginM, 80000000, 1e-9)
	checkFloat(t, "NSExtensionM", m.NSExtensionM, 40000, 1e-9)
	checkFloat(t, "EWExtensionM", m.EWExtensionM, 30000, 1e-9)
	checkFloat(t, "MeanNorthingOffsetM", m.MeanNorthingOffsetM, 0.5, 1e-9)
	checkFloat(t, "MeanEastingOffsetM", m.MeanEastingOffsetM, -0.6, 1e-9)
	checkFloat(t, "MeanHeightOffsetM", m.MeanHeightOffsetM, 120.0, 1e-9)
	checkFloat(t, "GridPoints[0].NorthingResidualM", m.GridPoints[0].NorthingResidualM, 0.1, 1e-9)
	checkFloat(t, "GridPoints[0].EastingResidualM", m.GridPoints[0].EastingResidualM, -0.05, 1e-9)
	checkFloat(t, "GridPoints[0].HeightResidualM", m.GridPoints[0].HeightResidualM, 0.01, 1e-9)
	checkFloat(t, "GridPoints[1].NorthingResidualM", m.GridPoints[1].NorthingResidualM, 0.101, 1e-9)
	checkFloat(t, "GridPoints[1].EastingResidualM", m.GridPoints[1].EastingResidualM, -0.051000000000000004, 1e-9)
	checkFloat(t, "GridPoints[1].HeightResidualM", m.GridPoints[1].HeightResidualM, 0.011, 1e-9)
	checkFloat(t, "GridPoints[2].NorthingResidualM", m.GridPoints[2].NorthingResidualM, 0.10200000000000001, 1e-9)
	checkFloat(t, "GridPoints[2].EastingResidualM", m.GridPoints[2].EastingResidualM, -0.052000000000000005, 1e-9)
	checkFloat(t, "GridPoints[2].HeightResidualM", m.GridPoints[2].HeightResidualM, 0.012, 1e-9)
	checkFloat(t, "GridPoints[3].NorthingResidualM", m.GridPoints[3].NorthingResidualM, 0.10300000000000001, 1e-9)
	checkFloat(t, "GridPoints[3].EastingResidualM", m.GridPoints[3].EastingResidualM, -0.053, 1e-9)
	checkFloat(t, "GridPoints[3].HeightResidualM", m.GridPoints[3].HeightResidualM, 0.013000000000000001, 1e-9)
	checkFloat(t, "GridPoints[4].NorthingResidualM", m.GridPoints[4].NorthingResidualM, 0.10400000000000001, 1e-9)
	checkFloat(t, "GridPoints[4].EastingResidualM", m.GridPoints[4].EastingResidualM, -0.054, 1e-9)
	checkFloat(t, "GridPoints[4].HeightResidualM", m.GridPoints[4].HeightResidualM, 0.014, 1e-9)
	checkFloat(t, "GridPoints[5].NorthingResidualM", m.GridPoints[5].NorthingResidualM, 0.105, 1e-9)
	checkFloat(t, "GridPoints[5].EastingResidualM", m.GridPoints[5].EastingResidualM, -0.055, 1e-9)
	checkFloat(t, "GridPoints[5].HeightResidualM", m.GridPoints[5].HeightResidualM, 0.015, 1e-9)
	checkFloat(t, "GridPoints[6].NorthingResidualM", m.GridPoints[6].NorthingResidualM, 0.106, 1e-9)
	checkFloat(t, "GridPoints[6].EastingResidualM", m.GridPoints[6].EastingResidualM, -0.056, 1e-9)
	checkFloat(t, "GridPoints[6].HeightResidualM", m.GridPoints[6].HeightResidualM, 0.016, 1e-9)
	checkFloat(t, "GridPoints[7].NorthingResidualM", m.GridPoints[7].NorthingResidualM, 0.107, 1e-9)
	checkFloat(t, "GridPoints[7].EastingResidualM", m.GridPoints[7].EastingResidualM, -0.057, 1e-9)
	checkFloat(t, "GridPoints[7].HeightResidualM", m.GridPoints[7].HeightResidualM, 0.017, 1e-9)
	checkFloat(t, "GridPoints[8].NorthingResidualM", m.GridPoints[8].NorthingResidualM, 0.108, 1e-9)
	checkFloat(t, "GridPoints[8].EastingResidualM", m.GridPoints[8].EastingResidualM, -0.058, 1e-9)
	checkFloat(t, "GridPoints[8].HeightResidualM", m.GridPoints[8].HeightResidualM, 0.018000000000000002, 1e-9)
	checkFloat(t, "GridPoints[9].NorthingResidualM", m.GridPoints[9].NorthingResidualM, 0.109, 1e-9)
	checkFloat(t, "GridPoints[9].EastingResidualM", m.GridPoints[9].EastingResidualM, -0.059000000000000004, 1e-9)
	checkFloat(t, "GridPoints[9].HeightResidualM", m.GridPoints[9].HeightResidualM, 0.019, 1e-9)
	checkFloat(t, "GridPoints[10].NorthingResidualM", m.GridPoints[10].NorthingResidualM, 0.11, 1e-9)
	checkFloat(t, "GridPoints[10].EastingResidualM", m.GridPoints[10].EastingResidualM, -0.06, 1e-9)
	checkFloat(t, "GridPoints[10].HeightResidualM", m.GridPoints[10].HeightResidualM, 0.02, 1e-9)
	checkFloat(t, "GridPoints[11].NorthingResidualM", m.GridPoints[11].NorthingResidualM, 0.111, 1e-9)
	checkFloat(t, "GridPoints[11].EastingResidualM", m.GridPoints[11].EastingResidualM, -0.061, 1e-9)
	checkFloat(t, "GridPoints[11].HeightResidualM", m.GridPoints[11].HeightResidualM, 0.021, 1e-9)
	checkFloat(t, "GridPoints[12].NorthingResidualM", m.GridPoints[12].NorthingResidualM, 0.112, 1e-9)
	checkFloat(t, "GridPoints[12].EastingResidualM", m.GridPoints[12].EastingResidualM, -0.062, 1e-9)
	checkFloat(t, "GridPoints[12].HeightResidualM", m.GridPoints[12].HeightResidualM, 0.022, 1e-9)
	checkFloat(t, "GridPoints[13].NorthingResidualM", m.GridPoints[13].NorthingResidualM, 0.113, 1e-9)
	checkFloat(t, "GridPoints[13].EastingResidualM", m.GridPoints[13].EastingResidualM, -0.063, 1e-9)
	checkFloat(t, "GridPoints[13].HeightResidualM", m.GridPoints[13].HeightResidualM, 0.023, 1e-9)
	checkFloat(t, "GridPoints[14].NorthingResidualM", m.GridPoints[14].NorthingResidualM, 0.114, 1e-9)
	checkFloat(t, "GridPoints[14].EastingResidualM", m.GridPoints[14].EastingResidualM, -0.064, 1e-9)
	checkFloat(t, "GridPoints[14].HeightResidualM", m.GridPoints[14].HeightResidualM, 0.024, 1e-9)
	checkFloat(t, "GridPoints[15].NorthingResidualM", m.GridPoints[15].NorthingResidualM, 0.115, 1e-9)
	checkFloat(t, "GridPoints[15].EastingResidualM", m.GridPoints[15].EastingResidualM, -0.065, 1e-9)
	checkFloat(t, "GridPoints[15].HeightResidualM", m.GridPoints[15].HeightResidualM, 0.025, 1e-9)
}
