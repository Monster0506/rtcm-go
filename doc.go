// Package rtcm decodes RTCM3 (Radio Technical Commission for Maritime
// Services, version 3) GNSS correction messages.
//
// Decode and PeekMessageType are the entry points. Decode
// dispatches on the message type and returns the decoded struct as an
// any value. Message types are grouped by category into subpackages:
// core, legacyobs, station, networkrtk, ephemeris, transform, msm,
// ssr, and misc.
//
// This package ports test fixtures from pyrtcm
// (https://github.com/semuconsulting/pyrtcm), a Python RTCM3 parser
// (BSD-3-Clause, Copyright (c) 2020, semuadmin). Its test suite uses
// data from real receivers.
//
// The bit-layout reference for each message comes from RTKLIB
// (https://github.com/tomojitakasu/RTKLIB)'s decoder source.
package rtcm
