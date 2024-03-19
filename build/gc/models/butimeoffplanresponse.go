package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButimeoffplanresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButimeoffplanresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Butimeoffplanresponse
type Butimeoffplanresponse struct { 
    


    // Name - The name of this time-off plan
    Name string `json:"name"`


    // ActivityCodeIds - The IDs of activity codes associated with this time-off plan
    ActivityCodeIds []string `json:"activityCodeIds"`


    // TimeOffLimits - The IDs of time-off limits associated with this time-off plan
    TimeOffLimits []Butimeofflimitreference `json:"timeOffLimits"`


    // AutoApprovalRule - Auto approval rule for this time-off plan
    AutoApprovalRule string `json:"autoApprovalRule"`


    // DaysBeforeStartToExpireFromWaitlist - The number of days before the time-off request start date for when the request will be expired from the waitlist
    DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`


    // HrisTimeOffType - Time-off type, if this time-off plan is associated with the integration
    HrisTimeOffType Hristimeofftype `json:"hrisTimeOffType"`


    // Enabled - Whether this time-off plan is currently being used by agents
    Enabled bool `json:"enabled"`


    // CountAgainstTimeOffLimits - Whether this time-off plan counts against time-off limits
    CountAgainstTimeOffLimits bool `json:"countAgainstTimeOffLimits"`


    // BusinessUnitAssociation - Business unit association, if the time-off plan belongs to a business unit. managementUnitAssociation must not be set if this is populated
    BusinessUnitAssociation Timeoffplanbusinessunitassociation `json:"businessUnitAssociation"`


    // ManagementUnitAssociation - Management Unit association, if the time-off plan belongs to a management unit. businessUnitAssociation must not be set if this is populated
    ManagementUnitAssociation Timeoffplanmanagementunitassociation `json:"managementUnitAssociation"`


    // Metadata - Version metadata for the time-off plan
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Butimeoffplanresponse) String() string {
    
     o.ActivityCodeIds = []string{""} 
     o.TimeOffLimits = []Butimeofflimitreference{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Butimeoffplanresponse) MarshalJSON() ([]byte, error) {
    type Alias Butimeoffplanresponse

    if ButimeoffplanresponseMarshalled {
        return []byte("{}"), nil
    }
    ButimeoffplanresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ActivityCodeIds []string `json:"activityCodeIds"`
        
        TimeOffLimits []Butimeofflimitreference `json:"timeOffLimits"`
        
        AutoApprovalRule string `json:"autoApprovalRule"`
        
        DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`
        
        HrisTimeOffType Hristimeofftype `json:"hrisTimeOffType"`
        
        Enabled bool `json:"enabled"`
        
        CountAgainstTimeOffLimits bool `json:"countAgainstTimeOffLimits"`
        
        BusinessUnitAssociation Timeoffplanbusinessunitassociation `json:"businessUnitAssociation"`
        
        ManagementUnitAssociation Timeoffplanmanagementunitassociation `json:"managementUnitAssociation"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        
        ActivityCodeIds: []string{""},
        


        
        TimeOffLimits: []Butimeofflimitreference{{}},
        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

