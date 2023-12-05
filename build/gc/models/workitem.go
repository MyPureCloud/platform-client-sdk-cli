package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workitem
type Workitem struct { 
    


    // Name - The name of the Workitem.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // VarType - The Worktype of the Workitem.
    VarType Worktypereference `json:"type"`


    // Description - The description of the Workitem.
    Description string `json:"description"`


    // Language - The language of the Workitem.
    Language Languagereference `json:"language"`


    // Priority - The priority of the Workitem. The valid range is between -25,000,000 and 25,000,000.
    Priority int `json:"priority"`


    // DateCreated - The creation date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The modified date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DateDue - The due date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateDue time.Time `json:"dateDue"`


    // DateExpires - The expiry date of the Workitem. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpires time.Time `json:"dateExpires"`


    // DurationSeconds - The estimated duration in seconds to complete the workitem.
    DurationSeconds int `json:"durationSeconds"`


    // Ttl - The time to live of the Workitem in seconds.
    Ttl int `json:"ttl"`


    // Status - The current Status of the Workitem.
    Status Workitemstatusreference `json:"status"`


    // StatusCategory - The Category of the current Status of the Workitem.
    StatusCategory string `json:"statusCategory"`


    // DateStatusChanged - The State change date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStatusChanged time.Time `json:"dateStatusChanged"`


    // DateClosed - The date the Workitem was closed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateClosed time.Time `json:"dateClosed"`


    // Workbin - The Workbin that contains the Workitem.
    Workbin Workbinreference `json:"workbin"`


    // Reporter - The reporter of the Workitem.
    Reporter Userreferencewithname `json:"reporter"`


    // Assignee - The assignee of the Workitem.
    Assignee Userreferencewithname `json:"assignee"`


    // ExternalContact - The external contact of the Workitem.
    ExternalContact Externalcontactreference `json:"externalContact"`


    // ExternalTag - The external tag of the Workitem.
    ExternalTag string `json:"externalTag"`


    // ModifiedBy - The User who modified the Workitem.
    ModifiedBy Userreference `json:"modifiedBy"`


    // Queue - The Workitems queue.
    Queue Workitemqueuereference `json:"queue"`


    // AssignmentState - The assignment state of the workitem.
    AssignmentState string `json:"assignmentState"`


    // DateAssignmentStateChanged - The assignment state change date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateAssignmentStateChanged time.Time `json:"dateAssignmentStateChanged"`


    // AlertTimeoutSeconds - The duration in seconds before an alert will timeout.
    AlertTimeoutSeconds int `json:"alertTimeoutSeconds"`


    // Skills - The skills of the Workitem.
    Skills []Routingskillreference `json:"skills"`


    // PreferredAgents - The preferred agents of the Workitem.
    PreferredAgents []Userreference `json:"preferredAgents"`


    // AutoStatusTransition - Set it to false to disable auto status transition. By default, it is enabled.
    AutoStatusTransition bool `json:"autoStatusTransition"`


    // Schema - The schema defining the custom fields of the Workitem. The schema is inherited from the Workitems Worktype at creation time.
    Schema Workitemschema `json:"schema"`


    // CustomFields - Custom fields defined in the schema referenced by the Workitem.
    CustomFields map[string]interface{} `json:"customFields"`


    // AutoStatusTransitionDetail - Auto status transition details of Workitem.
    AutoStatusTransitionDetail Autostatustransitiondetail `json:"autoStatusTransitionDetail"`


    // ScoredAgents - A list of scored agents for the Workitem.
    ScoredAgents []Workitemscoredagent `json:"scoredAgents"`


    

}

// String returns a JSON representation of the model
func (o *Workitem) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Skills = []Routingskillreference{{}} 
     o.PreferredAgents = []Userreference{{}} 
    
    
     o.CustomFields = map[string]interface{}{"": Interface{}} 
    
     o.ScoredAgents = []Workitemscoredagent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitem) MarshalJSON() ([]byte, error) {
    type Alias Workitem

    if WorkitemMarshalled {
        return []byte("{}"), nil
    }
    WorkitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        VarType Worktypereference `json:"type"`
        
        Description string `json:"description"`
        
        Language Languagereference `json:"language"`
        
        Priority int `json:"priority"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DateDue time.Time `json:"dateDue"`
        
        DateExpires time.Time `json:"dateExpires"`
        
        DurationSeconds int `json:"durationSeconds"`
        
        Ttl int `json:"ttl"`
        
        Status Workitemstatusreference `json:"status"`
        
        StatusCategory string `json:"statusCategory"`
        
        DateStatusChanged time.Time `json:"dateStatusChanged"`
        
        DateClosed time.Time `json:"dateClosed"`
        
        Workbin Workbinreference `json:"workbin"`
        
        Reporter Userreferencewithname `json:"reporter"`
        
        Assignee Userreferencewithname `json:"assignee"`
        
        ExternalContact Externalcontactreference `json:"externalContact"`
        
        ExternalTag string `json:"externalTag"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        Queue Workitemqueuereference `json:"queue"`
        
        AssignmentState string `json:"assignmentState"`
        
        DateAssignmentStateChanged time.Time `json:"dateAssignmentStateChanged"`
        
        AlertTimeoutSeconds int `json:"alertTimeoutSeconds"`
        
        Skills []Routingskillreference `json:"skills"`
        
        PreferredAgents []Userreference `json:"preferredAgents"`
        
        AutoStatusTransition bool `json:"autoStatusTransition"`
        
        Schema Workitemschema `json:"schema"`
        
        CustomFields map[string]interface{} `json:"customFields"`
        
        AutoStatusTransitionDetail Autostatustransitiondetail `json:"autoStatusTransitionDetail"`
        
        ScoredAgents []Workitemscoredagent `json:"scoredAgents"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Skills: []Routingskillreference{{}},
        


        
        PreferredAgents: []Userreference{{}},
        


        


        


        
        CustomFields: map[string]interface{}{"": Interface{}},
        


        


        
        ScoredAgents: []Workitemscoredagent{{}},
        


        

        Alias: (*Alias)(u),
    })
}

