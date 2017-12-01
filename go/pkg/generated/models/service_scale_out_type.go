package models

// ServiceScaleOutType

import "encoding/json"

type ServiceScaleOutType struct {
	AutoScale    bool `json:"auto_scale"`
	MaxInstances int  `json:"max_instances"`
}

func (model *ServiceScaleOutType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeServiceScaleOutType() *ServiceScaleOutType {
	return &ServiceScaleOutType{
		//TODO(nati): Apply default
		AutoScale:    false,
		MaxInstances: 0,
	}
}

func InterfaceToServiceScaleOutType(iData interface{}) *ServiceScaleOutType {
	data := iData.(map[string]interface{})
	return &ServiceScaleOutType{
		AutoScale: data["auto_scale"].(bool),

		//{"Title":"","Description":"Automatically change the number of virtual machines. Not implemented","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AutoScale","GoType":"bool"}
		MaxInstances: data["max_instances"].(int),

		//{"Title":"","Description":"Maximum number of scale out factor(virtual machines). can be changed dynamically","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"MaxInstances","GoType":"int"}

	}
}

func InterfaceToServiceScaleOutTypeSlice(data interface{}) []*ServiceScaleOutType {
	list := data.([]interface{})
	result := MakeServiceScaleOutTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToServiceScaleOutType(item))
	}
	return result
}

func MakeServiceScaleOutTypeSlice() []*ServiceScaleOutType {
	return []*ServiceScaleOutType{}
}
