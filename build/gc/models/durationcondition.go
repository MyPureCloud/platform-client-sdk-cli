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
    // DurationTarget - Whether durationRange holds a single duration to compare using durationOperator (DURATION), or a start/end range (DURATION_RANGE).
    DurationTarget string `json:"durationTarget"`


    // DurationOperator - Comparison operator, one of >, <, >=, <=, == or !=. Required when durationTarget is DURATION, and must not be set when durationTarget is DURATION_RANGE.
    DurationOperator string `json:"durationOperator"`


    // DurationRange - Duration in milliseconds. For DURATION this is a single value, for example 60000. For DURATION_RANGE this is a start and end separated by a slash, for example 60000/120000, where the start is inclusive, the end is inclusive and the start is less than the end.
    DurationRange string `json:"durationRange"`


    // DurationMode - How the range was entered: Between, Over or Under. It describes the range for display and does not change which conversations match; durationTarget, durationOperator and durationRange decide that.
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

