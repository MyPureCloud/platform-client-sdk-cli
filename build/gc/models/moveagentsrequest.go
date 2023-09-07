package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MoveagentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MoveagentsrequestDud struct { 
    


    

}

// Moveagentsrequest
type Moveagentsrequest struct { 
    // UserIds - The list of user ids to move
    UserIds []string `json:"userIds"`


    // DestinationManagementUnitId - The id of the management unit for which the move will be performed. To remove users from their management unit this should be set to null.
    DestinationManagementUnitId string `json:"destinationManagementUnitId"`

}

// String returns a JSON representation of the model
func (o *Moveagentsrequest) String() string {
     o.UserIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Moveagentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Moveagentsrequest

    if MoveagentsrequestMarshalled {
        return []byte("{}"), nil
    }
    MoveagentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        UserIds []string `json:"userIds"`
        
        DestinationManagementUnitId string `json:"destinationManagementUnitId"`
        *Alias
    }{

        
        UserIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

