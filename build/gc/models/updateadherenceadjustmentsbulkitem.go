package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateadherenceadjustmentsbulkitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateadherenceadjustmentsbulkitemDud struct { 
    


    


    


    

}

// Updateadherenceadjustmentsbulkitem
type Updateadherenceadjustmentsbulkitem struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // ReviewerNotes - Notes provided by the reviewer for this adherence adjustment
    ReviewerNotes string `json:"reviewerNotes"`


    // Status - The new status for the adherence adjustment
    Status string `json:"status"`


    // Metadata - Version metadata for the adherence adjustment
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updateadherenceadjustmentsbulkitem) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateadherenceadjustmentsbulkitem) MarshalJSON() ([]byte, error) {
    type Alias Updateadherenceadjustmentsbulkitem

    if UpdateadherenceadjustmentsbulkitemMarshalled {
        return []byte("{}"), nil
    }
    UpdateadherenceadjustmentsbulkitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        ReviewerNotes string `json:"reviewerNotes"`
        
        Status string `json:"status"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

