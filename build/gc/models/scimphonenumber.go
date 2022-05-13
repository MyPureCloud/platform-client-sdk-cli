package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScimphonenumberMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScimphonenumberDud struct { 
    


    


    

}

// Scimphonenumber - Defines a SCIM phone number.
type Scimphonenumber struct { 
    // Value - The phone number in E.164 or tel URI format, for example, tel:+nnnnnnnn; ext=xxxxx.
    Value string `json:"value"`


    // VarType - The type of phone number.
    VarType string `json:"type"`


    // Primary - Indicates whether the phone number is the primary phone number.
    Primary bool `json:"primary"`

}

// String returns a JSON representation of the model
func (o *Scimphonenumber) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scimphonenumber) MarshalJSON() ([]byte, error) {
    type Alias Scimphonenumber

    if ScimphonenumberMarshalled {
        return []byte("{}"), nil
    }
    ScimphonenumberMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        
        VarType string `json:"type"`
        
        Primary bool `json:"primary"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

