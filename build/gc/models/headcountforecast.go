package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HeadcountforecastMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HeadcountforecastDud struct { 
    


    

}

// Headcountforecast - Headcount interval information for schedule
type Headcountforecast struct { 
    // Required - Headcount information with shrinkage
    Required []Headcountinterval `json:"required"`


    // RequiredWithoutShrinkage - Headcount information without shrinkage
    RequiredWithoutShrinkage []Headcountinterval `json:"requiredWithoutShrinkage"`

}

// String returns a JSON representation of the model
func (o *Headcountforecast) String() string {
    
    
     o.Required = []Headcountinterval{{}} 
    
    
    
     o.RequiredWithoutShrinkage = []Headcountinterval{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Headcountforecast) MarshalJSON() ([]byte, error) {
    type Alias Headcountforecast

    if HeadcountforecastMarshalled {
        return []byte("{}"), nil
    }
    HeadcountforecastMarshalled = true

    return json.Marshal(&struct { 
        Required []Headcountinterval `json:"required"`
        
        RequiredWithoutShrinkage []Headcountinterval `json:"requiredWithoutShrinkage"`
        
        *Alias
    }{
        

        
        Required: []Headcountinterval{{}},
        

        

        
        RequiredWithoutShrinkage: []Headcountinterval{{}},
        

        
        Alias: (*Alias)(u),
    })
}

