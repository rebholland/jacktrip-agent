// Copyright 2020 20hz, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx/types"
)

// DeviceConfig defines configuration for a particular device
type DeviceConfig struct {
	// DevicePort is the bindport used by the device
	DevicePort int `json:"devicePort" db:"port"`

	// Reverb level from 0 to 100, only used if compressor is enabled
	Reverb int `json:"reverb" db:"reverb"`

	// If true, a limiter will be applied to input and output volume
	Limiter types.BitBool `json:"limiter" db:"limiter"`

	// If true, a compressor will be applied to device output
	Compressor types.BitBool `json:"compressor" db:"compressor"`

	// connection quality
	// 0: low quality Jamulus (low)
	// 1: high quality Jamulus (medium)
	// 2: JackTrip (high)
	Quality int `json:"quality" db:"quality"`
}

// ALSAConfig defines configuration for a device's ALSA sound card
type ALSAConfig struct {
	// If true, apply volume boost for audio capture or input
	CaptureBoost types.BitBool `json:"captureBoost" db:"capture_boost"`

	// If true, apply volume boost for audio playback or output
	PlaybackBoost types.BitBool `json:"playbackBoost" db:"playback_boost"`

	// Volume level percent (0-100) for audio capture or input
	CaptureVolume int `json:"captureVolume" db:"capture_volume"`

	// Volume level percent (0-100) for audio playback or output
	PlaybackVolume int `json:"playbackVolume" db:"playback_volume"`
}

// PingStats is used to represent ping statistics
type PingStats struct {
	// PacketsRecv is the number of packets received.
	PacketsRecv int `json:"pkts_recv" db:"pkts_recv"`

	// PacketsSent is the number of packets sent.
	PacketsSent int `json:"pkts_sent" db:"pkts_sent"`

	// MinRtt is the minimum round-trip time sent via this pinger.
	MinRtt time.Duration `json:"min_rtt" db:"min_rtt"`

	// MaxRtt is the maximum round-trip time sent via this pinger.
	MaxRtt time.Duration `json:"max_rtt" db:"max_rtt"`

	// AvgRtt is the average round-trip time sent via this pinger.
	AvgRtt time.Duration `json:"avg_rtt" db:"avg_rtt"`

	// StdDevRtt is the standard deviation of the round-trip times sent via
	// this pinger.
	StdDevRtt time.Duration `json:"stddev_rtt" db:"stddev_rtt"`

	// timestamp when the device stats were last updated
	StatsUpdatedAt time.Time `json:"stats_updated_at" db:"stats_updated_at"`
}

// AgentConfig defines active configuration for an agent
type AgentConfig struct {
	DeviceConfig
	ALSAConfig
	ServerConfig

	// frames per period
	Period int `json:"period" db:"period"`

	// size of jitter queue buffer
	QueueBuffer int `json:"queueBuffer" db:"queue_buffer"`
}

// AgentCredentials defines authentication credentials for an agent
type AgentCredentials struct {
	// API key prefix
	APIPrefix string `json:"apiPrefix"`

	// API key secret value (used to generate APIHash)
	APISecret string `json:"apiSecret"`
}

// GetAPIHash returns hashed value for a given api secret
func GetAPIHash(apiSecret string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(apiSecret)))
}

// AgentPing is used to receive ping from a device
type AgentPing struct {
	AgentCredentials
	PingStats

	// Cloud identifier for server (used when running on cloud audio server)
	CloudID string `json:"cloudId"`

	// MAC address for ethernet device (used when running on raspberry pi device)
	MAC string `json:"mac"`

	// Current image version for the device
	Version string `json:"version"`

	// Type of sound device ("snd_rpi_hifiberry_dacplusadcpro")
	Type string `json:"type" db:"type"`
}
