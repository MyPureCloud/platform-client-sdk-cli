package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CopyschedulebidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CopyschedulebidDud struct { 
    

}

// Copyschedulebid
type Copyschedulebid struct { 
    // Name - The name of the new schedule bid
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Copyschedulebid) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Copyschedulebid) MarshalJSON() ([]byte, error) {
    type Alias Copyschedulebid

    if CopyschedulebidMarshalled {
        return []byte("{}"), nil
    }
    CopyschedulebidMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

