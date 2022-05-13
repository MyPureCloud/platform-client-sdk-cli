package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationscoringsetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationscoringsetDud struct { 
    


    


    


    


    


    


    


    

}

// Evaluationscoringset
type Evaluationscoringset struct { 
    // TotalScore - Score of all questions
    TotalScore float32 `json:"totalScore"`


    // TotalCriticalScore - Score of only the critical questions
    TotalCriticalScore float32 `json:"totalCriticalScore"`


    // TotalNonCriticalScore - Score of only the non-critical questions
    TotalNonCriticalScore float32 `json:"totalNonCriticalScore"`


    // QuestionGroupScores
    QuestionGroupScores []Evaluationquestiongroupscore `json:"questionGroupScores"`


    // AnyFailedKillQuestions - Indicates that at least one fatal question was answered without having the highest score available for the question
    AnyFailedKillQuestions bool `json:"anyFailedKillQuestions"`


    // Comments - Overall comments from the evaluator
    Comments string `json:"comments"`


    // AgentComments - Comments from the agent while reviewing evaluation results
    AgentComments string `json:"agentComments"`


    // TranscriptTopics - List of topics found within the conversation's transcripts
    TranscriptTopics []Transcripttopic `json:"transcriptTopics"`

}

// String returns a JSON representation of the model
func (o *Evaluationscoringset) String() string {
    
    
    
     o.QuestionGroupScores = []Evaluationquestiongroupscore{{}} 
    
    
    
     o.TranscriptTopics = []Transcripttopic{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationscoringset) MarshalJSON() ([]byte, error) {
    type Alias Evaluationscoringset

    if EvaluationscoringsetMarshalled {
        return []byte("{}"), nil
    }
    EvaluationscoringsetMarshalled = true

    return json.Marshal(&struct {
        
        TotalScore float32 `json:"totalScore"`
        
        TotalCriticalScore float32 `json:"totalCriticalScore"`
        
        TotalNonCriticalScore float32 `json:"totalNonCriticalScore"`
        
        QuestionGroupScores []Evaluationquestiongroupscore `json:"questionGroupScores"`
        
        AnyFailedKillQuestions bool `json:"anyFailedKillQuestions"`
        
        Comments string `json:"comments"`
        
        AgentComments string `json:"agentComments"`
        
        TranscriptTopics []Transcripttopic `json:"transcriptTopics"`
        *Alias
    }{

        


        


        


        
        QuestionGroupScores: []Evaluationquestiongroupscore{{}},
        


        


        


        


        
        TranscriptTopics: []Transcripttopic{{}},
        

        Alias: (*Alias)(u),
    })
}

