package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WrapupdeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WrapupdeltaDud struct { 
    


    


    

}

// Wrapupdelta
type Wrapupdelta struct { 
    // Action
    Action string `json:"action"`


    // Code
    Code string `json:"code"`


    // UserId
    UserId string `json:"userId"`

}

// String returns a JSON representation of the model
func (o *Wrapupdelta) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wrapupdelta) MarshalJSON() ([]byte, error) {
    type Alias Wrapupdelta

    if WrapupdeltaMarshalled {
        return []byte("{}"), nil
    }
    WrapupdeltaMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        Code string `json:"code"`
        
        UserId string `json:"userId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

