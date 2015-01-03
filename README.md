# ktoblzcheck for go

tiny golang wrapper around libktoblzcheck

## Example usage

``` go
func main() {
  fmt.Printf("%v\n%v\n%v\n", LibraryVersion(), BankDataDir(), StringEncoding())

  check := NewDefaultAccountNumberCheck()
  // or
  check := NewAccountNumberCheck(BankDataDir() + "/bankdata_20141208.txt")

  fmt.Printf("%d\n", check.BankCount())
  fmt.Printf("%d\n", check.Check("21090900", "123456789"))

  bank, _ := check.FindBank("21090900")
  fmt.Printf("%v\n", bank)
  if _, err := check.FindBank("2109090000"); err != nil {
    fmt.Printf("%v\n", err.Error())
  }

  check.Free()
}

```