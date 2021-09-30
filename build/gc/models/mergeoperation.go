package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MergeoperationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MergeoperationDud struct { 
    SourceContact Addressableentityref `json:"sourceContact"`


    TargetContact Addressableentityref `json:"targetContact"`


    ResultingContact Addressableentityref `json:"resultingContact"`

}

// Mergeoperation
type Mergeoperation struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Mergeoperation) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mergeoperation) MarshalJSON() ([]byte, error) {
    type Alias Mergeoperation

    if MergeoperationMarshalled {
        return []byte("{}"), nil
    }
    MergeoperationMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

