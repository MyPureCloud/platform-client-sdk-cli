package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftasyncresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftasyncresponseDud struct { 
    

}

// Alternativeshiftasyncresponse
type Alternativeshiftasyncresponse struct { 
    // Job - The job related to the async request
    Job Alternativeshiftjobreference `json:"job"`

}

// String returns a JSON representation of the model
func (o *Alternativeshiftasyncresponse) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftasyncresponse) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftasyncresponse

    if AlternativeshiftasyncresponseMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftasyncresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Alternativeshiftjobreference `json:"job"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

