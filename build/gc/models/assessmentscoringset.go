package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssessmentscoringsetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssessmentscoringsetDud struct { 
    TotalScore float32 `json:"totalScore"`


    TotalCriticalScore float32 `json:"totalCriticalScore"`


    TotalNonCriticalScore float32 `json:"totalNonCriticalScore"`


    


    FailureReasons []string `json:"failureReasons"`


    


    


    IsPassed bool `json:"isPassed"`

}

// Assessmentscoringset
type Assessmentscoringset struct { 
    


    


    


    // QuestionGroupScores - The individual scores for each question group
    QuestionGroupScores []Assessmentquestiongroupscore `json:"questionGroupScores"`


    


    // Comments - Comments provided for these answers.
    Comments string `json:"comments"`


    // AgentComments - Comments provided by agent.
    AgentComments string `json:"agentComments"`


    

}

// String returns a JSON representation of the model
func (o *Assessmentscoringset) String() string {
    
    
    
    
    
    
    
    
     o.QuestionGroupScores = []Assessmentquestiongroupscore{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assessmentscoringset) MarshalJSON() ([]byte, error) {
    type Alias Assessmentscoringset

    if AssessmentscoringsetMarshalled {
        return []byte("{}"), nil
    }
    AssessmentscoringsetMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        QuestionGroupScores []Assessmentquestiongroupscore `json:"questionGroupScores"`
        
        
        
        Comments string `json:"comments"`
        
        AgentComments string `json:"agentComments"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        QuestionGroupScores: []Assessmentquestiongroupscore{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

