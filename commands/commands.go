package commands

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Cmd struct {
	root *cobra.Command

	Err io.Writer
	Out io.Writer

	consul *consul

	Template string
}

func Init() *Cmd {
	c := Cmd{
		Err: os.Stderr,
		Out: os.Stdout,
		consul: &consul{
			auth: new(auth),
		},
	}

	c.root = &cobra.Command{
		Use:   "consul-cli",
		Short: "Command line interface for Consul HTTP API",
		Long:  "Command line interface for Consul HTTP API",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c.root.Help()
			return nil
		},
	}

	c.root.PersistentFlags().StringVar(&c.consul.address, "consul", "127.0.0.1:8500", "Consul address:port")
	c.root.PersistentFlags().BoolVar(&c.consul.sslEnabled, "ssl", false, "Use HTTPS when talking to Consul")
	c.root.PersistentFlags().BoolVar(&c.consul.sslVerify, "ssl-verify", true, "Verify certificates when connecting via SSL")
	c.root.PersistentFlags().StringVar(&c.consul.sslCert, "ssl-cert", "", "Path to an SSL client certificate for authentication")
	c.root.PersistentFlags().StringVar(&c.consul.sslCaCert, "ssl-ca-cert", "", "Path to a CA certificate file to validate the Consul server")
	c.root.PersistentFlags().Var((*auth)(c.consul.auth), "auth", "The HTTP basic authentication username (and optional password) separated by a colon")
	c.root.PersistentFlags().StringVar(&c.consul.token, "token", "", "The Consul ACL token")

	c.initAcl()
	c.initAgent()
	c.initCatalog()
	c.initCheck()
	c.initHealth()
	c.initKv()
	c.initService()
	c.initSession()
	c.initStatus()

	return &c
}

func (c *Cmd) Execute() error {
	return c.root.Execute()
}

func (c *Cmd) Help() error {
	return c.root.Help()
}

func (c *Cmd) AddCommand(cmd *cobra.Command) {
	c.root.AddCommand(cmd)
}

func (c *Cmd) AddTemplateOption(cmd *cobra.Command) {
	cmd.Flags().StringVar(&c.Template, "template", "", "Output template")
}

type funcVar func(s string) error

func (f funcVar) Set(s string) error { return f(s) }
func (f funcVar) String() string { return "" }
func (f funcVar) IsBoolFlag() bool { return false }
func (f funcVar) Type() string { return "funcVar" }