package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChecklistitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChecklistitemDud struct { 
    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Checklistitem
type Checklistitem struct { 
    // Id - ID of the checklist item.
    Id string `json:"id"`


    // Name - Name of the checklist item.
    Name string `json:"name"`


    // Description - Description of the checklist item.
    Description string `json:"description"`


    // AutomatedCheckEnabled - Flag to indicate whether automated check is enabled for this checklist item.
    AutomatedCheckEnabled bool `json:"automatedCheckEnabled"`


    // Important - Flag to indicate whether this checklist item is marked as important.
    Important bool `json:"important"`


    // StateFromModel - Checklist state as evaluated by the model.
    StateFromModel string `json:"stateFromModel"`


    // StateFromAgent - Checklist state as evaluated by the agent.
    StateFromAgent string `json:"stateFromAgent"`


    // DateLastModifiedByModel - Date when the checklist item was last modified by the model. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateLastModifiedByModel time.Time `json:"dateLastModifiedByModel"`


    // DateLastModifiedByAgent - Date when the checklist item was last modified by the agent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateLastModifiedByAgent time.Time `json:"dateLastModifiedByAgent"`


    // LastModifiedInAcw - Flag to indicate whether this checklist item was modified in ACW.
    LastModifiedInAcw bool `json:"lastModifiedInAcw"`


    

}

// String returns a JSON representation of the model
func (o *Checklistitem) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Checklistitem) MarshalJSON() ([]byte, error) {
    type Alias Checklistitem

    if ChecklistitemMarshalled {
        return []byte("{}"), nil
    }
    ChecklistitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        AutomatedCheckEnabled bool `json:"automatedCheckEnabled"`
        
        Important bool `json:"important"`
        
        StateFromModel string `json:"stateFromModel"`
        
        StateFromAgent string `json:"stateFromAgent"`
        
        DateLastModifiedByModel time.Time `json:"dateLastModifiedByModel"`
        
        DateLastModifiedByAgent time.Time `json:"dateLastModifiedByAgent"`
        
        LastModifiedInAcw bool `json:"lastModifiedInAcw"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

