package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemstatusfilterMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemstatusfilterDud struct { 
    


    

}

// Workitemstatusfilter
type Workitemstatusfilter struct { 
    // WorktypeId - Worktype ID belonging to the selected workitem status
    WorktypeId string `json:"worktypeId"`


    // WorkitemStatusId - Workitem status ID
    WorkitemStatusId string `json:"workitemStatusId"`

}

// String returns a JSON representation of the model
func (o *Workitemstatusfilter) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemstatusfilter) MarshalJSON() ([]byte, error) {
    type Alias Workitemstatusfilter

    if WorkitemstatusfilterMarshalled {
        return []byte("{}"), nil
    }
    WorkitemstatusfilterMarshalled = true

    return json.Marshal(&struct {
        
        WorktypeId string `json:"worktypeId"`
        
        WorkitemStatusId string `json:"workitemStatusId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

