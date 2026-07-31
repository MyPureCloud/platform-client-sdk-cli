package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisionmetricsdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisionmetricsdataDud struct { 
    


    


    


    

}

// Decisionmetricsdata
type Decisionmetricsdata struct { 
    // User - The user associated with the decision metrics
    User Userreference `json:"user"`


    // PerformanceRank - The performance ranking value of the user for decision metrics. The value ranges from 0 to 9999, with the highest value indicating the best performer
    PerformanceRank int `json:"performanceRank"`


    // TieBreakerValue - A numeric tie-breaker value used to resolve ties in performance rankings. Values are sorted in ascending order, with lower values taking precedence
    TieBreakerValue int `json:"tieBreakerValue"`


    // Metadata - The metadata associated to the users decision metric, which will be null if the user has no associated decision metrics
    Metadata Wfmentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Decisionmetricsdata) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisionmetricsdata) MarshalJSON() ([]byte, error) {
    type Alias Decisionmetricsdata

    if DecisionmetricsdataMarshalled {
        return []byte("{}"), nil
    }
    DecisionmetricsdataMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        PerformanceRank int `json:"performanceRank"`
        
        TieBreakerValue int `json:"tieBreakerValue"`
        
        Metadata Wfmentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

