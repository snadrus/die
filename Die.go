package die

// Get Usage: a, b := die.G2(some run)
func Get[T any](a T, err error) T {
  if err != nil {
    panic(err)
  }
  return a
}

func Gets[T, U any](a T, b U, err error) (T, U) {
  if err != nil {
    panic(err)
  }
  return a, b
}

alias Get2 = Gets

func Get3[T, U, V any](a T, b U, c V, err error) (T, U, V) {
  if err != nil {
    panic(err)
  }
  return a, b, c
}

func Get4[T, U, V, W any](a T, b U, c V, d W, err error) (T, U, V, W) {
  if err != nil {
    panic(err)
  }
  return a, b, c, d
}

alias G2 = Get2
alias G3 = Get3
alias G4 = Get4


func Wrap[T any](a T, err error) wrapped1 {
  if err != nil {
    panic(err)
  }
  return a
}

type wrapped1[T] struct {
  A T
  err error
}

// If usage: die.If(strconv.Atoi("1")).Say("Bad number input")
func If[T any](a T, err error) wrapped1 {
  return wrapped1{a, err}
}

func (w wrapped[T]) Say(format string) T {
  if err != nil {
    panic(fmt.Errorf(format, w.err))
  }
  return w.a
}

func (w wrapped[T]) Ok(format string) {
  if err != nil {
    panic(w.err)
  }
  return w.a
}

// Append usage: defer die.Append("In MakeFoo")
// Presume errors are a string.
// TODO save original stack. Restore a save
func Append(format string) {
  if v:= recover(); v!= nil {
    s := fmt.Sprint(v)
    panic(fmt.Errorf(format, s))
  }
}

// TODO: Get err as a defer
func ToErr(e *error) {
  if v:= recover(); v!= nil {
    *e = v
  }
}
