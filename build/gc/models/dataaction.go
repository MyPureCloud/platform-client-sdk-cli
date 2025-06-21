package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataactionDud struct { 
    


    


    


    SelfUri string `json:"selfUri"`

}

// Dataaction
type Dataaction struct { 
    // Id - The id of the data action.
    Id string `json:"id"`


    // Label - The label of the GC data action as referenced in the guide instruction.
    Label string `json:"label"`


    // Description - The optional description of the data action.
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Dataaction) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataaction) MarshalJSON() ([]byte, error) {
    type Alias Dataaction

    if DataactionMarshalled {
        return []byte("{}"), nil
    }
    DataactionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Label string `json:"label"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

