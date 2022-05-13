package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActioncontractMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActioncontractDud struct { 
    


    

}

// Actioncontract - This resource contains all of the schemas needed to define the inputs and outputs, of a single Action.
type Actioncontract struct { 
    // Output - The output to expect when executing this action.
    Output Actionoutput `json:"output"`


    // Input - The input required when executing this action.
    Input Actioninput `json:"input"`

}

// String returns a JSON representation of the model
func (o *Actioncontract) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actioncontract) MarshalJSON() ([]byte, error) {
    type Alias Actioncontract

    if ActioncontractMarshalled {
        return []byte("{}"), nil
    }
    ActioncontractMarshalled = true

    return json.Marshal(&struct {
        
        Output Actionoutput `json:"output"`
        
        Input Actioninput `json:"input"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

