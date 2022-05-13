package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LineidMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LineidDud struct { 
    


    

}

// Lineid - User information for a Line account
type Lineid struct { 
    // Ids - The set of Line userIds that this person has. Each userId is specific to the Line channel that the user interacts with.
    Ids []Lineuserid `json:"ids"`


    // DisplayName - The displayName of this person's account in Line
    DisplayName string `json:"displayName"`

}

// String returns a JSON representation of the model
func (o *Lineid) String() string {
     o.Ids = []Lineuserid{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lineid) MarshalJSON() ([]byte, error) {
    type Alias Lineid

    if LineidMarshalled {
        return []byte("{}"), nil
    }
    LineidMarshalled = true

    return json.Marshal(&struct {
        
        Ids []Lineuserid `json:"ids"`
        
        DisplayName string `json:"displayName"`
        *Alias
    }{

        
        Ids: []Lineuserid{{}},
        


        

        Alias: (*Alias)(u),
    })
}

