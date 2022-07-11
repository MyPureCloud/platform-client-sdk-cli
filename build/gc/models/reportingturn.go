package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReportingturnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReportingturnDud struct { 
    


    


    


    


    


    


    


    


    


    Conversation Addressableentityref `json:"conversation"`

}

// Reportingturn
type Reportingturn struct { 
    // UserInput - The chosen user input associated with this reporting turn.
    UserInput string `json:"userInput"`


    // BotPrompts - The bot prompts associated with this reporting turn.
    BotPrompts []string `json:"botPrompts"`


    // SessionId - The bot session ID that this reporting turn is grouped under.
    SessionId string `json:"sessionId"`


    // AskAction - The bot flow 'ask' action associated with this reporting turn (e.g. AskForIntent).
    AskAction Reportingturnaction `json:"askAction"`


    // Intent - The intent and associated slots detected during this reporting turn.
    Intent Reportingturnintent `json:"intent"`


    // Knowledge - The knowledge data captured during this reporting turn.
    Knowledge Reportingturnknowledge `json:"knowledge"`


    // DateCreated - Timestamp indicating when the original turn was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // AskActionResult - Result of the bot flow 'ask' action.
    AskActionResult string `json:"askActionResult"`


    // SessionEndDetails - The details related to end of bot flow session.
    SessionEndDetails Sessionenddetails `json:"sessionEndDetails"`


    

}

// String returns a JSON representation of the model
func (o *Reportingturn) String() string {
    
     o.BotPrompts = []string{""} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reportingturn) MarshalJSON() ([]byte, error) {
    type Alias Reportingturn

    if ReportingturnMarshalled {
        return []byte("{}"), nil
    }
    ReportingturnMarshalled = true

    return json.Marshal(&struct {
        
        UserInput string `json:"userInput"`
        
        BotPrompts []string `json:"botPrompts"`
        
        SessionId string `json:"sessionId"`
        
        AskAction Reportingturnaction `json:"askAction"`
        
        Intent Reportingturnintent `json:"intent"`
        
        Knowledge Reportingturnknowledge `json:"knowledge"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        AskActionResult string `json:"askActionResult"`
        
        SessionEndDetails Sessionenddetails `json:"sessionEndDetails"`
        *Alias
    }{

        


        
        BotPrompts: []string{""},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

