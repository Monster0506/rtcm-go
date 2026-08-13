package rtcm

import (
	"fmt"

	"github.com/Monster0506/rtcm-go/msm"
)

var messageTypeNames = map[int]string{
	1001: "GPS L1 Observations",
	1002: "GPS L1 Extended Observations",
	1003: "GPS L1/L2 Observations",
	1004: "GPS L1/L2 Extended Observations",
	1005: "Stationary RTK Reference Station ARP",
	1006: "Stationary RTK Reference Station ARP with Antenna Height",
	1007: "Antenna Descriptor",
	1008: "Antenna Descriptor and Serial Number",
	1009: "GLONASS L1 Observations",
	1010: "GLONASS L1 Extended Observations",
	1011: "GLONASS L1/L2 Observations",
	1012: "GLONASS L1/L2 Extended Observations",
	1013: "System Parameters",
	1014: "Network Auxiliary Station Data",
	1015: "GPS Ionospheric Correction Differences",
	1016: "GPS Geometric Correction Differences",
	1017: "GPS Combined Geometric and Ionospheric Correction Differences",
	1019: "GPS Ephemeris",
	1020: "GLONASS Ephemeris",
	1021: "Helmert/Abridged Molodenski Transformation Parameters",
	1022: "Molodenski-Badekas Transformation Parameters",
	1023: "Residuals, Ellipsoidal Grid Representation",
	1024: "Residuals, Plane Grid Representation",
	1025: "Projection Parameters (excluding LCC2SP, OM)",
	1026: "Projection Parameters, Lambert Conformal Conic",
	1027: "Projection Parameters, Oblique Mercator",
	1029: "Unicode Text String",
	1030: "GPS Network RTK Residuals",
	1031: "GLONASS Network RTK Residuals",
	1032: "Physical Reference Station Position",
	1033: "Receiver and Antenna Descriptors",
	1034: "GPS Network FKP Gradient",
	1035: "GLONASS Network FKP Gradient",
	1037: "GLONASS Ionospheric Correction Differences",
	1038: "GLONASS Geometric Correction Differences",
	1039: "GLONASS Combined Geometric and Ionospheric Correction Differences",
	1042: "BeiDou Ephemeris",
	1044: "QZSS Ephemeris",
	1045: "Galileo F/NAV Ephemeris",
	1046: "Galileo I/NAV Ephemeris",
	1057: "GPS SSR Orbit Correction",
	1058: "GPS SSR Clock Correction",
	1059: "GPS SSR Code Bias",
	1060: "GPS SSR Combined Orbit and Clock Correction",
	1061: "GPS SSR URA",
	1062: "GPS SSR High Rate Clock Correction",
	1063: "GLONASS SSR Orbit Correction",
	1064: "GLONASS SSR Clock Correction",
	1065: "GLONASS SSR Code Bias",
	1066: "GLONASS SSR Combined Orbit and Clock Correction",
	1067: "GLONASS SSR URA",
	1068: "GLONASS SSR High Rate Clock Correction",
	1230: "GLONASS Code-Phase Biases",
}

func MessageTypeName(msgType int) string {
	if name, ok := messageTypeNames[msgType]; ok {
		return name
	}
	if constellation, ok := msm.ConstellationForType(msgType); ok {
		return fmt.Sprintf("%s MSM%d", constellation, msgType%10)
	}
	return fmt.Sprintf("RTCM %d", msgType)
}
