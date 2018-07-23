package gonionoo

import (
	"testing"
)

var fingerprint = "D5F2C65F4131A1468D5B67A8838A9B7ED8C049E2"
var maxLimitForTest = "100"

func TestGetSummaryNoQuery(t *testing.T) {
	summary, _, err := GetSummary(nil, "")
	if err != nil {
		t.Error(err)
	}

	if summary == nil {
		t.Error("We didn't get a result, even though we should have")
	}
}

func TestGetSummaryInvalidQueryParameter(t *testing.T) {
	_, _, err := GetSummary(map[string]string{"invalidParameter": "invalidValue"}, "")
	if err == nil {
		t.Error(err)
	}
}

func TestGetSummaryWithFingerprintQueryParameter(t *testing.T) {
	// We are using the fingerprint of che (https://metrics.torproject.org/rs.html#details/D5F2C65F4131A1468D5B67A8838A9B7ED8C049E2)
	// a very solid running Tor node
	summary, _, err := GetSummary(map[string]string{"fingerprint": fingerprint}, "")
	if err != nil {
		t.Error(err)
	}

	if summary == nil {
		t.Errorf("Failed to get a summary object")
	}

	if len(summary.Relays) < 1 {
		t.Errorf("Got response, but got no relays")
	}

	if summary.Relays[0].Fingerprint != fingerprint {
		t.Errorf("Got response but result finger print '%s' does not equal expected finger print '%s'", summary.Relays[0].Fingerprint, fingerprint)
	}
}

func TestGetDetailsNoQuery(t *testing.T) {
	details, _, err := GetDetails(map[string]string{"limit": maxLimitForTest}, "")
	if err != nil {
		t.Error(err)
	}

	if details == nil {
		t.Errorf("Failed to get details object")
	}

	if len(details.Relays) < 1 {
		t.Errorf("Got response, but got no relays")
	}
}

func TestGetDetailsWithFingerprintQueryParameter(t *testing.T) {
	details, _, err := GetDetails(map[string]string{"fingerprint": fingerprint}, "")
	if err != nil {
		t.Error(err)
	}

	if details == nil {
		t.Errorf("Failed to get details object")
	}

	if len(details.Relays) < 1 {
		t.Errorf("Got response, but got no relays")
	}

	if details.Relays[0].Fingerprint != fingerprint {
		t.Errorf("Got response but result finger print '%s' does not equal expected finger print '%s'", details.Relays[0].Fingerprint, fingerprint)
	}
}

func TestGetBandwidthWithFingerprintQueryParameter(t *testing.T) {
	bandwidth, _, err := GetBandwidth(map[string]string{"fingerprint": fingerprint}, "")
	if err != nil {
		t.Error(err)
	}

	if bandwidth == nil {
		t.Errorf("Failed to get bandwidth object")
	}

	if len(bandwidth.Relays) < 1 {
		t.Errorf("Got response, but got no relays")
	}

	if bandwidth.Relays[0].Fingerprint != fingerprint {
		t.Errorf("Got response but result finger print '%s' does not equal expected finger print '%s'", bandwidth.Relays[0].Fingerprint, fingerprint)
	}
}

func TestGetWeightsWithFingerprintQueryParameter(t *testing.T) {
	weights, _, err := GetWeights(map[string]string{"fingerprint": fingerprint}, "")
	if err != nil {
		t.Error(err)
	}

	if weights == nil {
		t.Errorf("Failed to get weights object")
	}

	if len(weights.Relays) < 1 {
		t.Errorf("Got response, but got no relays")
	}

	if weights.Relays[0].Fingerprint != fingerprint {
		t.Errorf("Got response but result finger print '%s' does not equal expected finger print '%s'", weights.Relays[0].Fingerprint, fingerprint)
	}
}

func TestGetClientsNoQuery(t *testing.T) {
	clients, _, err := GetClients(map[string]string{"limit": maxLimitForTest}, "")
	if err != nil {
		t.Error(err)
	}

	if clients == nil {
		t.Errorf("Failed to get clients object")
	}

	if len(clients.Bridges) < 1 {
		t.Errorf("Got response, but got no bridges")
	}
}

func TestGetUptimeNoQuery(t *testing.T) {
	uptime, _, err := GetUptime(map[string]string{"limit": maxLimitForTest}, "")
	if err != nil {
		t.Error(err)
	}

	if uptime == nil {
		t.Errorf("Failed to get uptime object")
	}

	if len(uptime.Relays) < 1 {
		t.Errorf("Got response, but got no relays")
	}
}

func TestValidatMethodNoMethod(t *testing.T) {
	err := validateMethod("")
	if err == nil {
		t.Errorf("validateMethod with no method didn't return an error")
	}
}

func TestValidatMethodUnknownMethod(t *testing.T) {
	err := validateMethod("Unknown")
	if err == nil {
		t.Errorf("validateMethod with an unknown method didn't return an error")
	}
}

func TestExecuteRequestWithUnknownMethod(t *testing.T) {
	var result = new(Uptime)
	_, err := executeRequest("Unknown", nil, "", &result)
	if err == nil {
		t.Errorf("executeRequest with an unknown method didn't return an error")
	}
}
