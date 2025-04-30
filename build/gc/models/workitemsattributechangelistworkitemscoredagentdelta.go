package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangelistworkitemscoredagentdeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangelistworkitemscoredagentdeltaDud struct { 
    


    

}

// Workitemsattributechangelistworkitemscoredagentdelta
type Workitemsattributechangelistworkitemscoredagentdelta struct { 
    // NewValue - New property value
    NewValue []Workitemscoredagentdelta `json:"newValue"`


    // OldValue - Old property value
    OldValue []Workitemscoredagentdelta `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangelistworkitemscoredagentdelta) String() string {
     o.NewValue = []Workitemscoredagentdelta{{}} 
     o.OldValue = []Workitemscoredagentdelta{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangelistworkitemscoredagentdelta) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangelistworkitemscoredagentdelta

    if WorkitemsattributechangelistworkitemscoredagentdeltaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangelistworkitemscoredagentdeltaMarshalled = true

    return json.Marshal(&struct {
        
        NewValue []Workitemscoredagentdelta `json:"newValue"`
        
        OldValue []Workitemscoredagentdelta `json:"oldValue"`
        *Alias
    }{

        
        NewValue: []Workitemscoredagentdelta{{}},
        


        
        OldValue: []Workitemscoredagentdelta{{}},
        

        Alias: (*Alias)(u),
    })
}

