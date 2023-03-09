package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationparticipantsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationparticipantsearchrequestDud struct { 
    


    


    


    


    


    

}

// Conversationparticipantsearchrequest
type Conversationparticipantsearchrequest struct { 
    // SortOrder - The sort order for results
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the resource that you want to sort the results by
    SortBy string `json:"sortBy"`


    // Sort - Multi-value sort order, list of multiple sort values
    Sort []Searchsort `json:"sort"`


    // ReturnFields
    ReturnFields []string `json:"returnFields"`


    // Query
    Query []Conversationparticipantsearchcriteria `json:"query"`


    // Cursor
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Conversationparticipantsearchrequest) String() string {
    
    
     o.Sort = []Searchsort{{}} 
     o.ReturnFields = []string{""} 
     o.Query = []Conversationparticipantsearchcriteria{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationparticipantsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Conversationparticipantsearchrequest

    if ConversationparticipantsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    ConversationparticipantsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        Sort []Searchsort `json:"sort"`
        
        ReturnFields []string `json:"returnFields"`
        
        Query []Conversationparticipantsearchcriteria `json:"query"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        


        


        
        Sort: []Searchsort{{}},
        


        
        ReturnFields: []string{""},
        


        
        Query: []Conversationparticipantsearchcriteria{{}},
        


        

        Alias: (*Alias)(u),
    })
}

