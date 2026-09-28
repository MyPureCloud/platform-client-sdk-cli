package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramprocessingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramprocessingsettingsDud struct { 
    


    


    

}

// Programprocessingsettings
type Programprocessingsettings struct { 
    // Program - The ID of the program
    Program Baseprogramentity `json:"program"`


    // SentimentAnalysisEnabled - Whether sentiment analysis is enabled for the program
    SentimentAnalysisEnabled bool `json:"sentimentAnalysisEnabled"`


    // AgentEmpathyAnalysisEnabled - Whether agent empathy analysis is enabled for the program
    AgentEmpathyAnalysisEnabled bool `json:"agentEmpathyAnalysisEnabled"`

}

// String returns a JSON representation of the model
func (o *Programprocessingsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programprocessingsettings) MarshalJSON() ([]byte, error) {
    type Alias Programprocessingsettings

    if ProgramprocessingsettingsMarshalled {
        return []byte("{}"), nil
    }
    ProgramprocessingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        Program Baseprogramentity `json:"program"`
        
        SentimentAnalysisEnabled bool `json:"sentimentAnalysisEnabled"`
        
        AgentEmpathyAnalysisEnabled bool `json:"agentEmpathyAnalysisEnabled"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

