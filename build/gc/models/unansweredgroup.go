package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UnansweredgroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UnansweredgroupDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Unansweredgroup
type Unansweredgroup struct { 
    


    // Label - Knowledge base unanswered group label
    Label string `json:"label"`


    // PhraseGroups - Represents a list of phrase groups inside an unanswered group
    PhraseGroups []Unansweredphrasegroup `json:"phraseGroups"`


    // SuggestedDocuments - Represents a list of documents that may be linked to an unanswered group
    SuggestedDocuments []Unansweredgroupsuggesteddocument `json:"suggestedDocuments"`


    // Statistics - Statistics object containing the various hit counts for an unanswered group
    Statistics Knowledgegroupstatistics `json:"statistics"`


    

}

// String returns a JSON representation of the model
func (o *Unansweredgroup) String() string {
    
     o.PhraseGroups = []Unansweredphrasegroup{{}} 
     o.SuggestedDocuments = []Unansweredgroupsuggesteddocument{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Unansweredgroup) MarshalJSON() ([]byte, error) {
    type Alias Unansweredgroup

    if UnansweredgroupMarshalled {
        return []byte("{}"), nil
    }
    UnansweredgroupMarshalled = true

    return json.Marshal(&struct {
        
        Label string `json:"label"`
        
        PhraseGroups []Unansweredphrasegroup `json:"phraseGroups"`
        
        SuggestedDocuments []Unansweredgroupsuggesteddocument `json:"suggestedDocuments"`
        
        Statistics Knowledgegroupstatistics `json:"statistics"`
        *Alias
    }{

        


        


        
        PhraseGroups: []Unansweredphrasegroup{{}},
        


        
        SuggestedDocuments: []Unansweredgroupsuggesteddocument{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

