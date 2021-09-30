package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingjobDud struct { 
    Id string `json:"id"`


    


    RecordingJobsQuery Recordingjobsquery `json:"recordingJobsQuery"`


    DateCreated time.Time `json:"dateCreated"`


    TotalConversations int `json:"totalConversations"`


    TotalRecordings int `json:"totalRecordings"`


    TotalSkippedRecordings int `json:"totalSkippedRecordings"`


    TotalFailedRecordings int `json:"totalFailedRecordings"`


    TotalProcessedRecordings int `json:"totalProcessedRecordings"`


    PercentProgress int `json:"percentProgress"`


    ErrorMessage string `json:"errorMessage"`


    FailedRecordings string `json:"failedRecordings"`


    SelfUri string `json:"selfUri"`


    User Addressableentityref `json:"user"`

}

// Recordingjob
type Recordingjob struct { 
    


    // State - The current state of the job.
    State string `json:"state"`


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Recordingjob) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingjob) MarshalJSON() ([]byte, error) {
    type Alias Recordingjob

    if RecordingjobMarshalled {
        return []byte("{}"), nil
    }
    RecordingjobMarshalled = true

    return json.Marshal(&struct { 
        
        
        State string `json:"state"`
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

