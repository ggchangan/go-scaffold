package model

import "time"

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// ObjectMeta is also used by gorm.
type ObjectMeta struct {
	// ID is the unique in time and space value for this object. It is typically generated by
	// the storage on successful creation of a resource and is not allowed to change on PUT
	// operations.
	//
	// Populated by the system.
	// Read-only.
	//ID uint64 `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	ID uint64 `json:"id" gorm:"primary_key;AUTO_INCREMENT;column:id"`

	// InstanceID defines a string type resource identifier,
	// use prefixed to distinguish resource types, easy to remember, Url-friendly.
	//InstanceID string `json:"instanceID,omitempty" gorm:"unique;column:instanceID;type:varchar(32);not null"`

	// Name defines the space within each name must be unique.
	// Not all objects are required to be scoped to a username - the value of this field for
	// those objects will be empty.
	//
	// Must be a DNS_LABEL.
	// Cannot be updated.
	// Username string `json:"username,omitempty" gorm:"column:username" validate:"omitempty"`

	// Required: true
	// Name must be unique. Is required when creating resources.
	// Name is primarily intended for creation idempotence and configuration
	// definition.
	// It will be generated automated only if Name is not specified.
	// Cannot be updated.
	//Name string `json:"name,omitempty" gorm:"column:name;type:varchar(64);not null" validate:"name"`
	Name string `json:"name" gorm:"column:name;type:varchar(64);not null" validate:"name"`

	// CreatedAt is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system.
	// Read-only.
	// Null for lists.
	//CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:created_at"`

	// UpdatedAt is a timestamp representing the server time when this object was updated.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	//
	// Populated by the system.
	// Read-only.
	// Null for lists.
	//UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`

	// DeletedAt is RFC 3339 date and time at which this resource will be deleted. This
	// field is set by the server when a graceful deletion is requested by the user, and is not
	// directly settable by a client.
	//
	// Populated by the system when a graceful deletion is requested.
	// Read-only.
	// DeletedAt *time.Time `json:"-" gorm:"column:deletedAt;index:idx_deletedAt"`
}
