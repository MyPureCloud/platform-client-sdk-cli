package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HeaderentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HeaderentryDud struct { 
    


    

}

// Headerentry
type Headerentry struct { 
    // Key - The key of the header (e.g., 'Subject', 'From', 'X-Custom-Header').
    Key string `json:"key"`


    // Value - The value of the header.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Headerentry) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Headerentry) MarshalJSON() ([]byte, error) {
    type Alias Headerentry

    if HeaderentryMarshalled {
        return []byte("{}"), nil
    }
    HeaderentryMarshalled = true

    return json.Marshal(&struct {
        
        Key string `json:"key"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

