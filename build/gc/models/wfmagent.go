package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmagentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Wfmagent
type Wfmagent struct { 
    


    // User - The user associated with this data
    User Userreference `json:"user"`


    // WorkPlan - The work plan associated with this agent, if applicable
    WorkPlan Workplanreference `json:"workPlan"`


    // WorkPlanRotation - The work plan rotation associated with this agent, if applicable
    WorkPlanRotation Workplanrotationreference `json:"workPlanRotation"`


    // AcceptDirectShiftTrades - Whether the agent accepts direct shift trade requests
    AcceptDirectShiftTrades bool `json:"acceptDirectShiftTrades"`


    // Queues - List of queues to which this agent is capable of handling
    Queues []Queuereference `json:"queues"`


    // Languages - The list of languages this agent is capable of handling
    Languages []Languagereference `json:"languages"`


    // Skills - The list of skills this agent is capable of handling
    Skills []Routingskillreference `json:"skills"`


    // Schedulable - Whether the agent can be included in schedule generation
    Schedulable bool `json:"schedulable"`


    // Metadata - Metadata for this agent
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Wfmagent) String() string {
    
    
    
    
     o.Queues = []Queuereference{{}} 
     o.Languages = []Languagereference{{}} 
     o.Skills = []Routingskillreference{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmagent) MarshalJSON() ([]byte, error) {
    type Alias Wfmagent

    if WfmagentMarshalled {
        return []byte("{}"), nil
    }
    WfmagentMarshalled = true

    return json.Marshal(&struct {
        
        User Userreference `json:"user"`
        
        WorkPlan Workplanreference `json:"workPlan"`
        
        WorkPlanRotation Workplanrotationreference `json:"workPlanRotation"`
        
        AcceptDirectShiftTrades bool `json:"acceptDirectShiftTrades"`
        
        Queues []Queuereference `json:"queues"`
        
        Languages []Languagereference `json:"languages"`
        
        Skills []Routingskillreference `json:"skills"`
        
        Schedulable bool `json:"schedulable"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        
        Queues: []Queuereference{{}},
        


        
        Languages: []Languagereference{{}},
        


        
        Skills: []Routingskillreference{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

