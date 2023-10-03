package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutcomescoreMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutcomescoreDud struct { 
    


    


    


    


    


    


    

}

// Outcomescore
type Outcomescore struct { 
    // Outcome - The outcome that the score was calculated for.
    Outcome Addressableentityref `json:"outcome"`


    // SessionMaxProbability - Represents the max probability reached in the session.
    SessionMaxProbability float32 `json:"sessionMaxProbability"`


    // Probability - Represents the likelihood of a customer reaching or achieving a given outcome.
    Probability float32 `json:"probability"`


    // Percentile - (Deprecated: use the 'quantile' field instead) Represents the predicted probability's percentile score when compared with all other generated probabilities for a given outcome.
    Percentile int `json:"percentile"`


    // SessionMaxPercentile - (Deprecated: use the 'quantile' field instead) Represents the maximum likelihood percentile score reached for a given outcome by the current session.
    SessionMaxPercentile int `json:"sessionMaxPercentile"`


    // Quantile - Represents the quantity of sessions that have a maximum probability less than the predicted probability.
    Quantile float32 `json:"quantile"`


    // SessionMaxQuantile - Represents the quantity of sessions that have a maximum probability less than the predicted session max probability.
    SessionMaxQuantile float32 `json:"sessionMaxQuantile"`

}

// String returns a JSON representation of the model
func (o *Outcomescore) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outcomescore) MarshalJSON() ([]byte, error) {
    type Alias Outcomescore

    if OutcomescoreMarshalled {
        return []byte("{}"), nil
    }
    OutcomescoreMarshalled = true

    return json.Marshal(&struct {
        
        Outcome Addressableentityref `json:"outcome"`
        
        SessionMaxProbability float32 `json:"sessionMaxProbability"`
        
        Probability float32 `json:"probability"`
        
        Percentile int `json:"percentile"`
        
        SessionMaxPercentile int `json:"sessionMaxPercentile"`
        
        Quantile float32 `json:"quantile"`
        
        SessionMaxQuantile float32 `json:"sessionMaxQuantile"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

