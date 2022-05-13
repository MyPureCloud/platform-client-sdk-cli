package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActioncontractinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActioncontractinputDud struct { 
    


    

}

// Actioncontractinput - Contract definition.
type Actioncontractinput struct { 
    // Input - Execution input contract
    Input Postinputcontract `json:"input"`


    // Output - Execution output contract
    Output Postoutputcontract `json:"output"`

}

// String returns a JSON representation of the model
func (o *Actioncontractinput) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actioncontractinput) MarshalJSON() ([]byte, error) {
    type Alias Actioncontractinput

    if ActioncontractinputMarshalled {
        return []byte("{}"), nil
    }
    ActioncontractinputMarshalled = true

    return json.Marshal(&struct {
        
        Input Postinputcontract `json:"input"`
        
        Output Postoutputcontract `json:"output"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

