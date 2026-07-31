package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AgenticvirtualagentversionpublishjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AgenticvirtualagentversionpublishjobrequestDud struct { 
    

}

// Agenticvirtualagentversionpublishjobrequest
type Agenticvirtualagentversionpublishjobrequest struct { 
    // VirtualAgentVersion - The attributes of the virtual agent version to update as part of this publish job.
    VirtualAgentVersion Agenticvirtualagentversionpublish `json:"virtualAgentVersion"`

}

// String returns a JSON representation of the model
func (o *Agenticvirtualagentversionpublishjobrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Agenticvirtualagentversionpublishjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Agenticvirtualagentversionpublishjobrequest

    if AgenticvirtualagentversionpublishjobrequestMarshalled {
        return []byte("{}"), nil
    }
    AgenticvirtualagentversionpublishjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        VirtualAgentVersion Agenticvirtualagentversionpublish `json:"virtualAgentVersion"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

