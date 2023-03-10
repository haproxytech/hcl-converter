package configuration

import (
	"github.com/haproxytech/client-native/v4/models"
	"github.com/jessevdk/go-flags"
)

type StorageDataplaneAPIConfiguration struct {
	Version                *int                        `yaml:"config_version,omitempty" hcl:"config_version,omitempty"`
	Name                   *string                     `yaml:"name,omitempty" hcl:"name,omitempty"`
	Mode                   *string                     `yaml:"mode,omitempty" hcl:"mode,omitempty"`
	DeprecatedBootstrapKey *string                     `yaml:"bootstrap_key,omitempty" hcl:"bootstrap_key,omitempty"`
	Status                 *string                     `yaml:"status,omitempty" hcl:"status,omitempty"`
	Dataplaneapi           *configTypeDataplaneapi     `yaml:"dataplaneapi,omitempty" hcl:"dataplaneapi,omitempty"`
	Haproxy                *configTypeHaproxy          `yaml:"haproxy,omitempty" hcl:"haproxy,omitempty"`
	Cluster                *configTypeCluster          `yaml:"cluster,omitempty" hcl:"cluster,omitempty"`
	ServiceDiscovery       *configTypeServiceDiscovery `yaml:"service_discovery,omitempty" hcl:"service_discovery,omitempty"`
	Log                    *configTypeLog              `yaml:"log,omitempty" hcl:"log,omitempty"`
	LogTargets             *Targets                    `yaml:"log_targets,omitempty" hcl:"log_targets,omitempty"`
}

type configTypeDataplaneapi struct {
	WriteTimeout     *string                `yaml:"write_timeout,omitempty" hcl:"write_timeout,omitempty"`
	GracefulTimeout  *string                `yaml:"graceful_timeout,omitempty" hcl:"graceful_timeout,omitempty"`
	ShowSystemInfo   *bool                  `yaml:"show_system_info,omitempty" hcl:"show_system_info,omitempty"`
	MaxHeaderSize    *string                `yaml:"max_header_size,omitempty" hcl:"max_header_size,omitempty"`
	SocketPath       *flags.Filename        `yaml:"socket_path,omitempty" hcl:"socket_path,omitempty"`
	Host             *string                `yaml:"host,omitempty" hcl:"host,omitempty"`
	Port             *int                   `yaml:"port,omitempty" hcl:"port,omitempty"`
	ListenLimit      *int                   `yaml:"listen_limit,omitempty" hcl:"listen_limit,omitempty"`
	DisableInotify   *bool                  `yaml:"disable_inotify,omitempty" hcl:"disable_inotify,omitempty"`
	ReadTimeout      *string                `yaml:"read_timeout,omitempty" hcl:"read_timeout,omitempty"`
	Advertised       *configTypeAdvertised  `yaml:"advertised,omitempty" hcl:"advertised,omitempty"`
	CleanupTimeout   *string                `yaml:"cleanup_timeout,omitempty" hcl:"cleanup_timeout,omitempty"`
	KeepAlive        *string                `yaml:"keep_alive,omitempty" hcl:"keep_alive,omitempty"`
	PIDFile          *string                `yaml:"pid_file,omitempty" hcl:"pid_file,omitempty"`
	UID              *int                   `yaml:"uid,omitempty" hcl:"uid,omitempty"`
	GID              *int                   `yaml:"gid,omitempty" hcl:"gid,omitempty"`
	TLS              *configTypeTLS         `yaml:"tls,omitempty" hcl:"tls,omitempty"`
	EnabledListeners *[]string              `yaml:"scheme,omitempty" hcl:"scheme,omitempty"`
	Userlist         *configTypeUserlist    `yaml:"userlist,omitempty" hcl:"userlist,omitempty"`
	Transaction      *configTypeTransaction `yaml:"transaction,omitempty" hcl:"transaction,omitempty"`
	Resources        *configTypeResources   `yaml:"resources,omitempty" hcl:"resources,omitempty"`
	User             []configTypeUser       `yaml:"user,omitempty" hcl:"user,omitempty"`
}

type configTypeHaproxy struct {
	ConfigFile       *string           `yaml:"config_file,omitempty" hcl:"config_file,omitempty"`
	HAProxy          *string           `yaml:"haproxy_bin,omitempty" hcl:"haproxy_bin,omitempty"`
	MasterRuntime    *string           `yaml:"master_runtime,omitempty" hcl:"master_runtime,omitempty"`
	NodeIDFile       *string           `yaml:"fid,omitempty" hcl:"fid,omitempty"`
	MasterWorkerMode *bool             `yaml:"master_worker_mode,omitempty" hcl:"master_worker_mode,omitempty"`
	Reload           *configTypeReload `yaml:"reload,omitempty" hcl:"reload,omitempty"`
}

type configTypeCluster struct {
	APINodesPath       *string                    `yaml:"api_nodes_path,omitempty" hcl:"api_nodes_path,omitempty"`
	Token              *string                    `yaml:"token,omitempty" hcl:"token,omitempty"`
	ClusterTLSCertDir  *string                    `yaml:"cluster_tls_dir,omitempty" hcl:"cluster_tls_dir,omitempty"`
	ActiveBootstrapKey *string                    `yaml:"active_bootstrap_key,omitempty" hcl:"active_bootstrap_key,omitempty"`
	APIRegisterPath    *string                    `yaml:"api_register_path,omitempty" hcl:"api_register_path,omitempty"`
	URL                *string                    `yaml:"url,omitempty" hcl:"url,omitempty"`
	Port               *int                       `yaml:"port,omitempty" hcl:"port,omitempty"`
	StorageDir         *string                    `yaml:"storage_dir,omitempty" hcl:"storage_dir,omitempty"`
	BootstrapKey       *string                    `yaml:"bootstrap_key,omitempty" hcl:"bootstrap_key,omitempty"`
	ID                 *string                    `yaml:"id,omitempty" hcl:"id,omitempty"`
	APIBasePath        *string                    `yaml:"api_base_path,omitempty" hcl:"api_base_path,omitempty"`
	CertificateDir     *string                    `yaml:"cert_path,omitempty" hcl:"cert_path,omitempty"`
	CertificateFetched *bool                      `yaml:"cert_fetched,omitempty" hcl:"cert_fetched,omitempty"`
	Name               *string                    `yaml:"name,omitempty" hcl:"name,omitempty"`
	Description        *string                    `yaml:"description,omitempty" hcl:"description,omitempty"`
	ClusterID          *string                    `yaml:"cluster_id,omitempty" hcl:"cluster_id,omitempty" group:"cluster" save:"true"`
	ClusterLogTargets  []*models.ClusterLogTarget `yaml:"cluster_log_targets,omitempty" hcl:"cluster_log_targets,omitempty" group:"cluster" save:"true"`
}

type configTypeServiceDiscovery struct {
	Consuls    *[]*models.Consul    `yaml:"consuls,omitempty" hcl:"consuls,omitempty"`
	AWSRegions *[]*models.AwsRegion `yaml:"aws_regions,omitempty" hcl:"aws_regions,omitempty"`
}

type configTypeLog struct {
	LogTo     *string           `yaml:"log_to,omitempty" hcl:"log_to,omitempty"`
	LogFile   *string           `yaml:"log_file,omitempty" hcl:"log_file,omitempty"`
	LogLevel  *string           `yaml:"log_level,omitempty" hcl:"log_level,omitempty"`
	LogFormat *string           `yaml:"log_format,omitempty" hcl:"log_format,omitempty"`
	ACLFormat *string           `yaml:"apache_common_log_format,omitempty" hcl:"apache_common_log_format,omitempty"`
	Syslog    *configTypeSyslog `yaml:"syslog,omitempty" hcl:"syslog,omitempty"`
}

type configTypeAdvertised struct {
	APIAddress *string `yaml:"api_address,omitempty" hcl:"api_address,omitempty"`
	APIPort    *int64  `yaml:"api_port,omitempty" hcl:"api_port,omitempty"`
}

type configTypeTLS struct {
	TLSHost           *string         `yaml:"tls_host,omitempty" hcl:"tls_host,omitempty"`
	TLSPort           *int            `yaml:"tls_port,omitempty" hcl:"tls_port,omitempty"`
	TLSCertificate    *flags.Filename `yaml:"tls_certificate,omitempty" hcl:"tls_certificate,omitempty"`
	TLSCertificateKey *flags.Filename `yaml:"tls_key,omitempty" hcl:"tls_key,omitempty"`
	TLSCACertificate  *flags.Filename `yaml:"tls_ca,omitempty" hcl:"tls_ca,omitempty"`
	TLSListenLimit    *int            `yaml:"tls_listen_limit,omitempty" hcl:"tls_listen_limit,omitempty"`
	TLSKeepAlive      *string         `yaml:"tls_keep_alive,omitempty" hcl:"tls_keep_alive,omitempty"`
	TLSReadTimeout    *string         `yaml:"tls_read_timeout,omitempty" hcl:"tls_read_timeout,omitempty"`
	TLSWriteTimeout   *string         `yaml:"tls_write_timeout,omitempty" hcl:"tls_write_timeout,omitempty"`
}

type configTypeUserlist struct {
	Userlist     *string `yaml:"userlist,omitempty" hcl:"userlist,omitempty"`
	UserListFile *string `yaml:"userlist_file,omitempty" hcl:"userlist_file,omitempty"`
}

type configTypeTransaction struct {
	TransactionDir      *string `yaml:"transaction_dir,omitempty" hcl:"transaction_dir,omitempty"`
	BackupsNumber       *int    `yaml:"backups_number,omitempty" hcl:"backups_number,omitempty"`
	BackupsDir          *string `yaml:"backups_dir,omitempty" hcl:"backups_dir,omitempty"`
	MaxOpenTransactions *int64  `yaml:"max_open_transactions,omitempty" hcl:"max_open_transactions,omitempty"`
}

type configTypeResources struct {
	MapsDir              *string `yaml:"maps_dir,omitempty" hcl:"maps_dir,omitempty"`
	SSLCertsDir          *string `yaml:"ssl_certs_dir,omitempty" hcl:"ssl_certs_dir,omitempty"`
	GeneralStorageDir    *string `yaml:"general_storage_dir,omitempty" hcl:"general_storage_dir,omitempty"`
	UpdateMapFiles       *bool   `yaml:"update_map_files,omitempty" hcl:"update_map_files,omitempty"`
	UpdateMapFilesPeriod *int64  `yaml:"update_map_files_period,omitempty" hcl:"update_map_files_period,omitempty"`
	SpoeDir              *string `yaml:"spoe_dir,omitempty" hcl:"spoe_dir,omitempty"`
	SpoeTransactionDir   *string `yaml:"spoe_transaction_dir,omitempty" hcl:"spoe_transaction_dir,omitempty"`
}

type configTypeUser struct {
	Insecure *bool   `yaml:"insecure,omitempty" hcl:"insecure,omitempty"`
	Password *string `yaml:"password,omitempty" hcl:"password,omitempty"`
	Name     string  `yaml:"name" hcl:"name,key"`
}

type configTypeReload struct {
	ReloadDelay     *int    `yaml:"reload_delay,omitempty" hcl:"reload_delay,omitempty"`
	ReloadCmd       *string `yaml:"reload_cmd,omitempty" hcl:"reload_cmd,omitempty"`
	RestartCmd      *string `yaml:"restart_cmd,omitempty" hcl:"restart_cmd,omitempty"`
	StatusCmd       *string `yaml:"status_cmd,omitempty" hcl:"status_cmd,omitempty"`
	ServiceName     *string `yaml:"service_name,omitempty" hcl:"service_name,omitempty"`
	ReloadRetention *int    `yaml:"reload_retention,omitempty" hcl:"reload_retention,omitempty"`
	ReloadStrategy  *string `yaml:"reload_strategy,omitempty" hcl:"reload_strategy,omitempty"`
	ValidateCmd     *string `yaml:"validate_cmd,omitempty" hcl:"validate_cmd,omitempty"`
}

type configTypeSyslog struct {
	SyslogAddr     *string `yaml:"syslog_address,omitempty" hcl:"syslog_address,omitempty"`
	SyslogProto    *string `yaml:"syslog_protocol,omitempty" hcl:"syslog_protocol,omitempty"`
	SyslogTag      *string `yaml:"syslog_tag,omitempty" hcl:"syslog_tag,omitempty"`
	SyslogLevel    *string `yaml:"syslog_level,omitempty" hcl:"syslog_level,omitempty"`
	SyslogFacility *string `yaml:"syslog_facility,omitempty" hcl:"syslog_facility,omitempty"`
}

type Target struct {
	LogTo          string   `yaml:"log_to,omitempty" description:"Log target, can be stdout, file, or syslog" hcl:"log_to"`
	LogFile        string   `yaml:"log_file,omitempty" description:"Location of the log file" hcl:"log_file"`
	LogLevel       string   `yaml:"log_level,omitempty" description:"Logging level, allowed values: trace|debug|info|warning|error" hcl:"log_level"`
	LogFormat      string   `yaml:"log_format,omitempty" description:"Logging format, allowed values: text|JSON" hcl:"log_format"`
	ACLFormat      string   `yaml:"acl_format,omitempty" description:"Apache Common Log Format to format the access log entries, default:\"%h %l %u %t \\\"%r\\\" %>s %b \\\"%{Referer}i\\\" \\\"%{User-agent}i\\\" %{us}T"`
	SyslogAddr     string   `yaml:"syslog_address,omitempty" description:"Syslog address (with port declaration in case of TCP type) where logs should be forwarded: accepting socket path in case of unix or unixgram" hcl:"syslog_address"`
	SyslogProto    string   `yaml:"syslog_protocol,omitempty" description:"Syslog server protocol, allowed values: tcp|tcp4|tcp6|unix|unixgram" hcl:"syslog_protocol"`
	SyslogTag      string   `yaml:"syslog_tag,omitempty" description:"String to tag the syslog messages" hcl:"syslog_tag"`
	SyslogLevel    string   `yaml:"syslog_level,omitempty" description:"Define the required syslog messages level, allowed values: debug|info|notice|warning|error|critical|alert|emergency " hcl:"syslog_level"`
	SyslogFacility string   `yaml:"syslog_facility,omitempty" description:"Define the Syslog facility number, allowed values: kern|user|mail|daemon|auth|syslog|lpr|news|uucp|cron|authpriv|ftp|local0|local1|local2|local3|local4|local5|local6|local7" hcl:"syslog_facillity"`
	SyslogMsgID    string   `yaml:"-"`
	LogTypes       []string `yaml:"log_types,omitempty" description:"Define which log types to log to this target, allowed values: app|access" save:"true" hcl:"log_types"`
}

type Targets []Target
