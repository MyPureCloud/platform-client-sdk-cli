package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutboundrouteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutboundrouteDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    Site Site `json:"site"`


    Managed bool `json:"managed"`


    SelfUri string `json:"selfUri"`

}

// Outboundroute
type Outboundroute struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


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
func (o *Outboundroute) String() string {
    
    
    
    
     o.ClassificationTypes = []string{""} 
    
    
     o.ExternalTrunkBases = []Domainentityref{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outboundroute) MarshalJSON() ([]byte, error) {
    type Alias Outboundroute

    if OutboundrouteMarshalled {
        return []byte("{}"), nil
    }
    OutboundrouteMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
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

