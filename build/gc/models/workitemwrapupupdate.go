package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemwrapupupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemwrapupupdateDud struct { 
    


    

}

// Workitemwrapupupdate
type Workitemwrapupupdate struct { 
    // Action - Action to be performed for the wrapup code.
    Action string `json:"action"`


    // WrapupCode - The wrapup code which will be added/removed.
    WrapupCode string `json:"wrapupCode"`

}

// String returns a JSON representation of the model
func (o *Workitemwrapupupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemwrapupupdate) MarshalJSON() ([]byte, error) {
    type Alias Workitemwrapupupdate

    if WorkitemwrapupupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemwrapupupdateMarshalled = true

    return json.Marshal(&struct {
        
        Action string `json:"action"`
        
        WrapupCode string `json:"wrapupCode"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

