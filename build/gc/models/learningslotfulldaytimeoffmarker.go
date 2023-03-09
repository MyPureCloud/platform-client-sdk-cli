package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LearningslotfulldaytimeoffmarkerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LearningslotfulldaytimeoffmarkerDud struct { 
    


    


    


    


    


    

}

// Learningslotfulldaytimeoffmarker
type Learningslotfulldaytimeoffmarker struct { 
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


    // TimeOffRequestId - The ID of the time off request
    TimeOffRequestId string `json:"timeOffRequestId"`

}

// String returns a JSON representation of the model
func (o *Learningslotfulldaytimeoffmarker) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Learningslotfulldaytimeoffmarker) MarshalJSON() ([]byte, error) {
    type Alias Learningslotfulldaytimeoffmarker

    if LearningslotfulldaytimeoffmarkerMarshalled {
        return []byte("{}"), nil
    }
    LearningslotfulldaytimeoffmarkerMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnitDate time.Time `json:"businessUnitDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Description string `json:"description"`
        
        ActivityCodeId string `json:"activityCodeId"`
        
        Paid bool `json:"paid"`
        
        TimeOffRequestId string `json:"timeOffRequestId"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

