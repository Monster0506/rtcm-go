package rtcm

import "testing"

func TestMessageTypeName(t *testing.T) {
	cases := []struct {
		msgType int
		want    string
	}{
		{1005, "Stationary RTK Reference Station ARP"},
		{1077, "GPS MSM7"},
		{1087, "GLONASS MSM7"},
		{1124, "BeiDou MSM4"},
		{1230, "GLONASS Code-Phase Biases"},
		{9999, "RTCM 9999"},
	}
	for _, tt := range cases {
		if got := MessageTypeName(tt.msgType); got != tt.want {
			t.Errorf("MessageTypeName(%d) = %q, want %q", tt.msgType, got, tt.want)
		}
	}
}
