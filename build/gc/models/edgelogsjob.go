package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgelogsjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgelogsjobDud struct { 
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

// Edgelogsjob
type Edgelogsjob struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // Files - The files available to upload from the Edge to the cloud.
    Files []Edgelogsjobfile `json:"files"`


    

}

// String returns a JSON representation of the model
func (o *Edgelogsjob) String() string {
    
    
    
    
     o.Files = []Edgelogsjobfile{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgelogsjob) MarshalJSON() ([]byte, error) {
    type Alias Edgelogsjob

    if EdgelogsjobMarshalled {
        return []byte("{}"), nil
    }
    EdgelogsjobMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        Files []Edgelogsjobfile `json:"files"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        Files: []Edgelogsjobfile{{}},
        


        

        Alias: (*Alias)(u),
    })
}

