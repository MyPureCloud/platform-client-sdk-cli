package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisionmetricsuploaddataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisionmetricsuploaddataDud struct { 
    


    


    

}

// Decisionmetricsuploaddata
type Decisionmetricsuploaddata struct { 
    // UserId - The ID of the user associated with this decision metrics data
    UserId string `json:"userId"`


    // PerformanceRank - The performance ranking value of the user for decision metrics. The value ranges from 0 to 9999, with the highest value indicating the best performer
    PerformanceRank Valuewrapperinteger `json:"performanceRank"`


    // TieBreakerValue - A numeric tie-breaker value used to resolve ties in performance rankings. Values are sorted in ascending order, with lower values taking precedence
    TieBreakerValue Valuewrapperinteger `json:"tieBreakerValue"`

}

// String returns a JSON representation of the model
func (o *Decisionmetricsuploaddata) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisionmetricsuploaddata) MarshalJSON() ([]byte, error) {
    type Alias Decisionmetricsuploaddata

    if DecisionmetricsuploaddataMarshalled {
        return []byte("{}"), nil
    }
    DecisionmetricsuploaddataMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        PerformanceRank Valuewrapperinteger `json:"performanceRank"`
        
        TieBreakerValue Valuewrapperinteger `json:"tieBreakerValue"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

