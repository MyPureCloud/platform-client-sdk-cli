package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StepplanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StepplanDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Stepplan
type Stepplan struct { 
    


    // Name - The name of the Stepplan.
    Name string `json:"name"`


    // Description - The description of the Stepplan.
    Description string `json:"description"`


    // Caseplan - The Caseplan of the Stepplan.
    Caseplan Caseplanreference `json:"caseplan"`


    // Stageplan - The Stageplan of the Stepplan.
    Stageplan Stageplanreference `json:"stageplan"`


    // DateCreated - The Stepplan creation date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The Stepplan modification date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // ModifiedBy - The ID of the User who modified the Stepplan.
    ModifiedBy Userreference `json:"modifiedBy"`


    // ActivityType - The activityType of the Stepplan.
    ActivityType string `json:"activityType"`


    // WorkitemSettings - The workitemSettings of the Stepplan.
    WorkitemSettings Workitemsettingsresponse `json:"workitemSettings"`


    

}

// String returns a JSON representation of the model
func (o *Stepplan) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stepplan) MarshalJSON() ([]byte, error) {
    type Alias Stepplan

    if StepplanMarshalled {
        return []byte("{}"), nil
    }
    StepplanMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Caseplan Caseplanreference `json:"caseplan"`
        
        Stageplan Stageplanreference `json:"stageplan"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        ActivityType string `json:"activityType"`
        
        WorkitemSettings Workitemsettingsresponse `json:"workitemSettings"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

