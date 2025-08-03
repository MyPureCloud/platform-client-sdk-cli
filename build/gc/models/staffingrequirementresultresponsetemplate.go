package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffingrequirementresultresponsetemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffingrequirementresultresponsetemplateDud struct { 
    

}

// Staffingrequirementresultresponsetemplate
type Staffingrequirementresultresponsetemplate struct { 
    // RequirementResults - List of staffing requirement results
    RequirementResults []Planninggrouprequirementoutput `json:"requirementResults"`

}

// String returns a JSON representation of the model
func (o *Staffingrequirementresultresponsetemplate) String() string {
     o.RequirementResults = []Planninggrouprequirementoutput{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffingrequirementresultresponsetemplate) MarshalJSON() ([]byte, error) {
    type Alias Staffingrequirementresultresponsetemplate

    if StaffingrequirementresultresponsetemplateMarshalled {
        return []byte("{}"), nil
    }
    StaffingrequirementresultresponsetemplateMarshalled = true

    return json.Marshal(&struct {
        
        RequirementResults []Planninggrouprequirementoutput `json:"requirementResults"`
        *Alias
    }{

        
        RequirementResults: []Planninggrouprequirementoutput{{}},
        

        Alias: (*Alias)(u),
    })
}

