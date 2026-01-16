package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ChunksearchregisterrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ChunksearchregisterrequestDud struct { 
    


    


    

}

// Chunksearchregisterrequest
type Chunksearchregisterrequest struct { 
    // SessionId - The unique identifier of this session
    SessionId string `json:"sessionId"`


    // Answered - Mark the chunks search as answered/unanswered
    Answered bool `json:"answered"`


    // SelectedAnswers - The search results selected as answers
    SelectedAnswers []Selectedanswer `json:"selectedAnswers"`

}

// String returns a JSON representation of the model
func (o *Chunksearchregisterrequest) String() string {
    
    
     o.SelectedAnswers = []Selectedanswer{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Chunksearchregisterrequest) MarshalJSON() ([]byte, error) {
    type Alias Chunksearchregisterrequest

    if ChunksearchregisterrequestMarshalled {
        return []byte("{}"), nil
    }
    ChunksearchregisterrequestMarshalled = true

    return json.Marshal(&struct {
        
        SessionId string `json:"sessionId"`
        
        Answered bool `json:"answered"`
        
        SelectedAnswers []Selectedanswer `json:"selectedAnswers"`
        *Alias
    }{

        


        


        
        SelectedAnswers: []Selectedanswer{{}},
        

        Alias: (*Alias)(u),
    })
}

