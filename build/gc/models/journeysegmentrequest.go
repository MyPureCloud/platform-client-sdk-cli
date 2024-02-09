package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    JourneysegmentrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type JourneysegmentrequestDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Journeysegmentrequest
type Journeysegmentrequest struct { 
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


    // Scope - The target entity that a segment applies to.
    Scope string `json:"scope"`


    // ShouldDisplayToAgent - Whether or not the segment should be displayed to agent/supervisor users.
    ShouldDisplayToAgent bool `json:"shouldDisplayToAgent"`


    // Context - The context of the segment.
    Context Requestcontext `json:"context"`


    // Journey - The pattern of rules defining the segment.
    Journey Requestjourney `json:"journey"`


    // ExternalSegment - Details of an entity corresponding to this segment in an external system.
    ExternalSegment Requestexternalsegment `json:"externalSegment"`


    // AssignmentExpirationDays - Time, in days, from when the segment is assigned until it is automatically unassigned.
    AssignmentExpirationDays int `json:"assignmentExpirationDays"`

}

// String returns a JSON representation of the model
func (o *Journeysegmentrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Journeysegmentrequest) MarshalJSON() ([]byte, error) {
    type Alias Journeysegmentrequest

    if JourneysegmentrequestMarshalled {
        return []byte("{}"), nil
    }
    JourneysegmentrequestMarshalled = true

    return json.Marshal(&struct {
        
        IsActive bool `json:"isActive"`
        
        DisplayName string `json:"displayName"`
        
        Version int `json:"version"`
        
        Description string `json:"description"`
        
        Color string `json:"color"`
        
        Scope string `json:"scope"`
        
        ShouldDisplayToAgent bool `json:"shouldDisplayToAgent"`
        
        Context Requestcontext `json:"context"`
        
        Journey Requestjourney `json:"journey"`
        
        ExternalSegment Requestexternalsegment `json:"externalSegment"`
        
        AssignmentExpirationDays int `json:"assignmentExpirationDays"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

