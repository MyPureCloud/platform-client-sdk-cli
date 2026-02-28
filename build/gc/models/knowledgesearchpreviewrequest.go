package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgesearchpreviewrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgesearchpreviewrequestDud struct { 
    


    


    


    


    

}

// Knowledgesearchpreviewrequest
type Knowledgesearchpreviewrequest struct { 
    // Query - Query to search content in the knowledge sources.
    Query string `json:"query"`


    // Sources - Source information to search upon.
    Sources []V3sourceref `json:"sources"`


    // GenerationSetting - Setting for answer generation.
    GenerationSetting Knowledgegenerationsetting `json:"generationSetting"`


    // Stateful - Indicates if stateful search and generation is enabled for the knowledge setting.
    Stateful bool `json:"stateful"`


    // ConversationTurns - List of conversation turns to use for stateful search.
    ConversationTurns []Knowledgeconversationturn `json:"conversationTurns"`

}

// String returns a JSON representation of the model
func (o *Knowledgesearchpreviewrequest) String() string {
    
     o.Sources = []V3sourceref{{}} 
    
    
     o.ConversationTurns = []Knowledgeconversationturn{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgesearchpreviewrequest) MarshalJSON() ([]byte, error) {
    type Alias Knowledgesearchpreviewrequest

    if KnowledgesearchpreviewrequestMarshalled {
        return []byte("{}"), nil
    }
    KnowledgesearchpreviewrequestMarshalled = true

    return json.Marshal(&struct {
        
        Query string `json:"query"`
        
        Sources []V3sourceref `json:"sources"`
        
        GenerationSetting Knowledgegenerationsetting `json:"generationSetting"`
        
        Stateful bool `json:"stateful"`
        
        ConversationTurns []Knowledgeconversationturn `json:"conversationTurns"`
        *Alias
    }{

        


        
        Sources: []V3sourceref{{}},
        


        


        


        
        ConversationTurns: []Knowledgeconversationturn{{}},
        

        Alias: (*Alias)(u),
    })
}

