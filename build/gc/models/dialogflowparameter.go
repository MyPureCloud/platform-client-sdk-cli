package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowparameterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowparameterDud struct { 
    


    

}

// Dialogflowparameter
type Dialogflowparameter struct { 
    // Name - The parameter name
    Name string `json:"name"`


    // VarType - The parameter type
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Dialogflowparameter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowparameter) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowparameter

    if DialogflowparameterMarshalled {
        return []byte("{}"), nil
    }
    DialogflowparameterMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

