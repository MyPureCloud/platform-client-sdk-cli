package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemdeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemdeltaDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Workitemdelta
type Workitemdelta struct { 
    // Name
    Name Workitemsattributechangestring `json:"name"`


    // Description
    Description Workitemsattributechangestring `json:"description"`


    // LanguageId
    LanguageId Workitemsattributechangestring `json:"languageId"`


    // UtilizationLabelId
    UtilizationLabelId Workitemsattributechangestring `json:"utilizationLabelId"`


    // Priority
    Priority Workitemsattributechangeinteger `json:"priority"`


    // SkillIds
    SkillIds Workitemsattributechangelist `json:"skillIds"`


    // PreferredAgentIds
    PreferredAgentIds Workitemsattributechangelist `json:"preferredAgentIds"`


    // DateDue
    DateDue Workitemsattributechangeinstant `json:"dateDue"`


    // DateExpires
    DateExpires Workitemsattributechangeinstant `json:"dateExpires"`


    // DurationSeconds
    DurationSeconds Workitemsattributechangeinteger `json:"durationSeconds"`


    // StatusId
    StatusId Workitemsattributechangestring `json:"statusId"`


    // ReporterId
    ReporterId Workitemsattributechangestring `json:"reporterId"`


    // ExternalContactId
    ExternalContactId Workitemsattributechangestring `json:"externalContactId"`


    // AssigneeId
    AssigneeId Workitemsattributechangestring `json:"assigneeId"`


    // WorkbinId
    WorkbinId Workitemsattributechangestring `json:"workbinId"`


    // QueueId
    QueueId Workitemsattributechangestring `json:"queueId"`


    // ExternalTag
    ExternalTag Workitemsattributechangestring `json:"externalTag"`


    // WrapupId
    WrapupId Workitemsattributechangestring `json:"wrapupId"`


    // Ttl
    Ttl Workitemsattributechangeinteger `json:"ttl"`


    // DateClosed
    DateClosed Workitemsattributechangeinstant `json:"dateClosed"`


    // AssignmentState
    AssignmentState Workitemsattributechangestring `json:"assignmentState"`


    // AutoStatusTransition
    AutoStatusTransition Workitemsattributechangeboolean `json:"autoStatusTransition"`


    // CustomFields
    CustomFields Workitemsattributechangemap `json:"customFields"`


    // DateModified
    DateModified Workitemsattributechangeinstant `json:"dateModified"`


    // ModifiedBy
    ModifiedBy Workitemsattributechangestring `json:"modifiedBy"`


    // StatusCategory
    StatusCategory Workitemsattributechangeworkitemstatuscategory `json:"statusCategory"`


    // ScriptId
    ScriptId Workitemsattributechangestring `json:"scriptId"`

}

// String returns a JSON representation of the model
func (o *Workitemdelta) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemdelta) MarshalJSON() ([]byte, error) {
    type Alias Workitemdelta

    if WorkitemdeltaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemdeltaMarshalled = true

    return json.Marshal(&struct {
        
        Name Workitemsattributechangestring `json:"name"`
        
        Description Workitemsattributechangestring `json:"description"`
        
        LanguageId Workitemsattributechangestring `json:"languageId"`
        
        UtilizationLabelId Workitemsattributechangestring `json:"utilizationLabelId"`
        
        Priority Workitemsattributechangeinteger `json:"priority"`
        
        SkillIds Workitemsattributechangelist `json:"skillIds"`
        
        PreferredAgentIds Workitemsattributechangelist `json:"preferredAgentIds"`
        
        DateDue Workitemsattributechangeinstant `json:"dateDue"`
        
        DateExpires Workitemsattributechangeinstant `json:"dateExpires"`
        
        DurationSeconds Workitemsattributechangeinteger `json:"durationSeconds"`
        
        StatusId Workitemsattributechangestring `json:"statusId"`
        
        ReporterId Workitemsattributechangestring `json:"reporterId"`
        
        ExternalContactId Workitemsattributechangestring `json:"externalContactId"`
        
        AssigneeId Workitemsattributechangestring `json:"assigneeId"`
        
        WorkbinId Workitemsattributechangestring `json:"workbinId"`
        
        QueueId Workitemsattributechangestring `json:"queueId"`
        
        ExternalTag Workitemsattributechangestring `json:"externalTag"`
        
        WrapupId Workitemsattributechangestring `json:"wrapupId"`
        
        Ttl Workitemsattributechangeinteger `json:"ttl"`
        
        DateClosed Workitemsattributechangeinstant `json:"dateClosed"`
        
        AssignmentState Workitemsattributechangestring `json:"assignmentState"`
        
        AutoStatusTransition Workitemsattributechangeboolean `json:"autoStatusTransition"`
        
        CustomFields Workitemsattributechangemap `json:"customFields"`
        
        DateModified Workitemsattributechangeinstant `json:"dateModified"`
        
        ModifiedBy Workitemsattributechangestring `json:"modifiedBy"`
        
        StatusCategory Workitemsattributechangeworkitemstatuscategory `json:"statusCategory"`
        
        ScriptId Workitemsattributechangestring `json:"scriptId"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

