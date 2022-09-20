package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatebusinessunitrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatebusinessunitrequestDud struct { 
    


    


    

}

// Createbusinessunitrequest
type Createbusinessunitrequest struct { 
    // Name - The name of the business unit
    Name string `json:"name"`


    // DivisionId - The ID of the division to which the business unit should be added
    DivisionId string `json:"divisionId"`


    // Settings - Configuration for the business unit
    Settings Createbusinessunitsettingsrequest `json:"settings"`

}

// String returns a JSON representation of the model
func (o *Createbusinessunitrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createbusinessunitrequest) MarshalJSON() ([]byte, error) {
    type Alias Createbusinessunitrequest

    if CreatebusinessunitrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatebusinessunitrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DivisionId string `json:"divisionId"`
        
        Settings Createbusinessunitsettingsrequest `json:"settings"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

