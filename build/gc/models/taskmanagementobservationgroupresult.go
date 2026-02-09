package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementobservationgroupresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementobservationgroupresultDud struct { 
    


    


    


    

}

// Taskmanagementobservationgroupresult
type Taskmanagementobservationgroupresult struct { 
    // QueueId - The queueId for this group.
    QueueId string `json:"queueId"`


    // TypeId - The typeId for this group. Present when group includes typeId.
    TypeId string `json:"typeId"`


    // AssigneeId - The assigneeId for this group. Present when group includes assigneeId.
    AssigneeId string `json:"assigneeId"`


    // StatusCategory - The status category for this group. Present when group includes statusCategory.
    StatusCategory string `json:"statusCategory"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementobservationgroupresult) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementobservationgroupresult) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementobservationgroupresult

    if TaskmanagementobservationgroupresultMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementobservationgroupresultMarshalled = true

    return json.Marshal(&struct {
        
        QueueId string `json:"queueId"`
        
        TypeId string `json:"typeId"`
        
        AssigneeId string `json:"assigneeId"`
        
        StatusCategory string `json:"statusCategory"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

