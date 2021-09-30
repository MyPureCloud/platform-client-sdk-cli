package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomaincapabilitiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomaincapabilitiesDud struct { 
    


    


    


    


    SupportsMetric bool `json:"supportsMetric"`


    

}

// Domaincapabilities
type Domaincapabilities struct { 
    // Enabled - True if this address family on the interface is enabled.
    Enabled bool `json:"enabled"`


    // Dhcp - True if this address family on the interface is using DHCP.
    Dhcp bool `json:"dhcp"`


    // Metric - The metric being used for the address family on this interface. Lower values will have a higher priority. If autoMetric is true, this value will be the automatically calculated metric. To set this value be sure autoMetric is false. If no value is returned, metric configuration is not supported on this Edge.
    Metric int `json:"metric"`


    // AutoMetric - True if the metric is being calculated automatically for the address family on this interface.
    AutoMetric bool `json:"autoMetric"`


    


    // PingEnabled - Set to true to enable this address family on this interface to respond to ping requests.
    PingEnabled bool `json:"pingEnabled"`

}

// String returns a JSON representation of the model
func (o *Domaincapabilities) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domaincapabilities) MarshalJSON() ([]byte, error) {
    type Alias Domaincapabilities

    if DomaincapabilitiesMarshalled {
        return []byte("{}"), nil
    }
    DomaincapabilitiesMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        Dhcp bool `json:"dhcp"`
        
        Metric int `json:"metric"`
        
        AutoMetric bool `json:"autoMetric"`
        
        
        
        PingEnabled bool `json:"pingEnabled"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

