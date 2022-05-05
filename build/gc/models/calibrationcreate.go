package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CalibrationcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CalibrationcreateDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Calibrationcreate
type Calibrationcreate struct { 
    


    // Name
    Name string `json:"name"`


    // Calibrator
    Calibrator User `json:"calibrator"`


    // Agent
    Agent User `json:"agent"`


    // Conversation - The conversation to use for the calibration.
    Conversation Conversationreference `json:"conversation"`


    // EvaluationForm
    EvaluationForm Evaluationform `json:"evaluationForm"`


    // ContextId
    ContextId string `json:"contextId"`


    // AverageScore
    AverageScore int `json:"averageScore"`


    // HighScore
    HighScore int `json:"highScore"`


    // LowScore
    LowScore int `json:"lowScore"`


    // CreatedDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreatedDate time.Time `json:"createdDate"`


    // Evaluations
    Evaluations []Evaluation `json:"evaluations"`


    // Evaluators
    Evaluators []User `json:"evaluators"`


    // ScoringIndex
    ScoringIndex Evaluation `json:"scoringIndex"`


    // ExpertEvaluator
    ExpertEvaluator User `json:"expertEvaluator"`


    

}

// String returns a JSON representation of the model
func (o *Calibrationcreate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Evaluations = []Evaluation{{}} 
    
    
    
     o.Evaluators = []User{{}} 
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Calibrationcreate) MarshalJSON() ([]byte, error) {
    type Alias Calibrationcreate

    if CalibrationcreateMarshalled {
        return []byte("{}"), nil
    }
    CalibrationcreateMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Calibrator User `json:"calibrator"`
        
        Agent User `json:"agent"`
        
        Conversation Conversationreference `json:"conversation"`
        
        EvaluationForm Evaluationform `json:"evaluationForm"`
        
        ContextId string `json:"contextId"`
        
        AverageScore int `json:"averageScore"`
        
        HighScore int `json:"highScore"`
        
        LowScore int `json:"lowScore"`
        
        CreatedDate time.Time `json:"createdDate"`
        
        Evaluations []Evaluation `json:"evaluations"`
        
        Evaluators []User `json:"evaluators"`
        
        ScoringIndex Evaluation `json:"scoringIndex"`
        
        ExpertEvaluator User `json:"expertEvaluator"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Evaluations: []Evaluation{{}},
        

        

        
        Evaluators: []User{{}},
        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

