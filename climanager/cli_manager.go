package climanager

import (
	"github.com/akshaysingh1713/todoline/resources"
)

type CliManager struct {
	resources resources.Resources
}

func InitCliManager(auth_token string) CliManager {
	resources := resources.InitResources(auth_token)
	cli_manager := CliManager{
		resources: resources,
	}
	return cli_manager
}
