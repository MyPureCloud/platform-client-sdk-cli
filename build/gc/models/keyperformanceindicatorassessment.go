package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KeyperformanceindicatorassessmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KeyperformanceindicatorassessmentDud struct { 
    Kpi string `json:"kpi"`


    AssessmentResult string `json:"assessmentResult"`


    Checks []Check `json:"checks"`

}

// Keyperformanceindicatorassessment
type Keyperformanceindicatorassessment struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Keyperformanceindicatorassessment) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Keyperformanceindicatorassessment) MarshalJSON() ([]byte, error) {
    type Alias Keyperformanceindicatorassessment

    if KeyperformanceindicatorassessmentMarshalled {
        return []byte("{}"), nil
    }
    KeyperformanceindicatorassessmentMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

