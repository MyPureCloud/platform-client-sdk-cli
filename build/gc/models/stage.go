package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StageDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    

}

// Stage
type Stage struct { 
    


    // Name - The name of the Stage.
    Name string `json:"name"`


    // Description - The description of the Stage.
    Description string `json:"description"`


    // DateCreated - The Stage creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The Stage modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DateCompleted - The Stage completion date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCompleted time.Time `json:"dateCompleted"`


    // DateStarted - The Stage start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStarted time.Time `json:"dateStarted"`


    // ModifiedBy - The ID of the User who modified the Stage.
    ModifiedBy Userreference `json:"modifiedBy"`


    // Version - The version of the Stage.
    Version int `json:"version"`


    // Status - The Status of the Stage.
    Status string `json:"status"`


    


    // VarCase - The parent case of the Stage.
    VarCase Casereference `json:"case"`

}

// String returns a JSON representation of the model
func (o *Stage) String() string {
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stage) MarshalJSON() ([]byte, error) {
    type Alias Stage

    if StageMarshalled {
        return []byte("{}"), nil
    }
    StageMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DateCompleted time.Time `json:"dateCompleted"`
        
        DateStarted time.Time `json:"dateStarted"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        Version int `json:"version"`
        
        Status string `json:"status"`
        
        VarCase Casereference `json:"case"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

