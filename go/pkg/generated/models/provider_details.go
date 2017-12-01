package models

// ProviderDetails

import "encoding/json"

type ProviderDetails struct {
	SegmentationID  VlanIdType `json:"segmentation_id"`
	PhysicalNetwork string     `json:"physical_network"`
}

func (model *ProviderDetails) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeProviderDetails() *ProviderDetails {
	return &ProviderDetails{
		//TODO(nati): Apply default
		PhysicalNetwork: "",
		SegmentationID:  MakeVlanIdType(),
	}
}

func InterfaceToProviderDetails(iData interface{}) *ProviderDetails {
	data := iData.(map[string]interface{})
	return &ProviderDetails{
		PhysicalNetwork: data["physical_network"].(string),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"PhysicalNetwork","GoType":"string"}
		SegmentationID: InterfaceToVlanIdType(data["segmentation_id"]),

		//{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"integer","Permission":null,"Properties":{},"Enum":null,"Minimum":1,"Maximum":4094,"Ref":"types.json#/definitions/VlanIdType","CollectionType":"","Column":"","Item":null,"GoName":"SegmentationID","GoType":"VlanIdType"}

	}
}

func InterfaceToProviderDetailsSlice(data interface{}) []*ProviderDetails {
	list := data.([]interface{})
	result := MakeProviderDetailsSlice()
	for _, item := range list {
		result = append(result, InterfaceToProviderDetails(item))
	}
	return result
}

func MakeProviderDetailsSlice() []*ProviderDetails {
	return []*ProviderDetails{}
}
