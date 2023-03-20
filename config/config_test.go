package config

import (
	"fmt"
	"testing"
)

func TestNewConfig(t *testing.T) {
	config := newConfig()
	fmt.Printf("config: %v\n", config)
}
