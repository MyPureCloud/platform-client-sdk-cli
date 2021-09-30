package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ProgramjobMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ProgramjobDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Programjob
type Programjob struct { 
    


    // State
    State string `json:"state"`


    // Programs
    Programs []Baseprogramentity `json:"programs"`


    // CreatedBy
    CreatedBy Addressableentityref `json:"createdBy"`


    // DateCreated - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Programjob) String() string {
    
    
    
    
    
    
    
    
     o.Programs = []Baseprogramentity{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Programjob) MarshalJSON() ([]byte, error) {
    type Alias Programjob

    if ProgramjobMarshalled {
        return []byte("{}"), nil
    }
    ProgramjobMarshalled = true

    return json.Marshal(&struct { 
        
        
        State string `json:"state"`
        
        Programs []Baseprogramentity `json:"programs"`
        
        CreatedBy Addressableentityref `json:"createdBy"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        
        Programs: []Baseprogramentity{{}},
        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

