package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkbinupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkbinupdateDud struct { 
    


    

}

// Workbinupdate
type Workbinupdate struct { 
    // Name - Workbin name. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // Description - Workbin description. Maximum length of 512 characters.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Workbinupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workbinupdate) MarshalJSON() ([]byte, error) {
    type Alias Workbinupdate

    if WorkbinupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkbinupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

