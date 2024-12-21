package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsagentstateagentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsagentstateagentresponseDud struct { 
    


    


    


    


    


    

}

// Analyticsagentstateagentresponse
type Analyticsagentstateagentresponse struct { 
    // UserId - User Id - only returned if division is covered by agentStateNames permission
    UserId string `json:"userId"`


    // DivisionId - Division Id
    DivisionId string `json:"divisionId"`


    // UserName - User name - only returned if division is covered by agentStateNames permission
    UserName string `json:"userName"`


    // ManagerId - The user that this user reports to
    ManagerId string `json:"managerId"`


    // SessionCount - The count of sessions
    SessionCount int `json:"sessionCount"`


    // Sessions - List of sessions
    Sessions []Analyticsagentstateagentsessionresult `json:"sessions"`

}

// String returns a JSON representation of the model
func (o *Analyticsagentstateagentresponse) String() string {
    
    
    
    
    
     o.Sessions = []Analyticsagentstateagentsessionresult{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsagentstateagentresponse) MarshalJSON() ([]byte, error) {
    type Alias Analyticsagentstateagentresponse

    if AnalyticsagentstateagentresponseMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsagentstateagentresponseMarshalled = true

    return json.Marshal(&struct {
        
        UserId string `json:"userId"`
        
        DivisionId string `json:"divisionId"`
        
        UserName string `json:"userName"`
        
        ManagerId string `json:"managerId"`
        
        SessionCount int `json:"sessionCount"`
        
        Sessions []Analyticsagentstateagentsessionresult `json:"sessions"`
        *Alias
    }{

        


        


        


        


        


        
        Sessions: []Analyticsagentstateagentsessionresult{{}},
        

        Alias: (*Alias)(u),
    })
}

