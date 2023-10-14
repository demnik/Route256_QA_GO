// Package model describes the properties of the device
package model

import (
	"time"
)

// Device structure with device properties
type Device struct {
	ID        uint64     `db:"id"         json:"id,omitempty"`
	Platform  string     `db:"platform"   json:"platform,omitempty"`
	UserID    uint64     `db:"user_id"    json:"user_id,omitempty"`
	EnteredAt *time.Time `db:"entered_at" json:"entered_at,omitempty"`
	Removed   bool       `db:"removed"    json:"removed,omitempty"`
	CreatedAt *time.Time `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// EventType type
type EventType uint8

const (
	// Created even type
	Created EventType = iota + 1
	// Updated even type
	Updated
	// Removed even type
	Removed
)

// EventStatus status event
type EventStatus uint8

const (
	// Deferred event status
	Deferred EventStatus = iota + 1
	// Processed event status
	Processed
)

// DeviceEvent struct
type DeviceEvent struct {
	ID        uint64      `db:"id"`
	DeviceID  uint64      `db:"device_id"`
	Type      EventType   `db:"type"`
	Status    EventStatus `db:"status"`
	Device    *Device     `db:"payload"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
}
