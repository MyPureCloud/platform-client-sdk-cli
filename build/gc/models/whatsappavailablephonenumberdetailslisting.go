package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WhatsappavailablephonenumberdetailslistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WhatsappavailablephonenumberdetailslistingDud struct { 
    

}

// Whatsappavailablephonenumberdetailslisting
type Whatsappavailablephonenumberdetailslisting struct { 
    // Entities
    Entities []Whatsappavailablephonenumberdetails `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Whatsappavailablephonenumberdetailslisting) String() string {
    
    
     o.Entities = []Whatsappavailablephonenumberdetails{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Whatsappavailablephonenumberdetailslisting) MarshalJSON() ([]byte, error) {
    type Alias Whatsappavailablephonenumberdetailslisting

    if WhatsappavailablephonenumberdetailslistingMarshalled {
        return []byte("{}"), nil
    }
    WhatsappavailablephonenumberdetailslistingMarshalled = true

    return json.Marshal(&struct { 
        Entities []Whatsappavailablephonenumberdetails `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Whatsappavailablephonenumberdetails{{}},
        

        
        Alias: (*Alias)(u),
    })
}

