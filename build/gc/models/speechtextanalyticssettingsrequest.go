package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SpeechtextanalyticssettingsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SpeechtextanalyticssettingsrequestDud struct { 
    


    


    


    

}

// Speechtextanalyticssettingsrequest
type Speechtextanalyticssettingsrequest struct { 
    // DefaultProgramId - Setting to choose name for the default program for topic detection
    DefaultProgramId string `json:"defaultProgramId"`


    // ExpectedDialects - Setting to choose expected dialects
    ExpectedDialects []string `json:"expectedDialects"`


    // TextAnalyticsEnabled - Setting to enable/disable text analytics
    TextAnalyticsEnabled bool `json:"textAnalyticsEnabled"`


    // AgentEmpathyEnabled - Setting to enable/disable Agent Empathy setting
    AgentEmpathyEnabled bool `json:"agentEmpathyEnabled"`

}

// String returns a JSON representation of the model
func (o *Speechtextanalyticssettingsrequest) String() string {
    
     o.ExpectedDialects = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Speechtextanalyticssettingsrequest) MarshalJSON() ([]byte, error) {
    type Alias Speechtextanalyticssettingsrequest

    if SpeechtextanalyticssettingsrequestMarshalled {
        return []byte("{}"), nil
    }
    SpeechtextanalyticssettingsrequestMarshalled = true

    return json.Marshal(&struct {
        
        DefaultProgramId string `json:"defaultProgramId"`
        
        ExpectedDialects []string `json:"expectedDialects"`
        
        TextAnalyticsEnabled bool `json:"textAnalyticsEnabled"`
        
        AgentEmpathyEnabled bool `json:"agentEmpathyEnabled"`
        *Alias
    }{

        


        
        ExpectedDialects: []string{""},
        


        


        

        Alias: (*Alias)(u),
    })
}

