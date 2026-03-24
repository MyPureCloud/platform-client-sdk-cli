package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalcontactspatchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalcontactspatchrequestDud struct { 
    

}

// Externalcontactspatchrequest
type Externalcontactspatchrequest struct { 
    // Changes - A list of changes to apply to the provided contact entity
    Changes []Contactspatchchange `json:"changes"`

}

// String returns a JSON representation of the model
func (o *Externalcontactspatchrequest) String() string {
     o.Changes = []Contactspatchchange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalcontactspatchrequest) MarshalJSON() ([]byte, error) {
    type Alias Externalcontactspatchrequest

    if ExternalcontactspatchrequestMarshalled {
        return []byte("{}"), nil
    }
    ExternalcontactspatchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Changes []Contactspatchchange `json:"changes"`
        *Alias
    }{

        
        Changes: []Contactspatchchange{{}},
        

        Alias: (*Alias)(u),
    })
}

