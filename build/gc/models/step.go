package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StepMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StepDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`


    

}

// Step
type Step struct { 
    


    // Name - The name of the Step.
    Name string `json:"name"`


    // Description - The description of the Step.
    Description string `json:"description"`


    // DateCreated - The Step creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The Step modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DateCompleted - The Step completion date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCompleted time.Time `json:"dateCompleted"`


    // DateStarted - The Step start date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateStarted time.Time `json:"dateStarted"`


    // ModifiedBy - The ID of the User who modified the Step.
    ModifiedBy Userreference `json:"modifiedBy"`


    // Version - The version of the Step.
    Version int `json:"version"`


    // Status - The Status of the Step.
    Status string `json:"status"`


    // Stage - The parent stage of the step.
    Stage Stagereference `json:"stage"`


    


    // VarCase - The parent case of the step.
    VarCase Casereference `json:"case"`

}

// String returns a JSON representation of the model
func (o *Step) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Step) MarshalJSON() ([]byte, error) {
    type Alias Step

    if StepMarshalled {
        return []byte("{}"), nil
    }
    StepMarshalled = true

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
        
        Stage Stagereference `json:"stage"`
        
        VarCase Casereference `json:"case"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

