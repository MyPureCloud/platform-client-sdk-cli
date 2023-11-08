package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdmintimeoffrequestpatchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdmintimeoffrequestpatchDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Admintimeoffrequestpatch
type Admintimeoffrequestpatch struct { 
    // Status - The status of this time off request
    Status string `json:"status"`


    // ActivityCodeId - The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
    ActivityCodeId string `json:"activityCodeId"`


    // Paid - Whether this is a paid time off request
    Paid bool `json:"paid"`


    // Notes - Notes about the time off request
    Notes string `json:"notes"`


    // FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format. Should be interpreted in the management unit's configured time zone.
    FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`


    // PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.
    PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`


    // DailyDurationMinutes - The daily duration of this time off request in minutes
    DailyDurationMinutes int `json:"dailyDurationMinutes"`


    // DurationMinutes - Daily durations for each day of this time off request in minutes
    DurationMinutes []int `json:"durationMinutes"`


    // PayableMinutes - Payable minutes for each day of this time off request
    PayableMinutes []int `json:"payableMinutes"`


    // Metadata - Version metadata for the time off request
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Admintimeoffrequestpatch) String() string {
    
    
    
    
     o.FullDayManagementUnitDates = []string{""} 
     o.PartialDayStartDateTimes = []time.Time{{}} 
    
     o.DurationMinutes = []int{0} 
     o.PayableMinutes = []int{0} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Admintimeoffrequestpatch) MarshalJSON() ([]byte, error) {
    type Alias Admintimeoffrequestpatch

    if AdmintimeoffrequestpatchMarshalled {
        return []byte("{}"), nil
    }
    AdmintimeoffrequestpatchMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        
        Notes string `json:"notes"`
        
        FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`
        
        PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`
        
        DailyDurationMinutes int `json:"dailyDurationMinutes"`
        
        DurationMinutes []int `json:"durationMinutes"`
        
        PayableMinutes []int `json:"payableMinutes"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        
        FullDayManagementUnitDates: []string{""},
        


        
        PartialDayStartDateTimes: []time.Time{{}},
        


        


        
        DurationMinutes: []int{0},
        


        
        PayableMinutes: []int{0},
        


        

        Alias: (*Alias)(u),
    })
}

