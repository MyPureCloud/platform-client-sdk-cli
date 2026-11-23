package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnomalyconsensusanalysisMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnomalyconsensusanalysisDud struct { 
    


    


    


    


    


    

}

// Anomalyconsensusanalysis
type Anomalyconsensusanalysis struct { 
    // ConsensusType - The name of the anomaly detection algorithm used for this analysis
    ConsensusType string `json:"consensusType"`


    // Anomalous - True if this method determined the detected value to be anomalous
    Anomalous bool `json:"anomalous"`


    // LowerBound - The lower bound of the normal range for this method. Populated when consensusType is Chronos or ZScore
    LowerBound float64 `json:"lowerBound"`


    // UpperBound - The upper bound of the normal range for this method. Populated when consensusType is Chronos or ZScore
    UpperBound float64 `json:"upperBound"`


    // ExpectedValue - The expected value according to this method. Populated when consensusType is Chronos or ZScore
    ExpectedValue float64 `json:"expectedValue"`


    // Deviation - Percentage value of the deviation from the expected value. Not specified when the expected value is 0
    Deviation float64 `json:"deviation"`

}

// String returns a JSON representation of the model
func (o *Anomalyconsensusanalysis) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Anomalyconsensusanalysis) MarshalJSON() ([]byte, error) {
    type Alias Anomalyconsensusanalysis

    if AnomalyconsensusanalysisMarshalled {
        return []byte("{}"), nil
    }
    AnomalyconsensusanalysisMarshalled = true

    return json.Marshal(&struct {
        
        ConsensusType string `json:"consensusType"`
        
        Anomalous bool `json:"anomalous"`
        
        LowerBound float64 `json:"lowerBound"`
        
        UpperBound float64 `json:"upperBound"`
        
        ExpectedValue float64 `json:"expectedValue"`
        
        Deviation float64 `json:"deviation"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

