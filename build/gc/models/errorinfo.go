package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ErrorinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ErrorinfoDud struct { 
    


    

}

// Errorinfo
type Errorinfo struct { 
    // Message
    Message string `json:"message"`


    // Code
    Code string `json:"code"`

}

// String returns a JSON representation of the model
func (o *Errorinfo) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Errorinfo) MarshalJSON() ([]byte, error) {
    type Alias Errorinfo

    if ErrorinfoMarshalled {
        return []byte("{}"), nil
    }
    ErrorinfoMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        Code string `json:"code"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

