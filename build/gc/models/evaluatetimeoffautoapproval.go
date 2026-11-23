package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluatetimeoffautoapprovalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluatetimeoffautoapprovalDud struct { 
    


    


    

}

// Evaluatetimeoffautoapproval
type Evaluatetimeoffautoapproval struct { 
    // TimeOffRequestId - The unique identifier of the time-off request to evaluate
    TimeOffRequestId string `json:"timeOffRequestId"`


    // UserId - The unique identifier of the user associated with the time-off request
    UserId string `json:"userId"`


    // Metadata - The version metadata of the time-off request
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Evaluatetimeoffautoapproval) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluatetimeoffautoapproval) MarshalJSON() ([]byte, error) {
    type Alias Evaluatetimeoffautoapproval

    if EvaluatetimeoffautoapprovalMarshalled {
        return []byte("{}"), nil
    }
    EvaluatetimeoffautoapprovalMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffRequestId string `json:"timeOffRequestId"`
        
        UserId string `json:"userId"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

