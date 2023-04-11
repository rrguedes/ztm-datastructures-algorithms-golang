package vin_test

import (
	"testing"

	"elv.com.br/ds/vin"
)

const testVin = "W09000051T2123456"

func TestVinManufacturer(t *testing.T) {
	manufacturer := vin.Manufacturer(testVin)
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVin)
	}
}
