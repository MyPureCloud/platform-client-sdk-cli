package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasesummaryupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasesummaryupdateDud struct { 
    

}

// Casesummaryupdate
type Casesummaryupdate struct { 
    // Summary - Overview information for the Case. Valid length between 3 and 512 characters.
    Summary string `json:"summary"`

}

// String returns a JSON representation of the model
func (o *Casesummaryupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casesummaryupdate) MarshalJSON() ([]byte, error) {
    type Alias Casesummaryupdate

    if CasesummaryupdateMarshalled {
        return []byte("{}"), nil
    }
    CasesummaryupdateMarshalled = true

    return json.Marshal(&struct {
        
        Summary string `json:"summary"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

