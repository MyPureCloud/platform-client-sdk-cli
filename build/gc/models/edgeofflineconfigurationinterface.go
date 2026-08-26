package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeofflineconfigurationinterfaceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeofflineconfigurationinterfaceDud struct { 
    


    


    


    

}

// Edgeofflineconfigurationinterface
type Edgeofflineconfigurationinterface struct { 
    // Routes - The list of routes assigned to this interface.
    Routes []Domainnetworkroute `json:"routes"`


    // Addresses - The list of IP addresses on this interface.  Priority of dns addresses are based on order in the list.
    Addresses []Domainnetworkaddress `json:"addresses"`


    // Ipv4Capabilities - IPv4 interface settings.
    Ipv4Capabilities Domaincapabilities `json:"ipv4Capabilities"`


    // Ipv6Capabilities - IPv6 interface settings.
    Ipv6Capabilities Domaincapabilities `json:"ipv6Capabilities"`

}

// String returns a JSON representation of the model
func (o *Edgeofflineconfigurationinterface) String() string {
     o.Routes = []Domainnetworkroute{{}} 
     o.Addresses = []Domainnetworkaddress{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeofflineconfigurationinterface) MarshalJSON() ([]byte, error) {
    type Alias Edgeofflineconfigurationinterface

    if EdgeofflineconfigurationinterfaceMarshalled {
        return []byte("{}"), nil
    }
    EdgeofflineconfigurationinterfaceMarshalled = true

    return json.Marshal(&struct {
        
        Routes []Domainnetworkroute `json:"routes"`
        
        Addresses []Domainnetworkaddress `json:"addresses"`
        
        Ipv4Capabilities Domaincapabilities `json:"ipv4Capabilities"`
        
        Ipv6Capabilities Domaincapabilities `json:"ipv6Capabilities"`
        *Alias
    }{

        
        Routes: []Domainnetworkroute{{}},
        


        
        Addresses: []Domainnetworkaddress{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

