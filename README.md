# scago
scago (or **s**ound **c**hange **a**pplier go) is a go library that can apply a sound change ruleset to a word. It's a useful tool in conlanging communities for artificially creating diachronic change in a constructed language.

The library is inspired by KathTheDragon's SCE, a similar tool implemented in Python.

## Usage
### Standalone
Make sure that you have the binary folder of your gopath in your shell path (usually $HOME/go/bin; or $GOPATH/bin if your gopath differs from default)
```
go install github.com/m5ka/scago/cmd/scago@latest
scago -r "a > e / #_" abacus
```

### Library
```go
import scago "github.com/m5ka/scago"

func main() {
    s := scago.New()
    err := s.AddCategory("P", []string{"p", "b", "t", "d", "k", "g"})
    if err != nil {
        fmt.Println("Couldn't add category!", err)
        return
    }
    err = s.AddRule("a > e / _P")
    if err != nil {
        fmt.Println("Couldn't add rule!", err)
        return
    }
    output, err := s.Apply("aba")
    if err != nil {
        fmt.Println("Couldn't apply ruleset!", err)
        return
    }
    fmt.Println(output) // Outputs: eba
}
```
