package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgetrunkbaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgetrunkbaseDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Edgetrunkbase
type Edgetrunkbase struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // TrunkMetabase - The meta-base this trunk is based on.
    TrunkMetabase Domainentityref `json:"trunkMetabase"`


    // Properties
    Properties map[string]interface{} `json:"properties"`


    // TrunkType - The type of this trunk base.
    TrunkType string `json:"trunkType"`


    // Site - Used to determine the media regions for inbound and outbound calls through a trunk. Also determines the dial plan to use for calls that came in on a trunk and have to be sent out on it as well.
    Site Domainentityref `json:"site"`


    // InboundSite - Allows a customer to set the site to which inbound calls will be routed
    InboundSite Domainentityref `json:"inboundSite"`


    

}

// String returns a JSON representation of the model
func (o *Edgetrunkbase) String() string {
    
    
    
    
    
     o.Properties = map[string]interface{}{"": Interface{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgetrunkbase) MarshalJSON() ([]byte, error) {
    type Alias Edgetrunkbase

    if EdgetrunkbaseMarshalled {
        return []byte("{}"), nil
    }
    EdgetrunkbaseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        TrunkMetabase Domainentityref `json:"trunkMetabase"`
        
        Properties map[string]interface{} `json:"properties"`
        
        TrunkType string `json:"trunkType"`
        
        Site Domainentityref `json:"site"`
        
        InboundSite Domainentityref `json:"inboundSite"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Properties: map[string]interface{}{"": Interface{}},
        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

