package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateutilizationtagrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateutilizationtagrequestDud struct { 
    

}

// Createutilizationtagrequest
type Createutilizationtagrequest struct { 
    // Name - The utilization tag name.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Createutilizationtagrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createutilizationtagrequest) MarshalJSON() ([]byte, error) {
    type Alias Createutilizationtagrequest

    if CreateutilizationtagrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateutilizationtagrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

