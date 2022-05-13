package client

import (
	"context"

	"github.com/open-telemetry/opamp-go/client/types"
	"github.com/open-telemetry/opamp-go/protobufs"
)

type OpAMPClient interface {

	// Start the client and begin attempts to connect to the Server. Once connection
	// is established the client will attempt to maintain it by reconnecting if
	// the connection is lost. All failed connection attempts will be reported via
	// OnConnectFailed callback.
	//
	// AgentDescription in settings MUST be set.
	//
	// Start may immediately return an error if the settings are incorrect (e.g. the
	// serverURL is not a valid URL).
	//
	// Start does not wait until the connection to the Server is established and will
	// likely return before the connection attempts are even made.
	//
	// It is guaranteed that after the Start() call returns without error one of the
	// following callbacks will be called eventually (unless Stop() is called earlier):
	//  - OnConnectFailed
	//  - OnError
	//  - OnRemoteConfig
	//
	// Start should be called only once. It should not be called concurrently with
	// any other OpAMPClient methods.
	Start(ctx context.Context, settings types.StartSettings) error

	// Stop the client. May be called only after Start() returns successfully.
	// May be called only once.
	// After this call returns successfully it is guaranteed that no
	// callbacks will be called. Stop() will cancel context of any in-fly
	// callbacks, but will wait until such in-fly callbacks are returned before
	// Stop returns, so make sure the callbacks don't block infinitely and react
	// promptly to context cancellations.
	// Once stopped OpAMPClient cannot be started again.
	Stop(ctx context.Context) error

	// SetAgentDescription sets attributes of the Agent. The attributes will be included
	// in the next status report sent to the Server.
	// May be called after Start(), in which case the attributes will be included
	// in the next outgoing status report. This is typically used by Agents which allow
	// their AgentDescription to change dynamically while the OpAMPClient is started.
	// To define the initial Agent description that is included in the first status report
	// set StartSettings.AgentDescription field.
	SetAgentDescription(descr *protobufs.AgentDescription) error

	// UpdateEffectiveConfig fetches the current local effective config using
	// GetEffectiveConfig callback and sends it to the Server.
	UpdateEffectiveConfig(ctx context.Context) error
}
