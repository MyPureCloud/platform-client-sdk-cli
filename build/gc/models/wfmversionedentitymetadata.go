package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WfmversionedentitymetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WfmversionedentitymetadataDud struct { 
    


    ModifiedBy Userreference `json:"modifiedBy"`


    DateModified time.Time `json:"dateModified"`


    CreatedBy Userreference `json:"createdBy"`


    DateCreated time.Time `json:"dateCreated"`

}

// Wfmversionedentitymetadata - Metadata to associate with a given entity
type Wfmversionedentitymetadata struct { 
    // Version - The version of the associated entity.  Used to prevent conflicts on concurrent edits
    Version int `json:"version"`


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Wfmversionedentitymetadata) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Wfmversionedentitymetadata) MarshalJSON() ([]byte, error) {
    type Alias Wfmversionedentitymetadata

    if WfmversionedentitymetadataMarshalled {
        return []byte("{}"), nil
    }
    WfmversionedentitymetadataMarshalled = true

    return json.Marshal(&struct { 
        Version int `json:"version"`
        
        
        
        
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

