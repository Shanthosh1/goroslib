package goroslib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestServiceProviderRegister(t *testing.T) {
	m, err := newContainerMaster()
	require.NoError(t, err)
	defer m.close()

	n, err := NewNode(NodeConf{
		Name:       "/goroslib",
		MasterHost: m.Ip(),
	})
	require.NoError(t, err)
	defer n.Close()

	sp, err := NewServiceProvider(ServiceProviderConf{
		Node:    n,
		Service: "/test_srv",
		Callback: func(req *TestServiceReq) *TestServiceRes {
			c := float64(0)
			if req.A == 123 && req.B == "456" {
				c = 123
			}

			return &TestServiceRes{
				C: c,
			}
		},
	})
	require.NoError(t, err)

	// test registration

	time.Sleep(1 * time.Second)

	services, err := n.GetServices()
	require.NoError(t, err)

	service, ok := services["/test_srv"]
	require.Equal(t, true, ok)

	_, ok = service.Providers["/goroslib"]
	require.Equal(t, true, ok)

	// test un-registration

	sp.Close()
	time.Sleep(1 * time.Second)

	services, err = n.GetServices()
	require.NoError(t, err)

	service, ok = services["/test_srv"]
	require.Equal(t, false, ok)
}

func TestServiceProviderResToGo(t *testing.T) {
	m, err := newContainerMaster()
	require.NoError(t, err)
	defer m.close()

	nsp, err := NewNode(NodeConf{
		Name:       "/goroslib_sp",
		MasterHost: m.Ip(),
	})
	require.NoError(t, err)
	defer nsp.Close()

	sp, err := NewServiceProvider(ServiceProviderConf{
		Node:    nsp,
		Service: "/test_srv",
		Callback: func(req *TestServiceReq) *TestServiceRes {
			c := float64(0)
			if req.A == 123 && req.B == "456" {
				c = 123
			}

			return &TestServiceRes{
				C: c,
			}
		},
	})
	require.NoError(t, err)
	defer sp.Close()

	nsc, err := NewNode(NodeConf{
		Name:       "/goroslib_sc",
		MasterHost: m.Ip(),
	})
	require.NoError(t, err)
	defer nsc.Close()

	sc, err := NewServiceClient(ServiceClientConf{
		Node:    nsc,
		Service: "/test_srv",
		Req:     &TestServiceReq{},
		Res:     &TestServiceRes{},
	})
	require.NoError(t, err)
	defer sc.Close()

	for i := 0; i < 2; i++ {
		req := TestServiceReq{
			A: 123,
			B: "456",
		}
		res := TestServiceRes{}
		err = sc.Call(&req, &res)
		require.NoError(t, err)

		expected := TestServiceRes{
			C: 123,
		}
		require.Equal(t, expected, res)
	}
}
