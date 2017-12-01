package models

// DomainLimitsType

import "encoding/json"

type DomainLimitsType struct {
	SecurityGroupLimit  int `json:"security_group_limit"`
	ProjectLimit        int `json:"project_limit"`
	VirtualNetworkLimit int `json:"virtual_network_limit"`
}

func (model *DomainLimitsType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeDomainLimitsType() *DomainLimitsType {
	return &DomainLimitsType{
		//TODO(nati): Apply default
		VirtualNetworkLimit: 0,
		SecurityGroupLimit:  0,
		ProjectLimit:        0,
	}
}

func InterfaceToDomainLimitsType(iData interface{}) *DomainLimitsType {
	data := iData.(map[string]interface{})
	return &DomainLimitsType{
		ProjectLimit: data["project_limit"].(int),

		//{"Title":"","Description":"Maximum number of projects allowed in this domain","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ProjectLimit","GoType":"int"}
		VirtualNetworkLimit: data["virtual_network_limit"].(int),

		//{"Title":"","Description":"Maximum number of virtual networks allowed in this domain","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualNetworkLimit","GoType":"int"}
		SecurityGroupLimit: data["security_group_limit"].(int),

		//{"Title":"","Description":"Maximum number of security groups allowed in this domain","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SecurityGroupLimit","GoType":"int"}

	}
}

func InterfaceToDomainLimitsTypeSlice(data interface{}) []*DomainLimitsType {
	list := data.([]interface{})
	result := MakeDomainLimitsTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToDomainLimitsType(item))
	}
	return result
}

func MakeDomainLimitsTypeSlice() []*DomainLimitsType {
	return []*DomainLimitsType{}
}
