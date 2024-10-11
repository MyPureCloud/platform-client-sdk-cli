package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemoncreaterulecreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemoncreaterulecreateDud struct { 
    

}

// Workitemoncreaterulecreate
type Workitemoncreaterulecreate struct { 
    // Name - The name of the rule.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Workitemoncreaterulecreate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemoncreaterulecreate) MarshalJSON() ([]byte, error) {
    type Alias Workitemoncreaterulecreate

    if WorkitemoncreaterulecreateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemoncreaterulecreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

