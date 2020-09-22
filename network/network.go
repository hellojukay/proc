package network

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"github.com/hellojukay/proc/fd"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	readLimit = 429496729
)

func newNet(file string) (Net, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	netUDP := Net{}

	lr := io.LimitReader(f, readLimit)
	s := bufio.NewScanner(lr)
	s.Scan() // skip first line with headers
	for s.Scan() {
		fields := strings.Fields(s.Text())
		line, err := parseNetLine(fields)
		if err != nil {
			return nil, err
		}
		netUDP = append(netUDP, line)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return netUDP, nil
}

// parseNetLine parses a single line, represented by a list of fields.
func parseNetLine(fields []string) (*NetLine, error) {
	line := &NetLine{}
	if len(fields) < 8 {
		return nil, fmt.Errorf(
			"cannot parse net udp socket line as it has less then 8 columns: %s",
			strings.Join(fields, " "),
		)
	}
	var err error // parse error

	// sl
	s := strings.Split(fields[0], ":")
	if len(s) != 2 {
		return nil, fmt.Errorf(
			"cannot parse sl field in socket line: %s", fields[0])
	}

	if line.Sl, err = strconv.ParseUint(s[0], 0, 64); err != nil {
		return nil, fmt.Errorf("cannot parse sl value in  socket line: %s", err)
	}
	// local_address
	l := strings.Split(fields[1], ":")
	if len(l) != 2 {
		return nil, fmt.Errorf(
			"cannot parse local_address field in socket line: %s", fields[1])
	}
	if line.LocalAddr, err = hex.DecodeString(l[0]); err != nil {
		return nil, fmt.Errorf(
			"cannot parse local_address value in  socket line: %s", err)
	}
	if line.LocalPort, err = strconv.ParseUint(l[1], 16, 64); err != nil {
		return nil, fmt.Errorf(
			"cannot parse local_address port value in socket line: %s", err)
	}

	// remote_address
	r := strings.Split(fields[2], ":")
	if len(r) != 2 {
		return nil, fmt.Errorf(
			"cannot parse rem_address field in socket line: %s", fields[1])
	}
	if line.RemAddr, err = hex.DecodeString(r[0]); err != nil {
		return nil, fmt.Errorf(
			"cannot parse rem_address value in socket line: %s", err)
	}
	if line.RemPort, err = strconv.ParseUint(r[1], 16, 64); err != nil {
		return nil, fmt.Errorf(
			"cannot parse rem_address port value in socket line: %s", err)
	}

	// st
	if line.St, err = strconv.ParseUint(fields[3], 16, 64); err != nil {
		return nil, fmt.Errorf(
			"cannot parse st value in socket line: %s", err)
	}

	// tx_queue and rx_queue
	q := strings.Split(fields[4], ":")
	if len(q) != 2 {
		return nil, fmt.Errorf(
			"cannot parse tx/rx queues in socket line as it has a missing colon: %s",
			fields[4],
		)
	}
	if line.TxQueue, err = strconv.ParseUint(q[0], 16, 64); err != nil {
		return nil, fmt.Errorf("cannot parse tx_queue value in socket line: %s", err)
	}
	if line.RxQueue, err = strconv.ParseUint(q[1], 16, 64); err != nil {
		return nil, fmt.Errorf("cannot parse rx_queue value in socket line: %s", err)
	}

	// uid
	if line.UID, err = strconv.ParseUint(fields[7], 0, 64); err != nil {
		return nil, fmt.Errorf(
			"cannot parse uid value in socket line: %s", err)
	}
	line.Inode = fields[9]
	return line, nil
}

type (
	Net []*NetLine

	NetSummary struct {
		// TxQueueLength shows the total queue length of all parsed tx_queue lengths.
		TxQueueLength uint64
		// RxQueueLength shows the total queue length of all parsed rx_queue lengths.
		RxQueueLength uint64
		// UsedSockets shows the total number of parsed lines representing the
		// number of used sockets.
		UsedSockets uint64
	}

	NetLine struct {
		Sl        uint64
		LocalAddr net.IP
		LocalPort uint64
		RemAddr   net.IP
		RemPort   uint64
		St        uint64
		TxQueue   uint64
		RxQueue   uint64
		UID       uint64
		Inode     string
	}
)

const (
	TCP   = "tcp"
	TCPv6 = "tcp6"
	UDP   = "udp"
	UDPv6 = "udp6"
)

type NetInfo struct {
	*NetLine
	// tcp , tcp6, udp,udp6
	Type string
}

func ReadTCP(pid int) ([]NetInfo, error) {
	var path = fmt.Sprintf("/proc/%d/net/tcp", pid)
	ns, err := newNet(path)
	if err != nil {
		return nil, err
	}
	var result []NetInfo
	for _, nt := range ns {
		var netInfo NetInfo
		netInfo.Type = TCP
		netInfo.NetLine = nt
		result = append(result, netInfo)
	}
	return result, nil
}
func ReadTCPv6(pid int) ([]NetInfo, error) {
	var path = fmt.Sprintf("/proc/%d/net/tcp6", pid)
	ns, err := newNet(path)
	if err != nil {
		return nil, err
	}
	var result []NetInfo
	for _, nt := range ns {
		var netInfo NetInfo
		netInfo.Type = TCPv6
		netInfo.NetLine = nt
		result = append(result, netInfo)
	}
	return result, nil
}
func ReadUDP(pid int) ([]NetInfo, error) {
	var path = fmt.Sprintf("/proc/%d/net/udp", pid)
	ns, err := newNet(path)
	if err != nil {
		return nil, err
	}
	var result []NetInfo
	for _, nt := range ns {
		var netInfo NetInfo
		netInfo.Type = UDP
		netInfo.NetLine = nt
		result = append(result, netInfo)
	}
	return result, nil
}
func ReadUDP6(pid int) ([]NetInfo, error) {
	var path = fmt.Sprintf("/proc/%d/net/udp6", pid)
	ns, err := newNet(path)
	if err != nil {
		return nil, err
	}
	var result []NetInfo
	for _, nt := range ns {
		var netInfo NetInfo
		netInfo.Type = TCPv6
		netInfo.NetLine = nt
		result = append(result, netInfo)
	}
	return result, nil
}

func ReadNetInfo(pid int) ([]NetInfo, error) {
	var netInfos []NetInfo
	tcpInfos, err := ReadTCP(pid)
	if err == nil {
		netInfos = append(netInfos, tcpInfos...)
	}
	tcp6Infos, err := ReadTCPv6(pid)
	if err == nil {
		netInfos = append(netInfos, tcp6Infos...)
	}
	udpInfos, err := ReadUDP(pid)
	if err == nil {
		netInfos = append(netInfos, udpInfos...)
	}
	udp6Infos, err := ReadUDP6(pid)
	if err == nil {
		netInfos = append(netInfos, udp6Infos...)
	}
	var inodes = socketInodes(pid)
	var result []NetInfo
	for _, info := range netInfos {
		if inArray(info.Inode, inodes) {
			result = append(result, info)
		}
	}
	return result, err
}
func PrintNetInfo(netInfos []NetInfo) {
	fmt.Printf("%-15s%-20s%-20s%-15s%-15s%10s\n", "PROTOCOL", "STATE", "LOCAL", "PORT", "REMOTE", "PORT")
	for _, info := range netInfos {
		fmt.Printf("%-15s%-20s%-20s%-15d%-15s%10d\n", info.Type, socketSatteString(info.St), reverseIp(info.LocalAddr.String()), info.LocalPort, reverseIp(info.RemAddr.String()), info.RemPort)
	}
}

func socketSatteString(st uint64) string {
	switch st {
	case 1:
		return "TCP_ESTABLISHED"
	case 2:
		return "TCP_SYN_SENT"
	case 3:
		return "TCP_SYN_RECV"
	case 4:
		return "TCP_FIN_WAIT1"
	case 5:
		return "TCP_FIN_WAIT2"
	case 6:
		return "TCP_TIME_WAIT"
	case 7:
		return "TCP_CLOSE"
	case 8:
		return "TCP_CLOSE_WAIT"
	case 9:
		return "TCP_LAST_ACL"
	case 10:
		return "TCP_LISTEN"
	case 11:
		return "TCP_CLOSING"
	default:
		return "UNKNOW"
	}
}
func reverseIp(ip string) string {
	return strings.Join(reverseString(strings.Split(ip, ".")), ".")
}
func reverseString(arr []string) []string {
	var result []string
	for _, s := range arr {
		result = append([]string{s}, result...)
	}
	return result
}

func socketInodes(pid int) []string {
	files, err := fd.ReadFd(pid)
	if err != nil {
		return nil
	}
	var inodes []string
	for _, file := range files {
		if file.IsSocket() {
			inode, _ := file.Inode()
			inodes = append(inodes, inode)
		}
	}
	return inodes
}

func inArray(s string, arr []string) bool {
	for _, str := range arr {
		if str == s {
			return true
		}
	}
	return false
}
