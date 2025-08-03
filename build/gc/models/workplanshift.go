package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanshiftDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Workplanshift
type Workplanshift struct { 
    // Name - Name of the shift
    Name string `json:"name"`


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


    // ConstrainStopTime - Whether the latest stop time constraint for the shift is enabled.  Deprecated, use constrainLatestStopTime instead
    ConstrainStopTime bool `json:"constrainStopTime"`


    // ConstrainLatestStopTime - Whether the latest stop time constraint for the shift is enabled
    ConstrainLatestStopTime bool `json:"constrainLatestStopTime"`


    // LatestStopTimeMinutesFromMidnight - Latest stop time of the shift defined as offset minutes from midnight. Used if constrainStopTime == true
    LatestStopTimeMinutesFromMidnight int `json:"latestStopTimeMinutesFromMidnight"`


    // ConstrainEarliestStopTime - Whether the earliest stop time constraint for the shift is enabled
    ConstrainEarliestStopTime bool `json:"constrainEarliestStopTime"`


    // EarliestStopTimeMinutesFromMidnight - This is the earliest time a shift can end
    EarliestStopTimeMinutesFromMidnight int `json:"earliestStopTimeMinutesFromMidnight"`


    // StartIncrementMinutes - Increment in offset minutes that would contribute to different possible start times for the shift. Used if flexibleStartTime == true
    StartIncrementMinutes int `json:"startIncrementMinutes"`


    // FlexiblePaidTime - Whether the paid time setting for the shift is flexible
    FlexiblePaidTime bool `json:"flexiblePaidTime"`


    // ExactPaidTimeMinutes - Exact paid time in minutes configured for the shift. Used if flexiblePaidTime == false
    ExactPaidTimeMinutes int `json:"exactPaidTimeMinutes"`


    // MinimumPaidTimeMinutes - Minimum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
    MinimumPaidTimeMinutes int `json:"minimumPaidTimeMinutes"`


    // MaximumPaidTimeMinutes - Maximum paid time in minutes configured for the shift. Used if flexiblePaidTime == true
    MaximumPaidTimeMinutes int `json:"maximumPaidTimeMinutes"`


    // ConstrainContiguousWorkTime - Whether the contiguous time constraint for the shift is enabled
    ConstrainContiguousWorkTime bool `json:"constrainContiguousWorkTime"`


    // MinimumContiguousWorkTimeMinutes - Minimum contiguous time in minutes configured for the shift. Used if constrainContiguousWorkTime == true
    MinimumContiguousWorkTimeMinutes int `json:"minimumContiguousWorkTimeMinutes"`


    // MaximumContiguousWorkTimeMinutes - Maximum contiguous time in minutes configured for the shift. Used if constrainContiguousWorkTime == true
    MaximumContiguousWorkTimeMinutes int `json:"maximumContiguousWorkTimeMinutes"`


    // ConstrainDayOff - Whether day off rule is enabled
    ConstrainDayOff bool `json:"constrainDayOff"`


    // DayOffRule - The day off rule for agents to have next day off or previous day off. used if constrainDayOff = true
    DayOffRule string `json:"dayOffRule"`


    // PlanningPeriodConstraints - Planning period constraints
    PlanningPeriodConstraints Planningperiodshiftconstraints `json:"planningPeriodConstraints"`


    // Activities - Activities configured for this shift
    Activities []Workplanactivity `json:"activities"`


    // Id - ID of the shift. This is required only for the case of updating an existing shift
    Id string `json:"id"`


    // Delete - If marked true for updating an existing shift, the shift will be permanently deleted
    Delete bool `json:"delete"`


    // ValidationId - ID of shift in the context of work plan validation
    ValidationId string `json:"validationId"`

}

// String returns a JSON representation of the model
func (o *Workplanshift) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Activities = []Workplanactivity{{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanshift) MarshalJSON() ([]byte, error) {
    type Alias Workplanshift

    if WorkplanshiftMarshalled {
        return []byte("{}"), nil
    }
    WorkplanshiftMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Days Setwrapperdayofweek `json:"days"`
        
        FlexibleStartTime bool `json:"flexibleStartTime"`
        
        ExactStartTimeMinutesFromMidnight int `json:"exactStartTimeMinutesFromMidnight"`
        
        EarliestStartTimeMinutesFromMidnight int `json:"earliestStartTimeMinutesFromMidnight"`
        
        LatestStartTimeMinutesFromMidnight int `json:"latestStartTimeMinutesFromMidnight"`
        
        ConstrainStopTime bool `json:"constrainStopTime"`
        
        ConstrainLatestStopTime bool `json:"constrainLatestStopTime"`
        
        LatestStopTimeMinutesFromMidnight int `json:"latestStopTimeMinutesFromMidnight"`
        
        ConstrainEarliestStopTime bool `json:"constrainEarliestStopTime"`
        
        EarliestStopTimeMinutesFromMidnight int `json:"earliestStopTimeMinutesFromMidnight"`
        
        StartIncrementMinutes int `json:"startIncrementMinutes"`
        
        FlexiblePaidTime bool `json:"flexiblePaidTime"`
        
        ExactPaidTimeMinutes int `json:"exactPaidTimeMinutes"`
        
        MinimumPaidTimeMinutes int `json:"minimumPaidTimeMinutes"`
        
        MaximumPaidTimeMinutes int `json:"maximumPaidTimeMinutes"`
        
        ConstrainContiguousWorkTime bool `json:"constrainContiguousWorkTime"`
        
        MinimumContiguousWorkTimeMinutes int `json:"minimumContiguousWorkTimeMinutes"`
        
        MaximumContiguousWorkTimeMinutes int `json:"maximumContiguousWorkTimeMinutes"`
        
        ConstrainDayOff bool `json:"constrainDayOff"`
        
        DayOffRule string `json:"dayOffRule"`
        
        PlanningPeriodConstraints Planningperiodshiftconstraints `json:"planningPeriodConstraints"`
        
        Activities []Workplanactivity `json:"activities"`
        
        Id string `json:"id"`
        
        Delete bool `json:"delete"`
        
        ValidationId string `json:"validationId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Activities: []Workplanactivity{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

