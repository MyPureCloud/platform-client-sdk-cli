package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IpaddressauthenticationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IpaddressauthenticationDud struct { 
    

}

// Ipaddressauthentication
type Ipaddressauthentication struct { 
    // NetworkWhitelist
    NetworkWhitelist []string `json:"networkWhitelist"`

}

// String returns a JSON representation of the model
func (o *Ipaddressauthentication) String() string {
    
    
     o.NetworkWhitelist = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ipaddressauthentication) MarshalJSON() ([]byte, error) {
    type Alias Ipaddressauthentication

    if IpaddressauthenticationMarshalled {
        return []byte("{}"), nil
    }
    IpaddressauthenticationMarshalled = true

    return json.Marshal(&struct { 
        NetworkWhitelist []string `json:"networkWhitelist"`
        
        *Alias
    }{
        

        
        NetworkWhitelist: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

