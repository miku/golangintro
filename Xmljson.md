# XML and JSON

Realized with struct tags. Marshaling and Unmarshaling interface, encoding and
decoding.

Tools to generate structs from data:

* [JSONGen](https://github.com/bemasher/JSONGen)
* [zek](https://github.com/miku/zek)
* and more ...

## Examples

```go
$ curl http://api.crossref.org/works/10.1186/s12862-017-0967-2 | JSONGen
type _ struct {
	Message struct {
		AlternativeId []string `json:"alternative-id"`
		ArticleNumber string   `json:"article-number"`
		Author        []struct {
			Affiliation        []interface{} `json:"affiliation"`
			AuthenticatedOrcid bool          `json:"authenticated-orcid"`
			Family             string        `json:"family"`
			Given              string        `json:"given"`
			ORCID              string
			Sequence           string `json:"sequence"`
		} `json:"author"`
		ContainerTitle []string `json:"container-title"`
		ContentDomain  struct {
			CrossmarkRestriction bool     `json:"crossmark-restriction"`
			Domain               []string `json:"domain"`
		} `json:"content-domain"`
		Created struct {
			DateParts []int64 `json:"date-parts"`
			DateTime  string  `json:"date-time"`
			Timestamp int64   `json:"timestamp"`
		} `json:"created"`
		DOI       string
		Deposited struct {
			DateParts []int64 `json:"date-parts"`
			DateTime  string  `json:"date-time"`
			Timestamp int64   `json:"timestamp"`
		} `json:"deposited"`
		Funder []struct {
			Award         []string `json:"award"`
			DOI           string
			DoiAssertedBy string `json:"doi-asserted-by"`
			Name          string `json:"name"`
		} `json:"funder"`
		ISSN    []string
		Indexed struct {
			DateParts []int64 `json:"date-parts"`
			DateTime  string  `json:"date-time"`
			Timestamp int64   `json:"timestamp"`
		} `json:"indexed"`
		IsReferencedByCount int64 `json:"is-referenced-by-count"`
		IssnType            []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"issn-type"`
		Issue  string `json:"issue"`
		Issued struct {
			DateParts []int64 `json:"date-parts"`
		} `json:"issued"`
		JournalIssue struct {
			Issue          string `json:"issue"`
			PublishedPrint struct {
				DateParts []int64 `json:"date-parts"`
			} `json:"published-print"`
		} `json:"journal-issue"`
		Language string `json:"language"`
		Link     []struct {
			ContentType         string `json:"content-type"`
			ContentVersion      string `json:"content-version"`
			IntendedApplication string `json:"intended-application"`
			URL                 string
		} `json:"link"`
		Member          string        `json:"member"`
		OriginalTitle   []interface{} `json:"original-title"`
		Prefix          string        `json:"prefix"`
		PublishedOnline struct {
			DateParts []int64 `json:"date-parts"`
		} `json:"published-online"`
		PublishedPrint struct {
			DateParts []int64 `json:"date-parts"`
		} `json:"published-print"`
		Publisher string `json:"publisher"`
		Reference []struct {
			Author        string `json:"author"`
			DOI           string
			DoiAssertedBy string `json:"doi-asserted-by"`
			Edition       string `json:"edition"`
			FirstPage     string `json:"first-page"`
			JournalTitle  string `json:"journal-title"`
			Key           string `json:"key"`
			Unstructured  string `json:"unstructured"`
			Volume        string `json:"volume"`
			VolumeTitle   string `json:"volume-title"`
			Year          string `json:"year"`
		} `json:"reference"`
		ReferenceCount  int64 `json:"reference-count"`
		ReferencesCount int64 `json:"references-count"`
		Relation        struct {
			Cites []interface{} `json:"cites"`
		} `json:"relation"`
		Score               float64       `json:"score"`
		ShortContainerTitle []string      `json:"short-container-title"`
		ShortTitle          []interface{} `json:"short-title"`
		Source              string        `json:"source"`
		Subtitle            []interface{} `json:"subtitle"`
		Title               []string      `json:"title"`
		Type                string        `json:"type"`
		URL                 string
		UpdatePolicy        string `json:"update-policy"`
		Volume              string `json:"volume"`
	} `json:"message"`
	MessageType    string `json:"message-type"`
	MessageVersion string `json:"message-version"`
	Status         string `json:"status"`
}
```

For XML.

```go
$ curl -sL https://git.io/fx80d | zek -e -c
// Records was generated 2018-10-12 00:38:12 by tir on apollo.local.
type Records struct {
	XMLName xml.Name `xml:"Records"`
	Text    string   `xml:",chardata"` // \n
	Xsi     string   `xml:"xsi,attr"`
	Record  []struct {
		Text   string `xml:",chardata"`
		Header struct {
			Text       string `xml:",chardata"`
			Status     string `xml:"status,attr"`
			Identifier string `xml:"identifier"` // oai:ojs.localhost:article...
			Datestamp  string `xml:"datestamp"`  // 2009-06-24T14:48:23Z, 200...
			SetSpec    string `xml:"setSpec"`    // eppp:ART, eppp:ART, eppp:...
		} `xml:"header"`
		Metadata struct {
			Text    string `xml:",chardata"`
			Rfc1807 struct {
				Text           string   `xml:",chardata"`
				Xmlns          string   `xml:"xmlns,attr"`
				Xsi            string   `xml:"xsi,attr"`
				SchemaLocation string   `xml:"schemaLocation,attr"`
				BibVersion     string   `xml:"bib-version"`  // v2, v2, v2, v2, v2, v2, v...
				ID             string   `xml:"id"`           // http://journals.zpid.de/i...
				Entry          string   `xml:"entry"`        // 2009-06-24T14:48:23Z, 200...
				Organization   []string `xml:"organization"` // Proceedings of the Worksh...
				Title          string   `xml:"title"`        // Introduction and some Ide...
				Type           string   `xml:"type"`
				Author         []string `xml:"author"`       // KRAMPEN, Günter, CARBON,...
				Copyright      string   `xml:"copyright"`    // Das Urheberrecht liegt be...
				OtherAccess    string   `xml:"other_access"` // url:http://journals.zpid....
				Keyword        string   `xml:"keyword"`
				Period         []string `xml:"period"`
				Monitoring     string   `xml:"monitoring"`
				Language       string   `xml:"language"` // en, en, en, en, en, en, e...
				Abstract       string   `xml:"abstract"` // After a short description...
				Date           string   `xml:"date"`     // 2009-06-22 12:12:00, 2009...
			} `xml:"rfc1807"`
		} `xml:"metadata"`
		About string `xml:"about"`
	} `xml:"Record"`
}
```
