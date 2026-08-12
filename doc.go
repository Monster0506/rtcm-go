// Package rtcm decodes RTCM3 (Radio Technical Commission for Maritime
// Services, version 3) GNSS correction messages.
//
// Test fixtures are ported from pyrtcm (https://github.com/semuconsulting/pyrtcm),
// a Python RTCM3 parser (BSD-3-Clause, Copyright (c) 2020, semuadmin). Its
// test suite uses data from real receivers.
//
// The bit-layout reference for each message comes from RTKLIB
// (https://github.com/tomojitakasu/RTKLIB)'s decoder source.
package rtcm
