package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DraftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DraftDud struct { 
    Id string `json:"id"`


    


    Miner Miner `json:"miner"`


    Intents []Draftintents `json:"intents"`


    Topics []Drafttopics `json:"topics"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Draft
type Draft struct { 
    


    // Name - Draft name
    Name string `json:"name"`


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Draft) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Draft) MarshalJSON() ([]byte, error) {
    type Alias Draft

    if DraftMarshalled {
        return []byte("{}"), nil
    }
    DraftMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

