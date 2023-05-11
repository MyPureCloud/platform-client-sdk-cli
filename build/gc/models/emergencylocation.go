package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmergencylocationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmergencylocationDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Emergencylocation
type Emergencylocation struct { 
    


    // Name
    Name string `json:"name"`


    // Address - Emergency address
    Address Locationaddress `json:"address"`


    // Did - Phone number in E164 format
    Did string `json:"did"`


    // Source - source
    Source string `json:"source"`


    

}

// String returns a JSON representation of the model
func (o *Emergencylocation) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emergencylocation) MarshalJSON() ([]byte, error) {
    type Alias Emergencylocation

    if EmergencylocationMarshalled {
        return []byte("{}"), nil
    }
    EmergencylocationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Address Locationaddress `json:"address"`
        
        Did string `json:"did"`
        
        Source string `json:"source"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

