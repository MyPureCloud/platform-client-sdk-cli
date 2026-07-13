package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NotesexportqueryconditionsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NotesexportqueryconditionsDud struct { 
    


    

}

// Notesexportqueryconditions
type Notesexportqueryconditions struct { 
    // Filters - Filters to apply on export
    Filters Notesexportfilter `json:"filters"`


    // Limit - Maximum result count in export, default is 180 000 000
    Limit int `json:"limit"`

}

// String returns a JSON representation of the model
func (o *Notesexportqueryconditions) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Notesexportqueryconditions) MarshalJSON() ([]byte, error) {
    type Alias Notesexportqueryconditions

    if NotesexportqueryconditionsMarshalled {
        return []byte("{}"), nil
    }
    NotesexportqueryconditionsMarshalled = true

    return json.Marshal(&struct {
        
        Filters Notesexportfilter `json:"filters"`
        
        Limit int `json:"limit"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

