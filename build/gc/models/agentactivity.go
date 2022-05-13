package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgentactivityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgentactivityDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Agentactivity
type Agentactivity struct { 
    


    // Name
    Name string `json:"name"`


    // Agent
    Agent User `json:"agent"`


    // NumEvaluations
    NumEvaluations int `json:"numEvaluations"`


    // AverageEvaluationScore
    AverageEvaluationScore int `json:"averageEvaluationScore"`


    // NumCriticalEvaluations
    NumCriticalEvaluations int `json:"numCriticalEvaluations"`


    // AverageCriticalScore
    AverageCriticalScore float32 `json:"averageCriticalScore"`


    // HighestEvaluationScore
    HighestEvaluationScore float32 `json:"highestEvaluationScore"`


    // LowestEvaluationScore
    LowestEvaluationScore float32 `json:"lowestEvaluationScore"`


    // HighestCriticalScore
    HighestCriticalScore float32 `json:"highestCriticalScore"`


    // LowestCriticalScore
    LowestCriticalScore float32 `json:"lowestCriticalScore"`


    // AgentEvaluatorActivityList
    AgentEvaluatorActivityList []Agentevaluatoractivity `json:"agentEvaluatorActivityList"`


    // NumEvaluationsWithoutViewPermission
    NumEvaluationsWithoutViewPermission int `json:"numEvaluationsWithoutViewPermission"`


    

}

// String returns a JSON representation of the model
func (o *Agentactivity) String() string {
    
    
    
    
    
    
    
    
    
    
     o.AgentEvaluatorActivityList = []Agentevaluatoractivity{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agentactivity) MarshalJSON() ([]byte, error) {
    type Alias Agentactivity

    if AgentactivityMarshalled {
        return []byte("{}"), nil
    }
    AgentactivityMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Agent User `json:"agent"`
        
        NumEvaluations int `json:"numEvaluations"`
        
        AverageEvaluationScore int `json:"averageEvaluationScore"`
        
        NumCriticalEvaluations int `json:"numCriticalEvaluations"`
        
        AverageCriticalScore float32 `json:"averageCriticalScore"`
        
        HighestEvaluationScore float32 `json:"highestEvaluationScore"`
        
        LowestEvaluationScore float32 `json:"lowestEvaluationScore"`
        
        HighestCriticalScore float32 `json:"highestCriticalScore"`
        
        LowestCriticalScore float32 `json:"lowestCriticalScore"`
        
        AgentEvaluatorActivityList []Agentevaluatoractivity `json:"agentEvaluatorActivityList"`
        
        NumEvaluationsWithoutViewPermission int `json:"numEvaluationsWithoutViewPermission"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        AgentEvaluatorActivityList: []Agentevaluatoractivity{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

