package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatedraftrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatedraftrequestDud struct { 
    


    

}

// Createdraftrequest
type Createdraftrequest struct { 
    // DraftType - The kind of draft to create (Reply, ReplyAll, Forward).
    DraftType string `json:"draftType"`


    // Translation - A set of definitions to translate email attributes and correctly display date and time, for a given language.
    Translation Historyheaderstranslation `json:"translation"`

}

// String returns a JSON representation of the model
func (o *Createdraftrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createdraftrequest) MarshalJSON() ([]byte, error) {
    type Alias Createdraftrequest

    if CreatedraftrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatedraftrequestMarshalled = true

    return json.Marshal(&struct {
        
        DraftType string `json:"draftType"`
        
        Translation Historyheaderstranslation `json:"translation"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

