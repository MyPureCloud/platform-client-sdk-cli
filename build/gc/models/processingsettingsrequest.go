package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProcessingsettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProcessingsettingsrequestDud struct { 
    


    

}

// Processingsettingsrequest
type Processingsettingsrequest struct { 
    // SentimentAnalysisEnabled - Whether sentiment analysis is enabled for the program
    SentimentAnalysisEnabled bool `json:"sentimentAnalysisEnabled"`


    // AgentEmpathyAnalysisEnabled - Whether agent empathy analysis is enabled for the program
    AgentEmpathyAnalysisEnabled bool `json:"agentEmpathyAnalysisEnabled"`

}

// String returns a JSON representation of the model
func (o *Processingsettingsrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Processingsettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Processingsettingsrequest

    if ProcessingsettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    ProcessingsettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        SentimentAnalysisEnabled bool `json:"sentimentAnalysisEnabled"`
        
        AgentEmpathyAnalysisEnabled bool `json:"agentEmpathyAnalysisEnabled"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

