package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ServerdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ServerdateDud struct { 
    

}

// Serverdate
type Serverdate struct { 
    // CurrentDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CurrentDate time.Time `json:"currentDate"`

}

// String returns a JSON representation of the model
func (o *Serverdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Serverdate) MarshalJSON() ([]byte, error) {
    type Alias Serverdate

    if ServerdateMarshalled {
        return []byte("{}"), nil
    }
    ServerdateMarshalled = true

    return json.Marshal(&struct {
        
        CurrentDate time.Time `json:"currentDate"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

