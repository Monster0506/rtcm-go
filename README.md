# rtcm-go

rtcm-go is a Go library. It decodes RTCM3 (Radio Technical Commission for
Maritime Services, version 3) GNSS correction messages. It has no external
dependencies.

## Install

```
go get github.com/Monster0506/rtcm-go
```

## Attribution

The test fixtures come from [pyrtcm](https://github.com/semuconsulting/pyrtcm).
pyrtcm is a Python RTCM3 parser (BSD-3-Clause, Copyright (c) 2020,
semuadmin). Its test suite uses data from real receivers.

The bit-layout reference for each message comes from
[RTKLIB](https://github.com/tomojitakasu/RTKLIB)'s decoder source.

## License

This project uses the BSD-3-Clause license. See [LICENSE](LICENSE).
