package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatecallrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatecallrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Createcallrequest
type Createcallrequest struct { 
    // PhoneNumber - The phone number to dial.
    PhoneNumber string `json:"phoneNumber"`


    // CallerId - The caller id phone number for this outbound call.
    CallerId string `json:"callerId"`


    // CallerIdName - The caller id name for this outbound call.
    CallerIdName string `json:"callerIdName"`


    // CallFromQueueId - The queue ID to call on behalf of.
    CallFromQueueId string `json:"callFromQueueId"`


    // CallQueueId - The queue ID to call.
    CallQueueId string `json:"callQueueId"`


    // CallUserId - The user ID to call.
    CallUserId string `json:"callUserId"`


    // Priority - The priority to assign to this call (if calling a queue).
    Priority int `json:"priority"`


    // LanguageId - The language skill ID to use for routing this call (if calling a queue).
    LanguageId string `json:"languageId"`


    // RoutingSkillsIds - The skill ID's to use for routing this call (if calling a queue).
    RoutingSkillsIds []string `json:"routingSkillsIds"`


    // ConversationIds - The list of existing call conversations to merge into a new ad-hoc conference.
    ConversationIds []string `json:"conversationIds"`


    // Participants - The list of participants to call to create a new ad-hoc conference.
    Participants []Destination `json:"participants"`


    // UuiData - User to User Information (UUI) data managed by SIP session application.
    UuiData string `json:"uuiData"`

}

// String returns a JSON representation of the model
func (o *Createcallrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.RoutingSkillsIds = []string{""} 
    
    
    
     o.ConversationIds = []string{""} 
    
    
    
     o.Participants = []Destination{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createcallrequest) MarshalJSON() ([]byte, error) {
    type Alias Createcallrequest

    if CreatecallrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatecallrequestMarshalled = true

    return json.Marshal(&struct { 
        PhoneNumber string `json:"phoneNumber"`
        
        CallerId string `json:"callerId"`
        
        CallerIdName string `json:"callerIdName"`
        
        CallFromQueueId string `json:"callFromQueueId"`
        
        CallQueueId string `json:"callQueueId"`
        
        CallUserId string `json:"callUserId"`
        
        Priority int `json:"priority"`
        
        LanguageId string `json:"languageId"`
        
        RoutingSkillsIds []string `json:"routingSkillsIds"`
        
        ConversationIds []string `json:"conversationIds"`
        
        Participants []Destination `json:"participants"`
        
        UuiData string `json:"uuiData"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        RoutingSkillsIds: []string{""},
        

        

        
        ConversationIds: []string{""},
        

        

        
        Participants: []Destination{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

