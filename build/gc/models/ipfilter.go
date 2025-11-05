package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IpfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IpfilterDud struct { 
    


    

}

// Ipfilter - Configuration for filtering tracking based on IP addresses.
type Ipfilter struct { 
    // IpAddress - IP address or CIDR range to filter (e.g. '192.168.1.0/24').
    IpAddress string `json:"ipAddress"`


    // Name - Descriptive name for the IP address filter.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Ipfilter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ipfilter) MarshalJSON() ([]byte, error) {
    type Alias Ipfilter

    if IpfilterMarshalled {
        return []byte("{}"), nil
    }
    IpfilterMarshalled = true

    return json.Marshal(&struct {
        
        IpAddress string `json:"ipAddress"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

