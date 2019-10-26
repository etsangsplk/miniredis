// +build int

package main

import (
	"testing"
)

func TestStream(t *testing.T) {
	testCommands(t,
		succ("XADD",
			"planets",
			"0-1",
			"name", "Mercury",
		),
		succLoosely("XADD",
			"planets",
			"*",
			"name", "Venus",
		),
		succ("XADD",
			"planets",
			"18446744073709551000-0",
			"name", "Earth",
		),
		fail("XADD",
			"planets",
			"18446744073709551000-0",
			"name", "Earth",
		),
		succ("XLEN", "planets"),
		succ("RENAME", "planets", "planets2"),
		succ("DEL", "planets2"),
		succ("XLEN", "planets"),
	)
}

func TestStreamRange(t *testing.T) {
	testCommands(t,
		succ("XADD",
			"ordplanets",
			"0-1",
			"name", "Mercury",
		),
		succ("XADD",
			"ordplanets",
			"1-0",
			"name", "Venus",
		),
		succ("XADD",
			"ordplanets",
			"2-1",
			"name", "Earth",
		),
		succ("XADD",
			"ordplanets",
			"3-0",
			"name", "Mars",
		),
		succ("XADD",
			"ordplanets",
			"4-1",
			"name", "Jupiter",
		),
		succ("XRANGE", "ordplanets", "-", "+"),
		succ("XRANGE", "ordplanets", "+", "-"),
		succ("XRANGE", "ordplanets", "-", "99"),
		succ("XRANGE", "ordplanets", "0", "4"),
		succ("XRANGE", "ordplanets", "0", "1-0"),
		succ("XRANGE", "ordplanets", "0", "1-99"),
		succ("XRANGE", "ordplanets", "0", "2", "COUNT", "1"),
		succ("XRANGE", "ordplanets", "1-42", "3-42", "COUNT", "1"),

		succ("XREVRANGE", "ordplanets", "+", "-"),
		succ("XREVRANGE", "ordplanets", "-", "+"),
		succ("XREVRANGE", "ordplanets", "4", "0"),
		succ("XREVRANGE", "ordplanets", "1-0", "0"),
		succ("XREVRANGE", "ordplanets", "3-42", "1-0", "COUNT", "2"),
		succ("DEL", "ordplanets"),

		// // failure cases
		fail("XRANGE"),
		fail("XRANGE", "foo"),
		fail("XRANGE", "foo", 1),
		fail("XRANGE", "foo", 2, 3, "toomany"),
		fail("XRANGE", "foo", 2, 3, "COUNT", "noint"),
		fail("XRANGE", "foo", 2, 3, "COUNT", 1, "toomany"),
		fail("XRANGE", "foo", "-", "noint"),
		succ("SET", "str", "I am a string"),
		fail("XRANGE", "str", "-", "+"),
	)
}