package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StepplanupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StepplanupdateDud struct { 
    


    


    


    

}

// Stepplanupdate
type Stepplanupdate struct { 
    // Name - The name of the Stepplan. Valid length between 3 and 256 characters.
    Name string `json:"name"`


    // Description - The description of the Stepplan. Maximum length of 512 characters.
    Description string `json:"description"`


    // WorkitemSettings - The workitemSettings of the Stepplan.
    WorkitemSettings Workitemsettings `json:"workitemSettings"`


    // ActivityType - The activityType of the Stepplan.
    ActivityType string `json:"activityType"`

}

// String returns a JSON representation of the model
func (o *Stepplanupdate) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Stepplanupdate) MarshalJSON() ([]byte, error) {
    type Alias Stepplanupdate

    if StepplanupdateMarshalled {
        return []byte("{}"), nil
    }
    StepplanupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        WorkitemSettings Workitemsettings `json:"workitemSettings"`
        
        ActivityType string `json:"activityType"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

