package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorktypeupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorktypeupdateDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Worktypeupdate
type Worktypeupdate struct { 
    // Name - The name of the Worktype. Valid length between 3 and 256 characters.
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


    // DefaultStatusId - The ID of the default status for Workitems created from the Worktype.
    DefaultStatusId string `json:"defaultStatusId"`


    // SchemaVersion - The version of the Worktypes custom attribute schema. The latest schema version will be used if this property is not set.
    SchemaVersion int `json:"schemaVersion"`


    // DefaultLanguageId - The ID of the default language for Workitems created from the Worktype. Must be a valid UUID.
    DefaultLanguageId string `json:"defaultLanguageId"`


    // DefaultSkillIds - The IDs of the default skills for Workitems created from the Worktype. Must be valid UUIDs. Maximum of 20 IDs
    DefaultSkillIds []string `json:"defaultSkillIds"`


    // DefaultQueueId - The ID of the default queue for Workitems created from the Worktype. Must be a valid UUID.
    DefaultQueueId string `json:"defaultQueueId"`

}

// String returns a JSON representation of the model
func (o *Worktypeupdate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.DefaultSkillIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Worktypeupdate) MarshalJSON() ([]byte, error) {
    type Alias Worktypeupdate

    if WorktypeupdateMarshalled {
        return []byte("{}"), nil
    }
    WorktypeupdateMarshalled = true

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
        
        DefaultStatusId string `json:"defaultStatusId"`
        
        SchemaVersion int `json:"schemaVersion"`
        
        DefaultLanguageId string `json:"defaultLanguageId"`
        
        DefaultSkillIds []string `json:"defaultSkillIds"`
        
        DefaultQueueId string `json:"defaultQueueId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        DefaultSkillIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

