package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotflowsessionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotflowsessionDud struct { 
    


    


    


    


    


    


    


    


    Conversation Addressableentityref `json:"conversation"`

}

// Botflowsession
type Botflowsession struct { 
    // Id - The ID of the bot session.
    Id string `json:"id"`


    // Flow - The flow associated to this bot session.
    Flow Entity `json:"flow"`


    // Channel - Channel-specific information that describes the message channel/provider.
    Channel Botchannel `json:"channel"`


    // Language - The initial language of operation for the session.
    Language string `json:"language"`


    // EndLanguage - The language of the session at the time the session ended
    EndLanguage string `json:"endLanguage"`


    // BotResult - The reason for session termination.
    BotResult string `json:"botResult"`


    // BotResultCategory - The category of result for the session.
    BotResultCategory string `json:"botResultCategory"`


    // DateCreated - Timestamp indicating when the session was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    

}

// String returns a JSON representation of the model
func (o *Botflowsession) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botflowsession) MarshalJSON() ([]byte, error) {
    type Alias Botflowsession

    if BotflowsessionMarshalled {
        return []byte("{}"), nil
    }
    BotflowsessionMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Flow Entity `json:"flow"`
        
        Channel Botchannel `json:"channel"`
        
        Language string `json:"language"`
        
        EndLanguage string `json:"endLanguage"`
        
        BotResult string `json:"botResult"`
        
        BotResultCategory string `json:"botResultCategory"`
        
        DateCreated time.Time `json:"dateCreated"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

