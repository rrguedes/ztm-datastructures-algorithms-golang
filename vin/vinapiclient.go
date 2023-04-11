package vin

/*

O problema exposto pelo artigo é ter o cliente dependendo totalmente do service
sem que possa ser substituído por um mock no desenvolvimento local ou para a re
dução de custos, por ex. Portanto, uma solução para isso seria injetar a depend
ência do cliente no serviço.

*/

type VinApiClientInterface interface {
	IsEuropean(code string) bool
}

type VinApiClient struct {
	apiUrl, apiKey string
}

func NewVinApiClient(apiUrl, apiKey string) VinApiClientInterface {
	return &VinApiClient{apiUrl: apiUrl, apiKey: apiKey}
}

func (client *VinApiClient) IsEuropean(code string) bool {
	return true
}
