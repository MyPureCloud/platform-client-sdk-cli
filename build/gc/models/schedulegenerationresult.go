package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulegenerationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulegenerationresultDud struct { 
    


    


    


    


    

}

// Schedulegenerationresult
type Schedulegenerationresult struct { 
    // Failed - Whether the schedule generation run failed
    Failed bool `json:"failed"`


    // RunId - The ID of the schedule generation run. Reference this when requesting support
    RunId string `json:"runId"`


    // MessageCount - The number of schedule generation messages for this schedule generation run
    MessageCount int `json:"messageCount"`


    // Messages - User facing messages related to the schedule generation run
    Messages []Schedulegenerationmessage `json:"messages"`


    // MessageSeverities - The list of messages by severity in this schedule generation run
    MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`

}

// String returns a JSON representation of the model
func (o *Schedulegenerationresult) String() string {
    
    
    
     o.Messages = []Schedulegenerationmessage{{}} 
     o.MessageSeverities = []Schedulermessagetypeseverity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulegenerationresult) MarshalJSON() ([]byte, error) {
    type Alias Schedulegenerationresult

    if SchedulegenerationresultMarshalled {
        return []byte("{}"), nil
    }
    SchedulegenerationresultMarshalled = true

    return json.Marshal(&struct {
        
        Failed bool `json:"failed"`
        
        RunId string `json:"runId"`
        
        MessageCount int `json:"messageCount"`
        
        Messages []Schedulegenerationmessage `json:"messages"`
        
        MessageSeverities []Schedulermessagetypeseverity `json:"messageSeverities"`
        *Alias
    }{

        


        


        


        
        Messages: []Schedulegenerationmessage{{}},
        


        
        MessageSeverities: []Schedulermessagetypeseverity{{}},
        

        Alias: (*Alias)(u),
    })
}

