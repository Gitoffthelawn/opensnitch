package firewall

import (
	"fmt"

	"github.com/evilsocket/opensnitch/daemon/firewall/common"
	"github.com/evilsocket/opensnitch/daemon/firewall/config"
	"github.com/evilsocket/opensnitch/daemon/firewall/iptables"
	"github.com/evilsocket/opensnitch/daemon/firewall/nftables"
	"github.com/evilsocket/opensnitch/daemon/log"
	"github.com/evilsocket/opensnitch/daemon/ui/protocol"
)

// Firewall is the interface that all firewalls (iptables, nftables) must implement.
type Firewall interface {
	Init(uint16, string, string, bool)
	Stop()
	Name() string
	IsRunning() bool
	SetQueueNum(num uint16)

	SaveConfiguration(rawConfig string) error

	EnableInterception()
	DisableInterception(bool)
	QueueDNSResponses(bool, bool) (error, error)
	QueueConnections(bool, bool) (error, error)
	CleanRules(bool)

	AddSystemRules(bool, bool)
	DeleteSystemRules(bool, bool, bool)

	Serialize() (*protocol.SysFirewall, error)
	Deserialize(sysfw *protocol.SysFirewall) ([]byte, error)

	ErrorsChan() <-chan string
	ErrChanEmpty() bool
}

var (
	fw       Firewall
	queueNum = uint16(0)
)

// Init initializes the firewall and loads firewall rules.
// We'll try to use the firewall configured in the configuration (iptables/nftables).
// If iptables is not installed, we can add nftables rules directly to the kernel,
// without relying on any binaries.
func Init(fwType, configPath, monitorInterval string, bypassQueue bool, qNum uint16) (err error) {
	confError := false
	if fwType == "" {
		confError = true
		fwType = nftables.Name
	}
	if configPath == "" {
		confError = true
		configPath = config.DefaultConfigFile
	}

	if fwType == iptables.Name {
		fw, err = iptables.Fw()
		if err != nil {
			log.Warning("iptables not available: %s", err)
		}
	}

	if fwType == nftables.Name || err != nil {
		fw, err = nftables.Fw()
		if err != nil {
			log.Warning("nftables not available: %s", err)
		}
	}

	if err != nil {
		return fmt.Errorf("firewall error: %s, not iptables nor nftables are available or are usable. Please, report it on github", err)
	}

	if fw == nil {
		return fmt.Errorf("Firewall not initialized. Be sure that you're using latest configuration file. Report it on github if needed.")
	}
	fw.Stop()
	fw.Init(qNum, configPath, monitorInterval, bypassQueue)
	if confError {
		log.Error("Firewall error: the default configuration seem to be outdated (default-config.json). Get latest configuration from github.")
	}

	queueNum = qNum

	log.Info("Using %s firewall", fw.Name())

	return
}

// IsRunning returns if the firewall is running or not.
func IsRunning() bool {
	return fw != nil && fw.IsRunning()
}

// ErrorsChan returns the channel where the errors are sent to.
func ErrorsChan() <-chan string {
	return fw.ErrorsChan()
}

// ErrChanEmpty checks if the errors channel is empty.
func ErrChanEmpty() bool {
	return fw.ErrChanEmpty()
}

// CleanRules deletes the rules we added.
func CleanRules(logErrors bool) {
	if fw == nil {
		return
	}
	fw.CleanRules(logErrors)
}

// Reload stops current firewall and initializes a new one.
func Reload(fwtype, configPath, monitorInterval string, bypassQueue bool, queueNum uint16) (err error) {
	Stop()
	err = Init(fwtype, configPath, monitorInterval, bypassQueue, queueNum)
	return
}

// ReloadSystemRules deletes existing rules, and add them again
func ReloadSystemRules() {
	fw.DeleteSystemRules(!common.ForcedDelRules, common.RestoreChains, true)
	fw.AddSystemRules(common.ReloadRules, common.BackupChains)
}

// EnableInterception removes the rules to intercept outbound connections.
func EnableInterception() error {
	if fw == nil {
		return fmt.Errorf("firewall not initialized when trying to enable interception, report please")
	}
	fw.EnableInterception()
	return nil
}

// DisableInterception removes the rules to intercept outbound connections.
func DisableInterception() error {
	if fw == nil {
		return fmt.Errorf("firewall not initialized when trying to disable interception, report please")
	}
	fw.DisableInterception(true)
	return nil
}

// Stop deletes the firewall rules, allowing network traffic.
func Stop() {
	if fw == nil {
		return
	}
	fw.Stop()
}

// SaveConfiguration saves configuration string to disk
func SaveConfiguration(rawConfig []byte) error {
	return fw.SaveConfiguration(string(rawConfig))
}

// Serialize transforms firewall json configuration to protobuf
func Serialize() (*protocol.SysFirewall, error) {
	if fw == nil {
		return nil, fmt.Errorf("firewall not initialized, report please")
	}
	return fw.Serialize()
}

// Deserialize transforms firewall json configuration to protobuf
func Deserialize(sysfw *protocol.SysFirewall) ([]byte, error) {
	if fw == nil {
		return nil, fmt.Errorf("firewall not initialized, report please")
	}
	return fw.Deserialize(sysfw)
}
