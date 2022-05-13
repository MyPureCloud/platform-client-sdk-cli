package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChangeDud struct { 
    


    


    


    

}

// Change
type Change struct { 
    // Entity
    Entity Auditentity `json:"entity"`


    // Property - The property that was changed
    Property string `json:"property"`


    // OldValues - The old values which were modified and/or removed by this action.
    OldValues []string `json:"oldValues"`


    // NewValues - The new values which were modified and/or added by this action.
    NewValues []string `json:"newValues"`

}

// String returns a JSON representation of the model
func (o *Change) String() string {
    
    
     o.OldValues = []string{""} 
     o.NewValues = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Change) MarshalJSON() ([]byte, error) {
    type Alias Change

    if ChangeMarshalled {
        return []byte("{}"), nil
    }
    ChangeMarshalled = true

    return json.Marshal(&struct {
        
        Entity Auditentity `json:"entity"`
        
        Property string `json:"property"`
        
        OldValues []string `json:"oldValues"`
        
        NewValues []string `json:"newValues"`
        *Alias
    }{

        


        


        
        OldValues: []string{""},
        


        
        NewValues: []string{""},
        

        Alias: (*Alias)(u),
    })
}

