package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchcallbackrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchcallbackrequestDud struct { 
    


    


    


    


    


    


    

}

// Patchcallbackrequest
type Patchcallbackrequest struct { 
    // ConversationId - The conversationId.
    ConversationId string `json:"conversationId"`


    // QueueId - The identifier of the queue to be used for the callback.
    QueueId string `json:"queueId"`


    // AgentId - The agentId.
    AgentId string `json:"agentId"`


    // CallbackScheduledTime - The scheduled date-time for the callback. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CallbackScheduledTime time.Time `json:"callbackScheduledTime"`


    // CountryCode - The countryCode
    CountryCode string `json:"countryCode"`


    // CallbackNumbers - The callbackNumbers
    CallbackNumbers []string `json:"callbackNumbers"`


    // ValidateCallbackNumbers - validateCallbackNumbers
    ValidateCallbackNumbers bool `json:"validateCallbackNumbers"`

}

// String returns a JSON representation of the model
func (o *Patchcallbackrequest) String() string {
    
    
    
    
    
     o.CallbackNumbers = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchcallbackrequest) MarshalJSON() ([]byte, error) {
    type Alias Patchcallbackrequest

    if PatchcallbackrequestMarshalled {
        return []byte("{}"), nil
    }
    PatchcallbackrequestMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        QueueId string `json:"queueId"`
        
        AgentId string `json:"agentId"`
        
        CallbackScheduledTime time.Time `json:"callbackScheduledTime"`
        
        CountryCode string `json:"countryCode"`
        
        CallbackNumbers []string `json:"callbackNumbers"`
        
        ValidateCallbackNumbers bool `json:"validateCallbackNumbers"`
        *Alias
    }{

        


        


        


        


        


        
        CallbackNumbers: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

