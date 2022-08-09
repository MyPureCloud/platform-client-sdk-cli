package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgegroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgegroupDud struct { 
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

// Edgegroup
type Edgegroup struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // Managed - Is this edge group being managed remotely.
    Managed bool `json:"managed"`


    // Hybrid - Is this edge group hybrid.
    Hybrid bool `json:"hybrid"`


    // EdgeTrunkBaseAssignment - A trunk base settings assignment of trunkType \"EDGE\" to use for edge-to-edge communication.
    EdgeTrunkBaseAssignment Trunkbaseassignment `json:"edgeTrunkBaseAssignment"`


    // PhoneTrunkBases - Trunk base settings of trunkType \"PHONE\" to inherit to edge logical interface for phone communication.
    PhoneTrunkBases []Trunkbase `json:"phoneTrunkBases"`


    

}

// String returns a JSON representation of the model
func (o *Edgegroup) String() string {
    
    
    
    
    
    
    
     o.PhoneTrunkBases = []Trunkbase{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgegroup) MarshalJSON() ([]byte, error) {
    type Alias Edgegroup

    if EdgegroupMarshalled {
        return []byte("{}"), nil
    }
    EdgegroupMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        Managed bool `json:"managed"`
        
        Hybrid bool `json:"hybrid"`
        
        EdgeTrunkBaseAssignment Trunkbaseassignment `json:"edgeTrunkBaseAssignment"`
        
        PhoneTrunkBases []Trunkbase `json:"phoneTrunkBases"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        


        
        PhoneTrunkBases: []Trunkbase{{}},
        


        

        Alias: (*Alias)(u),
    })
}

