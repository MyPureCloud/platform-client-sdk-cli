package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchassociatedvaluefieldMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchassociatedvaluefieldDud struct { 
    


    

}

// Patchassociatedvaluefield
type Patchassociatedvaluefield struct { 
    // DataType - The data type of the value field.
    DataType string `json:"dataType"`


    // Name - The field name for extracting value from event.
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Patchassociatedvaluefield) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchassociatedvaluefield) MarshalJSON() ([]byte, error) {
    type Alias Patchassociatedvaluefield

    if PatchassociatedvaluefieldMarshalled {
        return []byte("{}"), nil
    }
    PatchassociatedvaluefieldMarshalled = true

    return json.Marshal(&struct {
        
        DataType string `json:"dataType"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

