package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateadherenceadjustmentagentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateadherenceadjustmentagentrequestDud struct { 
    


    


    


    


    

}

// Updateadherenceadjustmentagentrequest
type Updateadherenceadjustmentagentrequest struct { 
    // ReasonCodeId - The ID of the reason code for this adherence adjustment
    ReasonCodeId string `json:"reasonCodeId"`


    // StartDate - The start timestamp of the adherence adjustment in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence adjustment in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Metadata - Version metadata for the adherence adjustment
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // SubmitterNotes - Notes provided by the submitter for this adherence adjustment
    SubmitterNotes string `json:"submitterNotes"`

}

// String returns a JSON representation of the model
func (o *Updateadherenceadjustmentagentrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateadherenceadjustmentagentrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateadherenceadjustmentagentrequest

    if UpdateadherenceadjustmentagentrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateadherenceadjustmentagentrequestMarshalled = true

    return json.Marshal(&struct {
        
        ReasonCodeId string `json:"reasonCodeId"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        SubmitterNotes string `json:"submitterNotes"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

