package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchsegmentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchsegmentDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Patchsegment
type Patchsegment struct { 
    // IsActive - Whether or not the segment is active.
    IsActive bool `json:"isActive"`


    // DisplayName - The display name of the segment.
    DisplayName string `json:"displayName"`


    // Version - The version of the segment.
    Version int `json:"version"`


    // Description - A description of the segment.
    Description string `json:"description"`


    // Color - The hexadecimal color value of the segment.
    Color string `json:"color"`


    // ShouldDisplayToAgent - Whether or not the segment should be displayed to agent/supervisor users.
    ShouldDisplayToAgent bool `json:"shouldDisplayToAgent"`


    // Context - The context of the segment.
    Context Patchcontext `json:"context"`


    // Journey - The pattern of rules defining the segment.
    Journey Patchjourney `json:"journey"`


    // ExternalSegment - Details of an entity corresponding to this segment in an external system.
    ExternalSegment Patchexternalsegment `json:"externalSegment"`


    // AssignmentExpirationDays - Time, in days, from when the segment is assigned until it is automatically unassigned.
    AssignmentExpirationDays int `json:"assignmentExpirationDays"`

}

// String returns a JSON representation of the model
func (o *Patchsegment) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchsegment) MarshalJSON() ([]byte, error) {
    type Alias Patchsegment

    if PatchsegmentMarshalled {
        return []byte("{}"), nil
    }
    PatchsegmentMarshalled = true

    return json.Marshal(&struct {
        
        IsActive bool `json:"isActive"`
        
        DisplayName string `json:"displayName"`
        
        Version int `json:"version"`
        
        Description string `json:"description"`
        
        Color string `json:"color"`
        
        ShouldDisplayToAgent bool `json:"shouldDisplayToAgent"`
        
        Context Patchcontext `json:"context"`
        
        Journey Patchjourney `json:"journey"`
        
        ExternalSegment Patchexternalsegment `json:"externalSegment"`
        
        AssignmentExpirationDays int `json:"assignmentExpirationDays"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

