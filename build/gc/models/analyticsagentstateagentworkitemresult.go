package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AnalyticsagentstateagentworkitemresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AnalyticsagentstateagentworkitemresultDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Analyticsagentstateagentworkitemresult
type Analyticsagentstateagentworkitemresult struct { 
    // Id - Workitem Id
    Id string `json:"id"`


    // QueueId - The Id of the queue used for routing the workitem
    QueueId string `json:"queueId"`


    // DivisionId - The Id of the division the workitem belongs to
    DivisionId string `json:"divisionId"`


    // TypeId - The Id of the worktype of the workitem
    TypeId string `json:"typeId"`


    // WorkbinId - The Id of the workbin the workitem is placed in
    WorkbinId string `json:"workbinId"`


    // DateDue - Due datetime of the workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateDue time.Time `json:"dateDue"`


    // DateExpires - Expiry datetime of the workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpires time.Time `json:"dateExpires"`


    // StatusId - The Id of the current status of the workitem
    StatusId string `json:"statusId"`


    // StatusCategory - The category of the current status of the workitem
    StatusCategory string `json:"statusCategory"`


    // DateStatusChanged - Datetime the status of the workitem last changed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStatusChanged time.Time `json:"dateStatusChanged"`


    // Priority - Priority of the workitem
    Priority int `json:"priority"`


    // AssignmentState - Assignment state of the workitem
    AssignmentState string `json:"assignmentState"`


    // DateAssignmentStateChanged - Datetime the assignment state of the workitem last changed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAssignmentStateChanged time.Time `json:"dateAssignmentStateChanged"`


    // DateCreated - Creation datetime of the workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`

}

// String returns a JSON representation of the model
func (o *Analyticsagentstateagentworkitemresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Analyticsagentstateagentworkitemresult) MarshalJSON() ([]byte, error) {
    type Alias Analyticsagentstateagentworkitemresult

    if AnalyticsagentstateagentworkitemresultMarshalled {
        return []byte("{}"), nil
    }
    AnalyticsagentstateagentworkitemresultMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        QueueId string `json:"queueId"`
        
        DivisionId string `json:"divisionId"`
        
        TypeId string `json:"typeId"`
        
        WorkbinId string `json:"workbinId"`
        
        DateDue time.Time `json:"dateDue"`
        
        DateExpires time.Time `json:"dateExpires"`
        
        StatusId string `json:"statusId"`
        
        StatusCategory string `json:"statusCategory"`
        
        DateStatusChanged time.Time `json:"dateStatusChanged"`
        
        Priority int `json:"priority"`
        
        AssignmentState string `json:"assignmentState"`
        
        DateAssignmentStateChanged time.Time `json:"dateAssignmentStateChanged"`
        
        DateCreated time.Time `json:"dateCreated"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

