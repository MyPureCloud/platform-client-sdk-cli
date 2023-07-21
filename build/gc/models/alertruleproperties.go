package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlertrulepropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlertrulepropertiesDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    VarType string `json:"type"`

}

// Alertruleproperties
type Alertruleproperties struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Alertruleproperties) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alertruleproperties) MarshalJSON() ([]byte, error) {
    type Alias Alertruleproperties

    if AlertrulepropertiesMarshalled {
        return []byte("{}"), nil
    }
    AlertrulepropertiesMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

