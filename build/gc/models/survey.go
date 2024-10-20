package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SurveyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SurveyDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Survey
type Survey struct { 
    


    // Name
    Name string `json:"name"`


    // Conversation
    Conversation Conversationreference `json:"conversation"`


    // SurveyForm - Survey form used for this survey.
    SurveyForm Surveyform `json:"surveyForm"`


    // Agent
    Agent Domainentityref `json:"agent"`


    // Status
    Status string `json:"status"`


    // Queue
    Queue Queuereference `json:"queue"`


    // Answers
    Answers Surveyscoringset `json:"answers"`


    // CompletedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CompletedDate time.Time `json:"completedDate"`


    // SurveyErrorDetails - Additional information about what happened when the survey is in Error status.
    SurveyErrorDetails Surveyerrordetails `json:"surveyErrorDetails"`


    // AgentTeam - The team that the agent belongs to
    AgentTeam Team `json:"agentTeam"`


    // SurveyType - Type of the survey
    SurveyType string `json:"surveyType"`


    // MissingRequiredAnswer - True if any of the required questions for the survey form have not been answered. Null if survey is not finished.
    MissingRequiredAnswer bool `json:"missingRequiredAnswer"`


    // Flow - An Architect flow that executed in order to collect the answers for this survey.
    Flow Addressableentityref `json:"flow"`


    

}

// String returns a JSON representation of the model
func (o *Survey) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Survey) MarshalJSON() ([]byte, error) {
    type Alias Survey

    if SurveyMarshalled {
        return []byte("{}"), nil
    }
    SurveyMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Conversation Conversationreference `json:"conversation"`
        
        SurveyForm Surveyform `json:"surveyForm"`
        
        Agent Domainentityref `json:"agent"`
        
        Status string `json:"status"`
        
        Queue Queuereference `json:"queue"`
        
        Answers Surveyscoringset `json:"answers"`
        
        CompletedDate time.Time `json:"completedDate"`
        
        SurveyErrorDetails Surveyerrordetails `json:"surveyErrorDetails"`
        
        AgentTeam Team `json:"agentTeam"`
        
        SurveyType string `json:"surveyType"`
        
        MissingRequiredAnswer bool `json:"missingRequiredAnswer"`
        
        Flow Addressableentityref `json:"flow"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

