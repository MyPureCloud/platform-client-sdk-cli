package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TrunkMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TrunkDud struct { 
    Id string `json:"id"`


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    


    


    


    InService bool `json:"inService"`


    


    LogicalInterface Domainentityref `json:"logicalInterface"`


    ConnectedStatus Trunkconnectedstatus `json:"connectedStatus"`


    OptionsStatus []Trunkmetricsoptions `json:"optionsStatus"`


    RegistersStatus []Trunkmetricsregisters `json:"registersStatus"`


    IpStatus Trunkmetricsnetworktypeip `json:"ipStatus"`


    OptionsEnabledStatus string `json:"optionsEnabledStatus"`


    RegistersEnabledStatus string `json:"registersEnabledStatus"`


    Family int `json:"family"`


    ProxyAddressList []string `json:"proxyAddressList"`


    SelfUri string `json:"selfUri"`

}

// Trunk
type Trunk struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // TrunkType - The type of this trunk.
    TrunkType string `json:"trunkType"`


    // Edge - The Edge using this trunk.
    Edge Domainentityref `json:"edge"`


    // TrunkBase - The trunk base configuration used on this trunk.
    TrunkBase Domainentityref `json:"trunkBase"`


    // TrunkMetabase - The metabase used to create this trunk.
    TrunkMetabase Domainentityref `json:"trunkMetabase"`


    // EdgeGroup - The edge group associated with this trunk.
    EdgeGroup Domainentityref `json:"edgeGroup"`


    


    // Enabled - True if the Edge used by this trunk is in-service
    Enabled bool `json:"enabled"`


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Trunk) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Trunk) MarshalJSON() ([]byte, error) {
    type Alias Trunk

    if TrunkMarshalled {
        return []byte("{}"), nil
    }
    TrunkMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        TrunkType string `json:"trunkType"`
        
        Edge Domainentityref `json:"edge"`
        
        TrunkBase Domainentityref `json:"trunkBase"`
        
        TrunkMetabase Domainentityref `json:"trunkMetabase"`
        
        EdgeGroup Domainentityref `json:"edgeGroup"`
        
        Enabled bool `json:"enabled"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

