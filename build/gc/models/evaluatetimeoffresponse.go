package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluatetimeoffresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluatetimeoffresponseDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Evaluatetimeoffresponse
type Evaluatetimeoffresponse struct { 
    


    // Status - The status of the evaluate time-off job
    Status string `json:"status"`


    // VarError - Error information. Set only when status is Error
    VarError Errorbody `json:"error"`


    // Result - The result of the job operation, containing details and statuses for time‑off requests linked to this job
    Result Timeoffrequestbulkevaluateresult `json:"result"`


    

}

// String returns a JSON representation of the model
func (o *Evaluatetimeoffresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluatetimeoffresponse) MarshalJSON() ([]byte, error) {
    type Alias Evaluatetimeoffresponse

    if EvaluatetimeoffresponseMarshalled {
        return []byte("{}"), nil
    }
    EvaluatetimeoffresponseMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        VarError Errorbody `json:"error"`
        
        Result Timeoffrequestbulkevaluateresult `json:"result"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

