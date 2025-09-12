package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IgnorephrasesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IgnorephrasesresponseDud struct { 
    


    


    

}

// Ignorephrasesresponse
type Ignorephrasesresponse struct { 
    // TotalPhrases - Total number of phrases in current org
    TotalPhrases int `json:"totalPhrases"`


    // AddedPhrases - Number of phrases added in current request
    AddedPhrases int `json:"addedPhrases"`


    // UpdatedPhrases - Number of phrases updated in current request
    UpdatedPhrases int `json:"updatedPhrases"`

}

// String returns a JSON representation of the model
func (o *Ignorephrasesresponse) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ignorephrasesresponse) MarshalJSON() ([]byte, error) {
    type Alias Ignorephrasesresponse

    if IgnorephrasesresponseMarshalled {
        return []byte("{}"), nil
    }
    IgnorephrasesresponseMarshalled = true

    return json.Marshal(&struct {
        
        TotalPhrases int `json:"totalPhrases"`
        
        AddedPhrases int `json:"addedPhrases"`
        
        UpdatedPhrases int `json:"updatedPhrases"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

