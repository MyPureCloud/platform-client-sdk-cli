package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConsulttransfertoqueueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConsulttransfertoqueueDud struct { 
    


    


    


    

}

// Consulttransfertoqueue
type Consulttransfertoqueue struct { 
    // SpeakTo - Determines to whom the initiating participant is speaking. Defaults to DESTINATION
    SpeakTo string `json:"speakTo"`


    // ConsultingUserId - The user ID of the person who wants to talk before completing the transfer. Could be the same of the context user ID
    ConsultingUserId string `json:"consultingUserId"`


    // QueueId - The id of the queue.
    QueueId string `json:"queueId"`


    // QueueName - The name of the queue.
    QueueName string `json:"queueName"`

}

// String returns a JSON representation of the model
func (o *Consulttransfertoqueue) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Consulttransfertoqueue) MarshalJSON() ([]byte, error) {
    type Alias Consulttransfertoqueue

    if ConsulttransfertoqueueMarshalled {
        return []byte("{}"), nil
    }
    ConsulttransfertoqueueMarshalled = true

    return json.Marshal(&struct {
        
        SpeakTo string `json:"speakTo"`
        
        ConsultingUserId string `json:"consultingUserId"`
        
        QueueId string `json:"queueId"`
        
        QueueName string `json:"queueName"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

