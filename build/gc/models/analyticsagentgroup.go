package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsagentgroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsagentgroupDud struct { 
    


    

}

// Analyticsagentgroup
type Analyticsagentgroup struct { 
    // AgentGroupId - Conditional group routing agent group identifier
    AgentGroupId string `json:"agentGroupId"`


    // AgentGroupType - Conditional group routing agent group type
    AgentGroupType string `json:"agentGroupType"`

}

// String returns a JSON representation of the model
func (o *Analyticsagentgroup) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsagentgroup) MarshalJSON() ([]byte, error) {
    type Alias Analyticsagentgroup

    if AnalyticsagentgroupMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsagentgroupMarshalled = true

    return json.Marshal(&struct {
        
        AgentGroupId string `json:"agentGroupId"`
        
        AgentGroupType string `json:"agentGroupType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

