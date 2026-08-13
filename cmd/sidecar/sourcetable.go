package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func fetchSourcetable(casterAddr string) ([]string, error) {
	conn, err := net.DialTimeout("tcp", casterAddr, 10*time.Second)
	if err != nil {
		return nil, fmt.Errorf("dial caster: %w", err)
	}
	defer func() { _ = conn.Close() }()

	if err := conn.SetDeadline(time.Now().Add(15 * time.Second)); err != nil {
		return nil, fmt.Errorf("set deadline: %w", err)
	}
	if _, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n")); err != nil {
		return nil, fmt.Errorf("send sourcetable request: %w", err)
	}

	var mounts []string
	scanner := bufio.NewScanner(conn)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	sawHeader := false
	for scanner.Scan() {
		line := strings.TrimRight(scanner.Text(), "\r")
		if line == "ENDSOURCETABLE" {
			break
		}
		if !sawHeader {
			sawHeader = true
			continue
		}
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "STR;") {
			fields := strings.Split(line, ";")
			if len(fields) > 1 && fields[1] != "" {
				mounts = append(mounts, fields[1])
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read sourcetable: %w", err)
	}
	return mounts, nil
}
