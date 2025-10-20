package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BucreatetimeoffplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BucreatetimeoffplanrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Bucreatetimeoffplanrequest
type Bucreatetimeoffplanrequest struct { 
    // Name - The name of this time-off plan
    Name string `json:"name"`


    // ActivityCodeIds - The IDs of activity codes to associate with this time-off plan
    ActivityCodeIds []string `json:"activityCodeIds"`


    // AutoApprovalRule - Auto approval rule for this time-off plan. Default is Never
    AutoApprovalRule string `json:"autoApprovalRule"`


    // DaysBeforeStartToExpireFromWaitlist - The number of days before the time-off request start date for when the request will be expired from the waitlist. Default is 0
    DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`


    // AutoPublishApprovedTimeOffRequests - Whether newly approved time-off requests with activity codes associated with this time-off plan should be automatically published to the schedule
    AutoPublishApprovedTimeOffRequests bool `json:"autoPublishApprovedTimeOffRequests"`


    // RestrictedActivityCodeIds - The IDs of non time-off activity codes to check for conflicts in case the auto approval rule specifies checking activity codes. If these activity codes are present in schedule and overlap with the time-off request duration, the request will not be auto approved
    RestrictedActivityCodeIds []string `json:"restrictedActivityCodeIds"`


    // HrisTimeOffType - Time-off type, if this time-off plan is associated with the integration
    HrisTimeOffType Hristimeofftype `json:"hrisTimeOffType"`


    // Enabled - Whether this time-off plan should be used by agents. Default is true
    Enabled bool `json:"enabled"`


    // CountAgainstTimeOffLimits - Whether this time-off plan should count against time-off limits. Default is false
    CountAgainstTimeOffLimits bool `json:"countAgainstTimeOffLimits"`


    // BusinessUnitAssociation - Business unit association, if the time-off plan belongs to a business unit. managementUnitAssociation must not be set if this is populated
    BusinessUnitAssociation Createtimeoffplanbusinessunitassociation `json:"businessUnitAssociation"`


    // ManagementUnitAssociation - Management unit association, if the time-off plan belongs to a management unit. businessUnitAssociation must not be set if this is populated
    ManagementUnitAssociation Createtimeoffplanmanagementunitassociation `json:"managementUnitAssociation"`

}

// String returns a JSON representation of the model
func (o *Bucreatetimeoffplanrequest) String() string {
    
     o.ActivityCodeIds = []string{""} 
    
    
    
     o.RestrictedActivityCodeIds = []string{""} 
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bucreatetimeoffplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Bucreatetimeoffplanrequest

    if BucreatetimeoffplanrequestMarshalled {
        return []byte("{}"), nil
    }
    BucreatetimeoffplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ActivityCodeIds []string `json:"activityCodeIds"`
        
        AutoApprovalRule string `json:"autoApprovalRule"`
        
        DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`
        
        AutoPublishApprovedTimeOffRequests bool `json:"autoPublishApprovedTimeOffRequests"`
        
        RestrictedActivityCodeIds []string `json:"restrictedActivityCodeIds"`
        
        HrisTimeOffType Hristimeofftype `json:"hrisTimeOffType"`
        
        Enabled bool `json:"enabled"`
        
        CountAgainstTimeOffLimits bool `json:"countAgainstTimeOffLimits"`
        
        BusinessUnitAssociation Createtimeoffplanbusinessunitassociation `json:"businessUnitAssociation"`
        
        ManagementUnitAssociation Createtimeoffplanmanagementunitassociation `json:"managementUnitAssociation"`
        *Alias
    }{

        


        
        ActivityCodeIds: []string{""},
        


        


        


        


        
        RestrictedActivityCodeIds: []string{""},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

