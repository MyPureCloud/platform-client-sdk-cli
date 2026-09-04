package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StageplancreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StageplancreateDud struct { 
    


    


    

}

// Stageplancreate
type Stageplancreate struct { 
    // Name - The name of the Stageplan. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // Description - The description of the Stageplan. Maximum length of 512 characters.
    Description string `json:"description"`


    // After - The ID of the Stageplan to place the new Stageplan after. Omit or null to place at the front.
    After string `json:"after"`

}

// String returns a JSON representation of the model
func (o *Stageplancreate) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stageplancreate) MarshalJSON() ([]byte, error) {
    type Alias Stageplancreate

    if StageplancreateMarshalled {
        return []byte("{}"), nil
    }
    StageplancreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        After string `json:"after"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

