negotiation
===========

**negotiation** is a standalone library that allows you to implement [content negotiation](http://www.w3.org/Protocols/rfc2616/rfc2616-sec12.html)
in your application, whatever framework you use.
This provides specific functions to negotiate `Accept` and `Accept-Language`
headers, although any kind of header can be parsed.

## Usage

### Language negotiation

```go
package main

import (
  "fmt"
  "github.com/K-Phoen/negotiation"
)

func main() {
  language, err := negotiation.NegotiateLanguage("da, en-gb;q=0.8, en;q=0.7", []string{"es", "fr", "en"})

  if err != nil {
    fmt.Println("Unable to negotiate the language")
  }

  fmt.Println("Negotiated language: ", language.Value) // outputs: "Negotiated language: en"
}
```

### Format negotiation

```go
package main

import (
  "fmt"
  "github.com/K-Phoen/negotiation"
)

func main() {
  format, err := negotiation.NegotiateAccept("application/rdf+xml;q=0.5,text/html;q=.3", []string{"text/html"})

  if err != nil {
    fmt.Println("Unable to negotiate the format")
  }

  fmt.Println("Negotiated format: ", format.Value) // outputs: "Negotiated format: text/html"
}
```

## Tests

```bash
$ go test
```

## License

This library is released under the MIT License. See the bundled LICENSE file for
details.
