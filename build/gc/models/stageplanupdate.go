package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StageplanupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StageplanupdateDud struct { 
    


    

}

// Stageplanupdate
type Stageplanupdate struct { 
    // Name - The name of the Stageplan. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // Description - The description of the Stageplan. Maximum length of 512 characters.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Stageplanupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stageplanupdate) MarshalJSON() ([]byte, error) {
    type Alias Stageplanupdate

    if StageplanupdateMarshalled {
        return []byte("{}"), nil
    }
    StageplanupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

