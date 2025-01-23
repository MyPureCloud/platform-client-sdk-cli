package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BufulldaytimeoffmarkerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BufulldaytimeoffmarkerDud struct { 
    


    


    


    


    


    


    


    


    

}

// Bufulldaytimeoffmarker
type Bufulldaytimeoffmarker struct { 
    // BusinessUnitDate - The date of the time off marker, interpreted in the business unit's time zone. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BusinessUnitDate time.Time `json:"businessUnitDate"`


    // LengthMinutes - The length of the time off marker in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Description - The description of the time off marker
    Description string `json:"description"`


    // ActivityCodeId - The ID of the activity code associated with the time off marker
    ActivityCodeId string `json:"activityCodeId"`


    // Paid - Whether the time off marker is paid
    Paid bool `json:"paid"`


    // PayableMinutes - Payable minutes for the time off marker
    PayableMinutes int `json:"payableMinutes"`


    // TimeOffRequestId - The ID of the time off request
    TimeOffRequestId string `json:"timeOffRequestId"`


    // TimeOffRequestSyncVersion - The sync version of the full day time off request for which the scheduled activity is associated
    TimeOffRequestSyncVersion int `json:"timeOffRequestSyncVersion"`


    // Delete - Set to 'true' to delete this time off marker. Will always be null on responses, only has an effect on schedule update
    Delete bool `json:"delete"`

}

// String returns a JSON representation of the model
func (o *Bufulldaytimeoffmarker) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bufulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
    type Alias Bufulldaytimeoffmarker

    if BufulldaytimeoffmarkerMarshalled {
        return []byte("{}"), nil
    }
    BufulldaytimeoffmarkerMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnitDate time.Time `json:"businessUnitDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Description string `json:"description"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        
        PayableMinutes int `json:"payableMinutes"`
        
        TimeOffRequestId string `json:"timeOffRequestId"`
        
        TimeOffRequestSyncVersion int `json:"timeOffRequestSyncVersion"`
        
        Delete bool `json:"delete"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

