// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package commands

import (
	"context"

	"github.com/databricks/databricks-sdk-go/retries"
)

// This is the high-level interface, that contains generated methods.
//
// Evolving: this interface is under development. Method signatures may change.
type CommandExecutionService interface {

	// Cancel a command
	Cancel(ctx context.Context, request CancelCommand) error

	// CancelAndWait calls Cancel() and waits to reach Cancelled state
	//
	// This method is generated by Databricks SDK Code Generator.
	CancelAndWait(ctx context.Context, request CancelCommand, options ...retries.Option[CommandStatusResponse]) (*CommandStatusResponse, error)

	// Get information about a command
	CommandStatus(ctx context.Context, request CommandStatusRequest) (*CommandStatusResponse, error)

	// Get information about an execution context
	ContextStatus(ctx context.Context, request ContextStatusRequest) (*ContextStatusResponse, error)

	// Create an execution context
	Create(ctx context.Context, request CreateContext) (*Created, error)

	// CreateAndWait calls Create() and waits to reach Running state
	//
	// This method is generated by Databricks SDK Code Generator.
	CreateAndWait(ctx context.Context, request CreateContext, options ...retries.Option[ContextStatusResponse]) (*ContextStatusResponse, error)

	// Delete an execution context
	Destroy(ctx context.Context, request DestroyContext) error

	// Run a command
	Execute(ctx context.Context, request Command) (*Created, error)

	// ExecuteAndWait calls Execute() and waits to reach Finished or Error state
	//
	// This method is generated by Databricks SDK Code Generator.
	ExecuteAndWait(ctx context.Context, request Command, options ...retries.Option[CommandStatusResponse]) (*CommandStatusResponse, error)
}
