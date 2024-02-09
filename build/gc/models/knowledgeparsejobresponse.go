package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeparsejobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeparsejobresponseDud struct { 
    Id string `json:"id"`


    


    


    Status string `json:"status"`


    ParseResults []Knowledgeparserecord `json:"parseResults"`


    ImportResult Knowledgeparseimportresult `json:"importResult"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    SelfUri string `json:"selfUri"`

}

// Knowledgeparsejobresponse
type Knowledgeparsejobresponse struct { 
    


    // DownloadURL - The URL of the location at which the caller can download the original html file.
    DownloadURL string `json:"downloadURL"`


    // Hints - Hinted titles for the parser.
    Hints []string `json:"hints"`


    


    


    


    // CreatedBy - The user who created the operation
    CreatedBy Userreference `json:"createdBy"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Knowledgeparsejobresponse) String() string {
    
     o.Hints = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeparsejobresponse) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeparsejobresponse

    if KnowledgeparsejobresponseMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeparsejobresponseMarshalled = true

    return json.Marshal(&struct {
        
        DownloadURL string `json:"downloadURL"`
        
        Hints []string `json:"hints"`
        
        CreatedBy Userreference `json:"createdBy"`
        *Alias
    }{

        


        


        
        Hints: []string{""},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

