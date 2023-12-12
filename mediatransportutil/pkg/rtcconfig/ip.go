package rtcconfig

import (
	"context"
	"fmt"
	"github.com/liuhailove/tc-base-go/protocol/logger"
	"github.com/pion/stun"
	"golang.org/x/exp/slices"
	"net"
	"time"

	"github.com/pkg/errors"
)

func (conf *RTCConfig) determineIP() (string, error) {
	if conf.UseExternalIP {
		stunServers := conf.STUNServers
		if len(stunServers) == 0 {
			stunServers = DefaultStunServers
		}
		var err error
		for i := 0; i < 3; i++ {
			var ip string
			ip, err = GetExternalIP(context.Background(), stunServers, nil)
			if err == nil {
				return ip, nil
			} else {
				time.Sleep(500 * time.Millisecond)
			}
		}
		return "", errors.Errorf("could not resolve external IP: %v", err)
	}

	// 使用本地ip代替
	addresses, err := GetLocalIPAddresses(false, nil)
	if len(addresses) > 0 {
		return addresses[0], err
	}
	return "", err
}

func GetLocalIPAddresses(includeLoopback bool, preferredInterfaces []string) ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	loopBacks := make([]string, 0)
	addresses := make([]string, 0)
	for _, iface := range ifaces {
		if len(preferredInterfaces) != 0 && !slices.Contains(preferredInterfaces, iface.Name) {
			continue
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch typedAddr := addr.(type) {
			case *net.IPNet:
				ip = typedAddr.IP.To16()
			case *net.IPAddr:
				ip = typedAddr.IP.To4()
			default:
				continue
			}
			if ip == nil {
				continue
			}
			if ip.IsLoopback() {
				loopBacks = append(loopBacks, ip.String())
			} else {
				addresses = append(addresses, ip.String())
			}
		}
	}

	if includeLoopback {
		addresses = append(addresses, loopBacks...)
	}

	if len(addresses) > 0 {
		return addresses, nil
	}
	if len(loopBacks) > 0 {
		return loopBacks, nil
	}
	return nil, fmt.Errorf("could not find local IP address")
}

// GetExternalIP 从 stun 服务器返回 localAddr 的外部 IP。如果 localAddr 为零，则自动选择本地地址，
// 否则该地址将用于验证外部 IP 是否可以从外部访问。
func GetExternalIP(ctx context.Context, stunServers []string, localAddr net.Addr) (string, error) {
	if len(stunServers) == 0 {
		return "", errors.New("STUN servers are required but not defined")
	}
	dialer := &net.Dialer{
		LocalAddr: localAddr,
	}
	conn, err := dialer.Dial("upd4", stunServers[0])
	if err != nil {
		return "", err
	}
	c, err := stun.NewClient(conn)
	if err != nil {
		return "", err
	}
	defer c.Close()

	message, err := stun.Build(stun.TransactionID, stun.BindingRequest)
	if err != nil {
		return "", err
	}

	var stunErr error
	// 足够大的缓冲区不会阻塞它
	ipChan := make(chan string, 20)
	err = c.Start(message, func(res stun.Event) {
		if res.Error != nil {
			stunErr = res.Error
			return
		}

		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			stunErr = err
			return
		}
		ip := xorAddr.IP.To4()
		if ip != nil {
			ipChan <- ip.String()
		}
	})
	if err != nil {
		return "", err
	}

	ctx1, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	select {
	case nodeIP := <-ipChan:
		if localAddr == nil {
			return nodeIP, nil
		}
		_ = c.Close()
		return nodeIP, validateExternalIP(ctx1, nodeIP, localAddr.(*net.UDPAddr))
	case <-ctx1.Done():
		msg := "could not determine public IP"
		if stunErr != nil {
			return "", errors.Wrap(stunErr, msg)
		} else {
			return "", fmt.Errorf(msg)
		}
	}
}

// validateExternalIP 通过侦听本地地址来验证外部 IP 是否可从外部访问，
// 它将发送一个魔术字符串到外部IP并检查本地地址是否接收到该字符串。
func validateExternalIP(ctx context.Context, nodeIP string, addr *net.UDPAddr) error {
	srv, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}
	defer srv.Close()

	magicString := "9#B8D2Nvg2xg5P$ZRwJ+f)*^Nne6*W3WamGY"

	validCh := make(chan struct{})
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := srv.Read(buf)
			if err != nil {
				logger.Debugw("error reading from UDP socket", "err", err)
				return
			}
			if string(buf[:n]) == magicString {
				close(validCh)
				return
			}
		}
	}()

	cli, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: net.ParseIP(nodeIP), Port: srv.LocalAddr().(*net.UDPAddr).Port})
	if err != nil {
		return err
	}
	defer cli.Close()

	if _, err = cli.Write([]byte(magicString)); err != nil {
		return err
	}

	ctx1, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	select {
	case <-validCh:
		return nil
	case <-ctx1.Done():
		break
	}
	return fmt.Errorf("could not validate external IP: %s", nodeIP)
}
