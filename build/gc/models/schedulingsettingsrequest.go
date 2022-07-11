package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulingsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulingsettingsrequestDud struct { 
    


    


    


    


    

}

// Schedulingsettingsrequest
type Schedulingsettingsrequest struct { 
    // MaxOccupancyPercentForDeferredWork - Max occupancy percent for deferred work
    MaxOccupancyPercentForDeferredWork int `json:"maxOccupancyPercentForDeferredWork"`


    // DefaultShrinkagePercent - Default shrinkage percent for scheduling
    DefaultShrinkagePercent float64 `json:"defaultShrinkagePercent"`


    // ShrinkageOverrides - Shrinkage overrides for scheduling
    ShrinkageOverrides Shrinkageoverrides `json:"shrinkageOverrides"`


    // PlanningPeriod - Planning period settings for scheduling
    PlanningPeriod Valuewrapperplanningperiodsettings `json:"planningPeriod"`


    // StartDayOfWeekend - Start day of weekend for scheduling
    StartDayOfWeekend string `json:"startDayOfWeekend"`

}

// String returns a JSON representation of the model
func (o *Schedulingsettingsrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulingsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Schedulingsettingsrequest

    if SchedulingsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    SchedulingsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        MaxOccupancyPercentForDeferredWork int `json:"maxOccupancyPercentForDeferredWork"`
        
        DefaultShrinkagePercent float64 `json:"defaultShrinkagePercent"`
        
        ShrinkageOverrides Shrinkageoverrides `json:"shrinkageOverrides"`
        
        PlanningPeriod Valuewrapperplanningperiodsettings `json:"planningPeriod"`
        
        StartDayOfWeekend string `json:"startDayOfWeekend"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

