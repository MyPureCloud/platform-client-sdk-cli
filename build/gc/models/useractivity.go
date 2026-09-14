package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UseractivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UseractivityDud struct { 
    


    


    


    


    


    

}

// Useractivity
type Useractivity struct { 
    // Id - The ID of the user
    Id string `json:"id"`


    // RoutingStatus - The current routing status of the user
    RoutingStatus Useractivityroutingstatus `json:"routingStatus"`


    // Presence - The current system presence of the user
    Presence Useractivityadherencepresence `json:"presence"`


    // OutOfOffice - The current out of office state of the user
    OutOfOffice Useractivityoutofoffice `json:"outOfOffice"`


    // ActiveQueueIds - The IDs of the queues for which the user is active
    ActiveQueueIds []string `json:"activeQueueIds"`


    // DateActiveQueuesChanged - The date the activeQueueIds list was last modified. For reference only - subject to eventual consistency. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateActiveQueuesChanged time.Time `json:"dateActiveQueuesChanged"`

}

// String returns a JSON representation of the model
func (o *Useractivity) String() string {
    
    
    
    
     o.ActiveQueueIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Useractivity) MarshalJSON() ([]byte, error) {
    type Alias Useractivity

    if UseractivityMarshalled {
        return []byte("{}"), nil
    }
    UseractivityMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        RoutingStatus Useractivityroutingstatus `json:"routingStatus"`
        
        Presence Useractivityadherencepresence `json:"presence"`
        
        OutOfOffice Useractivityoutofoffice `json:"outOfOffice"`
        
        ActiveQueueIds []string `json:"activeQueueIds"`
        
        DateActiveQueuesChanged time.Time `json:"dateActiveQueuesChanged"`
        *Alias
    }{

        


        


        


        


        
        ActiveQueueIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

