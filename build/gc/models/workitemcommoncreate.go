package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemcommoncreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemcommoncreateDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Workitemcommoncreate
type Workitemcommoncreate struct { 
    // Name - The name of the Workitem. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // Priority - The priority of the Workitem. The valid range is between -25,000,000 and 25,000,000.
    Priority int `json:"priority"`


    // DateDue - The due date of the Workitem. Can not be greater than 365 days from the current time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateDue time.Time `json:"dateDue"`


    // DateExpires - The expiry date of the Workitem. Can not be greater than 365 days from the current time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateExpires time.Time `json:"dateExpires"`


    // DurationSeconds - The estimated duration in seconds to complete the Workitem. Maximum of 365 days.
    DurationSeconds int `json:"durationSeconds"`


    // Ttl - The epoch timestamp in seconds specifying the time to live for the Workitem. Can not be greater than 365 days from the current time.
    Ttl int `json:"ttl"`


    // StatusId - The ID of the Status of the Workitem.
    StatusId string `json:"statusId"`


    // WorkbinId - The ID of Workbin that contains the Workitem.
    WorkbinId string `json:"workbinId"`


    // AutoStatusTransition - Set it to false to disable auto status transition. By default, it is enabled.
    AutoStatusTransition bool `json:"autoStatusTransition"`


    // Description - The description of the Workitem. Maximum length of 512 characters.
    Description string `json:"description"`


    // TypeId - The ID of the Worktype of the Workitem.
    TypeId string `json:"typeId"`


    // CustomFields - Custom fields defined in the schema referenced by the worktype of the workitem.
    CustomFields map[string]interface{} `json:"customFields"`


    // QueueId - The ID of the Workitems queue. Must be a valid UUID.
    QueueId string `json:"queueId"`


    // AssigneeId - The ID of the assignee of the Workitem. Must be a valid UUID.
    AssigneeId string `json:"assigneeId"`


    // LanguageId - The ID of language of the Workitem. Must be a valid UUID.
    LanguageId string `json:"languageId"`


    // ExternalContactId - The ID of the external contact of the Workitem. Must be a valid UUID.
    ExternalContactId string `json:"externalContactId"`


    // ExternalTag - The external tag of the Workitem.
    ExternalTag string `json:"externalTag"`


    // SkillIds - The skill IDs of the Workitem. Must be valid UUIDs.
    SkillIds []string `json:"skillIds"`

}

// String returns a JSON representation of the model
func (o *Workitemcommoncreate) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.CustomFields = map[string]interface{}{"": Interface{}} 
    
    
    
    
    
     o.SkillIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemcommoncreate) MarshalJSON() ([]byte, error) {
    type Alias Workitemcommoncreate

    if WorkitemcommoncreateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemcommoncreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Priority int `json:"priority"`
        
        DateDue time.Time `json:"dateDue"`
        
        DateExpires time.Time `json:"dateExpires"`
        
        DurationSeconds int `json:"durationSeconds"`
        
        Ttl int `json:"ttl"`
        
        StatusId string `json:"statusId"`
        
        WorkbinId string `json:"workbinId"`
        
        AutoStatusTransition bool `json:"autoStatusTransition"`
        
        Description string `json:"description"`
        
        TypeId string `json:"typeId"`
        
        CustomFields map[string]interface{} `json:"customFields"`
        
        QueueId string `json:"queueId"`
        
        AssigneeId string `json:"assigneeId"`
        
        LanguageId string `json:"languageId"`
        
        ExternalContactId string `json:"externalContactId"`
        
        ExternalTag string `json:"externalTag"`
        
        SkillIds []string `json:"skillIds"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        CustomFields: map[string]interface{}{"": Interface{}},
        


        


        


        


        


        


        
        SkillIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

