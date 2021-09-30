package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainnetworkrouteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainnetworkrouteDud struct { 
    


    


    


    


    

}

// Domainnetworkroute
type Domainnetworkroute struct { 
    // Prefix - The IPv4 or IPv6 route prefix in CIDR notation.
    Prefix string `json:"prefix"`


    // Nexthop - The IPv4 or IPv6 nexthop IP address.
    Nexthop string `json:"nexthop"`


    // Persistent - True if this route will persist on Edge restart.  Routes assigned by DHCP will be returned as false.
    Persistent bool `json:"persistent"`


    // Metric - The metric being used for route. Lower values will have a higher priority.
    Metric int `json:"metric"`


    // Family - The address family for this route.
    Family int `json:"family"`

}

// String returns a JSON representation of the model
func (o *Domainnetworkroute) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainnetworkroute) MarshalJSON() ([]byte, error) {
    type Alias Domainnetworkroute

    if DomainnetworkrouteMarshalled {
        return []byte("{}"), nil
    }
    DomainnetworkrouteMarshalled = true

    return json.Marshal(&struct { 
        Prefix string `json:"prefix"`
        
        Nexthop string `json:"nexthop"`
        
        Persistent bool `json:"persistent"`
        
        Metric int `json:"metric"`
        
        Family int `json:"family"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

