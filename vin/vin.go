package vin

import "fmt"

// Para ser verificável, necessita implementar Manufacturer tal como abaixo
type VinVerifiable interface {
	Manufacturer() string
}

type Vin string      // o primeiro objeto criar um tipo a partir de uma string
type EuropeanVin Vin // o segundo é um novo tipo que deriva do primeiro

// Construtor
func NewVin(code string) (Vin, error) {
	if len(code) != 17 {
		return "", fmt.Errorf("invalid vin %s: it has more or less than 17 chars", code)
	}
	return Vin(code), nil
}

// Construtor
func NewEuropeanVin(code string) (EuropeanVin, error) {
	// Chamando o construtor da superclasse
	euroVin, err := NewVin(code)
	return EuropeanVin(euroVin), err
}

/* Manufacturer */

func (vin Vin) Manufacturer() string {
	manufacturer := vin[:3]
	if manufacturer[2] == '9' {
		manufacturer += vin[11:14]
	}
	return string(manufacturer)
}

func (vin EuropeanVin) Manufacturer() string {
	manufacturer := vin[:3]
	if manufacturer[2] == '9' {
		manufacturer += vin[11:14]
	}
	return string(manufacturer)
}
