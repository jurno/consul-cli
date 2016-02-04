package commands

import (
	"github.com/spf13/cobra"
)

type Catalog struct {
	*Cmd
}

func (root *Cmd) initCatalog() {
	c := Catalog{ Cmd: root }

	catalogCmd := &cobra.Command{
		Use: "catalog",
		Short: "Consul Catalog endpoint interface",
		Long: "Consul Catalog endpoint interface",
		Run: func (cmd *cobra.Command, args []string) {
			root.Help()
		},
	}

	c.AddDatacentersSub(catalogCmd)
	c.AddNodeSub(catalogCmd)
	c.AddNodesSub(catalogCmd)
	c.AddServiceSub(catalogCmd)
	c.AddServicesSub(catalogCmd)

	c.AddCommand(catalogCmd)
}
