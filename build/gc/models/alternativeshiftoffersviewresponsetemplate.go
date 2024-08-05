package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AlternativeshiftoffersviewresponsetemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AlternativeshiftoffersviewresponsetemplateDud struct { 
    


    


    


    


    


    


    


    

}

// Alternativeshiftoffersviewresponsetemplate
type Alternativeshiftoffersviewresponsetemplate struct { 
    // JobId - The unique identifier of the async list job that created this file
    JobId string `json:"jobId"`


    // BusinessUnitId - The unique identifier of the business unit to which the user (agent) belongs at the time the offer is created
    BusinessUnitId string `json:"businessUnitId"`


    // AgentId - The unique identifier of the agent for whom the offer was made
    AgentId string `json:"agentId"`


    // ManagementUnitId - The unique identifier of the management unit to which the user (agent) belongs at the time the offer is created
    ManagementUnitId string `json:"managementUnitId"`


    // Schedule - The existing schedule information associated with the offer
    Schedule Alternativeshiftschedulelookup `json:"schedule"`


    // OfferWeekDate - The first date of the week for the schedule we are querying in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    OfferWeekDate time.Time `json:"offerWeekDate"`


    // Shifts - The shifts the agent is scheduled for at the time the offer is created
    Shifts []Alternativeshiftagentscheduledshift `json:"shifts"`


    // AlternativeDays - The offered alternative shift days in this week at the time the offer is created
    AlternativeDays []Alternativeshiftagentscheduledshift `json:"alternativeDays"`

}

// String returns a JSON representation of the model
func (o *Alternativeshiftoffersviewresponsetemplate) String() string {
    
    
    
    
    
    
     o.Shifts = []Alternativeshiftagentscheduledshift{{}} 
     o.AlternativeDays = []Alternativeshiftagentscheduledshift{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Alternativeshiftoffersviewresponsetemplate) MarshalJSON() ([]byte, error) {
    type Alias Alternativeshiftoffersviewresponsetemplate

    if AlternativeshiftoffersviewresponsetemplateMarshalled {
        return []byte("{}"), nil
    }
    AlternativeshiftoffersviewresponsetemplateMarshalled = true

    return json.Marshal(&struct {
        
        JobId string `json:"jobId"`
        
        BusinessUnitId string `json:"businessUnitId"`
        
        AgentId string `json:"agentId"`
        
        ManagementUnitId string `json:"managementUnitId"`
        
        Schedule Alternativeshiftschedulelookup `json:"schedule"`
        
        OfferWeekDate time.Time `json:"offerWeekDate"`
        
        Shifts []Alternativeshiftagentscheduledshift `json:"shifts"`
        
        AlternativeDays []Alternativeshiftagentscheduledshift `json:"alternativeDays"`
        *Alias
    }{

        


        


        


        


        


        


        
        Shifts: []Alternativeshiftagentscheduledshift{{}},
        


        
        AlternativeDays: []Alternativeshiftagentscheduledshift{{}},
        

        Alias: (*Alias)(u),
    })
}

