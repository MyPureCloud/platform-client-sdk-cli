package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NtpsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NtpsettingsDud struct { 
    

}

// Ntpsettings
type Ntpsettings struct { 
    // Servers - List of NTP servers, in priority order
    Servers []string `json:"servers"`

}

// String returns a JSON representation of the model
func (o *Ntpsettings) String() string {
     o.Servers = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ntpsettings) MarshalJSON() ([]byte, error) {
    type Alias Ntpsettings

    if NtpsettingsMarshalled {
        return []byte("{}"), nil
    }
    NtpsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Servers []string `json:"servers"`
        *Alias
    }{

        
        Servers: []string{""},
        

        Alias: (*Alias)(u),
    })
}

