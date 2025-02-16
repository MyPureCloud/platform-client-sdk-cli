package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorktypeversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorktypeversionDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Worktypeversion
type Worktypeversion struct { 
    


    // Name - The name of the Worktype.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The description of the Worktype.
    Description string `json:"description"`


    // DateCreated - The creation date of the Worktype. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The modified date of the Worktype. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DefaultWorkbin - The default Workbin for Workitems created from the Worktype.
    DefaultWorkbin Workbinreference `json:"defaultWorkbin"`


    // DefaultStatus - The default status for Workitems created from the Worktype.
    DefaultStatus Workitemstatusreference `json:"defaultStatus"`


    // Statuses - The list of possible statuses for Workitems created from the Worktype.
    Statuses []Workitemstatus `json:"statuses"`


    // DefaultDurationSeconds - The default duration in seconds for Workitems created from the Worktype.
    DefaultDurationSeconds int `json:"defaultDurationSeconds"`


    // DefaultExpirationSeconds - The default expiration time in seconds for Workitems created from the Worktype.
    DefaultExpirationSeconds int `json:"defaultExpirationSeconds"`


    // DefaultDueDurationSeconds - The default due duration in seconds for Workitems created from the Worktype.
    DefaultDueDurationSeconds int `json:"defaultDueDurationSeconds"`


    // DefaultPriority - The default priority for Workitems created from the Worktype. The valid range is between -25,000,000 and 25,000,000.
    DefaultPriority int `json:"defaultPriority"`


    // DefaultLanguage - The default language for Workitems created from the Worktype.
    DefaultLanguage Languagereference `json:"defaultLanguage"`


    // DefaultTtlSeconds - The default time to time to live in seconds for Workitems created from the Worktype.
    DefaultTtlSeconds int `json:"defaultTtlSeconds"`


    // ModifiedBy - The id of the User who modified the Worktype.
    ModifiedBy Userreference `json:"modifiedBy"`


    // DefaultQueue - The default queue for Workitems created from the Worktype.
    DefaultQueue Workitemqueuereference `json:"defaultQueue"`


    // DefaultSkills - The default skills for Workitems created from the Worktype.
    DefaultSkills []Routingskillreference `json:"defaultSkills"`


    // AssignmentEnabled - When set to true, Workitems will be sent to the queue of the Worktype as they are created. Default value is false.
    AssignmentEnabled bool `json:"assignmentEnabled"`


    // Schema - The schema defining the custom attributes for Workitems created from the Worktype.
    Schema Workitemschema `json:"schema"`


    // ServiceLevelTarget - The target service level for Workitems created from the Worktype. The default value is 100.
    ServiceLevelTarget int `json:"serviceLevelTarget"`


    // RuleSettings - Settings for the worktypes rules.
    RuleSettings Workitemrulesettings `json:"ruleSettings"`


    // Flow - The flow associated with the Worktype.
    Flow Workitemflowreference `json:"flow"`


    // DefaultScript - The default script for Workitems created from the Worktype.
    DefaultScript Workitemscriptreference `json:"defaultScript"`


    // Version - Version
    Version int `json:"version"`


    

}

// String returns a JSON representation of the model
func (o *Worktypeversion) String() string {
    
    
    
    
    
    
    
     o.Statuses = []Workitemstatus{{}} 
    
    
    
    
    
    
    
    
     o.DefaultSkills = []Routingskillreference{{}} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Worktypeversion) MarshalJSON() ([]byte, error) {
    type Alias Worktypeversion

    if WorktypeversionMarshalled {
        return []byte("{}"), nil
    }
    WorktypeversionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DefaultWorkbin Workbinreference `json:"defaultWorkbin"`
        
        DefaultStatus Workitemstatusreference `json:"defaultStatus"`
        
        Statuses []Workitemstatus `json:"statuses"`
        
        DefaultDurationSeconds int `json:"defaultDurationSeconds"`
        
        DefaultExpirationSeconds int `json:"defaultExpirationSeconds"`
        
        DefaultDueDurationSeconds int `json:"defaultDueDurationSeconds"`
        
        DefaultPriority int `json:"defaultPriority"`
        
        DefaultLanguage Languagereference `json:"defaultLanguage"`
        
        DefaultTtlSeconds int `json:"defaultTtlSeconds"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        DefaultQueue Workitemqueuereference `json:"defaultQueue"`
        
        DefaultSkills []Routingskillreference `json:"defaultSkills"`
        
        AssignmentEnabled bool `json:"assignmentEnabled"`
        
        Schema Workitemschema `json:"schema"`
        
        ServiceLevelTarget int `json:"serviceLevelTarget"`
        
        RuleSettings Workitemrulesettings `json:"ruleSettings"`
        
        Flow Workitemflowreference `json:"flow"`
        
        DefaultScript Workitemscriptreference `json:"defaultScript"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        
        Statuses: []Workitemstatus{{}},
        


        


        


        


        


        


        


        


        


        
        DefaultSkills: []Routingskillreference{{}},
        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

