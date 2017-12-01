package models

// IpamDnsAddressType

import "encoding/json"

type IpamDnsAddressType struct {
	VirtualDNSServerName   string           `json:"virtual_dns_server_name"`
	TenantDNSServerAddress *IpAddressesType `json:"tenant_dns_server_address"`
}

func (model *IpamDnsAddressType) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

func MakeIpamDnsAddressType() *IpamDnsAddressType {
	return &IpamDnsAddressType{
		//TODO(nati): Apply default
		TenantDNSServerAddress: MakeIpAddressesType(),
		VirtualDNSServerName:   "",
	}
}

func InterfaceToIpamDnsAddressType(iData interface{}) *IpamDnsAddressType {
	data := iData.(map[string]interface{})
	return &IpamDnsAddressType{
		VirtualDNSServerName: data["virtual_dns_server_name"].(string),

		//{"Title":"","Description":"In case of virtual DNS server, name of virtual DNS server","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"string","Permission":null,"Properties":null,"Enum":null,"Minimum":null,"Maximum":null,"Ref":"","CollectionType":"","Column":"","Item":null,"GoName":"VirtualDNSServerName","GoType":"string"}
		TenantDNSServerAddress: InterfaceToIpAddressesType(data["tenant_dns_server_address"]),

		//{"Title":"","Description":"In case of tenant DNS server method, Ip address of DNS server. This will be given in DHCP","SQL":"","Default":null,"Operation":"","Presence":"optional","Type":"object","Permission":null,"Properties":{"ip_address":{"Title":"","Description":"","SQL":"","Default":null,"Operation":"","Presence":"","Type":"string","Permission":null,"Properties":{},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressType","CollectionType":"","Column":"","Item":null,"GoName":"IPAddress","GoType":"IpAddressType"}},"Enum":null,"Minimum":null,"Maximum":null,"Ref":"types.json#/definitions/IpAddressesType","CollectionType":"","Column":"","Item":null,"GoName":"TenantDNSServerAddress","GoType":"IpAddressesType"}

	}
}

func InterfaceToIpamDnsAddressTypeSlice(data interface{}) []*IpamDnsAddressType {
	list := data.([]interface{})
	result := MakeIpamDnsAddressTypeSlice()
	for _, item := range list {
		result = append(result, InterfaceToIpamDnsAddressType(item))
	}
	return result
}

func MakeIpamDnsAddressTypeSlice() []*IpamDnsAddressType {
	return []*IpamDnsAddressType{}
}
