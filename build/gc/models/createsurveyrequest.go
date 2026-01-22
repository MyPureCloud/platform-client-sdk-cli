package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatesurveyrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatesurveyrequestDud struct { 
    


    


    


    

}

// Createsurveyrequest
type Createsurveyrequest struct { 
    // ConversationId - The conversation to create the survey for.
    ConversationId string `json:"conversationId"`


    // SurveyFormContextId - The survey form context to use for the survey.
    SurveyFormContextId string `json:"surveyFormContextId"`


    // AgentId - The agent to associate with the survey. If not specified, the last agent on the conversation will be selected.
    AgentId string `json:"agentId"`


    // QueueId - The queue to associate with the survey. If not specified, the queue associated with the selected agent will be used.
    QueueId string `json:"queueId"`

}

// String returns a JSON representation of the model
func (o *Createsurveyrequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createsurveyrequest) MarshalJSON() ([]byte, error) {
    type Alias Createsurveyrequest

    if CreatesurveyrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatesurveyrequestMarshalled = true

    return json.Marshal(&struct {
        
        ConversationId string `json:"conversationId"`
        
        SurveyFormContextId string `json:"surveyFormContextId"`
        
        AgentId string `json:"agentId"`
        
        QueueId string `json:"queueId"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

