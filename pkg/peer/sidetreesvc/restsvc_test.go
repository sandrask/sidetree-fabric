/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package sidetreesvc

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	extroles "github.com/trustbloc/fabric-peer-ext/pkg/roles"

	obmocks "github.com/trustbloc/sidetree-fabric/pkg/observer/mocks"
	"github.com/trustbloc/sidetree-fabric/pkg/peer/config"
	peermocks "github.com/trustbloc/sidetree-fabric/pkg/peer/mocks"
	"github.com/trustbloc/sidetree-fabric/pkg/role"
)

//go:generate counterfeiter -o ../mocks/batchwriter.gen.go --fake-name BatchWriter . batchWriter
//go:generate counterfeiter -o ../mocks/protocolprovider.gen.go --fake-name ProtocolProvider . protocolProvider
//go:generate counterfeiter -o ../mocks/restconfig.gen.go --fake-name RestConfig . restConfig
//go:generate counterfeiter -o ../mocks/httphandler.gen.go --fake-name HTTPHandler github.com/trustbloc/sidetree-core-go/pkg/restapi/common.HTTPHandler

// Ensure that the roles are loaded
var _ = extroles.GetRoles()

func TestRESTService(t *testing.T) {
	cfg := &peermocks.RestConfig{}
	cfg.SidetreeListenURLReturns("localhost:7393", nil)

	handler := &peermocks.HTTPHandler{}

	t.Run("Success", func(t *testing.T) {
		rs, err := newRESTService(cfg, handler)
		require.NoError(t, err)
		require.NotNil(t, rs)

		require.NotPanics(t, rs.Stop)

		require.NoError(t, rs.Start())
		time.Sleep(100 * time.Millisecond)
		rs.Stop()
	})

	t.Run("No handlers", func(t *testing.T) {
		rs, err := newRESTService(cfg)
		require.NoError(t, err)
		require.NotNil(t, rs)

		require.NoError(t, rs.Start())
		rs.Stop()
	})

	t.Run("sidetreeService error", func(t *testing.T) {
		errExpected := errors.New("injected sidetreeCfgService error")
		cfg.SidetreeListenURLReturns("", errExpected)
		rs, err := newRESTService(cfg, handler)
		require.EqualError(t, err, errExpected.Error())
		require.Nil(t, rs)
	})
}

func TestRESTHandlers(t *testing.T) {
	nsCfg := config.Namespace{}
	bw := &peermocks.BatchWriter{}
	pp := &peermocks.ProtocolProvider{}
	osc := &obmocks.OperationStoreClient{}

	t.Run("Resolver and batch-writer role -> not empty", func(t *testing.T) {
		rolesValue := make(map[extroles.Role]struct{})
		rolesValue[role.Resolver] = struct{}{}
		rolesValue[role.BatchWriter] = struct{}{}
		extroles.SetRoles(rolesValue)
		defer func() {
			extroles.SetRoles(nil)
		}()

		rh, err := newRESTHandlers(channel1, nsCfg, bw, pp, osc)
		require.NoError(t, err)
		require.NotNil(t, rh)
		require.Len(t, rh.HTTPHandlers(), 2)
	})

	t.Run("No resolver or batch-writer role -> no handlers", func(t *testing.T) {
		rolesValue := make(map[extroles.Role]struct{})
		rolesValue[role.Observer] = struct{}{}
		extroles.SetRoles(rolesValue)
		defer func() {
			extroles.SetRoles(nil)
		}()

		rh, err := newRESTHandlers(channel1, nsCfg, bw, pp, osc)
		require.NoError(t, err)
		require.NotNil(t, rh)
		require.Empty(t, rh.HTTPHandlers())
	})
}
