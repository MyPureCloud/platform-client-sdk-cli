package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorktypedeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorktypedeltaDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Worktypedelta
type Worktypedelta struct { 
    // Name
    Name Workitemsattributechangestring `json:"name"`


    // Description
    Description Workitemsattributechangestring `json:"description"`


    // Statuses
    Statuses Workitemsattributechangelist `json:"statuses"`


    // DefaultWorkbinId
    DefaultWorkbinId Workitemsattributechangestring `json:"defaultWorkbinId"`


    // DefaultDurationSeconds
    DefaultDurationSeconds Workitemsattributechangeinteger `json:"defaultDurationSeconds"`


    // DefaultExpirationSeconds
    DefaultExpirationSeconds Workitemsattributechangeinteger `json:"defaultExpirationSeconds"`


    // DefaultDueDurationSeconds
    DefaultDueDurationSeconds Workitemsattributechangeinteger `json:"defaultDueDurationSeconds"`


    // DefaultPriority
    DefaultPriority Workitemsattributechangeinteger `json:"defaultPriority"`


    // DefaultSkillIds
    DefaultSkillIds Workitemsattributechangelist `json:"defaultSkillIds"`


    // DefaultStatusId
    DefaultStatusId Workitemsattributechangestring `json:"defaultStatusId"`


    // DefaultLanguageId
    DefaultLanguageId Workitemsattributechangestring `json:"defaultLanguageId"`


    // DefaultTtlSeconds
    DefaultTtlSeconds Workitemsattributechangeinteger `json:"defaultTtlSeconds"`


    // AssignmentEnabled
    AssignmentEnabled Workitemsattributechangeboolean `json:"assignmentEnabled"`


    // DefaultQueueId
    DefaultQueueId Workitemsattributechangestring `json:"defaultQueueId"`


    // SchemaId
    SchemaId Workitemsattributechangestring `json:"schemaId"`


    // SchemaVersion
    SchemaVersion Workitemsattributechangestring `json:"schemaVersion"`


    // ServiceLevelTarget
    ServiceLevelTarget Workitemsattributechangeinteger `json:"serviceLevelTarget"`


    // DateModified
    DateModified Workitemsattributechangeinstant `json:"dateModified"`


    // ModifiedBy
    ModifiedBy Workitemsattributechangestring `json:"modifiedBy"`


    // DefaultScriptId
    DefaultScriptId Workitemsattributechangestring `json:"defaultScriptId"`


    // FlowId
    FlowId Workitemsattributechangestring `json:"flowId"`


    // RuleSettings
    RuleSettings Workitemsattributechangeworkitemrulesettings `json:"ruleSettings"`


    // UnassignedDivisionContactsEnabled
    UnassignedDivisionContactsEnabled Workitemsattributechangeboolean `json:"unassignedDivisionContactsEnabled"`

}

// String returns a JSON representation of the model
func (o *Worktypedelta) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Worktypedelta) MarshalJSON() ([]byte, error) {
    type Alias Worktypedelta

    if WorktypedeltaMarshalled {
        return []byte("{}"), nil
    }
    WorktypedeltaMarshalled = true

    return json.Marshal(&struct {
        
        Name Workitemsattributechangestring `json:"name"`
        
        Description Workitemsattributechangestring `json:"description"`
        
        Statuses Workitemsattributechangelist `json:"statuses"`
        
        DefaultWorkbinId Workitemsattributechangestring `json:"defaultWorkbinId"`
        
        DefaultDurationSeconds Workitemsattributechangeinteger `json:"defaultDurationSeconds"`
        
        DefaultExpirationSeconds Workitemsattributechangeinteger `json:"defaultExpirationSeconds"`
        
        DefaultDueDurationSeconds Workitemsattributechangeinteger `json:"defaultDueDurationSeconds"`
        
        DefaultPriority Workitemsattributechangeinteger `json:"defaultPriority"`
        
        DefaultSkillIds Workitemsattributechangelist `json:"defaultSkillIds"`
        
        DefaultStatusId Workitemsattributechangestring `json:"defaultStatusId"`
        
        DefaultLanguageId Workitemsattributechangestring `json:"defaultLanguageId"`
        
        DefaultTtlSeconds Workitemsattributechangeinteger `json:"defaultTtlSeconds"`
        
        AssignmentEnabled Workitemsattributechangeboolean `json:"assignmentEnabled"`
        
        DefaultQueueId Workitemsattributechangestring `json:"defaultQueueId"`
        
        SchemaId Workitemsattributechangestring `json:"schemaId"`
        
        SchemaVersion Workitemsattributechangestring `json:"schemaVersion"`
        
        ServiceLevelTarget Workitemsattributechangeinteger `json:"serviceLevelTarget"`
        
        DateModified Workitemsattributechangeinstant `json:"dateModified"`
        
        ModifiedBy Workitemsattributechangestring `json:"modifiedBy"`
        
        DefaultScriptId Workitemsattributechangestring `json:"defaultScriptId"`
        
        FlowId Workitemsattributechangestring `json:"flowId"`
        
        RuleSettings Workitemsattributechangeworkitemrulesettings `json:"ruleSettings"`
        
        UnassignedDivisionContactsEnabled Workitemsattributechangeboolean `json:"unassignedDivisionContactsEnabled"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

