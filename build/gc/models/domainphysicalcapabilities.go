package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainphysicalcapabilitiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainphysicalcapabilitiesDud struct { 
    


    

}

// Domainphysicalcapabilities
type Domainphysicalcapabilities struct { 
    // Vlan
    Vlan bool `json:"vlan"`


    // Team
    Team bool `json:"team"`

}

// String returns a JSON representation of the model
func (o *Domainphysicalcapabilities) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainphysicalcapabilities) MarshalJSON() ([]byte, error) {
    type Alias Domainphysicalcapabilities

    if DomainphysicalcapabilitiesMarshalled {
        return []byte("{}"), nil
    }
    DomainphysicalcapabilitiesMarshalled = true

    return json.Marshal(&struct {
        
        Vlan bool `json:"vlan"`
        
        Team bool `json:"team"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

