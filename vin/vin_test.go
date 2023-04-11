package vin_test

import (
	"fmt"
	"testing"

	"elv.com.br/ds/vin"
)

const testVin = "W09000051T2123456"

/* func TestVinManufacturer(t *testing.T) {
	manufacturer := vin.Manufacturer(testVin)
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVin)
	}
} */

const (
	validVin   = vin.Vin("W09000051T2123456")
	invalidVin = vin.Vin("W0")
)

func TestVinManufacturer(t *testing.T) {
	manufacturer := validVin.Manufacturer()
	if manufacturer != "W0L" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, validVin)
	}
}

func TestVinNew(t *testing.T) {
	validReturn, err := vin.NewVin(string(validVin))
	if err == nil {
		t.Errorf("creating an valid Vin generated error: %s", err.Error())
	} else {
		fmt.Println(validReturn)
	}

	invalidReturn, err := vin.NewVin(string(invalidVin))
	if err == nil {
		t.Errorf("creating invalid Vin did not generated error")
	} else {
		fmt.Println(invalidReturn)
	}
}

const EUROVIN = "W09000051T2123456"

func TestVinPolymorphism(t *testing.T) {
	var testMass []vin.VinVerifiable // Interface que permite implementação através de duck typing
	euroVinToTest, _ := vin.NewEuropeanVin(EUROVIN)
	testMass = append(testMass, euroVinToTest)

	for _, vin := range testMass {
		manufacturer := vin.Manufacturer()
		if manufacturer != "W09123" {
			t.Errorf("unexpected manufacturer %s for VIN %s ", manufacturer, vin)
		}
	}
}

type mockApiClient struct {
	apiCalls int
}

func NewMockApiClient() *mockApiClient {
	return &mockApiClient{}
}

func (client *mockApiClient) IsEuropean(code string) bool {
	client.apiCalls++
	return true
}

func TestVinEuropeanUsingApi(t *testing.T) {
	apiClient := NewMockApiClient()
	service := vin.NewVinService(&vin.VinServiceConfig{}, apiClient)
	testVinInterface, _ := service.CreateFromCode(EUROVIN)

	manufacturer := testVinInterface.Manufacturer()
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVinInterface)
	}

	if apiClient.apiCalls != 1 {
		t.Errorf("unexpected number of API calls: %d", apiClient.apiCalls)
	}
}
