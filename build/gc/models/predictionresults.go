package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PredictionresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PredictionresultsDud struct { 
    


    


    

}

// Predictionresults
type Predictionresults struct { 
    // Intent - Indicates the media type scope of this estimated wait time
    Intent string `json:"intent"`


    // Formula - Indicates the estimated wait time Formula
    Formula string `json:"formula"`


    // EstimatedWaitTimeSeconds - Estimated wait time in seconds
    EstimatedWaitTimeSeconds int `json:"estimatedWaitTimeSeconds"`

}

// String returns a JSON representation of the model
func (o *Predictionresults) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Predictionresults) MarshalJSON() ([]byte, error) {
    type Alias Predictionresults

    if PredictionresultsMarshalled {
        return []byte("{}"), nil
    }
    PredictionresultsMarshalled = true

    return json.Marshal(&struct { 
        Intent string `json:"intent"`
        
        Formula string `json:"formula"`
        
        EstimatedWaitTimeSeconds int `json:"estimatedWaitTimeSeconds"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

