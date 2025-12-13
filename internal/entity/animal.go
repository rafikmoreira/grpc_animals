// Package entity provides the entity definitions for the animals service.
package entity

import "time"

type Animal struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	AnimalType string    `json:"animal_type"`
	CreatedAt  time.Time `json:"created_at"`
}
