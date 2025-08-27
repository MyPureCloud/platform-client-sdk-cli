package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AdditionalmatchcriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AdditionalmatchcriteriaDud struct { 
    


    

}

// Additionalmatchcriteria
type Additionalmatchcriteria struct { 
    // Topics - List of topics with specific data ingestion rules to filter messages for escalation.
    Topics []Topiccriteria `json:"topics"`


    // MediaFilter - Escalate message based on media presence. Not setting any value will escalate all types of msg.
    MediaFilter string `json:"mediaFilter"`

}

// String returns a JSON representation of the model
func (o *Additionalmatchcriteria) String() string {
     o.Topics = []Topiccriteria{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Additionalmatchcriteria) MarshalJSON() ([]byte, error) {
    type Alias Additionalmatchcriteria

    if AdditionalmatchcriteriaMarshalled {
        return []byte("{}"), nil
    }
    AdditionalmatchcriteriaMarshalled = true

    return json.Marshal(&struct {
        
        Topics []Topiccriteria `json:"topics"`
        
        MediaFilter string `json:"mediaFilter"`
        *Alias
    }{

        
        Topics: []Topiccriteria{{}},
        


        

        Alias: (*Alias)(u),
    })
}

