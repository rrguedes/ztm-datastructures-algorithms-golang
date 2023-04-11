package vin

type VinService struct {
	client VinApiClientInterface
}

type VinServiceConfig struct {
	apiUrl, apiKey string
}

// Injeção de dependência como forma de implementar a funcionalidade
func NewVinService(config *VinServiceConfig, apiClient VinApiClientInterface) *VinService {
	return &VinService{apiClient}
}

func (s *VinService) CreateFromCode(code string) (VinVerifiable, error) {
	if s.client.IsEuropean(code) {
		return NewEuropeanVin(code)
	}
	return NewVin(code)
}
