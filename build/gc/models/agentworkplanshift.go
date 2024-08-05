package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentworkplanshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentworkplanshiftDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Agentworkplanshift
type Agentworkplanshift struct { 
    // Days - Days of the week applicable for this shift
    Days Setwrapperdayofweek `json:"days"`


    // FlexibleStartTime - Whether the start time of the shift is flexible
    FlexibleStartTime bool `json:"flexibleStartTime"`


    // ExactStartTimeMinutesFromMidnight - Exact start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == false
    ExactStartTimeMinutesFromMidnight int `json:"exactStartTimeMinutesFromMidnight"`


    // EarliestStartTimeMinutesFromMidnight - Earliest start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == true
    EarliestStartTimeMinutesFromMidnight int `json:"earliestStartTimeMinutesFromMidnight"`


    // LatestStartTimeMinutesFromMidnight - Latest start time of the shift defined as offset minutes from midnight. Used if flexibleStartTime == true
    LatestStartTimeMinutesFromMidnight int `json:"latestStartTimeMinutesFromMidnight"`


    // EarliestStopTimeMinutesFromMidnight - This is the earliest time a shift can end
    EarliestStopTimeMinutesFromMidnight int `json:"earliestStopTimeMinutesFromMidnight"`


    // ConstrainLatestStopTime - Whether the latest stop time constraint for the shift is enabled
    ConstrainLatestStopTime bool `json:"constrainLatestStopTime"`


    // LatestStopTimeMinutesFromMidnight - Latest stop time of the shift defined as offset minutes from midnight. Used if constrainStopTime == true
    LatestStopTimeMinutesFromMidnight int `json:"latestStopTimeMinutesFromMidnight"`


    // FlexiblePaidTime - Whether the paid time setting for the shift is flexible
    FlexiblePaidTime bool `json:"flexiblePaidTime"`


    // ExactPaidTimeMinutes - Exact paid time in minutes configured for the shift. Used if flexiblePaidTime == false
    ExactPaidTimeMinutes int `json:"exactPaidTimeMinutes"`


    // MinimumPaidTimeMinutes - Minimum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
    MinimumPaidTimeMinutes int `json:"minimumPaidTimeMinutes"`


    // MaximumPaidTimeMinutes - Maximum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
    MaximumPaidTimeMinutes int `json:"maximumPaidTimeMinutes"`


    // Activities - Activities configured for this shift
    Activities []Agentworkplanactivity `json:"activities"`

}

// String returns a JSON representation of the model
func (o *Agentworkplanshift) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.Activities = []Agentworkplanactivity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentworkplanshift) MarshalJSON() ([]byte, error) {
    type Alias Agentworkplanshift

    if AgentworkplanshiftMarshalled {
        return []byte("{}"), nil
    }
    AgentworkplanshiftMarshalled = true

    return json.Marshal(&struct {
        
        Days Setwrapperdayofweek `json:"days"`
        
        FlexibleStartTime bool `json:"flexibleStartTime"`
        
        ExactStartTimeMinutesFromMidnight int `json:"exactStartTimeMinutesFromMidnight"`
        
        EarliestStartTimeMinutesFromMidnight int `json:"earliestStartTimeMinutesFromMidnight"`
        
        LatestStartTimeMinutesFromMidnight int `json:"latestStartTimeMinutesFromMidnight"`
        
        EarliestStopTimeMinutesFromMidnight int `json:"earliestStopTimeMinutesFromMidnight"`
        
        ConstrainLatestStopTime bool `json:"constrainLatestStopTime"`
        
        LatestStopTimeMinutesFromMidnight int `json:"latestStopTimeMinutesFromMidnight"`
        
        FlexiblePaidTime bool `json:"flexiblePaidTime"`
        
        ExactPaidTimeMinutes int `json:"exactPaidTimeMinutes"`
        
        MinimumPaidTimeMinutes int `json:"minimumPaidTimeMinutes"`
        
        MaximumPaidTimeMinutes int `json:"maximumPaidTimeMinutes"`
        
        Activities []Agentworkplanactivity `json:"activities"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        Activities: []Agentworkplanactivity{{}},
        

        Alias: (*Alias)(u),
    })
}

