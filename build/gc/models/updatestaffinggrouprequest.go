package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatestaffinggrouprequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatestaffinggrouprequestDud struct { 
    


    


    


    

}

// Updatestaffinggrouprequest
type Updatestaffinggrouprequest struct { 
    // Name - The name of the staffing group
    Name string `json:"name"`


    // UserIds - The set of user Ids to associate with the staffing group
    UserIds Setwrapperstring `json:"userIds"`


    // PlanningGroupIds - The set of planning group Ids to associate with the staffing group
    PlanningGroupIds Setwrapperstring `json:"planningGroupIds"`


    // Metadata - Version metadata for the staffing group
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updatestaffinggrouprequest) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatestaffinggrouprequest) MarshalJSON() ([]byte, error) {
    type Alias Updatestaffinggrouprequest

    if UpdatestaffinggrouprequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatestaffinggrouprequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        UserIds Setwrapperstring `json:"userIds"`
        
        PlanningGroupIds Setwrapperstring `json:"planningGroupIds"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

