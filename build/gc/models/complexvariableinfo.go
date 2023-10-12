package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ComplexvariableinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ComplexvariableinfoDud struct { 
    


    

}

// Complexvariableinfo - Variable type information about a complex type from the bot's definition
type Complexvariableinfo struct { 
    // Id - The variable type ID
    Id string `json:"id"`


    // Name - The variable type display name
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Complexvariableinfo) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Complexvariableinfo) MarshalJSON() ([]byte, error) {
    type Alias Complexvariableinfo

    if ComplexvariableinfoMarshalled {
        return []byte("{}"), nil
    }
    ComplexvariableinfoMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

