package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MoveagentsresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MoveagentsresponseDud struct { 
    


    


    

}

// Moveagentsresponse
type Moveagentsresponse struct { 
    // RequestingUser - The user that made the request
    RequestingUser Userreference `json:"requestingUser"`


    // DestinationManagementUnit - The management unit specified on the request
    DestinationManagementUnit Managementunitreference `json:"destinationManagementUnit"`


    // Results - The list containing the agent and result of the move operation
    Results []Moveagentresponse `json:"results"`

}

// String returns a JSON representation of the model
func (o *Moveagentsresponse) String() string {
    
    
     o.Results = []Moveagentresponse{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Moveagentsresponse) MarshalJSON() ([]byte, error) {
    type Alias Moveagentsresponse

    if MoveagentsresponseMarshalled {
        return []byte("{}"), nil
    }
    MoveagentsresponseMarshalled = true

    return json.Marshal(&struct {
        
        RequestingUser Userreference `json:"requestingUser"`
        
        DestinationManagementUnit Managementunitreference `json:"destinationManagementUnit"`
        
        Results []Moveagentresponse `json:"results"`
        *Alias
    }{

        


        


        
        Results: []Moveagentresponse{{}},
        

        Alias: (*Alias)(u),
    })
}

