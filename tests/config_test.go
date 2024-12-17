package tests

import (
	"testing"

	"tunnel-hel/config"
)

func TestConfig(t *testing.T) {
	config.LoadConfig()
	t.Log(config.Get())
}
