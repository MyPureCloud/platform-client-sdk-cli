package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkbincreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkbincreateDud struct { 
    


    


    

}

// Workbincreate
type Workbincreate struct { 
    // Name - Workbin name. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // Description - Workbin description. Maximum length of 512 characters.
    Description string `json:"description"`


    // DivisionId - The ID of the division the Workbin belongs to. Defaults to home division ID.
    DivisionId string `json:"divisionId"`

}

// String returns a JSON representation of the model
func (o *Workbincreate) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workbincreate) MarshalJSON() ([]byte, error) {
    type Alias Workbincreate

    if WorkbincreateMarshalled {
        return []byte("{}"), nil
    }
    WorkbincreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        DivisionId string `json:"divisionId"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

