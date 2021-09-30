package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IpaddressrangelistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IpaddressrangelistingDud struct { 
    

}

// Ipaddressrangelisting
type Ipaddressrangelisting struct { 
    // Entities
    Entities []Ipaddressrange `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Ipaddressrangelisting) String() string {
    
    
     o.Entities = []Ipaddressrange{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ipaddressrangelisting) MarshalJSON() ([]byte, error) {
    type Alias Ipaddressrangelisting

    if IpaddressrangelistingMarshalled {
        return []byte("{}"), nil
    }
    IpaddressrangelistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Ipaddressrange `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Ipaddressrange{{}},
        

        
        Alias: (*Alias)(u),
    })
}

