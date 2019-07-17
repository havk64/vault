package reader

import (
	"github.com/hashicorp/vault/api"
	"os"
)

// Data comments
type Data map[string]interface{}

var (
	vaultAddress string = os.Getenv("VAULT_ADDR")
	vaultToken   string = os.Getenv("VAULT_TOKEN")
)

// Read comments
func Read(path string) (Data, error) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	if err := client.SetAddress(vaultAddress); err != nil {
		return nil, err
	}
	client.SetToken(vaultToken)
	secretValue, err := client.Logical().Read(path)
	if err != nil {
		return nil, err
	}
	return (Data)(secretValue.Data), nil
}

// GetValue comments
func (data Data) GetValue(field string) string {
	return data[field].(string)
}
