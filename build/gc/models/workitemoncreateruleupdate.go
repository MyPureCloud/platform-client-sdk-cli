package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemoncreateruleupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemoncreateruleupdateDud struct { 
    

}

// Workitemoncreateruleupdate
type Workitemoncreateruleupdate struct { 
    // Name - The name of the rule.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Workitemoncreateruleupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemoncreateruleupdate) MarshalJSON() ([]byte, error) {
    type Alias Workitemoncreateruleupdate

    if WorkitemoncreateruleupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemoncreateruleupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

