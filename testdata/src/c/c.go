// want package:`c.TestKind:\[c.TestKindHoge c.TestKindFuga c.TestKindPiyo\] c.TestKind2:\[c.TestKind2Hoge c.TestKind2Fuga c.TestKind2Piyo\]`

package c

import (
	"a"
	"b"
	xa "x/a"
)

type TestKind int

const (
	TestKindHoge TestKind = iota
	TestKindFuga
	TestKindPiyo
)

type TestKind2 int

const (
	TestKind2Hoge TestKind2 = iota
	TestKind2Fuga
	TestKind2Piyo
)

func c() {
	{
		var v TestKind

		// no annotation
		switch v {
		case TestKindHoge:
			// do something
		}

		// switchchecker
		switch v {
		case TestKindHoge:
		case TestKindFuga:
		case TestKindPiyo:
		default:
			// do something
		}

		// switchchecker
		switch v { // want "no case of c.TestKindPiyo"
		case TestKindHoge:
		case TestKindFuga:
		default:
			// do something
		}

		// switchchecker -exclude TestKindHoge,TestKindFuga
		switch v { // want "no case of c.TestKindPiyo"
		default:
			// do something
		}

		// switchchecker -e TestKindHoge,TestKindFuga
		switch v { // want "no case of c.TestKindPiyo"
		default:
			// do something
		}
	}

	{
		var v a.TestKind

		// no annotation
		switch v {
		case a.TestKindHoge:
			// do something
		}

		// switchchecker
		switch v {
		case a.TestKindHoge, a.TestKindFuga:
		case a.TestKindPiyo:
		default:
			// do something
		}

		// switchchecker
		switch v { // want "no case of a.TestKindPiyo"
		case a.TestKindHoge:
		case a.TestKindFuga:
		default:
			// do something
		}

		// switchchecker -exclude TestKindHoge,TestKindFuga
		switch v { // want "no case of a.TestKindPiyo"
		default:
			// do something
		}

		// switchchecker -e TestKindHoge,TestKindFuga
		switch v { // want "no case of a.TestKindPiyo"
		default:
			// do something
		}
	}

	{
		var v b.TestKind

		// switchchecker
		switch v { // want "no case of b.TestKindFuga, and b.TestKindPiyo"
		case b.TestKindHoge:
		default:
			// do something
		}
	}

	{
		var v xa.TestKind

		// switchchecker
		switch v { // want "no case of x/a.TestKindDoga, x/a.TestKindBosukete, x/a.TestKindXxx, and more"
		case xa.TestKindHoge:
		case xa.TestKindFuga:
		case xa.TestKindPiyo:
		default:
			// do something
		}
	}
}
