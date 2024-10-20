package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HeaderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HeaderDud struct { 
    


    

}

// Header
type Header struct { 
    // Name - Name of the header
    Name string `json:"name"`


    // Value - Value of the header
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Header) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Header) MarshalJSON() ([]byte, error) {
    type Alias Header

    if HeaderMarshalled {
        return []byte("{}"), nil
    }
    HeaderMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

