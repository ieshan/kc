package namespace

import (
	"github.com/spf13/cobra"
)

var NsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Work with kubernates namespace",
}
