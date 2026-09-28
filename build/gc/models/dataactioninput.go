package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataactioninputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataactioninputDud struct { 
    


    

}

// Dataactioninput
type Dataactioninput struct { 
    // ParameterName - The name of the data action input parameter to map a guide variable to.
    ParameterName string `json:"parameterName"`


    // VariableName - The guide variable whose value will be passed as the input to the paired data action parameter.
    VariableName string `json:"variableName"`

}

// String returns a JSON representation of the model
func (o *Dataactioninput) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataactioninput) MarshalJSON() ([]byte, error) {
    type Alias Dataactioninput

    if DataactioninputMarshalled {
        return []byte("{}"), nil
    }
    DataactioninputMarshalled = true

    return json.Marshal(&struct {
        
        ParameterName string `json:"parameterName"`
        
        VariableName string `json:"variableName"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

