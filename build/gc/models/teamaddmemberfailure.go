package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamaddmemberfailureMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamaddmemberfailureDud struct { 
    


    Reason string `json:"reason"`

}

// Teamaddmemberfailure
type Teamaddmemberfailure struct { 
    // Id
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Teamaddmemberfailure) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamaddmemberfailure) MarshalJSON() ([]byte, error) {
    type Alias Teamaddmemberfailure

    if TeamaddmemberfailureMarshalled {
        return []byte("{}"), nil
    }
    TeamaddmemberfailureMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

