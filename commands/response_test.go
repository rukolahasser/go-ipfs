package commands

import "testing"

type TestOutput struct {
  Foo, Bar string
  Baz int
}

func TestMarshalling(t *testing.T) {
  req := NewRequest()

  res := Response{
    req: req,
    Value: TestOutput{ "beep", "boop", 1337 },
  }

  _, err := res.Marshal()
  if err == nil {
    t.Error("Should have failed (no encoding type specified in request)")
  }

  req.options["enc"] = Json
  bytes, err := res.Marshal()
  if err != nil {
    t.Error("Should have passed")
  }
  output := string(bytes)
  if output != "{\"Foo\":\"beep\",\"Bar\":\"boop\",\"Baz\":1337}" {
    t.Error("Incorrect JSON output")
  }
}
