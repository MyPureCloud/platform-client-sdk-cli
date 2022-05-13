package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrustupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrustupdateDud struct { 
    


    

}

// Trustupdate
type Trustupdate struct { 
    // Enabled - If disabled no trustee user will have access, even if they were previously added.
    Enabled bool `json:"enabled"`


    // DateExpired - The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpired time.Time `json:"dateExpired"`

}

// String returns a JSON representation of the model
func (o *Trustupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trustupdate) MarshalJSON() ([]byte, error) {
    type Alias Trustupdate

    if TrustupdateMarshalled {
        return []byte("{}"), nil
    }
    TrustupdateMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        DateExpired time.Time `json:"dateExpired"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

