package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceexplanationnotificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceexplanationnotificationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Adherenceexplanationnotification
type Adherenceexplanationnotification struct { 
    


    // Agent - The agent for whom the adherence explanation applies
    Agent Userreference `json:"agent"`


    // ManagementUnit - The management unit to which the agent belonged at the time the adherence explanation was submitted
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // BusinessUnit - The business unit to which the agent belonged at the time the adherence explanation was submitted
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // StartDate - The start date of the adherence explanation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence explanation in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Status - The status of the adherence explanation
    Status string `json:"status"`


    // VarType - The type of the adherence explanation
    VarType string `json:"type"`


    // Notes - Notes about the adherence explanation
    Notes string `json:"notes"`


    

}

// String returns a JSON representation of the model
func (o *Adherenceexplanationnotification) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceexplanationnotification) MarshalJSON() ([]byte, error) {
    type Alias Adherenceexplanationnotification

    if AdherenceexplanationnotificationMarshalled {
        return []byte("{}"), nil
    }
    AdherenceexplanationnotificationMarshalled = true

    return json.Marshal(&struct {
        
        Agent Userreference `json:"agent"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Status string `json:"status"`
        
        VarType string `json:"type"`
        
        Notes string `json:"notes"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

