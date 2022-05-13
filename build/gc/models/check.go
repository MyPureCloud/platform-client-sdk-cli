package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CheckMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CheckDud struct { 
    Result string `json:"result"`


    VarType string `json:"type"`

}

// Check
type Check struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Check) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Check) MarshalJSON() ([]byte, error) {
    type Alias Check

    if CheckMarshalled {
        return []byte("{}"), nil
    }
    CheckMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

