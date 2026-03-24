package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasedatedueupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasedatedueupdateDud struct { 
    

}

// Casedatedueupdate
type Casedatedueupdate struct { 
    // DateDue - The due date of the Case. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateDue time.Time `json:"dateDue"`

}

// String returns a JSON representation of the model
func (o *Casedatedueupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casedatedueupdate) MarshalJSON() ([]byte, error) {
    type Alias Casedatedueupdate

    if CasedatedueupdateMarshalled {
        return []byte("{}"), nil
    }
    CasedatedueupdateMarshalled = true

    return json.Marshal(&struct {
        
        DateDue time.Time `json:"dateDue"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

