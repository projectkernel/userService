package controller

import (
	"auth/src/operation"
	"auth/src/format"
)

type ProviderPersistence struct {
	prov *operation.Provider
	persist *operation.Persistent
	format    format.Formatter
}

func NewProviderPersistence(persist *operation.Persistent, prov *operation.Provider, form format.Formatter) *ProviderPersistence {
	return &ProviderPersistence{
		persist: persist,
		prov:      prov,
		format:    form,
	}
}