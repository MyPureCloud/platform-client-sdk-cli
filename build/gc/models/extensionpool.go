package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExtensionpoolMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExtensionpoolDud struct { 
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

// Extensionpool
type Extensionpool struct { 
    


    // Name - The name of the entity.
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Division `json:"division"`


    // Description - The resource's description.
    Description string `json:"description"`


    // Version - The current version of the resource.
    Version int `json:"version"`


    


    


    


    


    


    


    


    // StartNumber - The starting phone number for the range of this Extension pool. The number must be between 3 and 9 digits in length and the same length as the endNumber.
    StartNumber string `json:"startNumber"`


    // EndNumber - The ending phone number for the range of this Extension pool. The number must be between 3 and 9 digits in length and the same length as the startNumber.
    EndNumber string `json:"endNumber"`


    

}

// String returns a JSON representation of the model
func (o *Extensionpool) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Extensionpool) MarshalJSON() ([]byte, error) {
    type Alias Extensionpool

    if ExtensionpoolMarshalled {
        return []byte("{}"), nil
    }
    ExtensionpoolMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Division `json:"division"`
        
        Description string `json:"description"`
        
        Version int `json:"version"`
        
        StartNumber string `json:"startNumber"`
        
        EndNumber string `json:"endNumber"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

