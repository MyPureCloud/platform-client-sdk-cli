package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcentercustommessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcentercustommessageDud struct { 
    


    

}

// Supportcentercustommessage
type Supportcentercustommessage struct { 
    // DefaultValue - Default value for the message. Required for each custom message entry
    DefaultValue string `json:"defaultValue"`


    // VarType - Type of the message. Required for each custom message entry
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Supportcentercustommessage) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcentercustommessage) MarshalJSON() ([]byte, error) {
    type Alias Supportcentercustommessage

    if SupportcentercustommessageMarshalled {
        return []byte("{}"), nil
    }
    SupportcentercustommessageMarshalled = true

    return json.Marshal(&struct {
        
        DefaultValue string `json:"defaultValue"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

