package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuupdatetimeoffplanrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuupdatetimeoffplanrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    

}

// Buupdatetimeoffplanrequest
type Buupdatetimeoffplanrequest struct { 
    // Name - The name of this time-off plan
    Name string `json:"name"`


    // ActivityCodeIds - The IDs of activity codes to associate with this time-off plan
    ActivityCodeIds Setwrapperstring `json:"activityCodeIds"`


    // AutoApprovalRule - Auto approval rule for this time-off plan
    AutoApprovalRule string `json:"autoApprovalRule"`


    // DaysBeforeStartToExpireFromWaitlist - The number of days before the time-off request start date for when the request will be expired from the waitlist
    DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`


    // AutoPublishApprovedTimeOffRequests - Whether newly approved time-off requests with activity codes associated with this time-off plan should be automatically published to the schedule
    AutoPublishApprovedTimeOffRequests bool `json:"autoPublishApprovedTimeOffRequests"`


    // RestrictedActivityCodeIds - The IDs of non time-off activity codes to check for conflicts in case the auto approval rule specifies checking activity codes. If these activity codes are present in schedule and overlap with the time-off request duration, the request will not be auto approved
    RestrictedActivityCodeIds Setwrapperstring `json:"restrictedActivityCodeIds"`


    // HrisTimeOffType - Time-off type, if this time-off plan is associated with the integration
    HrisTimeOffType Valuewrapperhristimeofftype `json:"hrisTimeOffType"`


    // Enabled - Whether this time-off plan should be used by agents
    Enabled bool `json:"enabled"`


    // CountAgainstTimeOffLimits - Whether this time-off plan should count against time-off limits
    CountAgainstTimeOffLimits bool `json:"countAgainstTimeOffLimits"`


    // BusinessUnitAssociation - Business unit association, if the time-off plan belongs to a business unit. managementUnitAssociation must not be set if this is populated
    BusinessUnitAssociation Updatetimeoffplanbusinessunitassociation `json:"businessUnitAssociation"`


    // ManagementUnitAssociation - Management unit association, if the time-off plan belongs to a management unit. businessUnitAssociation must not be set if this is populated
    ManagementUnitAssociation Updatetimeoffplanmanagementunitassociation `json:"managementUnitAssociation"`


    // Metadata - Version metadata for this time-off plan
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Buupdatetimeoffplanrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buupdatetimeoffplanrequest) MarshalJSON() ([]byte, error) {
    type Alias Buupdatetimeoffplanrequest

    if BuupdatetimeoffplanrequestMarshalled {
        return []byte("{}"), nil
    }
    BuupdatetimeoffplanrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ActivityCodeIds Setwrapperstring `json:"activityCodeIds"`
        
        AutoApprovalRule string `json:"autoApprovalRule"`
        
        DaysBeforeStartToExpireFromWaitlist int `json:"daysBeforeStartToExpireFromWaitlist"`
        
        AutoPublishApprovedTimeOffRequests bool `json:"autoPublishApprovedTimeOffRequests"`
        
        RestrictedActivityCodeIds Setwrapperstring `json:"restrictedActivityCodeIds"`
        
        HrisTimeOffType Valuewrapperhristimeofftype `json:"hrisTimeOffType"`
        
        Enabled bool `json:"enabled"`
        
        CountAgainstTimeOffLimits bool `json:"countAgainstTimeOffLimits"`
        
        BusinessUnitAssociation Updatetimeoffplanbusinessunitassociation `json:"businessUnitAssociation"`
        
        ManagementUnitAssociation Updatetimeoffplanmanagementunitassociation `json:"managementUnitAssociation"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

