package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchupdaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchupdaterequestDud struct { 
    SessionId string `json:"sessionId"`


    


    


    

}

// Searchupdaterequest
type Searchupdaterequest struct { 
    


    // Answered - Mark the search as answered/unanswered
    Answered bool `json:"answered"`


    // SelectedAnswer - The selected search result chosen as the answer.
    SelectedAnswer Selectedanswer `json:"selectedAnswer"`


    // SelectedAnswers - The search results selected as answers
    SelectedAnswers []Selectedanswer `json:"selectedAnswers"`

}

// String returns a JSON representation of the model
func (o *Searchupdaterequest) String() string {
    
    
     o.SelectedAnswers = []Selectedanswer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchupdaterequest) MarshalJSON() ([]byte, error) {
    type Alias Searchupdaterequest

    if SearchupdaterequestMarshalled {
        return []byte("{}"), nil
    }
    SearchupdaterequestMarshalled = true

    return json.Marshal(&struct {
        
        Answered bool `json:"answered"`
        
        SelectedAnswer Selectedanswer `json:"selectedAnswer"`
        
        SelectedAnswers []Selectedanswer `json:"selectedAnswers"`
        *Alias
    }{

        


        


        


        
        SelectedAnswers: []Selectedanswer{{}},
        

        Alias: (*Alias)(u),
    })
}

