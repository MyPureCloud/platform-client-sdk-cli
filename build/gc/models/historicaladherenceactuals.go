package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricaladherenceactualsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricaladherenceactualsDud struct { 
    


    


    


    

}

// Historicaladherenceactuals
type Historicaladherenceactuals struct { 
    // ActualActivityCategory - Activity in which the user is actually engaged
    ActualActivityCategory string `json:"actualActivityCategory"`


    // ActualSecondaryPresenceLookupId - The lookup ID used to retrieve the actual secondary status from map of lookup ID to corresponding secondary presence ID
    ActualSecondaryPresenceLookupId string `json:"actualSecondaryPresenceLookupId"`


    // StartOffsetSeconds - Actual start offset in seconds relative to query start time
    StartOffsetSeconds int `json:"startOffsetSeconds"`


    // EndOffsetSeconds - Actual end offset in seconds relative to query start time
    EndOffsetSeconds int `json:"endOffsetSeconds"`

}

// String returns a JSON representation of the model
func (o *Historicaladherenceactuals) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicaladherenceactuals) MarshalJSON() ([]byte, error) {
    type Alias Historicaladherenceactuals

    if HistoricaladherenceactualsMarshalled {
        return []byte("{}"), nil
    }
    HistoricaladherenceactualsMarshalled = true

    return json.Marshal(&struct {
        
        ActualActivityCategory string `json:"actualActivityCategory"`
        
        ActualSecondaryPresenceLookupId string `json:"actualSecondaryPresenceLookupId"`
        
        StartOffsetSeconds int `json:"startOffsetSeconds"`
        
        EndOffsetSeconds int `json:"endOffsetSeconds"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

