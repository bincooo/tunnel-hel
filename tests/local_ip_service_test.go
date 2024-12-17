package tests

import (
	"testing"

	"tunnel-hel/service/ip"
)

func TestGetIpInfoLocal(t *testing.T) {
	result := ip.GetIpInfo("139.9.64.238", "")
	t.Log(result)
}
