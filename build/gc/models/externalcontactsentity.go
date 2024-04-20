package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExternalcontactsentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExternalcontactsentityDud struct { 
    Id string `json:"id"`

}

// Externalcontactsentity
type Externalcontactsentity struct { 
    

}

// String returns a JSON representation of the model
func (o *Externalcontactsentity) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Externalcontactsentity) MarshalJSON() ([]byte, error) {
    type Alias Externalcontactsentity

    if ExternalcontactsentityMarshalled {
        return []byte("{}"), nil
    }
    ExternalcontactsentityMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

