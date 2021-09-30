package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainnetworkaddressMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainnetworkaddressDud struct { 
    


    


    


    

}

// Domainnetworkaddress
type Domainnetworkaddress struct { 
    // VarType - The type of address.
    VarType string `json:"type"`


    // Address - An IPv4 or IPv6 IP address. When specifying an address of type \"ip\", use CIDR format for the subnet mask.
    Address string `json:"address"`


    // Persistent - True if this address will persist on Edge restart.  Addresses assigned by DHCP will be returned as false.
    Persistent bool `json:"persistent"`


    // Family - The address family for this address.
    Family int `json:"family"`

}

// String returns a JSON representation of the model
func (o *Domainnetworkaddress) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainnetworkaddress) MarshalJSON() ([]byte, error) {
    type Alias Domainnetworkaddress

    if DomainnetworkaddressMarshalled {
        return []byte("{}"), nil
    }
    DomainnetworkaddressMarshalled = true

    return json.Marshal(&struct { 
        VarType string `json:"type"`
        
        Address string `json:"address"`
        
        Persistent bool `json:"persistent"`
        
        Family int `json:"family"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

