package wifi_test

import (
	"errors"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	myWifiImpl "github.com/JingolBong/task-6/internal/wifi"
)

//go:generate mockery --name=WiFiHandle --testonly --quiet --outpkg wifi_test --output .
var errorMocker = errors.New("error mock interface")

const errorGettingInterface = "getting interfaces: "

type mockWiFi struct {
	interfacesFunc func() ([]*wifi.Interface, error)
}

func (m *mockWiFi) Interfaces() ([]*wifi.Interface, error) {
	return m.interfacesFunc()
}

func TestGetAddressesSuccess(t *testing.T) {
	address, _ := net.ParseMAC("00:11:22:33:44:55")

	mock := &mockWiFi{
		interfacesFunc: func() ([]*wifi.Interface, error) {
			return []*wifi.Interface{
				{Name: "wlan0", HardwareAddr: address},
			}, nil
		},
	}

	service := myWifiImpl.New(mock)
	addressGot, err := service.GetAddresses()

	require.NoError(t, err)
	require.Equal(t, []net.HardwareAddr{address}, addressGot)
}

func TestGetAddressesError(t *testing.T) {
	mock := &mockWiFi{
		interfacesFunc: func() ([]*wifi.Interface, error) {
			return nil, errorMocker
		},
	}

	service := myWifiImpl.New(mock)
	addrs, err := service.GetAddresses()

	require.Error(t, err)
	require.Nil(t, addrs)
	require.ErrorContains(t, err, errorGettingInterface)

}

func TestGetNamesSuccess(t *testing.T) {
	mock := &mockWiFi{
		interfacesFunc: func() ([]*wifi.Interface, error) {
			return []*wifi.Interface{
				{Name: "wlan"},
			}, nil
		},
	}

	service := myWifiImpl.New(mock)
	nameGot, err := service.GetNames()

	require.NoError(t, err)
	require.Equal(t, []string{"wlan"}, nameGot)
}

func TestGetNamesError(t *testing.T) {
	mock := &mockWiFi{
		interfacesFunc: func() ([]*wifi.Interface, error) {
			return nil, errorMocker
		},
	}

	service := myWifiImpl.New(mock)
	names, err := service.GetNames()

	require.Error(t, err)
	require.Nil(t, names)
	require.ErrorContains(t, err, errorGettingInterface)
}
