package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffrequestnotificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffrequestnotificationDud struct { 
    


    


    


    


    


    

}

// Timeoffrequestnotification
type Timeoffrequestnotification struct { 
    // TimeOffRequestId - The ID of this time off request
    TimeOffRequestId string `json:"timeOffRequestId"`


    // User - The user associated with this time off request
    User Userreference `json:"user"`


    // IsFullDayRequest - Whether this is a full day request (false means partial day)
    IsFullDayRequest bool `json:"isFullDayRequest"`


    // Status - The status of this time off request
    Status string `json:"status"`


    // PartialDayStartDateTimes - A set of start date-times in ISO-8601 format for partial day requests.  Will be not empty if isFullDayRequest == false
    PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`


    // FullDayManagementUnitDates - A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.  Will be not empty if isFullDayRequest == true
    FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`

}

// String returns a JSON representation of the model
func (o *Timeoffrequestnotification) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.PartialDayStartDateTimes = []time.Time{{}} 
    
    
    
     o.FullDayManagementUnitDates = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffrequestnotification) MarshalJSON() ([]byte, error) {
    type Alias Timeoffrequestnotification

    if TimeoffrequestnotificationMarshalled {
        return []byte("{}"), nil
    }
    TimeoffrequestnotificationMarshalled = true

    return json.Marshal(&struct { 
        TimeOffRequestId string `json:"timeOffRequestId"`
        
        User Userreference `json:"user"`
        
        IsFullDayRequest bool `json:"isFullDayRequest"`
        
        Status string `json:"status"`
        
        PartialDayStartDateTimes []time.Time `json:"partialDayStartDateTimes"`
        
        FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        PartialDayStartDateTimes: []time.Time{{}},
        

        

        
        FullDayManagementUnitDates: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

