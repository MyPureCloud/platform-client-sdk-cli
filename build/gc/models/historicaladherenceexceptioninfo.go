package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricaladherenceexceptioninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricaladherenceexceptioninfoDud struct { 
    


    


    


    


    


    


    


    


    

}

// Historicaladherenceexceptioninfo
type Historicaladherenceexceptioninfo struct { 
    // StartOffsetSeconds - Exception start offset in seconds relative to query start time
    StartOffsetSeconds int `json:"startOffsetSeconds"`


    // EndOffsetSeconds - Exception end offset in seconds relative to query start time
    EndOffsetSeconds int `json:"endOffsetSeconds"`


    // ScheduledActivityCodeId - The ID of the scheduled activity code for this user
    ScheduledActivityCodeId string `json:"scheduledActivityCodeId"`


    // ScheduledActivityCategory - Activity for which the user is scheduled
    ScheduledActivityCategory string `json:"scheduledActivityCategory"`


    // ActualActivityCategory - Activity for which the user is actually engaged
    ActualActivityCategory string `json:"actualActivityCategory"`


    // SystemPresence - Actual underlying system presence value
    SystemPresence string `json:"systemPresence"`


    // RoutingStatus - Actual underlying routing status, used to determine whether a user is actually in adherence when OnQueue
    RoutingStatus Routingstatus `json:"routingStatus"`


    // Impact - The impact of the current adherence state for this user
    Impact string `json:"impact"`


    // SecondaryPresenceLookupId - The lookup ID used to retrieve the actual secondary status from map of lookup ID to corresponding secondary presence ID
    SecondaryPresenceLookupId string `json:"secondaryPresenceLookupId"`

}

// String returns a JSON representation of the model
func (o *Historicaladherenceexceptioninfo) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicaladherenceexceptioninfo) MarshalJSON() ([]byte, error) {
    type Alias Historicaladherenceexceptioninfo

    if HistoricaladherenceexceptioninfoMarshalled {
        return []byte("{}"), nil
    }
    HistoricaladherenceexceptioninfoMarshalled = true

    return json.Marshal(&struct { 
        StartOffsetSeconds int `json:"startOffsetSeconds"`
        
        EndOffsetSeconds int `json:"endOffsetSeconds"`
        
        ScheduledActivityCodeId string `json:"scheduledActivityCodeId"`
        
        ScheduledActivityCategory string `json:"scheduledActivityCategory"`
        
        ActualActivityCategory string `json:"actualActivityCategory"`
        
        SystemPresence string `json:"systemPresence"`
        
        RoutingStatus Routingstatus `json:"routingStatus"`
        
        Impact string `json:"impact"`
        
        SecondaryPresenceLookupId string `json:"secondaryPresenceLookupId"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

