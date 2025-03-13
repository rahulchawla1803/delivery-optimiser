package output

import (
	"fmt"

	"github.com/rahulchawla1803/delivery-optimiser/internal/utils"
)

// Fail handles failure scenarios by logging errors
func Fail(err error) error {
	// Step 1: Prepare Output Directory
	_ = PrepareOutputDirectory() // Ensure output directory exists

	// Step 2: Write Error Log
	errorLog := fmt.Sprintf("Execution failed: %v\n", err)
	_ = utils.WriteText("output/error.log", errorLog)

	fmt.Println("Error encountered. Check error.log for details.")
	return err
}
