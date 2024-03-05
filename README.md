## Go Questions.

---

- where is the guarantee that tells us the main body for loop waits for `go func` to finish?
- what is the relationship between functions and their immediately called `go routines`?
```go
go func() {
    for i := range 3 {
        msg := fmt.Sprintf("message: #%d", i+1)
        ch <- msg
    }
    close(ch)
}()

for msg := range ch {
    fmt.Println("got: ", msg)
}
```
possible explanations:
- [ ] send and receive will wait until opposite operation happens

---

