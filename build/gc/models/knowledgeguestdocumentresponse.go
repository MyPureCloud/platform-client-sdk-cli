package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeguestdocumentresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeguestdocumentresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    CreatedBy Userreference `json:"createdBy"`


    ModifiedBy Userreference `json:"modifiedBy"`


    


    SessionId string `json:"sessionId"`


    Category Guestcategoryreference `json:"category"`


    


    SelfUri string `json:"selfUri"`

}

// Knowledgeguestdocumentresponse
type Knowledgeguestdocumentresponse struct { 
    


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
    Variations []Knowledgeguestdocumentvariation `json:"variations"`


    

}

// String returns a JSON representation of the model
func (o *Knowledgeguestdocumentresponse) String() string {
    
    
     o.Alternatives = []Knowledgedocumentalternative{{}} 
    
    
    
    
    
    
    
     o.Variations = []Knowledgeguestdocumentvariation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeguestdocumentresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeguestdocumentresponse

    if KnowledgeguestdocumentresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeguestdocumentresponseMarshalled = true

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
        
        Variations []Knowledgeguestdocumentvariation `json:"variations"`
        *Alias
    }{

        


        


        


        
        Alternatives: []Knowledgedocumentalternative{{}},
        


        


        


        


        


        


        


        


        


        


        


        


        
        Variations: []Knowledgeguestdocumentvariation{{}},
        


        

        Alias: (*Alias)(u),
    })
}

