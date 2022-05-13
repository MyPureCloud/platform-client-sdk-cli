package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SocialhandleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SocialhandleDud struct { 
    


    

}

// Socialhandle
type Socialhandle struct { 
    // VarType
    VarType string `json:"type"`


    // Value
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Socialhandle) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Socialhandle) MarshalJSON() ([]byte, error) {
    type Alias Socialhandle

    if SocialhandleMarshalled {
        return []byte("{}"), nil
    }
    SocialhandleMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

