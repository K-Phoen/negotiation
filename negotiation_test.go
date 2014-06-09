package negotiation

import "testing"

type alternativeTestData struct {
  header       string
  alternatives []string
  choice       string
}

var pearAcceptHeader = "text/html,application/xhtml+xml,application/xml;q=0.9,text/*;q=0.7,*/*,image/gif; q=0.8, image/jpeg; q=0.6, image/*"

var negotiateAcceptTestData = []alternativeTestData{
  // valid cases
  {pearAcceptHeader, []string{"image/gif", "image/png", "application/xhtml+xml", "application/xml", "text/html", "image/jpeg", "text/plain"}, "text/html"},
  {pearAcceptHeader, []string{"image/gif", "image/png", "application/xhtml+xml", "application/xml", "image/jpeg", "text/plain"}, "application/xhtml+xml"},
  {pearAcceptHeader, []string{"image/gif", "image/png", "application/xml", "image/jpeg", "text/plain"}, "application/xml"},
  {pearAcceptHeader, []string{"image/gif", "image/png", "image/jpeg", "text/plain"}, "image/gif"},
  {pearAcceptHeader, []string{"image/png", "image/jpeg", "text/plain"}, "text/plain"},
  {pearAcceptHeader, []string{"image/png"}, "image/png"},
  {pearAcceptHeader, []string{"audio/midi"}, "audio/midi"},
  {"text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8", []string{"application/rss+xml", "*/*"}, "application/rss+xml"},
  {"text/html", []string{"application/rss"}, ""},
  {"application/*", []string{"application/rss"}, "application/rss"},
  {"application/rdf+xml;q=0.5,text/html;q=.3", []string{"application/rdf+xml", "text/html"}, "application/rdf+xml"},
  {"application/rdf+xml;q=0.5,text/html;q=.3", []string{"application/rdf+xml"}, "application/rdf+xml"},
  {"application/rdf+xml;q=0.5,text/html;q=.3", []string{"text/html"}, "text/html"},
  {"application/rdf+xml;q=0.5,text/html;q=.5", []string{"text/html"}, "text/html"},
  {"application/rdf+xml;q=0.5,text/html;q=.5", []string{"application/rdf+xml"}, "application/rdf+xml"},
  {"image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, */*", []string{"text/html", "application/xhtml+xml", "*/*"}, "text/html"},
}

func TestNegotiateAccept(t *testing.T) {
  for _, pair := range negotiateAcceptTestData {
    var result string
    alternative, _ := NegotiateAccept(pair.header, pair.alternatives)

    if alternative == nil {
      result = ""
    } else {
      result = alternative.Value
    }

    if result != pair.choice {
      t.Errorf("For \"%v\" expected \"%v\" but got \"%s\"", pair.header, pair.choice, result)
    }
  }
}

var negotiateInvalidTestData = []alternativeTestData{
  // invalid data
  {"text", []string{"text/plain"}, ""},
  {"text/", []string{"text/plain"}, ""},

  // no match found
  {"text/plain", []string{"text/html"}, ""},
  {"text/*", []string{"image/png"}, ""},
}

func TestNegotiateWithInvalidData(t *testing.T) {
  for _, pair := range negotiateInvalidTestData {
    _, err := NegotiateAccept(pair.header, pair.alternatives)

    if err == nil {
      t.Error("error expected")
    }
  }
}

var negotiateLanguageTestData = []alternativeTestData{
  {"en,fr", []string{"en", "fr"}, "en"},
  {"en; q=0.1, fr; q=0.4, fu; q=0.9, de; q=0.2", []string{"fu", "fr"}, "fu"},
  {"da, en-gb;q=0.8, en;q=0.7", []string{"da", "en"}, "da"},
  {"da, en-gb;q=0.8, en;q=0.7, *", []string{"da", "en"}, "da"},
  {"da, en-gb;q=0.8, en;q=0.7, *", []string{"en"}, "en"},
  {"fr-FR,en-US;q=0.6,en;q=0.4", []string{"fr"}, "fr"},
}

func TestNegotiateLanguage(t *testing.T) {
  for _, pair := range negotiateLanguageTestData {
    var result string
    alternative, _ := NegotiateLanguage(pair.header, pair.alternatives)

    if alternative == nil {
      result = ""
    } else {
      result = alternative.Value
    }

    if result != pair.choice {
      t.Errorf("For \"%v\" expected \"%v\" but got \"%s\"", pair.header, pair.choice, result)
    }
  }
}
