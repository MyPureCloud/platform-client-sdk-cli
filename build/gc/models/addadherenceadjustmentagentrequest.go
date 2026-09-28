package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AddadherenceadjustmentagentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AddadherenceadjustmentagentrequestDud struct { 
    


    


    


    

}

// Addadherenceadjustmentagentrequest
type Addadherenceadjustmentagentrequest struct { 
    // ReasonCodeId - The ID of the reason code for this adherence adjustment
    ReasonCodeId string `json:"reasonCodeId"`


    // StartDate - The start timestamp of the adherence adjustment in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence adjustment in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // SubmitterNotes - Notes provided by the submitter for this adherence adjustment
    SubmitterNotes string `json:"submitterNotes"`

}

// String returns a JSON representation of the model
func (o *Addadherenceadjustmentagentrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Addadherenceadjustmentagentrequest) MarshalJSON() ([]byte, error) {
    type Alias Addadherenceadjustmentagentrequest

    if AddadherenceadjustmentagentrequestMarshalled {
        return []byte("{}"), nil
    }
    AddadherenceadjustmentagentrequestMarshalled = true

    return json.Marshal(&struct {
        
        ReasonCodeId string `json:"reasonCodeId"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        SubmitterNotes string `json:"submitterNotes"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

