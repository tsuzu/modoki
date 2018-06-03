// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Modoki API": CLI Commands
//
// Command:
// $ goagen
// --design=github.com/cs3238-tsuzu/modoki/design
// --out=$(GOPATH)/src/github.com/cs3238-tsuzu/modoki
// --version=v1.3.1

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cs3238-tsuzu/modoki/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// CreateContainerCommand is the command line data structure for the create action of container
	CreateContainerCommand struct {
		// Command to run specified as a string or an array of strings.
		Cmd []string
		// The entry point for the container as a string or an array of strings
		Entrypoint []string
		// Environment variables
		Env []string
		// Name of image
		Image string
		// Name of container
		Name string
		// Path to volumes in a container
		Volumes []string
		// Current directory (PWD) in the command will be launched
		WorkingDir  string
		PrettyPrint bool
	}

	// RemoveContainerCommand is the command line data structure for the remove action of container
	RemoveContainerCommand struct {
		// If the container is running, kill it before removing it.
		Force string
		// id or name
		ID          string
		PrettyPrint bool
	}

	// StartContainerCommand is the command line data structure for the start action of container
	StartContainerCommand struct {
		// id or name
		ID          string
		PrettyPrint bool
	}

	// StopContainerCommand is the command line data structure for the stop action of container
	StopContainerCommand struct {
		// id or name
		ID          string
		PrettyPrint bool
	}

	// AuthtypeVironCommand is the command line data structure for the authtype action of viron
	AuthtypeVironCommand struct {
		PrettyPrint bool
	}

	// GetVironCommand is the command line data structure for the get action of viron
	GetVironCommand struct {
		PrettyPrint bool
	}

	// SigninVironCommand is the command line data structure for the signin action of viron
	SigninVironCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "authtype",
		Short: `Get viron authtype`,
	}
	tmp1 := new(AuthtypeVironCommand)
	sub = &cobra.Command{
		Use:   `viron ["/api/v1/viron_authtype"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "create",
		Short: `create a new container`,
	}
	tmp2 := new(CreateContainerCommand)
	sub = &cobra.Command{
		Use:   `container ["/api/v1/container/create"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "get",
		Short: `Get viron menu`,
	}
	tmp3 := new(GetVironCommand)
	sub = &cobra.Command{
		Use:   `viron ["/api/v1/viron"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "remove",
		Short: `remove a container`,
	}
	tmp4 := new(RemoveContainerCommand)
	sub = &cobra.Command{
		Use:   `container ["/api/v1/container/remove"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "signin",
		Short: `Creates a valid JWT`,
	}
	tmp5 := new(SigninVironCommand)
	sub = &cobra.Command{
		Use:   `viron ["/api/v1/signin"]`,
		Short: ``,
		Long: `

Payload example:

{
   "id": "identify key",
   "password": "3zvdvz59a3"
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "start",
		Short: `start a container`,
	}
	tmp6 := new(StartContainerCommand)
	sub = &cobra.Command{
		Use:   `container ["/api/v1/container/start"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "stop",
		Short: `stop a container`,
	}
	tmp7 := new(StopContainerCommand)
	sub = &cobra.Command{
		Use:   `container ["/api/v1/container/stop"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the CreateContainerCommand command.
func (cmd *CreateContainerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v1/container/create"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateContainer(ctx, path, cmd.Cmd, cmd.Entrypoint, cmd.Env, cmd.Image, cmd.Name, cmd.Volumes, stringFlagVal("workingDir", cmd.WorkingDir))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateContainerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var cmd []string
	cc.Flags().StringSliceVar(&cmd.Cmd, "cmd", cmd, `Command to run specified as a string or an array of strings.`)
	var entrypoint []string
	cc.Flags().StringSliceVar(&cmd.Entrypoint, "entrypoint", entrypoint, `The entry point for the container as a string or an array of strings`)
	var env []string
	cc.Flags().StringSliceVar(&cmd.Env, "env", env, `Environment variables`)
	var image string
	cc.Flags().StringVar(&cmd.Image, "image", image, `Name of image`)
	var name string
	cc.Flags().StringVar(&cmd.Name, "name", name, `Name of container`)
	var volumes []string
	cc.Flags().StringSliceVar(&cmd.Volumes, "volumes", volumes, `Path to volumes in a container`)
	var workingDir string
	cc.Flags().StringVar(&cmd.WorkingDir, "workingDir", workingDir, `Current directory (PWD) in the command will be launched`)
}

// Run makes the HTTP request corresponding to the RemoveContainerCommand command.
func (cmd *RemoveContainerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v1/container/remove"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp8 *bool
	if cmd.Force != "" {
		var err error
		tmp8, err = boolVal(cmd.Force)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *bool value", "flag", "--force", "err", err)
			return err
		}
	}
	if tmp8 == nil {
		goa.LogError(ctx, "required flag is missing", "flag", "--force")
		return fmt.Errorf("required flag force is missing")
	}
	resp, err := c.RemoveContainer(ctx, path, *tmp8, cmd.ID)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RemoveContainerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var force string
	cc.Flags().StringVar(&cmd.Force, "force", force, `If the container is running, kill it before removing it.`)
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `id or name`)
}

// Run makes the HTTP request corresponding to the StartContainerCommand command.
func (cmd *StartContainerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v1/container/start"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.StartContainer(ctx, path, cmd.ID)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *StartContainerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `id or name`)
}

// Run makes the HTTP request corresponding to the StopContainerCommand command.
func (cmd *StopContainerCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v1/container/stop"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.StopContainer(ctx, path, cmd.ID)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *StopContainerCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, `id or name`)
}

// Run makes the HTTP request corresponding to the AuthtypeVironCommand command.
func (cmd *AuthtypeVironCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v1/viron_authtype"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AuthtypeViron(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AuthtypeVironCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the GetVironCommand command.
func (cmd *GetVironCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v1/viron"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.GetViron(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *GetVironCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the SigninVironCommand command.
func (cmd *SigninVironCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v1/signin"
	}
	var payload client.SigninPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.SigninViron(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *SigninVironCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}
