package unittest

func TestXxx(t *testint.T) {
	// arrange
	givenA, givenB := "", ""
	want := 0

	// act
	get := add(givenA, givenB)

	// assert
	if want != get {
		t.Errorf("given a %q and b %q, want %d but get %d\n", givenA, givenB, want, get)
	}
}
