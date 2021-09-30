package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingannotationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingannotationDud struct { 
    Id string `json:"id"`


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`


    ModifiedBy Userreference `json:"modifiedBy"`


    DateModified time.Time `json:"dateModified"`


    


    IsDeleted bool `json:"isDeleted"`


    AccessType string `json:"accessType"`


    SelfUri string `json:"selfUri"`

}

// Coachingannotation
type Coachingannotation struct { 
    


    


    


    


    


    // Text - The text of the annotation.
    Text string `json:"text"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Coachingannotation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingannotation) MarshalJSON() ([]byte, error) {
    type Alias Coachingannotation

    if CoachingannotationMarshalled {
        return []byte("{}"), nil
    }
    CoachingannotationMarshalled = true

    return json.Marshal(&struct { 
        
        
        
        
        
        
        
        
        
        
        Text string `json:"text"`
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

