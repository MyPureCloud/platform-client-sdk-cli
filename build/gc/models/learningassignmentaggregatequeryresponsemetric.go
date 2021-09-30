package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningassignmentaggregatequeryresponsemetricMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningassignmentaggregatequeryresponsemetricDud struct { 
    


    

}

// Learningassignmentaggregatequeryresponsemetric
type Learningassignmentaggregatequeryresponsemetric struct { 
    // Metric - The metric this applies to
    Metric string `json:"metric"`


    // Stats - The aggregated values for this metric
    Stats Learningassignmentaggregatequeryresponsestats `json:"stats"`

}

// String returns a JSON representation of the model
func (o *Learningassignmentaggregatequeryresponsemetric) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningassignmentaggregatequeryresponsemetric) MarshalJSON() ([]byte, error) {
    type Alias Learningassignmentaggregatequeryresponsemetric

    if LearningassignmentaggregatequeryresponsemetricMarshalled {
        return []byte("{}"), nil
    }
    LearningassignmentaggregatequeryresponsemetricMarshalled = true

    return json.Marshal(&struct { 
        Metric string `json:"metric"`
        
        Stats Learningassignmentaggregatequeryresponsestats `json:"stats"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

