package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenericdataingestionruleresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenericdataingestionruleresponseDud struct { 
    


    


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    Platform string `json:"platform"`


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Genericdataingestionruleresponse
type Genericdataingestionruleresponse struct { 
    // Id - ID of the data ingestion rule.
    Id string `json:"id"`


    // Name - The name of the data ingestion rule.
    Name string `json:"name"`


    // Description - A description of the data ingestion rule.
    Description string `json:"description"`


    // Status - The status of the data ingestion rule.
    Status string `json:"status"`


    // Version - The version number of the data ingestion rule.
    Version int `json:"version"`


    


    


    


    // IngestionRuleInfo - The Info about ingestion rule.
    IngestionRuleInfo Messageinfo `json:"ingestionRuleInfo"`


    // Countries - ISO 3166-1 alpha-2 country codes where Data Ingestion Rules should apply. Defaults to worldwide.
    Countries []string `json:"countries"`


    // NonOwnedAccount - The non-owned Instagram account for listening. Can be present when source is External.
    NonOwnedAccount interface{} `json:"nonOwnedAccount"`


    // Hashtags - The hashtags configuration for listening. Can be present when source is External.
    Hashtags interface{} `json:"hashtags"`


    // IntegrationId - The Integration Id from which public social posts are ingested.
    IntegrationId string `json:"integrationId"`


    // SearchTerms - Search terms for X (formally Twitter).
    SearchTerms string `json:"searchTerms"`


    // MessageTypes - Array of message types to fetch when source type is external.
    MessageTypes []string `json:"messageTypes"`


    // PageId - The ID of the Facebook page when source is external.
    PageId string `json:"pageId"`


    // ExternalSource - The external source associated with this data ingestion rule, which is used when performing identity resolution.
    ExternalSource Domainentityref `json:"externalSource"`


    

}

// String returns a JSON representation of the model
func (o *Genericdataingestionruleresponse) String() string {
    
    
    
    
    
    
     o.Countries = []string{""} 
     o.NonOwnedAccount = Interface{} 
     o.Hashtags = Interface{} 
    
    
     o.MessageTypes = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Genericdataingestionruleresponse) MarshalJSON() ([]byte, error) {
    type Alias Genericdataingestionruleresponse

    if GenericdataingestionruleresponseMarshalled {
        return []byte("{}"), nil
    }
    GenericdataingestionruleresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Status string `json:"status"`
        
        Version int `json:"version"`
        
        IngestionRuleInfo Messageinfo `json:"ingestionRuleInfo"`
        
        Countries []string `json:"countries"`
        
        NonOwnedAccount interface{} `json:"nonOwnedAccount"`
        
        Hashtags interface{} `json:"hashtags"`
        
        IntegrationId string `json:"integrationId"`
        
        SearchTerms string `json:"searchTerms"`
        
        MessageTypes []string `json:"messageTypes"`
        
        PageId string `json:"pageId"`
        
        ExternalSource Domainentityref `json:"externalSource"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        Countries: []string{""},
        


        
        NonOwnedAccount: Interface{},
        


        
        Hashtags: Interface{},
        


        


        


        
        MessageTypes: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

