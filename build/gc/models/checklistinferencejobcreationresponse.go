package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChecklistinferencejobcreationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChecklistinferencejobcreationresponseDud struct { 
    


    

}

// Checklistinferencejobcreationresponse
type Checklistinferencejobcreationresponse struct { 
    // JobId - ID of the inference job.
    JobId string `json:"jobId"`


    // Message - Message with reason in case an inference job is not created.
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Checklistinferencejobcreationresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Checklistinferencejobcreationresponse) MarshalJSON() ([]byte, error) {
    type Alias Checklistinferencejobcreationresponse

    if ChecklistinferencejobcreationresponseMarshalled {
        return []byte("{}"), nil
    }
    ChecklistinferencejobcreationresponseMarshalled = true

    return json.Marshal(&struct {
        
        JobId string `json:"jobId"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

