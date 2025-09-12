package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SummarysettingparticipantlabelsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SummarysettingparticipantlabelsDud struct { 
    


    

}

// Summarysettingparticipantlabels
type Summarysettingparticipantlabels struct { 
    // Internal - Specify how to refer the internal participant of the interaction.
    Internal string `json:"internal"`


    // External - Specify how to refer the external participant of the interaction.
    External string `json:"external"`

}

// String returns a JSON representation of the model
func (o *Summarysettingparticipantlabels) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Summarysettingparticipantlabels) MarshalJSON() ([]byte, error) {
    type Alias Summarysettingparticipantlabels

    if SummarysettingparticipantlabelsMarshalled {
        return []byte("{}"), nil
    }
    SummarysettingparticipantlabelsMarshalled = true

    return json.Marshal(&struct {
        
        Internal string `json:"internal"`
        
        External string `json:"external"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

