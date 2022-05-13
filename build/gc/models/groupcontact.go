package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GroupcontactMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GroupcontactDud struct { 
    


    


    Display string `json:"display"`


    


    

}

// Groupcontact
type Groupcontact struct { 
    // Address - Phone number for this contact type
    Address string `json:"address"`


    // Extension - Extension is set if the number is e164 valid
    Extension string `json:"extension"`


    


    // VarType - Contact type of the address
    VarType string `json:"type"`


    // MediaType - Media type of the address
    MediaType string `json:"mediaType"`

}

// String returns a JSON representation of the model
func (o *Groupcontact) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Groupcontact) MarshalJSON() ([]byte, error) {
    type Alias Groupcontact

    if GroupcontactMarshalled {
        return []byte("{}"), nil
    }
    GroupcontactMarshalled = true

    return json.Marshal(&struct {
        
        Address string `json:"address"`
        
        Extension string `json:"extension"`
        
        VarType string `json:"type"`
        
        MediaType string `json:"mediaType"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

