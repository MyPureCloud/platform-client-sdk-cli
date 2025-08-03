package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReprocessinteractionsbyjobidresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReprocessinteractionsbyjobidresponseDud struct { 
    

}

// Reprocessinteractionsbyjobidresponse
type Reprocessinteractionsbyjobidresponse struct { 
    // Url
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Reprocessinteractionsbyjobidresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reprocessinteractionsbyjobidresponse) MarshalJSON() ([]byte, error) {
    type Alias Reprocessinteractionsbyjobidresponse

    if ReprocessinteractionsbyjobidresponseMarshalled {
        return []byte("{}"), nil
    }
    ReprocessinteractionsbyjobidresponseMarshalled = true

    return json.Marshal(&struct {
        
        Url string `json:"url"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

