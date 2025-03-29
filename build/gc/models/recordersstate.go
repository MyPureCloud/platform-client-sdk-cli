package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordersstateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordersstateDud struct { 
    


    


    

}

// Recordersstate
type Recordersstate struct { 
    // AdhocState - Indicates the state of the adhoc recorder.
    AdhocState string `json:"adhocState"`


    // CustomerExperienceState - Indicates the state of the customer experience recorder.
    CustomerExperienceState string `json:"customerExperienceState"`


    // AgentExperienceState - Indicates the state of the agent experience recorder.
    AgentExperienceState string `json:"agentExperienceState"`

}

// String returns a JSON representation of the model
func (o *Recordersstate) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordersstate) MarshalJSON() ([]byte, error) {
    type Alias Recordersstate

    if RecordersstateMarshalled {
        return []byte("{}"), nil
    }
    RecordersstateMarshalled = true

    return json.Marshal(&struct {
        
        AdhocState string `json:"adhocState"`
        
        CustomerExperienceState string `json:"customerExperienceState"`
        
        AgentExperienceState string `json:"agentExperienceState"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

