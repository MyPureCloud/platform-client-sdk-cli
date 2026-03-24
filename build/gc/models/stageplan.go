package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StageplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StageplanDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Stageplan
type Stageplan struct { 
    


    // Name - The name of the Stageplan.
    Name string `json:"name"`


    // Description - The description of the Stageplan.
    Description string `json:"description"`


    // Caseplan - The Caseplan of the Stageplan.
    Caseplan Caseplanreference `json:"caseplan"`


    // DateCreated - The Stageplan creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The Stageplan modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the User who modified the Stageplan.
    ModifiedBy Userreference `json:"modifiedBy"`


    

}

// String returns a JSON representation of the model
func (o *Stageplan) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stageplan) MarshalJSON() ([]byte, error) {
    type Alias Stageplan

    if StageplanMarshalled {
        return []byte("{}"), nil
    }
    StageplanMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Caseplan Caseplanreference `json:"caseplan"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

