package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluatedshifttraderesponseitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluatedshifttraderesponseitemDud struct { 
    


    

}

// Evaluatedshifttraderesponseitem
type Evaluatedshifttraderesponseitem struct { 
    // Trade - The evaluated shift trade
    Trade Shifttraderesponseitem `json:"trade"`


    // MatchReview - A preview of what the schedule would look like if the shift trade is approved plus any violations, or null if the shift is in a one-sided trade
    MatchReview Shifttradematchreviewresponse `json:"matchReview"`

}

// String returns a JSON representation of the model
func (o *Evaluatedshifttraderesponseitem) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluatedshifttraderesponseitem) MarshalJSON() ([]byte, error) {
    type Alias Evaluatedshifttraderesponseitem

    if EvaluatedshifttraderesponseitemMarshalled {
        return []byte("{}"), nil
    }
    EvaluatedshifttraderesponseitemMarshalled = true

    return json.Marshal(&struct {
        
        Trade Shifttraderesponseitem `json:"trade"`
        
        MatchReview Shifttradematchreviewresponse `json:"matchReview"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

