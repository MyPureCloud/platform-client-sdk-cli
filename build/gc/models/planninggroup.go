package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PlanninggroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PlanninggroupDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Planninggroup
type Planninggroup struct { 
    


    // Name
    Name string `json:"name"`


    // ServiceGoalTemplate - The ID of the service goal template associated with this planning group
    ServiceGoalTemplate Servicegoaltemplatereference `json:"serviceGoalTemplate"`


    // RoutePaths - Set of route paths associated with the planning group
    RoutePaths []Routepathresponse `json:"routePaths"`


    // Metadata - Version metadata for the planning group
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Planninggroup) String() string {
    
    
     o.RoutePaths = []Routepathresponse{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Planninggroup) MarshalJSON() ([]byte, error) {
    type Alias Planninggroup

    if PlanninggroupMarshalled {
        return []byte("{}"), nil
    }
    PlanninggroupMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        ServiceGoalTemplate Servicegoaltemplatereference `json:"serviceGoalTemplate"`
        
        RoutePaths []Routepathresponse `json:"routePaths"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        


        
        RoutePaths: []Routepathresponse{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

