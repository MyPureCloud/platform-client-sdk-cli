package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasedescriptionupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasedescriptionupdateDud struct { 
    

}

// Casedescriptionupdate
type Casedescriptionupdate struct { 
    // Description - The description of the Case. Maximum length of 512 characters.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Casedescriptionupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casedescriptionupdate) MarshalJSON() ([]byte, error) {
    type Alias Casedescriptionupdate

    if CasedescriptionupdateMarshalled {
        return []byte("{}"), nil
    }
    CasedescriptionupdateMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

