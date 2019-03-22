# Go Maybe

Maybes for Go. This package contains definitions of maybe wrappers for a few basic and often-used types in Go. In case you don't know what a maybe is, you can read about it for example [here](https://en.wikipedia.org/wiki/Option_type).

Wait! Go already has maybes, right? It has pointers.

Yes, it does (and that's what the maybe types use internally too).

So what's the motivation behind this?

1) Working with pointers is pretty seemless in Go. Maybe types do not introduce much new in this. Instead of doing ...

```go
maybeString := someFunctionReturningAPointerToString()

if maybeString == nil {
  // handle nil case
} else {
  // deref and do stuff with the string
}
```

... it allows you to do ...

```go
maybeString := someFunctionReturningAMaybeString()

if maybeString.HasValue() {
  str := maybeString.Get()
  // do stuff with the string
} else {
  // handle nil case
}
```

It might or might not be more expressive, depending on your opinion, nonetheless if there were no other reason it would pretty much be just extra code.

2) It's a custom type and therefore allows you to implement the [Scanner](https://golang.org/pkg/database/sql/#Scanner) and [Valuer](https://golang.org/pkg/database/sql/driver/#Valuer) interfaces. And that's what the maybe types do. Therefore you can use them directly to unmarshal data from your SQL queries without the need of an intermediary struct.

```go
type User struct{
  OptionalName maybe.String
}

func main() {
  user := User{}
  row := db.QueryRow("SELECT name FROM users WHERE id = 1")
  row.Scan(&user.OptionalName)
}
```

Okay, but why not just use `pq`'s [NullTime](https://github.com/lib/pq/blob/8c6ee72f3e6bcb1542298dd5f76cb74af9742cec/encode.go#L586-L589) and other nullable structs?

You certainly can if you're fine with having db-specific types in your structs. Also IMHO this family of types allows you to produce invalid states which is what I like to avoid. For example ...

```go
func main() {
  // Ooops. It's not valid, but it does have a value
  value := pq.NullTime{Valid: false, Time: time.Now()}

  // Ooops. It's valid, but where's the value?
  value := pq.NullTime{Valid: true}

  // Okay, it is there, but it's probably not what you like :-)
  fmt.Println(fmt.Sprintf("%v", value.Time))
}
```

## Usage

```go
func main() {
  m := maybe.NewInt(nil)

  if !m.HasValue() {
    fmt.Println("That's right, there's no value here")
  }

  value, err := m.SafeGet()

  if err != nil {
    fmt.Println("There's really no value here")
  }

  number := 12
  m := maybe.NewInt(&number)

  if m.HasValue() {
    fmt.Println("Yes, now it's safe to call Get(): %d", m.Get())
  }
}
```

Currently there are maybe types defined for `int`, `int8`, `int16`, `int32`, `int64`, `float32`, `float64`, `string` and `time.Time`. `time.Time` expects to get a unix timestamp in the `Scan` method, because I ran into some trouble parsing the postgres timestamp format. If you're reading from a table with the `timestamptz` column `stamp`, you can get it like this: `SELECT extract(epoch from stamp)::integer from mytable.`
