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
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Reportingturn
type Reportingturn struct { 
    // UserInput - The chosen user input associated with this reporting turn.
    UserInput string `json:"userInput"`


    // BotPrompts - The bot prompts associated with this reporting turn.
    BotPrompts []string `json:"botPrompts"`


    // SessionId - The bot session ID that this reporting turn is grouped under.
    SessionId string `json:"sessionId"`


    // Conversation - The conversation details, across potentially multiple Bot Flow sessions.
    Conversation Addressableentityref `json:"conversation"`


    // AskAction - The bot flow 'ask' action associated with this reporting turn (e.g. AskForIntent).
    AskAction Reportingturnaction `json:"askAction"`


    // Intent - The intent and associated slots detected during this reporting turn.
    Intent Reportingturnintent `json:"intent"`


    // Knowledge - The knowledge data captured during this reporting turn.
    Knowledge Reportingturnknowledge `json:"knowledge"`


    // KnowledgeBaseEvents - The knowledge data captured during this reporting turn.
    KnowledgeBaseEvents Reportingturnknowledgeevents `json:"knowledgeBaseEvents"`


    // BotFlowInvocationEvents - The details related to the invoking and invoked flows of the bot flow session.
    BotFlowInvocationEvents []Reportingturnbotflowinvocationevent `json:"botFlowInvocationEvents"`


    // DateInputStarted - Timestamp indicating the datetime when the input came from the user. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateInputStarted time.Time `json:"dateInputStarted"`


    // DateCreated - Timestamp indicating when the original turn was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateCompleted - Timestamp indicating when the original turn was completed. Note: The 'interval' query param uses this timestamp to filter the output. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCompleted time.Time `json:"dateCompleted"`


    // AskActionResult - Result of the bot flow 'ask' action.
    AskActionResult string `json:"askActionResult"`


    // SessionEndDetails - The details related to end of bot flow session.
    SessionEndDetails Sessionenddetails `json:"sessionEndDetails"`


    // ToolCalls - Metadata of tool calls carried out in the reporting turn.
    ToolCalls []Reportingturntoolcall `json:"toolCalls"`


    // GuardrailEvents - Metadata of any Guardrail events that occurred in the reporting turn.
    GuardrailEvents []Reportingturnguardrailevent `json:"guardrailEvents"`


    // ExecutionFrameId - The executionFrameId of the ask action within the reportingTurn.
    ExecutionFrameId string `json:"executionFrameId"`

}

// String returns a JSON representation of the model
func (o *Reportingturn) String() string {
    
     o.BotPrompts = []string{""} 
    
    
    
    
    
    
     o.BotFlowInvocationEvents = []Reportingturnbotflowinvocationevent{{}} 
    
    
    
    
    
     o.ToolCalls = []Reportingturntoolcall{{}} 
     o.GuardrailEvents = []Reportingturnguardrailevent{{}} 
    

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
        
        Conversation Addressableentityref `json:"conversation"`
        
        AskAction Reportingturnaction `json:"askAction"`
        
        Intent Reportingturnintent `json:"intent"`
        
        Knowledge Reportingturnknowledge `json:"knowledge"`
        
        KnowledgeBaseEvents Reportingturnknowledgeevents `json:"knowledgeBaseEvents"`
        
        BotFlowInvocationEvents []Reportingturnbotflowinvocationevent `json:"botFlowInvocationEvents"`
        
        DateInputStarted time.Time `json:"dateInputStarted"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateCompleted time.Time `json:"dateCompleted"`
        
        AskActionResult string `json:"askActionResult"`
        
        SessionEndDetails Sessionenddetails `json:"sessionEndDetails"`
        
        ToolCalls []Reportingturntoolcall `json:"toolCalls"`
        
        GuardrailEvents []Reportingturnguardrailevent `json:"guardrailEvents"`
        
        ExecutionFrameId string `json:"executionFrameId"`
        *Alias
    }{

        


        
        BotPrompts: []string{""},
        


        


        


        


        


        


        


        
        BotFlowInvocationEvents: []Reportingturnbotflowinvocationevent{{}},
        


        


        


        


        


        


        
        ToolCalls: []Reportingturntoolcall{{}},
        


        
        GuardrailEvents: []Reportingturnguardrailevent{{}},
        


        

        Alias: (*Alias)(u),
    })
}

