package topdown

import (
	"errors"
	"fmt"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/topdown/builtins"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

type multiplier int64

const (
	base = 1024

	none multiplier = 1
	kb              = base
	mb              = kb * base
	gb              = mb * base
	tb              = gb * base
)

func builtinNumBytes(a ast.Value) (ast.Value, error) {
	var m multiplier

	numPattern := `([0-9]+)`
	strPattern := `([a-z]+)`

	raw, err := builtins.StringOperand(a, 1)
	if err != nil {
		return nil, err
	}

	s := strings.ToLower(raw.String())

	numRe, err := regexp.Compile(numPattern)
	if err != nil {
		return nil, err
	}

	strRe, err := regexp.Compile(strPattern)
	if err != nil {
		return nil, err
	}

	numStr := numRe.FindString(s)
	unit := strRe.FindString(s)

	if numStr == "" {
		return nil, errors.New("no number provided")
	}

	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return nil, err
	}

	switch unit {
	case "":
		m = none
	case "kb", "k":
		m = kb
	case "mb", "m":
		m = mb
	case "gb", "g":
		m = gb
	case "tb", "t":
		m = tb
	default:
		return nil, errors.New(fmt.Sprintf("unit %s not recognized", unit))
	}

	total := num * int64(m)

	return builtins.IntToNumber(big.NewInt(total)), nil
}

func init() {
	RegisterFunctionalBuiltin1(ast.NumBytes.Name, builtinNumBytes)
}
