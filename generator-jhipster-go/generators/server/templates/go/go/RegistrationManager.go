package main

type RegistrationManager interface {
	ManageDiscovery(configs RegistrationVariables)
	RegisterWithSerivceRegistry()
	SendHeartBeat(configs RegistrationVariables)
	DeRegisterFromServiceRegistry(configs RegistrationVariables)
}

type RegistrationVariables struct {
	registryType       string
	serviceRegistryURL string
	instanceId         string
}

func (rv RegistrationVariables) RegistryType() string {
	return rv.registryType
}

func (rv RegistrationVariables) ServiceRegistryURL() string {
	return rv.serviceRegistryURL
}

