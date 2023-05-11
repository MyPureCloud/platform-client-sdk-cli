package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DraftmanipulationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DraftmanipulationrequestDud struct { 
    


    

}

// Draftmanipulationrequest
type Draftmanipulationrequest struct { 
    // Translate - A set of definitions to translate email attributes and correctly display date and time, for a given language
    Translate Historyheaderstranslation `json:"translate"`


    // DraftType - The kind of draft that as to be treated. Used to prefix response subject or auto-include information
    DraftType string `json:"draftType"`

}

// String returns a JSON representation of the model
func (o *Draftmanipulationrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Draftmanipulationrequest) MarshalJSON() ([]byte, error) {
    type Alias Draftmanipulationrequest

    if DraftmanipulationrequestMarshalled {
        return []byte("{}"), nil
    }
    DraftmanipulationrequestMarshalled = true

    return json.Marshal(&struct {
        
        Translate Historyheaderstranslation `json:"translate"`
        
        DraftType string `json:"draftType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

