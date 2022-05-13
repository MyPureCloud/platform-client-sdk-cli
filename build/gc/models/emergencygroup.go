package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmergencygroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmergencygroupDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    State string `json:"state"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Emergencygroup - A group of emergency call flows to use in an emergency.
type Emergencygroup struct { 
    


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


    // Enabled - True if an emergency is occurring and the associated emergency call flow(s) should be used.  False otherwise.
    Enabled bool `json:"enabled"`


    // EmergencyCallFlows - The emergency call flow(s) to use during an emergency.
    EmergencyCallFlows []Emergencycallflow `json:"emergencyCallFlows"`


    

}

// String returns a JSON representation of the model
func (o *Emergencygroup) String() string {
    
    
    
    
    
    
    
    
    
    
    
     o.EmergencyCallFlows = []Emergencycallflow{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emergencygroup) MarshalJSON() ([]byte, error) {
    type Alias Emergencygroup

    if EmergencygroupMarshalled {
        return []byte("{}"), nil
    }
    EmergencygroupMarshalled = true

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
        
        Enabled bool `json:"enabled"`
        
        EmergencyCallFlows []Emergencycallflow `json:"emergencyCallFlows"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        EmergencyCallFlows: []Emergencycallflow{{}},
        


        

        Alias: (*Alias)(u),
    })
}

