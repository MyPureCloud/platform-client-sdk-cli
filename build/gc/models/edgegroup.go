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


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Edgegroup
type Edgegroup struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


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
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy string `json:"modifiedBy"`
        
        CreatedBy string `json:"createdBy"`
        
        
        
        ModifiedByApp string `json:"modifiedByApp"`
        
        CreatedByApp string `json:"createdByApp"`
        
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

