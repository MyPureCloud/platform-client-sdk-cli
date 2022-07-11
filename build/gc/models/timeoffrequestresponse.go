package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Timeoffrequestresponse
type Timeoffrequestresponse struct { 
    


    // User - The user associated with this time off request
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


    // PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests. Will be not empty if isFullDayRequest == false
    PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`


    // FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone. Will be not empty if isFullDayRequest == true
    FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`


    // DailyDurationMinutes - The daily duration of this time off request in minutes
    DailyDurationMinutes int `json:"dailyDurationMinutes"`


    // Notes - Notes about the time off request
    Notes string `json:"notes"`


    // SubmittedBy - The user who submitted this time off request
    SubmittedBy Userreference `json:"submittedBy"`


    // SubmittedDate - The timestamp when this request was submitted. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    SubmittedDate time.Time `json:"submittedDate"`


    // ReviewedBy - The user who reviewed this time off request
    ReviewedBy Userreference `json:"reviewedBy"`


    // ReviewedDate - The timestamp when this request was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReviewedDate time.Time `json:"reviewedDate"`


    // ModifiedBy - The user who last modified this TimeOffRequestResponse
    ModifiedBy Userreference `json:"modifiedBy"`


    // ModifiedDate - The timestamp when this request was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifiedDate time.Time `json:"modifiedDate"`


    // Metadata - The version metadata of the time off request
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Timeoffrequestresponse) String() string {
    
    
    
    
    
    
    
     o.PartialDayStartDateTimes = []time.Time{{}} 
     o.FullDayManagementUnitDates = []string{""} 
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestresponse) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestresponse

    if TimeoffrequestresponseMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestresponseMarshalled = true

    return json.Marshal(&struct {
        
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
        
        Notes string `json:"notes"`
        
        SubmittedBy Userreference `json:"submittedBy"`
        
        SubmittedDate time.Time `json:"submittedDate"`
        
        ReviewedBy Userreference `json:"reviewedBy"`
        
        ReviewedDate time.Time `json:"reviewedDate"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        ModifiedDate time.Time `json:"modifiedDate"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        PartialDayStartDateTimes: []time.Time{{}},
        


        
        FullDayManagementUnitDates: []string{""},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

