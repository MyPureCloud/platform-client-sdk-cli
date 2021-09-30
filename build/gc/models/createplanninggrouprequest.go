package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreateplanninggrouprequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreateplanninggrouprequestDud struct { 
    


    


    

}

// Createplanninggrouprequest
type Createplanninggrouprequest struct { 
    // Name - The name of the planning group
    Name string `json:"name"`


    // RoutePaths - Set of route paths to associate with the planning group
    RoutePaths []Routepathrequest `json:"routePaths"`


    // ServiceGoalTemplateId - The ID of the service goal template to associate with this planning group
    ServiceGoalTemplateId string `json:"serviceGoalTemplateId"`

}

// String returns a JSON representation of the model
func (o *Createplanninggrouprequest) String() string {
    
    
    
    
    
    
     o.RoutePaths = []Routepathrequest{{}} 
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createplanninggrouprequest) MarshalJSON() ([]byte, error) {
    type Alias Createplanninggrouprequest

    if CreateplanninggrouprequestMarshalled {
        return []byte("{}"), nil
    }
    CreateplanninggrouprequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        RoutePaths []Routepathrequest `json:"routePaths"`
        
        ServiceGoalTemplateId string `json:"serviceGoalTemplateId"`
        
        *Alias
    }{
        

        

        

        
        RoutePaths: []Routepathrequest{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

