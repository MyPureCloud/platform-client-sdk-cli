package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantqueueusersjobsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantqueueusersjobsresponseDud struct { 
    


    


    


    

}

// Assistantqueueusersjobsresponse
type Assistantqueueusersjobsresponse struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // SelfUri - The URI for this object.
    SelfUri string `json:"selfUri"`


    // Status - Status of the job.
    Status string `json:"status"`


    // ErrorInfo - Error, if any, related to the job.
    ErrorInfo Assistantqueueusersjoberrorinfo `json:"errorInfo"`

}

// String returns a JSON representation of the model
func (o *Assistantqueueusersjobsresponse) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantqueueusersjobsresponse) MarshalJSON() ([]byte, error) {
    type Alias Assistantqueueusersjobsresponse

    if AssistantqueueusersjobsresponseMarshalled {
        return []byte("{}"), nil
    }
    AssistantqueueusersjobsresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        SelfUri string `json:"selfUri"`
        
        Status string `json:"status"`
        
        ErrorInfo Assistantqueueusersjoberrorinfo `json:"errorInfo"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

