package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OauthlasttokenissuedMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OauthlasttokenissuedDud struct { 
    

}

// Oauthlasttokenissued
type Oauthlasttokenissued struct { 
    // DateIssued - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateIssued time.Time `json:"dateIssued"`

}

// String returns a JSON representation of the model
func (o *Oauthlasttokenissued) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Oauthlasttokenissued) MarshalJSON() ([]byte, error) {
    type Alias Oauthlasttokenissued

    if OauthlasttokenissuedMarshalled {
        return []byte("{}"), nil
    }
    OauthlasttokenissuedMarshalled = true

    return json.Marshal(&struct {
        
        DateIssued time.Time `json:"dateIssued"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

