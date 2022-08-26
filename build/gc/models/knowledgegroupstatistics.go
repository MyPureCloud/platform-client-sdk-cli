package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgegroupstatisticsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgegroupstatisticsDud struct { 
    


    


    

}

// Knowledgegroupstatistics
type Knowledgegroupstatistics struct { 
    // UnlinkedPhraseCount - Knowledge Group unique phrase count
    UnlinkedPhraseCount int `json:"unlinkedPhraseCount"`


    // UnlinkedPhraseHitCount - Knowledge Group unlinked phrases hit count
    UnlinkedPhraseHitCount int `json:"unlinkedPhraseHitCount"`


    // TotalPhraseHitCount - Total number of phrase hit counts of an unanswered group
    TotalPhraseHitCount int `json:"totalPhraseHitCount"`

}

// String returns a JSON representation of the model
func (o *Knowledgegroupstatistics) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgegroupstatistics) MarshalJSON() ([]byte, error) {
    type Alias Knowledgegroupstatistics

    if KnowledgegroupstatisticsMarshalled {
        return []byte("{}"), nil
    }
    KnowledgegroupstatisticsMarshalled = true

    return json.Marshal(&struct {
        
        UnlinkedPhraseCount int `json:"unlinkedPhraseCount"`
        
        UnlinkedPhraseHitCount int `json:"unlinkedPhraseHitCount"`
        
        TotalPhraseHitCount int `json:"totalPhraseHitCount"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

