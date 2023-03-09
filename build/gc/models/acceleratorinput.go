package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AcceleratorinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AcceleratorinputDud struct { 
    


    


    

}

// Acceleratorinput - Identity of a CX infrastructure as code accelerator to be run and its inputs
type Acceleratorinput struct { 
    // DryRun - Set this true to test the job without making any changes. Defaults to false.
    DryRun bool `json:"dryRun"`


    // AcceleratorId - Accelerator ID
    AcceleratorId string `json:"acceleratorId"`


    // Parameters - Parameters required for this accelerator
    Parameters []Acceleratorparameter `json:"parameters"`

}

// String returns a JSON representation of the model
func (o *Acceleratorinput) String() string {
    
    
     o.Parameters = []Acceleratorparameter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Acceleratorinput) MarshalJSON() ([]byte, error) {
    type Alias Acceleratorinput

    if AcceleratorinputMarshalled {
        return []byte("{}"), nil
    }
    AcceleratorinputMarshalled = true

    return json.Marshal(&struct {
        
        DryRun bool `json:"dryRun"`
        
        AcceleratorId string `json:"acceleratorId"`
        
        Parameters []Acceleratorparameter `json:"parameters"`
        *Alias
    }{

        


        


        
        Parameters: []Acceleratorparameter{{}},
        

        Alias: (*Alias)(u),
    })
}

