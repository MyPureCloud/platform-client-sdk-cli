package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgelogsjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgelogsjobresponseDud struct { 
    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy string `json:"modifiedBy"`


    CreatedBy string `json:"createdBy"`


    State string `json:"state"`


    ModifiedByApp string `json:"modifiedByApp"`


    CreatedByApp string `json:"createdByApp"`


    SelfUri string `json:"selfUri"`

}

// Edgelogsjobresponse
type Edgelogsjobresponse struct { 
    // Id - The created job id.
    Id string `json:"id"`


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Edgelogsjobresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgelogsjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Edgelogsjobresponse

    if EdgelogsjobresponseMarshalled {
        return []byte("{}"), nil
    }
    EdgelogsjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

