package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AcceleratorparameterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AcceleratorparameterDud struct { 
    


    

}

// Acceleratorparameter - Input parameter for a CX infrastructure as code accelerator
type Acceleratorparameter struct { 
    // Name - Parameter Name
    Name string `json:"name"`


    // Value - Parameter Value
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Acceleratorparameter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Acceleratorparameter) MarshalJSON() ([]byte, error) {
    type Alias Acceleratorparameter

    if AcceleratorparameterMarshalled {
        return []byte("{}"), nil
    }
    AcceleratorparameterMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

