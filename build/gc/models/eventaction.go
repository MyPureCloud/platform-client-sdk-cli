package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventactionDud struct { 
    


    


    


    


    

}

// Eventaction
type Eventaction struct { 
    // Id - ID of the action.
    Id string `json:"id"`


    // State - Current state of the action (e.g. qualified, succeeded, errored).
    State string `json:"state"`


    // MediaType - The media type used to deliver the action (e.g. email, webhook).
    MediaType string `json:"mediaType"`


    // Prompt - Prompt of the action to be displayed/sent to the visitor.
    Prompt string `json:"prompt"`


    // CreatedDate - Timestamp indicating when the action was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`

}

// String returns a JSON representation of the model
func (o *Eventaction) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventaction) MarshalJSON() ([]byte, error) {
    type Alias Eventaction

    if EventactionMarshalled {
        return []byte("{}"), nil
    }
    EventactionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        State string `json:"state"`
        
        MediaType string `json:"mediaType"`
        
        Prompt string `json:"prompt"`
        
        CreatedDate time.Time `json:"createdDate"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

