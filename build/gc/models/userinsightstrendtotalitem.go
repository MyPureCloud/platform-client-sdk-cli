package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UserinsightstrendtotalitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UserinsightstrendtotalitemDud struct { 
    

}

// Userinsightstrendtotalitem
type Userinsightstrendtotalitem struct { 
    // Trends - Trends for the metric
    Trends Userinsightstrends `json:"trends"`

}

// String returns a JSON representation of the model
func (o *Userinsightstrendtotalitem) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Userinsightstrendtotalitem) MarshalJSON() ([]byte, error) {
    type Alias Userinsightstrendtotalitem

    if UserinsightstrendtotalitemMarshalled {
        return []byte("{}"), nil
    }
    UserinsightstrendtotalitemMarshalled = true

    return json.Marshal(&struct {
        
        Trends Userinsightstrends `json:"trends"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

