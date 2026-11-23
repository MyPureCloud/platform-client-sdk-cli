package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SnippetrecordingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SnippetrecordingsettingsDud struct { 
    


    


    

}

// Snippetrecordingsettings
type Snippetrecordingsettings struct { 
    // DaysToArchive - The number of days after a snippet recording has ended before it is archived. Use -1 to never archive, or a positive number
    DaysToArchive int `json:"daysToArchive"`


    // DaysToDelete - The number of days after a snippet recording has ended before it is deleted. Use -1 to retain indefinitely, or a positive number
    DaysToDelete int `json:"daysToDelete"`


    // ExportIntegration - The aws-s3-recording-bulk-actions-integration used for snippet recording exports.
    ExportIntegration Domainentityref `json:"exportIntegration"`

}

// String returns a JSON representation of the model
func (o *Snippetrecordingsettings) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Snippetrecordingsettings) MarshalJSON() ([]byte, error) {
    type Alias Snippetrecordingsettings

    if SnippetrecordingsettingsMarshalled {
        return []byte("{}"), nil
    }
    SnippetrecordingsettingsMarshalled = true

    return json.Marshal(&struct {
        
        DaysToArchive int `json:"daysToArchive"`
        
        DaysToDelete int `json:"daysToDelete"`
        
        ExportIntegration Domainentityref `json:"exportIntegration"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

