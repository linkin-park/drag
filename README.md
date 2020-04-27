# Drag

    Drag is a web spider which list all the links for the keyword given.

    If you search "hell" it gives all the links for the given keyword (related).

This is an on going project [for me].

## Usage

```
var w crawl.WebSpider
	w.Start(crawl.Config{
		Timeout: 1 * time.Minute, // Timeout for depth ( > 0)
		Search:  "covid",
	})
fmt.Println(w.Results)
```

## Test

```
 cd test
 go test
```

### TODO

- Enhance efficiency
- Refactor
- Html parser
- Functionality & Features
