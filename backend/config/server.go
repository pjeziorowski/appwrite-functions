package config

import (
	"errors"
	"os"
)

var (
	Secret               = os.Getenv("SECRET")
	HasuraApiUrl         = os.Getenv("HASURA_API_URL")
	HasuraApiToken       = os.Getenv("HASURA_API_TOKEN")
	QoveryApiToken       = os.Getenv("QOVERY_API_TOKEN")
	QoveryOrganizationId = os.Getenv("QOVERY_ORGANIZATION_ID")
)

func CheckServerConfig() []error {
	var configErrors []error

	if Secret == "" {
		configErrors = append(configErrors, errors.New("SECRET env required"))
	}
	if HasuraApiUrl == "" {
		configErrors = append(configErrors, errors.New("HASURA_API_URL env required"))
	}
	if HasuraApiToken == "" {
		configErrors = append(configErrors, errors.New("HASURA_API_TOKEN env required"))
	}
	if QoveryApiToken == "" {
		configErrors = append(configErrors, errors.New("QOVERY_API_TOKEN env required"))
	}
	if QoveryOrganizationId == "" {
		configErrors = append(configErrors, errors.New("QOVERY_ORGANIZATION_ID env required"))
	}

	return configErrors
}
