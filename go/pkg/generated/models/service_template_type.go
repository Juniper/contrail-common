package models
// ServiceTemplateType



import "encoding/json"

type ServiceTemplateType struct {

    InstanceData string `json:"instance_data"`
    ServiceMode ServiceModeType `json:"service_mode"`
    Version int `json:"version"`
    ServiceType ServiceType `json:"service_type"`
    Flavor string `json:"flavor"`
    VrouterInstanceType VRouterInstanceType `json:"vrouter_instance_type"`
    AvailabilityZoneEnable bool `json:"availability_zone_enable"`
    OrderedInterfaces bool `json:"ordered_interfaces"`
    ServiceVirtualizationType ServiceVirtualizationType `json:"service_virtualization_type"`
    InterfaceType []*ServiceTemplateInterfaceType `json:"interface_type"`
    ImageName string `json:"image_name"`
    ServiceScaling bool `json:"service_scaling"`
}

func (model *ServiceTemplateType) String() string{
    b, _ := json.Marshal(model)
    return string(b)
}

func MakeServiceTemplateType() *ServiceTemplateType{
    return &ServiceTemplateType{
    //TODO(nati): Apply default
    Version: 0,
        ServiceType: MakeServiceType(),
        Flavor: "",
        VrouterInstanceType: MakeVRouterInstanceType(),
        InstanceData: "",
        ServiceMode: MakeServiceModeType(),
        ServiceVirtualizationType: MakeServiceVirtualizationType(),
        
            // {"Title":"","Description":"List of interfaces which decided number of interfaces and type","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"service_interface_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceInterfaceType","GoType":"ServiceInterfaceType"},"shared_ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SharedIP","GoType":"bool"},"static_route_enable":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StaticRouteEnable","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceTemplateInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"InterfaceType","GoType":"ServiceTemplateInterfaceType"},"GoName":"InterfaceType","GoType":"[]*ServiceTemplateInterfaceType"}
            
                InterfaceType:  MakeServiceTemplateInterfaceTypeSlice(),
            
        ImageName: "",
        ServiceScaling: false,
        AvailabilityZoneEnable: false,
        OrderedInterfaces: false,
        
    }
}

func InterfaceToServiceTemplateType(iData interface{}) *ServiceTemplateType {
    data := iData.(map[string]interface{})
    return &ServiceTemplateType{
    Flavor: data["flavor"].(string),
        
        //{"Title":"","Description":"Nova flavor used for service virtual machines, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Flavor","GoType":"string"}
        VrouterInstanceType: InterfaceToVRouterInstanceType(data["vrouter_instance_type"]),
        
        //{"Title":"","Description":"Mechanism used to spawn service instance, when vrouter is spawning instances.Allowed values libvirt-qemu, docker or netns","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["libvirt-qemu","docker"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/VRouterInstanceType","CollectionType":"","Column":"","Item":null,"GoName":"VrouterInstanceType","GoType":"VRouterInstanceType"}
        InstanceData: data["instance_data"].(string),
        
        //{"Title":"","Description":"Opaque string (typically in json format) used to spawn a vrouter-instance.","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"InstanceData","GoType":"string"}
        ServiceMode: InterfaceToServiceModeType(data["service_mode"]),
        
        //{"Title":"","Description":"Service instance mode decides how packets are forwarded across the service","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["transparent","in-network","in-network-nat"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceModeType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceMode","GoType":"ServiceModeType"}
        Version: data["version"].(int),
        
        //{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"integer","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"Version","GoType":"int"}
        ServiceType: InterfaceToServiceType(data["service_type"]),
        
        //{"Title":"","Description":"Service instance mode decides how routing happens across the service","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":["firewall","analyzer","source-nat","loadbalancer"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceType","GoType":"ServiceType"}
        ImageName: data["image_name"].(string),
        
        //{"Title":"","Description":"Glance image name for the service virtual machine, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ImageName","GoType":"string"}
        ServiceScaling: data["service_scaling"].(bool),
        
        //{"Title":"","Description":"Enable scaling of service virtual machines, Version 1 only","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"ServiceScaling","GoType":"bool"}
        AvailabilityZoneEnable: data["availability_zone_enable"].(bool),
        
        //{"Title":"","Description":"Enable availability zone for version 1 service instances","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"AvailabilityZoneEnable","GoType":"bool"}
        OrderedInterfaces: data["ordered_interfaces"].(bool),
        
        //{"Title":"","Description":"Deprecated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"OrderedInterfaces","GoType":"bool"}
        ServiceVirtualizationType: InterfaceToServiceVirtualizationType(data["service_virtualization_type"]),
        
        //{"Title":"","Description":"Service virtualization type decides how individual service instances are instantiated","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":{},"Enum":["virtual-machine","network-namespace","vrouter-instance","physical-device"],"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceVirtualizationType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceVirtualizationType","GoType":"ServiceVirtualizationType"}
        
            
                InterfaceType:  InterfaceToServiceTemplateInterfaceTypeSlice(data["interface_type"]),
            
        
        //{"Title":"","Description":"List of interfaces which decided number of interfaces and type","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"array","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"object","Permission":null,"Properties":{"service_interface_type":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"true","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"ServiceInterfaceType","GoType":"ServiceInterfaceType"},"shared_ip":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"SharedIP","GoType":"bool"},"static_route_enable":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"boolean","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"StaticRouteEnable","GoType":"bool"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/ServiceTemplateInterfaceType","CollectionType":"","Column":"","Item":null,"GoName":"InterfaceType","GoType":"ServiceTemplateInterfaceType"},"GoName":"InterfaceType","GoType":"[]*ServiceTemplateInterfaceType"}
        
    }
}


func InterfaceToServiceTemplateTypeSlice(data interface{}) []*ServiceTemplateType {
    list := data.([]interface{})
    result := MakeServiceTemplateTypeSlice()
    for _, item := range list {
        result = append(result, InterfaceToServiceTemplateType(item))
    }
    return result
}

func MakeServiceTemplateTypeSlice() []*ServiceTemplateType {
    return []*ServiceTemplateType{}
}
