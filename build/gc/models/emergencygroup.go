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


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    


    


    SelfUri string `json:"selfUri"`

}

// Emergencygroup - A group of emergency call flows to use in an emergency.
type Emergencygroup struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writabledivision `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


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
        
        Division Writabledivision `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        EmergencyCallFlows []Emergencycallflow `json:"emergencyCallFlows"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        EmergencyCallFlows: []Emergencycallflow{{}},
        


        

        Alias: (*Alias)(u),
    })
}

