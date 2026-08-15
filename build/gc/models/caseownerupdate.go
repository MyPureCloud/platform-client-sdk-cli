package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseownerupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseownerupdateDud struct { 
    

}

// Caseownerupdate
type Caseownerupdate struct { 
    // OwnerId - The ownerId of the Case.
    OwnerId string `json:"ownerId"`

}

// String returns a JSON representation of the model
func (o *Caseownerupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseownerupdate) MarshalJSON() ([]byte, error) {
    type Alias Caseownerupdate

    if CaseownerupdateMarshalled {
        return []byte("{}"), nil
    }
    CaseownerupdateMarshalled = true

    return json.Marshal(&struct {
        
        OwnerId string `json:"ownerId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

