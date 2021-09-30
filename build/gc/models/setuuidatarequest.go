package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SetuuidatarequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SetuuidatarequestDud struct { 
    

}

// Setuuidatarequest
type Setuuidatarequest struct { 
    // UuiData - The value of the uuiData to set.
    UuiData string `json:"uuiData"`

}

// String returns a JSON representation of the model
func (o *Setuuidatarequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Setuuidatarequest) MarshalJSON() ([]byte, error) {
    type Alias Setuuidatarequest

    if SetuuidatarequestMarshalled {
        return []byte("{}"), nil
    }
    SetuuidatarequestMarshalled = true

    return json.Marshal(&struct { 
        UuiData string `json:"uuiData"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

