package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorktypecreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorktypecreateDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Worktypecreate
type Worktypecreate struct { 
    // Name - The name of the Worktype. Length between 3 and 256 characters.
    Name string `json:"name"`


    // DefaultWorkbinId - The ID of the default Workbin for Workitems created from the Worktype.
    DefaultWorkbinId string `json:"defaultWorkbinId"`


    // DefaultDurationSeconds - The default duration in seconds for Workitems created from the Worktype. Maximum of 365 days.
    DefaultDurationSeconds int `json:"defaultDurationSeconds"`


    // DefaultExpirationSeconds - The default expiration time in seconds for Workitems created from the Worktype. Maximum of 365 days.
    DefaultExpirationSeconds int `json:"defaultExpirationSeconds"`


    // DefaultDueDurationSeconds - The default due duration in seconds for Workitems created from the Worktype. Maximum of 365 days.
    DefaultDueDurationSeconds int `json:"defaultDueDurationSeconds"`


    // DefaultPriority - The default priority for Workitems created from the Worktype. The valid range is between -25,000,000 and 25,000,000.
    DefaultPriority int `json:"defaultPriority"`


    // DefaultTtlSeconds - The default time to time to live in seconds for Workitems created from the Worktype. The valid range is between 1 and 365 days.
    DefaultTtlSeconds int `json:"defaultTtlSeconds"`


    // AssignmentEnabled - When set to true, Workitems will be sent to the queue of the Worktype as they are created. Default value is false.
    AssignmentEnabled bool `json:"assignmentEnabled"`


    // SchemaId - The ID of the custom attribute schema for Workitems created from the Worktype. Must be a valid UUID.
    SchemaId string `json:"schemaId"`


    // Description - The description of the Worktype. Maximum length of 4096 characters.
    Description string `json:"description"`


    // DivisionId - The ID of the division the Worktype belongs to. Defaults to the default Workbin division ID. The Worktype must be in the same division as its default Workbin.
    DivisionId string `json:"divisionId"`


    // DisableDefaultStatusCreation - Set to true to disable default status creation. Default statuses are created with the Worktype by default
    DisableDefaultStatusCreation bool `json:"disableDefaultStatusCreation"`


    // SchemaVersion - The version of the Worktypes custom attribute schema. The latest schema version will be used if this property is not set.
    SchemaVersion int `json:"schemaVersion"`


    // DefaultQueueId - The ID of the default queue for Workitems created from the Worktype. Must be a valid UUID.
    DefaultQueueId string `json:"defaultQueueId"`


    // DefaultLanguageId - The ID of the default language for Workitems created from the Worktype. Must be a valid UUID.
    DefaultLanguageId string `json:"defaultLanguageId"`


    // DefaultSkillIds - The IDs of the default skills for Workitems created from the Worktype. Must be valid UUIDs. Maximum of 20 IDs
    DefaultSkillIds []string `json:"defaultSkillIds"`

}

// String returns a JSON representation of the model
func (o *Worktypecreate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.DefaultSkillIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Worktypecreate) MarshalJSON() ([]byte, error) {
    type Alias Worktypecreate

    if WorktypecreateMarshalled {
        return []byte("{}"), nil
    }
    WorktypecreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DefaultWorkbinId string `json:"defaultWorkbinId"`
        
        DefaultDurationSeconds int `json:"defaultDurationSeconds"`
        
        DefaultExpirationSeconds int `json:"defaultExpirationSeconds"`
        
        DefaultDueDurationSeconds int `json:"defaultDueDurationSeconds"`
        
        DefaultPriority int `json:"defaultPriority"`
        
        DefaultTtlSeconds int `json:"defaultTtlSeconds"`
        
        AssignmentEnabled bool `json:"assignmentEnabled"`
        
        SchemaId string `json:"schemaId"`
        
        Description string `json:"description"`
        
        DivisionId string `json:"divisionId"`
        
        DisableDefaultStatusCreation bool `json:"disableDefaultStatusCreation"`
        
        SchemaVersion int `json:"schemaVersion"`
        
        DefaultQueueId string `json:"defaultQueueId"`
        
        DefaultLanguageId string `json:"defaultLanguageId"`
        
        DefaultSkillIds []string `json:"defaultSkillIds"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        DefaultSkillIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

