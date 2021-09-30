package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DurationconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DurationconditionDud struct { 
    


    


    


    

}

// Durationcondition
type Durationcondition struct { 
    // DurationTarget
    DurationTarget string `json:"durationTarget"`


    // DurationOperator
    DurationOperator string `json:"durationOperator"`


    // DurationRange
    DurationRange string `json:"durationRange"`


    // DurationMode
    DurationMode string `json:"durationMode"`

}

// String returns a JSON representation of the model
func (o *Durationcondition) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Durationcondition) MarshalJSON() ([]byte, error) {
    type Alias Durationcondition

    if DurationconditionMarshalled {
        return []byte("{}"), nil
    }
    DurationconditionMarshalled = true

    return json.Marshal(&struct { 
        DurationTarget string `json:"durationTarget"`
        
        DurationOperator string `json:"durationOperator"`
        
        DurationRange string `json:"durationRange"`
        
        DurationMode string `json:"durationMode"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

