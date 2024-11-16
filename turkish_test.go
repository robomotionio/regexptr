package regexp

import (
	"fmt"
)

func ExampleRegexp_TurkishCharacters_withI() {
	// Add tests for Turkish case-specific characters with Turkish mode
	reTurkishWords := MustCompileTurkish(`(?i)NumarasıIİi`)
	fmt.Println("\nTurkish word case-insensitive matches:")
	fmt.Println(reTurkishWords.MatchString("NUMARASIıiİ"))

	// Output:
	// Turkish word case-insensitive matches:
	// true
}
