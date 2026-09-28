package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateadherenceadjustmentadminrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateadherenceadjustmentadminrequestDud struct { 
    


    


    


    


    


    

}

// Updateadherenceadjustmentadminrequest
type Updateadherenceadjustmentadminrequest struct { 
    // ReasonCodeId - The ID of the reason code for this adherence adjustment
    ReasonCodeId string `json:"reasonCodeId"`


    // StartDate - The start timestamp of the adherence adjustment in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence adjustment in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Metadata - Version metadata for the adherence adjustment
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    // ReviewerNotes - Notes provided by the reviewer for this adherence adjustment
    ReviewerNotes string `json:"reviewerNotes"`


    // Status - The new status for the adherence adjustment
    Status string `json:"status"`

}

// String returns a JSON representation of the model
func (o *Updateadherenceadjustmentadminrequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateadherenceadjustmentadminrequest) MarshalJSON() ([]byte, error) {
    type Alias Updateadherenceadjustmentadminrequest

    if UpdateadherenceadjustmentadminrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateadherenceadjustmentadminrequestMarshalled = true

    return json.Marshal(&struct {
        
        ReasonCodeId string `json:"reasonCodeId"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        ReviewerNotes string `json:"reviewerNotes"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

