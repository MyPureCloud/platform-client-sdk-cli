package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateworkplanactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateworkplanactivityDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Createworkplanactivity
type Createworkplanactivity struct { 
    // ActivityCodeId - ID of the activity code associated with this activity
    ActivityCodeId string `json:"activityCodeId"`


    // Description - Description of the activity
    Description string `json:"description"`


    // LengthMinutes - Length of the activity in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // StartTimeIsRelativeToShiftStart - Whether the start time of the activity is relative to the start time of the shift it belongs to
    StartTimeIsRelativeToShiftStart bool `json:"startTimeIsRelativeToShiftStart"`


    // FlexibleStartTime - Whether the start time of the activity is flexible
    FlexibleStartTime bool `json:"flexibleStartTime"`


    // EarliestStartTimeMinutes - Earliest activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == true
    EarliestStartTimeMinutes int `json:"earliestStartTimeMinutes"`


    // LatestStartTimeMinutes - Latest activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == true
    LatestStartTimeMinutes int `json:"latestStartTimeMinutes"`


    // ExactStartTimeMinutes - Exact activity start in offset minutes relative to shift start time if startTimeIsRelativeToShiftStart == true else its based on midnight. Used if flexibleStartTime == false
    ExactStartTimeMinutes int `json:"exactStartTimeMinutes"`


    // StartTimeIncrementMinutes - Increment in offset minutes that would contribute to different possible start times for the activity
    StartTimeIncrementMinutes int `json:"startTimeIncrementMinutes"`


    // CountsAsPaidTime - Whether the activity is paid
    CountsAsPaidTime bool `json:"countsAsPaidTime"`


    // CountsAsContiguousWorkTime - Whether the activity duration is counted towards contiguous work time
    CountsAsContiguousWorkTime bool `json:"countsAsContiguousWorkTime"`


    // MinimumLengthFromShiftStartMinutes - The minimum duration between shift start and shift item (e.g., break or meal) start in minutes
    MinimumLengthFromShiftStartMinutes int `json:"minimumLengthFromShiftStartMinutes"`


    // MinimumLengthFromShiftEndMinutes - The minimum duration between shift item (e.g., break or meal) end and shift end in minutes
    MinimumLengthFromShiftEndMinutes int `json:"minimumLengthFromShiftEndMinutes"`

}

// String returns a JSON representation of the model
func (o *Createworkplanactivity) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createworkplanactivity) MarshalJSON() ([]byte, error) {
    type Alias Createworkplanactivity

    if CreateworkplanactivityMarshalled {
        return []byte("{}"), nil
    }
    CreateworkplanactivityMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Description string `json:"description"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        StartTimeIsRelativeToShiftStart bool `json:"startTimeIsRelativeToShiftStart"`
        
        FlexibleStartTime bool `json:"flexibleStartTime"`
        
        EarliestStartTimeMinutes int `json:"earliestStartTimeMinutes"`
        
        LatestStartTimeMinutes int `json:"latestStartTimeMinutes"`
        
        ExactStartTimeMinutes int `json:"exactStartTimeMinutes"`
        
        StartTimeIncrementMinutes int `json:"startTimeIncrementMinutes"`
        
        CountsAsPaidTime bool `json:"countsAsPaidTime"`
        
        CountsAsContiguousWorkTime bool `json:"countsAsContiguousWorkTime"`
        
        MinimumLengthFromShiftStartMinutes int `json:"minimumLengthFromShiftStartMinutes"`
        
        MinimumLengthFromShiftEndMinutes int `json:"minimumLengthFromShiftEndMinutes"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

