package transform

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1023(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x49, 0x3f, 0xf0, 0x39, 0xe8, 0x48, 0x10, 0xbd, 0xc0, 0x7d,
		0x05, 0xdc, 0x19, 0x62, 0x2e, 0xe0, 0x32, 0x73, 0x81, 0x46, 0x5e, 0x68,
		0x2c, 0xcd, 0xcc, 0x06, 0x19, 0xf9, 0x60, 0xd3, 0x47, 0x28, 0x1c, 0x69,
		0xe4, 0x83, 0xcd, 0x5c, 0x80, 0x81, 0xaf, 0x8e, 0x11, 0x36, 0x71, 0x82,
		0x46, 0xde, 0x28, 0x4c, 0xdd, 0xc4, 0x0a, 0x1b, 0xf8, 0x61, 0x53, 0x87,
		0x08, 0x2c, 0x71, 0xe0, 0x85, 0xce, 0x5c, 0x00, 0xc1, 0xcf, 0x7e, 0x19,
		0x6b, 0xb9, 0x78, 0x40, 0x25, 0xd5, 0x1c,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1023(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1023: unexpected error: %v", err)
	}
	if m.MessageType != 1023 || m.SystemIdentificationNumber != 3 ||
		m.HorizontalShiftIndicator != true || m.VerticalShiftIndicator != false ||
		m.HorizontalInterpolationMethod != 1 ||
		m.VerticalInterpolationMethod != 2 ||
		m.HorizontalGridQualityIndicator != 5 ||
		m.VerticalGridQualityIndicator != 6 || m.MJDNumber != 58849 {
		t.Fatalf("mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "LatitudeOfOriginArcSec", m.LatitudeOfOriginArcSec, 500000.0, 1e-9)
	testutil.CheckFloat(t, "LongitudeOfOriginArcSec", m.LongitudeOfOriginArcSec, -1000000.0, 1e-9)
	testutil.CheckFloat(t, "NSExtensionArcSec", m.NSExtensionArcSec, 2000.0, 1e-9)
	testutil.CheckFloat(t, "EWExtensionArcSec", m.EWExtensionArcSec, 1500.0, 1e-9)
	testutil.CheckFloat(t, "MeanLatitudeOffsetDeg", m.MeanLatitudeOffsetDeg, 0.05, 1e-9)
	testutil.CheckFloat(t, "MeanLongitudeOffsetDeg", m.MeanLongitudeOffsetDeg, -0.06, 1e-9)
	testutil.CheckFloat(t, "MeanHeightOffsetM", m.MeanHeightOffsetM, 120.0, 1e-9)
	testutil.CheckFloat(t, "GridPoints[0].LatitudeResidualDeg", m.GridPoints[0].LatitudeResidualDeg, 0.003, 1e-9)
	testutil.CheckFloat(t, "GridPoints[0].LongitudeResidualDeg", m.GridPoints[0].LongitudeResidualDeg, -0.0015, 1e-9)
	testutil.CheckFloat(t, "GridPoints[0].HeightResidualM", m.GridPoints[0].HeightResidualM, 0.01, 1e-9)
	testutil.CheckFloat(t, "GridPoints[1].LatitudeResidualDeg", m.GridPoints[1].LatitudeResidualDeg, 0.00303, 1e-9)
	testutil.CheckFloat(t, "GridPoints[1].LongitudeResidualDeg", m.GridPoints[1].LongitudeResidualDeg, -0.0015300000000000001, 1e-9)
	testutil.CheckFloat(t, "GridPoints[1].HeightResidualM", m.GridPoints[1].HeightResidualM, 0.011, 1e-9)
	testutil.CheckFloat(t, "GridPoints[2].LatitudeResidualDeg", m.GridPoints[2].LatitudeResidualDeg, 0.0030600000000000002, 1e-9)
	testutil.CheckFloat(t, "GridPoints[2].LongitudeResidualDeg", m.GridPoints[2].LongitudeResidualDeg, -0.00156, 1e-9)
	testutil.CheckFloat(t, "GridPoints[2].HeightResidualM", m.GridPoints[2].HeightResidualM, 0.012, 1e-9)
	testutil.CheckFloat(t, "GridPoints[3].LatitudeResidualDeg", m.GridPoints[3].LatitudeResidualDeg, 0.00309, 1e-9)
	testutil.CheckFloat(t, "GridPoints[3].LongitudeResidualDeg", m.GridPoints[3].LongitudeResidualDeg, -0.00159, 1e-9)
	testutil.CheckFloat(t, "GridPoints[3].HeightResidualM", m.GridPoints[3].HeightResidualM, 0.013000000000000001, 1e-9)
	testutil.CheckFloat(t, "GridPoints[4].LatitudeResidualDeg", m.GridPoints[4].LatitudeResidualDeg, 0.00312, 1e-9)
	testutil.CheckFloat(t, "GridPoints[4].LongitudeResidualDeg", m.GridPoints[4].LongitudeResidualDeg, -0.0016200000000000001, 1e-9)
	testutil.CheckFloat(t, "GridPoints[4].HeightResidualM", m.GridPoints[4].HeightResidualM, 0.014, 1e-9)
	testutil.CheckFloat(t, "GridPoints[5].LatitudeResidualDeg", m.GridPoints[5].LatitudeResidualDeg, 0.00315, 1e-9)
	testutil.CheckFloat(t, "GridPoints[5].LongitudeResidualDeg", m.GridPoints[5].LongitudeResidualDeg, -0.00165, 1e-9)
	testutil.CheckFloat(t, "GridPoints[5].HeightResidualM", m.GridPoints[5].HeightResidualM, 0.015, 1e-9)
	testutil.CheckFloat(t, "GridPoints[6].LatitudeResidualDeg", m.GridPoints[6].LatitudeResidualDeg, 0.00318, 1e-9)
	testutil.CheckFloat(t, "GridPoints[6].LongitudeResidualDeg", m.GridPoints[6].LongitudeResidualDeg, -0.00168, 1e-9)
	testutil.CheckFloat(t, "GridPoints[6].HeightResidualM", m.GridPoints[6].HeightResidualM, 0.016, 1e-9)
	testutil.CheckFloat(t, "GridPoints[7].LatitudeResidualDeg", m.GridPoints[7].LatitudeResidualDeg, 0.00321, 1e-9)
	testutil.CheckFloat(t, "GridPoints[7].LongitudeResidualDeg", m.GridPoints[7].LongitudeResidualDeg, -0.0017100000000000001, 1e-9)
	testutil.CheckFloat(t, "GridPoints[7].HeightResidualM", m.GridPoints[7].HeightResidualM, 0.017, 1e-9)
	testutil.CheckFloat(t, "GridPoints[8].LatitudeResidualDeg", m.GridPoints[8].LatitudeResidualDeg, 0.0032400000000000003, 1e-9)
	testutil.CheckFloat(t, "GridPoints[8].LongitudeResidualDeg", m.GridPoints[8].LongitudeResidualDeg, -0.00174, 1e-9)
	testutil.CheckFloat(t, "GridPoints[8].HeightResidualM", m.GridPoints[8].HeightResidualM, 0.018000000000000002, 1e-9)
	testutil.CheckFloat(t, "GridPoints[9].LatitudeResidualDeg", m.GridPoints[9].LatitudeResidualDeg, 0.00327, 1e-9)
	testutil.CheckFloat(t, "GridPoints[9].LongitudeResidualDeg", m.GridPoints[9].LongitudeResidualDeg, -0.00177, 1e-9)
	testutil.CheckFloat(t, "GridPoints[9].HeightResidualM", m.GridPoints[9].HeightResidualM, 0.019, 1e-9)
	testutil.CheckFloat(t, "GridPoints[10].LatitudeResidualDeg", m.GridPoints[10].LatitudeResidualDeg, 0.0033, 1e-9)
	testutil.CheckFloat(t, "GridPoints[10].LongitudeResidualDeg", m.GridPoints[10].LongitudeResidualDeg, -0.0018, 1e-9)
	testutil.CheckFloat(t, "GridPoints[10].HeightResidualM", m.GridPoints[10].HeightResidualM, 0.02, 1e-9)
	testutil.CheckFloat(t, "GridPoints[11].LatitudeResidualDeg", m.GridPoints[11].LatitudeResidualDeg, 0.00333, 1e-9)
	testutil.CheckFloat(t, "GridPoints[11].LongitudeResidualDeg", m.GridPoints[11].LongitudeResidualDeg, -0.00183, 1e-9)
	testutil.CheckFloat(t, "GridPoints[11].HeightResidualM", m.GridPoints[11].HeightResidualM, 0.021, 1e-9)
	testutil.CheckFloat(t, "GridPoints[12].LatitudeResidualDeg", m.GridPoints[12].LatitudeResidualDeg, 0.00336, 1e-9)
	testutil.CheckFloat(t, "GridPoints[12].LongitudeResidualDeg", m.GridPoints[12].LongitudeResidualDeg, -0.00186, 1e-9)
	testutil.CheckFloat(t, "GridPoints[12].HeightResidualM", m.GridPoints[12].HeightResidualM, 0.022, 1e-9)
	testutil.CheckFloat(t, "GridPoints[13].LatitudeResidualDeg", m.GridPoints[13].LatitudeResidualDeg, 0.0033900000000000002, 1e-9)
	testutil.CheckFloat(t, "GridPoints[13].LongitudeResidualDeg", m.GridPoints[13].LongitudeResidualDeg, -0.00189, 1e-9)
	testutil.CheckFloat(t, "GridPoints[13].HeightResidualM", m.GridPoints[13].HeightResidualM, 0.023, 1e-9)
	testutil.CheckFloat(t, "GridPoints[14].LatitudeResidualDeg", m.GridPoints[14].LatitudeResidualDeg, 0.0034200000000000003, 1e-9)
	testutil.CheckFloat(t, "GridPoints[14].LongitudeResidualDeg", m.GridPoints[14].LongitudeResidualDeg, -0.00192, 1e-9)
	testutil.CheckFloat(t, "GridPoints[14].HeightResidualM", m.GridPoints[14].HeightResidualM, 0.024, 1e-9)
	testutil.CheckFloat(t, "GridPoints[15].LatitudeResidualDeg", m.GridPoints[15].LatitudeResidualDeg, 0.00345, 1e-9)
	testutil.CheckFloat(t, "GridPoints[15].LongitudeResidualDeg", m.GridPoints[15].LongitudeResidualDeg, -0.0019500000000000001, 1e-9)
	testutil.CheckFloat(t, "GridPoints[15].HeightResidualM", m.GridPoints[15].HeightResidualM, 0.025, 1e-9)
}
