package cloudfoundry

import (
	"github.com/enaml-ops/enaml"
	"github.com/enaml-ops/ert-plugin/enaml-gen/proxy"
	"github.com/enaml-ops/ert-plugin/plugin/config"
)

//MySQLProxy -
type MySQLProxy struct {
	Config *config.Config
}

//NewMySQLProxyPartition -
func NewMySQLProxyPartition(config *config.Config) InstanceGroupCreator {
	return &MySQLProxy{
		Config: config,
	}
}

//ToInstanceGroup -
func (s *MySQLProxy) ToInstanceGroup() (ig *enaml.InstanceGroup) {
	ig = &enaml.InstanceGroup{
		Name:      "mysql_proxy-partition",
		Instances: len(s.Config.MySQLProxyIPs),
		VMType:    s.Config.MySQLProxyVMType,
		AZs:       s.Config.AZs,
		Stemcell:  s.Config.StemcellName,
		Jobs: []enaml.InstanceJob{
			s.newMySQLProxyJob(),
		},
		Networks: []enaml.Network{
			enaml.Network{Name: s.Config.NetworkName, StaticIPs: s.Config.MySQLProxyIPs},
		},
		Update: enaml.Update{
			MaxInFlight: 1,
		},
	}
	return
}

func (s *MySQLProxy) newMySQLProxyJob() enaml.InstanceJob {
	return enaml.InstanceJob{
		Name:    "proxy",
		Release: "cf-mysql",
		Properties: &proxy.ProxyJob{
			CfMysql: &proxy.CfMysql{
				ExternalHost: s.Config.MySQLProxyExternalHost,
				Proxy: &proxy.Proxy{
					ApiUsername: s.Config.MySQLProxyAPIUsername,
					ApiPassword: s.Config.MySQLProxyAPIPassword,
					ProxyIps:    s.Config.MySQLProxyIPs,
				},
				Mysql: &proxy.Mysql{
					ClusterIps: s.Config.MySQLIPs,
				},
			},
			SyslogAggregator: &proxy.SyslogAggregator{
				Address:   s.Config.SyslogAddress,
				Port:      s.Config.SyslogPort,
				Transport: s.Config.SyslogTransport,
			},
			Nats: &proxy.Nats{
				User:     s.Config.NATSUser,
				Password: s.Config.NATSPassword,
				Machines: s.Config.NATSMachines,
				Port:     s.Config.NATSPort,
			},
		},
	}
}
