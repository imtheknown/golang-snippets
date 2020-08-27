# ServeMux

 ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

```javascript 
type ServeMux struct {
    // contains filtered or unexported fields
}
```
### NewServeMux()

```javascript 
func NewServeMux() *ServeMux
```
 NewServeMux allocates and returns a new ServeMux. 