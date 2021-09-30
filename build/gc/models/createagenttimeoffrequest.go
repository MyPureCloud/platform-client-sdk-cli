package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateagenttimeoffrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateagenttimeoffrequestDud struct { 
    


    


    


    


    

}

// Createagenttimeoffrequest
type Createagenttimeoffrequest struct { 
    // ActivityCodeId - The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
    ActivityCodeId string `json:"activityCodeId"`


    // Notes - Notes about the time off request
    Notes string `json:"notes"`


    // FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.
    FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`


    // PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.
    PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`


    // DailyDurationMinutes - The daily duration of this time off request in minutes
    DailyDurationMinutes int `json:"dailyDurationMinutes"`

}

// String returns a JSON representation of the model
func (o *Createagenttimeoffrequest) String() string {
    
    
    
    
    
    
    
    
    
    
     o.FullDayManagementUnitDates = []string{""} 
    
    
    
     o.PartialDayStartDateTimes = []time.Time{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createagenttimeoffrequest) MarshalJSON() ([]byte, error) {
    type Alias Createagenttimeoffrequest

    if CreateagenttimeoffrequestMarshalled {
        return []byte("{}"), nil
    }
    CreateagenttimeoffrequestMarshalled = true

    return json.Marshal(&struct { 
        ActivityCodeId string `json:"activityCodeId"`
        
        Notes string `json:"notes"`
        
        FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`
        
        PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`
        
        DailyDurationMinutes int `json:"dailyDurationMinutes"`
        
        *Alias
    }{
        

        

        

        

        

        
        FullDayManagementUnitDates: []string{""},
        

        

        
        PartialDayStartDateTimes: []time.Time{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

