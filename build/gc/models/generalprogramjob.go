package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GeneralprogramjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GeneralprogramjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Generalprogramjob
type Generalprogramjob struct { 
    


    // State
    State string `json:"state"`


    // CreatedBy
    CreatedBy Addressableentityref `json:"createdBy"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Generalprogramjob) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Generalprogramjob) MarshalJSON() ([]byte, error) {
    type Alias Generalprogramjob

    if GeneralprogramjobMarshalled {
        return []byte("{}"), nil
    }
    GeneralprogramjobMarshalled = true

    return json.Marshal(&struct { 
        
        
        State string `json:"state"`
        
        CreatedBy Addressableentityref `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

