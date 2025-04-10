package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContestwinnersrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContestwinnersrequestDud struct { 
    


    

}

// Contestwinnersrequest
type Contestwinnersrequest struct { 
    // Tier - The Contest Winner tier
    Tier int `json:"tier"`


    // UserIds - The Contest Winner users at the tier
    UserIds []string `json:"userIds"`

}

// String returns a JSON representation of the model
func (o *Contestwinnersrequest) String() string {
    
     o.UserIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contestwinnersrequest) MarshalJSON() ([]byte, error) {
    type Alias Contestwinnersrequest

    if ContestwinnersrequestMarshalled {
        return []byte("{}"), nil
    }
    ContestwinnersrequestMarshalled = true

    return json.Marshal(&struct {
        
        Tier int `json:"tier"`
        
        UserIds []string `json:"userIds"`
        *Alias
    }{

        


        
        UserIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

