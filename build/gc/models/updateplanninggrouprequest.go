package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdateplanninggrouprequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdateplanninggrouprequestDud struct { 
    


    


    


    

}

// Updateplanninggrouprequest
type Updateplanninggrouprequest struct { 
    // Name - The name of the planning group
    Name string `json:"name"`


    // RoutePaths - Set of route paths to associate with the planning group
    RoutePaths Setwrapperroutepathrequest `json:"routePaths"`


    // ServiceGoalTemplateId - The ID of the service goal template to associate with this planning group
    ServiceGoalTemplateId string `json:"serviceGoalTemplateId"`


    // Metadata - Version metadata for the planning group
    Metadata Wfmversionedentitymetadata `json:"metadata"`

}

// String returns a JSON representation of the model
func (o *Updateplanninggrouprequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updateplanninggrouprequest) MarshalJSON() ([]byte, error) {
    type Alias Updateplanninggrouprequest

    if UpdateplanninggrouprequestMarshalled {
        return []byte("{}"), nil
    }
    UpdateplanninggrouprequestMarshalled = true

    return json.Marshal(&struct { 
        Name string `json:"name"`
        
        RoutePaths Setwrapperroutepathrequest `json:"routePaths"`
        
        ServiceGoalTemplateId string `json:"serviceGoalTemplateId"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

