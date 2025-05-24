package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IpaddressrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IpaddressrangeDud struct { 
    


    


    


    

}

// Ipaddressrange
type Ipaddressrange struct { 
    // Cidr
    Cidr string `json:"cidr"`


    // Service
    Service string `json:"service"`


    // Region
    Region string `json:"region"`


    // Direction - The direction of traffic for the IP range from the perspective of Genesys Cloud (e.g. inbound to Genesys; outbound from Genesys)
    Direction string `json:"direction"`

}

// String returns a JSON representation of the model
func (o *Ipaddressrange) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ipaddressrange) MarshalJSON() ([]byte, error) {
    type Alias Ipaddressrange

    if IpaddressrangeMarshalled {
        return []byte("{}"), nil
    }
    IpaddressrangeMarshalled = true

    return json.Marshal(&struct {
        
        Cidr string `json:"cidr"`
        
        Service string `json:"service"`
        
        Region string `json:"region"`
        
        Direction string `json:"direction"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

