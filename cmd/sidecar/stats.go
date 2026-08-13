package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type MountpointStats struct {
	Connected         bool           `json:"connected"`
	Constellations    map[string]int `json:"constellations,omitempty"`
	SatelliteCount    int            `json:"satellite_count"`
	AntennaDescriptor string         `json:"antenna_descriptor,omitempty"`
	AntennaSerial     string         `json:"antenna_serial,omitempty"`
	SetupID           int            `json:"setup_id,omitempty"`
	ReceiverType      string         `json:"receiver_type,omitempty"`
	FirmwareVersion   string         `json:"firmware_version,omitempty"`
	ReceiverSerial    string         `json:"receiver_serial,omitempty"`
	LastMessageType   int            `json:"last_message_type,omitempty"`
	LastUpdated       time.Time      `json:"last_updated"`
	LastError         string         `json:"last_error,omitempty"`
	LastMessages      map[int]any    `json:"last_messages,omitempty"`
}

func (s *MountpointStats) recomputeSatelliteCount() {
	total := 0
	for _, n := range s.Constellations {
		total += n
	}
	s.SatelliteCount = total
}

type StatsFile struct {
	mu          sync.Mutex
	path        string
	mountpoints map[string]*MountpointStats
}

type statsFileDocument struct {
	GeneratedAt time.Time                   `json:"generated_at"`
	Mountpoints map[string]*MountpointStats `json:"mountpoints"`
}

func NewStatsFile(path string) *StatsFile {
	return &StatsFile{
		path:        path,
		mountpoints: make(map[string]*MountpointStats),
	}
}

func (f *StatsFile) Update(mountpoint string, fn func(*MountpointStats)) error {
	f.mu.Lock()
	s, ok := f.mountpoints[mountpoint]
	if !ok {
		s = &MountpointStats{Constellations: make(map[string]int), LastMessages: make(map[int]any)}
		f.mountpoints[mountpoint] = s
	}
	fn(s)
	s.recomputeSatelliteCount()
	s.LastUpdated = time.Now().UTC()
	err := f.writeLocked()
	f.mu.Unlock()
	return err
}

func (f *StatsFile) writeLocked() error {
	doc := statsFileDocument{
		GeneratedAt: time.Now().UTC(),
		Mountpoints: f.mountpoints,
	}
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal stats: %w", err)
	}
	dir := filepath.Dir(f.path)
	tmp, err := os.CreateTemp(dir, ".mountpoints-*.tmp")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpName := tmp.Name()
	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		_ = os.Remove(tmpName)
		return fmt.Errorf("write temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmpName)
		return fmt.Errorf("close temp file: %w", err)
	}
	if err := os.Rename(tmpName, f.path); err != nil {
		_ = os.Remove(tmpName)
		return fmt.Errorf("rename temp file into place: %w", err)
	}
	return nil
}
