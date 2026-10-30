package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfscheduleactivitymoverequestitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfscheduleactivitymoverequestitemDud struct { 
    


    


    

}

// Selfscheduleactivitymoverequestitem
type Selfscheduleactivitymoverequestitem struct { 
    // OfferJobId - The ID of the self-schedule offers job
    OfferJobId string `json:"offerJobId"`


    // ReferenceKey - Unique identifier from the offer for an activity whose start time the user wants to change
    ReferenceKey string `json:"referenceKey"`


    // NewDate - The new date/time for the activity in ISO-8601 format
    NewDate time.Time `json:"newDate"`

}

// String returns a JSON representation of the model
func (o *Selfscheduleactivitymoverequestitem) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfscheduleactivitymoverequestitem) MarshalJSON() ([]byte, error) {
    type Alias Selfscheduleactivitymoverequestitem

    if SelfscheduleactivitymoverequestitemMarshalled {
        return []byte("{}"), nil
    }
    SelfscheduleactivitymoverequestitemMarshalled = true

    return json.Marshal(&struct {
        
        OfferJobId string `json:"offerJobId"`
        
        ReferenceKey string `json:"referenceKey"`
        
        NewDate time.Time `json:"newDate"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

