package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdherenceexplanationresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdherenceexplanationresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Adherenceexplanationresponse
type Adherenceexplanationresponse struct { 
    


    // Agent - The agent to whom this adherence explanation applies
    Agent Userreference `json:"agent"`


    // ManagementUnit - The management unit to which the agent belonged at the time the adherence explanation was submitted
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // BusinessUnit - The business unit to which the agent belonged at the time the adherence explanation was submitted
    BusinessUnit Businessunitreference `json:"businessUnit"`


    // VarType - The type of the adherence explanation
    VarType string `json:"type"`


    // Status - The status of the adherence explanation
    Status string `json:"status"`


    // StartDate - The start timestamp of the adherence explanation in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of the adherence explanation in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Notes - Notes about the adherence explanation
    Notes string `json:"notes"`


    // ReviewedBy - The user who reviewed the adherence explanation, if applicable. The id may be 'System' if it was an automated process
    ReviewedBy Userreference `json:"reviewedBy"`


    // ReviewedDate - The timestamp for when the adherence explanation was reviewed, if applicable. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ReviewedDate time.Time `json:"reviewedDate"`


    

}

// String returns a JSON representation of the model
func (o *Adherenceexplanationresponse) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Adherenceexplanationresponse) MarshalJSON() ([]byte, error) {
    type Alias Adherenceexplanationresponse

    if AdherenceexplanationresponseMarshalled {
        return []byte("{}"), nil
    }
    AdherenceexplanationresponseMarshalled = true

    return json.Marshal(&struct {
        
        Agent Userreference `json:"agent"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        BusinessUnit Businessunitreference `json:"businessUnit"`
        
        VarType string `json:"type"`
        
        Status string `json:"status"`
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Notes string `json:"notes"`
        
        ReviewedBy Userreference `json:"reviewedBy"`
        
        ReviewedDate time.Time `json:"reviewedDate"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

