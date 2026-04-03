package cmdescape_test

import (
    "bufio"
    "bytes"
    "testing"

    "github.com/qr243vbi/cmdescape"
)

func TestQuote(t *testing.T) {
    tests := []struct {
        in   string
        want string
    }{
        {`simple`, `"simple"`},
        {`hello world`, `"hello world"`},
        {`He said "Hi"`, `"He said \"Hi\""`},
        {`a&b`, `"a^&b"`},
        {`x|y`, `"x^|y"`},
        {`<tag>`, `"^<tag^>"`},
        {`(test)`, `"^(test^)"`},
        {`100%`, `"100^^^%"`},
        {`^caret`, `"^^caret"`},
    }

    for _, tt := range tests {
        got := cmdescape.Quote(tt.in)
        if got != tt.want {
            t.Errorf("Quote(%q) = %q, want %q", tt.in, got, tt.want)
        }
    }
}

func TestQuoteCommand(t *testing.T) {
    args := []string{`a b`, `c&d`, `x"y`}
    want := `"a b" "c^&d" "x\"y"`

    got := cmdescape.QuoteCommand(args)
    if got != want {
        t.Errorf("QuoteCommand = %q, want %q", got, want)
    }
}

func TestStripUnsafe(t *testing.T) {
    in := "Hello\x00World\x1b!"
    want := "HelloWorld!"

    got := cmdescape.StripUnsafe(in)
    if got != want {
        t.Errorf("StripUnsafe = %q, want %q", got, want)
    }
}

func TestStripSpaces(t *testing.T) {
    in := "A B\tC\nD"
    want := "ABCD"

    got := cmdescape.StripSpaces(in)
    if got != want {
        t.Errorf("StripSpaces = %q, want %q", got, want)
    }
}

func TestScanTokens(t *testing.T) {
    data := []byte("hello\x00world\x00last")
    scanner := bufio.NewScanner(bytes.NewReader(data))
    scanner.Split(cmdescape.ScanTokens)

    var tokens []string
    for scanner.Scan() {
        tokens = append(tokens, scanner.Text())
    }

    want := []string{"hello", "world", "last"}

    if len(tokens) != len(want) {
        t.Fatalf("expected %d tokens, got %d", len(want), len(tokens))
    }

    for i := range want {
        if tokens[i] != want[i] {
            t.Errorf("token[%d] = %q, want %q", i, tokens[i], want[i])
        }
    }
}

