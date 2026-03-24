package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsettingsDud struct { 
    

}

// Workitemsettings
type Workitemsettings struct { 
    // WorktypeId - The UUID of the Worktype.
    WorktypeId string `json:"worktypeId"`

}

// String returns a JSON representation of the model
func (o *Workitemsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsettings) MarshalJSON() ([]byte, error) {
    type Alias Workitemsettings

    if WorkitemsettingsMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsettingsMarshalled = true

    return json.Marshal(&struct {
        
        WorktypeId string `json:"worktypeId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

