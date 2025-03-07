package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Timeoffrequest
type Timeoffrequest struct { 
    // Id - The id of the time off request
    Id string `json:"id"`


    // User - The user that the time off request belongs to
    User Userreference `json:"user"`


    // IsFullDayRequest - Whether this is a full day request (false means partial day)
    IsFullDayRequest bool `json:"isFullDayRequest"`


    // MarkedAsRead - Whether this request has been marked as read by the agent
    MarkedAsRead bool `json:"markedAsRead"`


    // ActivityCodeId - The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
    ActivityCodeId string `json:"activityCodeId"`


    // Paid - Whether this is a paid time off request
    Paid bool `json:"paid"`


    // Status - The status of this time off request
    Status string `json:"status"`


    // Substatus - The substatus of this time off request
    Substatus string `json:"substatus"`


    // PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.  Will be not empty if isFullDayRequest == false
    PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`


    // FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.  Will be not empty if isFullDayRequest == true
    FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`


    // DailyDurationMinutes - The daily duration of this time off request in minutes
    DailyDurationMinutes int `json:"dailyDurationMinutes"`


    // DurationMinutes - Daily durations for each day of this time off request in minutes
    DurationMinutes []int `json:"durationMinutes"`


    // PayableMinutes - Payable minutes for each day of this time off request
    PayableMinutes []int `json:"payableMinutes"`


    // Notes - Notes about the time off request
    Notes string `json:"notes"`


    // SubmittedBy - The user who submitted this time off request. The id may be 'System' if it was an automated process
    SubmittedBy Userreference `json:"submittedBy"`


    // SubmittedDate - The timestamp when this request was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SubmittedDate time.Time `json:"submittedDate"`


    // ReviewedBy - The user who reviewed this time off request. The id may be 'System' if it was an automated process
    ReviewedBy Userreference `json:"reviewedBy"`


    // ReviewedDate - The timestamp when this request was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReviewedDate time.Time `json:"reviewedDate"`


    // SyncVersion - The sync version of this time off request for which the scheduled activity is associated
    SyncVersion int `json:"syncVersion"`


    // Metadata - The version metadata of the time off request
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Timeoffrequest) String() string {
    
    
    
    
    
    
    
    
     o.PartialDayStartDateTimes = []time.Time{{}} 
     o.FullDayManagementUnitDates = []string{""} 
    
     o.DurationMinutes = []int{0} 
     o.PayableMinutes = []int{0} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequest) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequest

    if TimeoffrequestMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        User Userreference `json:"user"`
        
        IsFullDayRequest bool `json:"isFullDayRequest"`
        
        MarkedAsRead bool `json:"markedAsRead"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        
        Status string `json:"status"`
        
        Substatus string `json:"substatus"`
        
        PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`
        
        FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`
        
        DailyDurationMinutes int `json:"dailyDurationMinutes"`
        
        DurationMinutes []int `json:"durationMinutes"`
        
        PayableMinutes []int `json:"payableMinutes"`
        
        Notes string `json:"notes"`
        
        SubmittedBy Userreference `json:"submittedBy"`
        
        SubmittedDate time.Time `json:"submittedDate"`
        
        ReviewedBy Userreference `json:"reviewedBy"`
        
        ReviewedDate time.Time `json:"reviewedDate"`
        
        SyncVersion int `json:"syncVersion"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        PartialDayStartDateTimes: []time.Time{{}},
        


        
        FullDayManagementUnitDates: []string{""},
        


        


        
        DurationMinutes: []int{0},
        


        
        PayableMinutes: []int{0},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

