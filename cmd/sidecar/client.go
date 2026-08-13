package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	rtcm "github.com/Monster0506/rtcm-go"
)

const maxFrameBuffer = 64 * 1024

func runMountpoint(ctx context.Context, casterAddr, mountpoint string, stats *StatsFile) {
	backoff := 1 * time.Second
	const maxBackoff = 30 * time.Second

	for {
		if ctx.Err() != nil {
			return
		}
		err := streamMountpoint(ctx, casterAddr, mountpoint, stats)
		if ctx.Err() != nil {
			return
		}
		updateStats(stats, mountpoint, func(s *MountpointStats) {
			s.Connected = false
			if err != nil {
				s.LastError = err.Error()
			}
		})
		log.Printf("[%s] disconnected: %v (reconnecting in %s)", mountpoint, err, backoff)
		select {
		case <-time.After(backoff):
		case <-ctx.Done():
			return
		}
		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

func updateStats(stats *StatsFile, mountpoint string, fn func(*MountpointStats)) {
	if err := stats.Update(mountpoint, fn); err != nil {
		log.Printf("[%s] writing stats file: %v", mountpoint, err)
	}
}

func streamMountpoint(ctx context.Context, casterAddr, mountpoint string, stats *StatsFile) error {
	conn, err := net.DialTimeout("tcp", casterAddr, 10*time.Second)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}
	defer func() { _ = conn.Close() }()

	go func() {
		<-ctx.Done()
		_ = conn.Close()
	}()

	if err := conn.SetDeadline(time.Now().Add(10 * time.Second)); err != nil {
		return fmt.Errorf("set deadline: %w", err)
	}
	req := fmt.Sprintf("GET /%s HTTP/1.1\r\n\r\n", mountpoint)
	if _, err := conn.Write([]byte(req)); err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	reader := bufio.NewReader(conn)
	status, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("read status line: %w", err)
	}
	status = strings.TrimRight(status, "\r\n")
	if !strings.Contains(status, "200") {
		return fmt.Errorf("caster rejected mountpoint %q: %q", mountpoint, status)
	}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("read headers: %w", err)
		}
		if strings.TrimRight(line, "\r\n") == "" {
			break
		}
	}

	updateStats(stats, mountpoint, func(s *MountpointStats) {
		s.Connected = true
		s.LastError = ""
	})
	if err := conn.SetDeadline(time.Time{}); err != nil {
		return fmt.Errorf("clear deadline: %w", err)
	}

	buf := make([]byte, 0, 4096)
	tmp := make([]byte, 4096)
	for {
		n, err := reader.Read(tmp)
		if n > 0 {
			buf = append(buf, tmp[:n]...)
			buf = drainFrames(buf, mountpoint, stats)
			if len(buf) > maxFrameBuffer {
				log.Printf("[%s] %d bytes without a resolvable frame, discarding", mountpoint, len(buf))
				buf = buf[:0]
			}
		}
		if err != nil {
			return fmt.Errorf("read stream: %w", err)
		}
	}
}

func drainFrames(buf []byte, mountpoint string, stats *StatsFile) []byte {
	for {
		payload, consumed, err := rtcm.ParseFrame(buf)
		if err != nil {
			if strings.Contains(err.Error(), "invalid payload size 0 bytes") {
				if len(buf) > 0 {
					buf = buf[1:]
					continue
				}
			}
			if consumed > 0 {
				buf = buf[consumed:]
				continue
			}
			return buf
		}
		handleMessage(payload, mountpoint, stats)
		buf = buf[consumed:]
	}
}

func handleMessage(payload []byte, mountpoint string, stats *StatsFile) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[%s] recovered from panic decoding a message: %v", mountpoint, r)
		}
	}()
	if len(payload) < 2 {
		return
	}
	msgType := int(rtcm.NewBitReader(payload).ReadUint(12))

	switch msgType {
	case 1008:
		m, err := rtcm.DecodeMsg1008(payload)
		if err != nil {
			return
		}
		updateStats(stats, mountpoint, func(s *MountpointStats) {
			s.LastMessageType = msgType
			s.AntennaDescriptor = m.AntennaDescriptor
			s.AntennaSerial = m.AntennaSerial
			s.SetupID = m.SetupID
		})
	case 1033:
		m, err := rtcm.DecodeMsg1033(payload)
		if err != nil {
			return
		}
		updateStats(stats, mountpoint, func(s *MountpointStats) {
			s.LastMessageType = msgType
			s.AntennaDescriptor = m.AntennaDescriptor
			s.AntennaSerial = m.AntennaSerial
			s.SetupID = m.SetupID
			s.ReceiverType = m.ReceiverType
			s.FirmwareVersion = m.FirmwareVersion
			s.ReceiverSerial = m.ReceiverSerial
		})
	default:
		if _, ok := rtcm.ConstellationForType(msgType); !ok {
			return
		}
		m, err := rtcm.DecodeMSM(payload)
		if err != nil {
			return
		}
		updateStats(stats, mountpoint, func(s *MountpointStats) {
			s.LastMessageType = msgType
			s.Constellations[m.Constellation] = m.SatelliteCount
		})
	}
}
