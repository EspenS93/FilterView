package provider

import (
	"github.com/EspenS93/filterview/internal/provider/peopleProvider"
	"github.com/EspenS93/filterview/internal/provider/providerInterface"
)

func New(providerName string) providerInterface.Provider {
	switch providerName {
	case "people":
		return &peopleProvider.Provider{}
	default:
		return &peopleProvider.Provider{}
	}
}

type Provider struct {
}
