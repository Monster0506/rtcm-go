package rtcm

import "testing"

func TestDecodeMsg1005(t *testing.T) {
	const tol = 1e-4

	tests := []struct {
		name             string
		frame            []byte
		wantStationID    int
		wantITRF         int
		wantGPS          bool
		wantGLONASS      bool
		wantGalileo      bool
		wantRefStation   bool
		wantOscillator   bool
		wantQuarterCycle int
		wantECEFXM       float64
		wantECEFYM       float64
		wantECEFZM       float64
	}{
		{
			name: "RTCM-standard example",
			frame: []byte{
				0xd3, 0x00, 0x13, 0x3e, 0xd7, 0xd3, 0x02, 0x02, 0x98, 0x0e, 0xde, 0xef,
				0x34, 0xb4, 0xbd, 0x62, 0xac, 0x09, 0x41, 0x98, 0x6f, 0x33, 0x36, 0x0b, 0x98,
			},
			wantStationID: 2003,
			wantGPS:       true,
			wantECEFXM:    1114104.5999,
			wantECEFYM:    -4850729.7108,
			wantECEFZM:    3975521.4643,
		},
		{
			name: "second vector",
			frame: []byte{
				0xd3, 0x00, 0x13, 0x3e, 0xd0, 0x00, 0x03, 0x8a, 0x58, 0xd9, 0x49,
				0x3c, 0x87, 0x2f, 0x34, 0x10, 0x9d, 0x07, 0xd6, 0xaf, 0x48, 0x20,
				0x5a, 0xd7, 0xf7,
			},
			wantStationID:  0,
			wantGPS:        true,
			wantGLONASS:    true,
			wantGalileo:    true,
			wantOscillator: true,
			wantECEFXM:     4444030.802800001,
			wantECEFYM:     3085671.2349,
			wantECEFZM:     3366658.256,
		},
		{
			name: "third vector (testUNKNOWNHDR)",
			frame: []byte{
				0xd3, 0x00, 0x13, 0x3e, 0xd0, 0x00, 0x03, 0x84, 0x1a, 0x86, 0x92,
				0xbf, 0xb4, 0x4b, 0x4b, 0xf4, 0xfa, 0xb7, 0xdc, 0x37, 0x62, 0x8a,
				0x33, 0x84, 0x79,
			},
			wantStationID:    0,
			wantGPS:          true,
			wantGLONASS:      true,
			wantGalileo:      true,
			wantOscillator:   true,
			wantQuarterCycle: 2,
			wantECEFXM:       1762489.6191,
			wantECEFYM:       -5027633.8438,
			wantECEFZM:       -3496008.8438000004,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, _, err := ParseFrame(tt.frame)
			if err != nil {
				t.Fatalf("ParseFrame: unexpected error: %v", err)
			}
			msg, err := DecodeMsg1005(payload)
			if err != nil {
				t.Fatalf("DecodeMsg1005: unexpected error: %v", err)
			}
			if msg.MessageType != 1005 {
				t.Fatalf("MessageType = %d, want 1005", msg.MessageType)
			}
			if msg.StationID != tt.wantStationID {
				t.Fatalf("StationID = %d, want %d", msg.StationID, tt.wantStationID)
			}
			if msg.ITRFRealizationYear != tt.wantITRF {
				t.Fatalf("ITRFRealizationYear = %d, want %d", msg.ITRFRealizationYear, tt.wantITRF)
			}
			if msg.GPSIndicator != tt.wantGPS || msg.GLONASSIndicator != tt.wantGLONASS ||
				msg.GalileoIndicator != tt.wantGalileo || msg.ReferenceStationIndicator != tt.wantRefStation {
				t.Fatalf("indicators = %v/%v/%v/%v, want %v/%v/%v/%v",
					msg.GPSIndicator, msg.GLONASSIndicator, msg.GalileoIndicator, msg.ReferenceStationIndicator,
					tt.wantGPS, tt.wantGLONASS, tt.wantGalileo, tt.wantRefStation)
			}
			if msg.OscillatorIndicator != tt.wantOscillator {
				t.Fatalf("OscillatorIndicator = %v, want %v", msg.OscillatorIndicator, tt.wantOscillator)
			}
			if msg.QuarterCycleIndicator != tt.wantQuarterCycle {
				t.Fatalf("QuarterCycleIndicator = %d, want %d", msg.QuarterCycleIndicator, tt.wantQuarterCycle)
			}
			checkFloat(t, "ECEFXM", msg.ECEFXM, tt.wantECEFXM, tol)
			checkFloat(t, "ECEFYM", msg.ECEFYM, tt.wantECEFYM, tol)
			checkFloat(t, "ECEFZM", msg.ECEFZM, tt.wantECEFZM, tol)
		})
	}
}

func checkFloat(t *testing.T, name string, got, want, tol float64) {
	t.Helper()
	diff := got - want
	if diff < 0 {
		diff = -diff
	}
	if diff > tol {
		t.Fatalf("%s = %.10f, want %.10f", name, got, want)
	}
}
