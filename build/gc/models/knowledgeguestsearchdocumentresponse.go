package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestsearchdocumentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestsearchdocumentresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    CreatedBy Userreference `json:"createdBy"`


    ModifiedBy Userreference `json:"modifiedBy"`


    


    SessionId string `json:"sessionId"`


    Category Guestcategoryreference `json:"category"`


    


    


    SelfUri string `json:"selfUri"`

}

// Knowledgeguestsearchdocumentresponse
type Knowledgeguestsearchdocumentresponse struct { 
    


    // Title - Document title, having a limit of 500 words.
    Title string `json:"title"`


    // Visible - Indicates if the knowledge document should be included in search results.
    Visible bool `json:"visible"`


    // Alternatives - List of alternate phrases related to the title which improves search results.
    Alternatives []Knowledgedocumentalternative `json:"alternatives"`


    // State - State of the document.
    State string `json:"state"`


    // DateCreated - Document creation date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Document last modification date-time. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // DateImported - Document import date-time, or null if was not imported. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateImported time.Time `json:"dateImported"`


    // LastPublishedVersionNumber - The last published version number of the document.
    LastPublishedVersionNumber int `json:"lastPublishedVersionNumber"`


    // DatePublished - The date on which the document was last published. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DatePublished time.Time `json:"datePublished"`


    


    


    // DocumentVersion - The version of the document.
    DocumentVersion Addressableentityref `json:"documentVersion"`


    


    


    // Variations - Variations of the document.
    Variations []Knowledgeguestdocumentvariationanswer `json:"variations"`


    // Answer - The answer to the query.
    Answer string `json:"answer"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgeguestsearchdocumentresponse) String() string {
    
    
     o.Alternatives = []Knowledgedocumentalternative{{}} 
    
    
    
    
    
    
    
     o.Variations = []Knowledgeguestdocumentvariationanswer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestsearchdocumentresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestsearchdocumentresponse

    if KnowledgeguestsearchdocumentresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestsearchdocumentresponseMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Visible bool `json:"visible"`
        
        Alternatives []Knowledgedocumentalternative `json:"alternatives"`
        
        State string `json:"state"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        DateImported time.Time `json:"dateImported"`
        
        LastPublishedVersionNumber int `json:"lastPublishedVersionNumber"`
        
        DatePublished time.Time `json:"datePublished"`
        
        DocumentVersion Addressableentityref `json:"documentVersion"`
        
        Variations []Knowledgeguestdocumentvariationanswer `json:"variations"`
        
        Answer string `json:"answer"`
        *Alias
    }{

        


        


        


        
        Alternatives: []Knowledgedocumentalternative{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        
        Variations: []Knowledgeguestdocumentvariationanswer{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

