package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EstimatedwaittimepredictionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EstimatedwaittimepredictionsDud struct { 
    

}

// Estimatedwaittimepredictions
type Estimatedwaittimepredictions struct { 
    // Results - Returned upon a successful estimated wait time request.
    Results []Predictionresults `json:"results"`

}

// String returns a JSON representation of the model
func (o *Estimatedwaittimepredictions) String() string {
    
    
     o.Results = []Predictionresults{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Estimatedwaittimepredictions) MarshalJSON() ([]byte, error) {
    type Alias Estimatedwaittimepredictions

    if EstimatedwaittimepredictionsMarshalled {
        return []byte("{}"), nil
    }
    EstimatedwaittimepredictionsMarshalled = true

    return json.Marshal(&struct { 
        Results []Predictionresults `json:"results"`
        
        *Alias
    }{
        

        
        Results: []Predictionresults{{}},
        

        
        Alias: (*Alias)(u),
    })
}

