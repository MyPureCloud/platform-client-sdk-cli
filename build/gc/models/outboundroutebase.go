package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutboundroutebaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutboundroutebaseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Outboundroutebase
type Outboundroutebase struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    // DateCreated - The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the user that last modified the resource.
    ModifiedBy string `json:"modifiedBy"`


    // CreatedBy - The ID of the user that created the resource.
    CreatedBy string `json:"createdBy"`


    


    // ModifiedByApp - The application that last modified the resource.
    ModifiedByApp string `json:"modifiedByApp"`


    // CreatedByApp - The application that created the resource.
    CreatedByApp string `json:"createdByApp"`


    // ClassificationTypes - The site associated to the outbound route.
    ClassificationTypes []string `json:"classificationTypes"`


    // Enabled
    Enabled bool `json:"enabled"`


    // Distribution
    Distribution string `json:"distribution"`


    // ExternalTrunkBases - Trunk base settings of trunkType \"EXTERNAL\".  This base must also be set on an edge logical interface for correct routing.
    ExternalTrunkBases []Domainentityref `json:"externalTrunkBases"`


    

}

// String returns a JSON representation of the model
func (o *Outboundroutebase) String() string {
    
    
    
    
    
    
    
    
    
    
     o.ClassificationTypes = []string{""} 
    
    
     o.ExternalTrunkBases = []Domainentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outboundroutebase) MarshalJSON() ([]byte, error) {
    type Alias Outboundroutebase

    if OutboundroutebaseMarshalled {
        return []byte("{}"), nil
    }
    OutboundroutebaseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
        ClassificationTypes []string `json:"classificationTypes"`
        
        Enabled bool `json:"enabled"`
        
        Distribution string `json:"distribution"`
        
        ExternalTrunkBases []Domainentityref `json:"externalTrunkBases"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        ClassificationTypes: []string{""},
        


        


        


        
        ExternalTrunkBases: []Domainentityref{{}},
        


        

        Alias: (*Alias)(u),
    })
}

