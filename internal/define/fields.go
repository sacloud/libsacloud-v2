package define

import (
	"github.com/sacloud/libsacloud-v2/internal/schema"
	"github.com/sacloud/libsacloud-v2/internal/schema/meta"
)

type fieldsDef struct{}

var fields = &fieldsDef{}

func (f *fieldsDef) New(name string, t meta.Type) *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: name,
		Type: t,
	}
}

func (f *fieldsDef) ID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ID",
		Type: meta.TypeID,
		ExtendAccessors: []*schema.ExtendAccessor{
			{
				Name: "StringID",
				Type: meta.TypeString,
			},
			{
				Name: "Int64ID",
				Type: meta.TypeInt64,
			},
		},
	}
}

func (f *fieldsDef) Name() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Name",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "required",
		},
	}
}

func (f *fieldsDef) InterfaceDriver() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "InterfaceDriver",
		Type: meta.TypeInterfaceDriver,
		Tags: &schema.FieldTags{
			MapConv: ",default=virtio",
		},
	}
}

func (f *fieldsDef) BridgeInfo() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "BridgeInfo",
		Type: models.bridgeInfoModel(),
		Tags: &schema.FieldTags{
			MapConv: "[]Switches,recursive",
		},
	}
}

func (f *fieldsDef) SwitchInZone() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SwitchInZone",
		Type: models.switchInZoneModel(),
	}
}

func (f *fieldsDef) DiskPlanID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DiskPlanID",
		Tags: &schema.FieldTags{
			MapConv: "Plan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) DiskPlanName() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DiskPlanName",
		Tags: &schema.FieldTags{
			MapConv: "Plan.Name",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) DiskPlanStorageClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DiskPlanStorageClass",
		Tags: &schema.FieldTags{
			MapConv: "Plan.StorageClass",
		},
		Type: meta.TypeString,
	}
}

// TODO CPUとServerPlanCPUのようにmapconvのタグだけ違う値をどう扱うか

func (f *fieldsDef) CPU() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CPU",
		Type: meta.TypeInt,
	}
}

/*
func (f *fieldsDef) ServerPlanCPU() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CPU",
		Tags: &schema.FieldTags{
			MapConv: "ServerPlan.CPU",
		},
		Type: meta.TypeInt,
	}
}
*/

func (f *fieldsDef) MemoryMB() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "MemoryMB",
		Type: meta.TypeInt,
		ExtendAccessors: []*schema.ExtendAccessor{
			{
				Name: "MemoryGB",
				Type: meta.TypeInt,
			},
		},
	}
}

func (f *fieldsDef) Generation() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ServerPlanGeneration",
		Type: meta.TypePlanGeneration,
	}
}

func (f *fieldsDef) Commitment() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ServerPlanCommitment",
		Tags: &schema.FieldTags{
			MapConv: "Commitment,default=standard",
		},
		Type: meta.TypeCommitment,
	}
}

func (f *fieldsDef) ServerPlanID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ServerPlanID",
		Tags: &schema.FieldTags{
			MapConv: "ServerPlan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ServerPlanName() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ServerPlanName",
		Tags: &schema.FieldTags{
			MapConv: "ServerPlan.Name",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) ServerPlanCPU() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CPU",
		Tags: &schema.FieldTags{
			MapConv: "ServerPlan.CPU",
		},
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) ServerPlanMemoryMB() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "MemoryMB",
		Tags: &schema.FieldTags{
			MapConv: "ServerPlan.MemoryMB",
		},
		Type: meta.TypeInt,
		ExtendAccessors: []*schema.ExtendAccessor{
			{
				Name: "MemoryGB",
				Type: meta.TypeInt,
			},
		},
	}
}

func (f *fieldsDef) ServerPlanGeneration() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ServerPlanGeneration",
		Type: meta.TypePlanGeneration,
		Tags: &schema.FieldTags{
			MapConv: "ServerPlan.Generation",
		},
	}
}

func (f *fieldsDef) ServerPlanCommitment() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ServerPlanCommitment",
		Tags: &schema.FieldTags{
			MapConv: "ServerPlan.Commitment,default=standard",
		},
		Type: meta.TypeCommitment,
	}
}

func (f *fieldsDef) PlanID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "PlanID",
		Tags: &schema.FieldTags{
			MapConv: "Plan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) stringEnabled() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Enabled",
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
		},
		Type: meta.TypeStringFlag,
	}
}

func (f *fieldsDef) SourceDiskID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SourceDiskID",
		Tags: &schema.FieldTags{
			MapConv: "SourceDisk.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) SourceDiskAvailability() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SourceDiskAvailability",
		Tags: &schema.FieldTags{
			MapConv: "SourceDisk.Availability,omitempty",
		},
		Type: meta.TypeAvailability,
	}
}

func (f *fieldsDef) SourceArchiveID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SourceArchiveID",
		Tags: &schema.FieldTags{
			MapConv: "SourceArchive.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) SourceArchiveAvailability() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SourceArchiveAvailability",
		Tags: &schema.FieldTags{
			MapConv: "SourceArchive.Availability,omitempty",
		},
		Type: meta.TypeAvailability,
	}
}

func (f *fieldsDef) OriginalArchiveID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "OriginalArchiveID",
		Tags: &schema.FieldTags{
			MapConv: "OriginalArchive.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) BridgeID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "BridgeID",
		Type: meta.TypeID,
		Tags: &schema.FieldTags{
			MapConv: "Bridge.ID,omitempty",
		},
	}
}

func (f *fieldsDef) SwitchID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SwitchID",
		Type: meta.TypeID,
		Tags: &schema.FieldTags{
			MapConv: "Switch.ID,omitempty",
		},
	}
}

func (f *fieldsDef) PacketFilterID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "PacketFilterID",
		Type: meta.TypeID,
		Tags: &schema.FieldTags{
			MapConv: "PacketFilter.ID,omitempty",
		},
	}
}

func (f *fieldsDef) ServerID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ServerID",
		Tags: &schema.FieldTags{
			MapConv: "Server.ID,omitempty",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) IconID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IconID",
		Tags: &schema.FieldTags{
			MapConv: "Icon.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ZoneID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ZoneID",
		Type: meta.TypeID,
		Tags: &schema.FieldTags{
			MapConv: "Zone.ID",
		},
	}
}

func (f *fieldsDef) CDROMID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CDROMID",
		Type: meta.TypeID,
		Tags: &schema.FieldTags{
			MapConv: "CDROM.ID",
		},
	}
}

func (f *fieldsDef) PrivateHostID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "PrivateHostID",
		Type: meta.TypeID,
		Tags: &schema.FieldTags{
			MapConv: "PrivateHost.ID",
		},
	}
}

func (f *fieldsDef) PrivateHostName() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "PrivateHostName",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "PrivateHost.Name",
		},
	}
}

func (f *fieldsDef) AppliancePlanID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "PlanID",
		Tags: &schema.FieldTags{
			MapConv: "Remark.Plan.ID/Plan.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ApplianceSwitchID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SwitchID",
		Tags: &schema.FieldTags{
			MapConv: "Remark.Switch.ID",
		},
		Type: meta.TypeID,
	}
}

func (f *fieldsDef) ApplianceIPAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddress",
		Tags: &schema.FieldTags{
			MapConv: "Remark.[]Servers.IPAddress",
		},
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) ApplianceIPAddresses() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddresses",
		Type: meta.TypeStringSlice,
		Tags: &schema.FieldTags{
			MapConv:  "Remark.[]Servers.IPAddress",
			Validate: "min=1,max=2,dive,ipv4",
		},
	}
}

func (f *fieldsDef) LoadBalancerVIP() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "VirtualIPAddresses",
		Type: &schema.Model{
			Name:    "LoadBalancerVirtualIPAddress",
			IsArray: true,
			Fields: []*schema.FieldDesc{
				{
					Name: "VirtualIPAddress",
					Type: meta.TypeString,
					Tags: &schema.FieldTags{
						Validate: "ipv4",
					},
				},
				{
					Name: "Port",
					Type: meta.TypeStringNumber,
				},
				{
					Name: "DelayLoop",
					Type: meta.TypeStringNumber,
					Tags: &schema.FieldTags{
						MapConv:  ",default=10",
						Validate: "min=0,max=60", // TODO 最大値確認
					},
				},
				{
					Name: "SorryServer",
					Type: meta.TypeString,
					Tags: &schema.FieldTags{
						Validate: "ipv4",
					},
				},
				f.Description(),
				{
					Name: "Servers",
					Type: &schema.Model{
						Name:    "LoadBalancerServer",
						IsArray: true,
						Fields: []*schema.FieldDesc{
							{
								Name: "IPAddress",
								Type: meta.TypeString,
								Tags: &schema.FieldTags{
									Validate: "ipv4",
								},
							},
							{
								Name: "Port",
								Type: meta.TypeStringNumber,
								Tags: &schema.FieldTags{
									Validate: "min=1,max=65535",
								},
							},
							{
								Name: "Enabled",
								Type: meta.TypeStringFlag,
							},
							{
								Name: "HealthCheckProtocol",
								Type: meta.TypeProtocol,
								Tags: &schema.FieldTags{
									MapConv:  "HealthCheck.Protocol",
									Validate: "oneof=http https ping tcp",
								},
							},
							{
								Name: "HealthCheckPath",
								Type: meta.TypeString,
								Tags: &schema.FieldTags{
									MapConv: "HealthCheck.Path",
								},
							},
							{
								Name: "HealthCheckResponseCode",
								Type: meta.TypeStringNumber,
								Tags: &schema.FieldTags{
									MapConv: "HealthCheck.Status",
								},
							},
						},
					},
					Tags: &schema.FieldTags{
						MapConv:  ",recursive",
						Validate: "min=0,max=40",
					},
				},
			},
		},
		Tags: &schema.FieldTags{
			MapConv:  "Settings.[]LoadBalancer,recursive",
			Validate: "min=0,max=10",
		},
	}
}

func (f *fieldsDef) Tags() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Tags",
		Type: meta.TypeStringSlice,
	}
}

func (f *fieldsDef) Class() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) NFSClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: ",default=nfs",
		},
	}
}

func (f *fieldsDef) LoadBalancerClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: ",default=loadbalancer",
		},
	}
}

func (f *fieldsDef) VPCRouterClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: ",default=vpcrouter",
		},
	}
}

func (f *fieldsDef) SIMProviderClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Provider.Class,default=sim",
		},
	}
}

func (f *fieldsDef) SIMICCID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ICCID",
		Type: meta.TypeStringNumber,
		Tags: &schema.FieldTags{
			MapConv:  "Status.ICCID",
			Validate: "numeric", // TODO 数値のみ15桁固定
		},
	}
}

func (f *fieldsDef) SIMPassCode() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "PassCode",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Remark.PassCode",
		},
	}
}

func (f *fieldsDef) GSLBProviderClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Provider.Class,default=gslb",
		},
	}
}

func (f *fieldsDef) GSLBFQDN() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "FQDN",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Status.FQDN",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckProtocol() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckProtocol",
		Type: meta.TypeProtocol,
		Tags: &schema.FieldTags{
			MapConv:  "Settings.GSLB.HealthCheck.Protocol",
			Validate: "oneof=http https ping tcp",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckHostHeader() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckHostHeader",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Host",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckPath() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckPath",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Path",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckResponseCode() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckResponseCode",
		Type: meta.TypeStringNumber,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Status",
		},
	}
}

func (f *fieldsDef) GSLBHealthCheckPort() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HealthCheckPort",
		Type: meta.TypeStringNumber,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.HealthCheck.Port",
		},
	}
}

func (f *fieldsDef) GSLBDelayLoop() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DelayLoop",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			Validate: "min=10,max=60",
			MapConv:  "Settings.GSLB.DelayLoop,default=10",
		},
	}
}

func (f *fieldsDef) GSLBWeighted() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Weighted",
		Type: meta.TypeStringFlag,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.Weighted",
		},
	}
}

func (f *fieldsDef) GSLBDestinationServers() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DestinationServers",
		Type: &schema.Model{
			Name:    "GSLBServer",
			IsArray: true,
			Fields: []*schema.FieldDesc{
				{
					Name: "IPAddress",
					Type: meta.TypeString,
					Tags: &schema.FieldTags{
						Validate: "ipv4",
					},
				},
				{
					Name: "Enabled",
					Type: meta.TypeStringFlag,
				},
				{
					Name: "Weight",
					Type: meta.TypeStringNumber,
					Tags: &schema.FieldTags{
						MapConv: ",default=1",
					},
				},
			},
		},
		Tags: &schema.FieldTags{
			MapConv:  "Settings.GSLB.[]Servers,recursive",
			Validate: "min=0,max=6",
		},
	}
}

func (f *fieldsDef) GSLBSorryServer() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SorryServer",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Settings.GSLB.SorryServer",
		},
	}
}

func (f *fieldsDef) SettingsHash() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SettingsHash",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) InstanceHostName() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "InstanceHostName",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Instance.Host.Name",
		},
	}
}

func (f *fieldsDef) InstanceHostInfoURL() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "InstanceHostInfoURL",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Instance.Host.InfoURL",
		},
	}
}

func (f *fieldsDef) InstanceStatus() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "InstanceStatus",
		Type: meta.TypeInstanceStatus,
		Tags: &schema.FieldTags{
			MapConv: "Instance.Status",
		},
	}
}

func (f *fieldsDef) InstanceBeforeStatus() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "InstanceBeforeStatus",
		Type: meta.TypeInstanceStatus,
		Tags: &schema.FieldTags{
			MapConv: "Instance.BeforeStatus",
		},
	}
}

func (f *fieldsDef) InstanceStatusChangedAt() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "InstanceStatusChangedAt",
		Type: meta.TypeTime,
		Tags: &schema.FieldTags{
			MapConv: "Instance.StatusChangedAt",
		},
	}
}

func (f *fieldsDef) InstanceWarnings() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "InstanceWarnings",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			MapConv: "Instance.Warnings",
		},
	}
}

func (f *fieldsDef) InstanceWarningsValue() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "InstanceWarningsValue",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			MapConv: "Instance.WarningsValue",
		},
	}
}

func (f *fieldsDef) Interfaces() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Interfaces",
		Type: models.interfaceModel(),
		Tags: &schema.FieldTags{
			JSON:    ",omitempty",
			MapConv: "[]Interfaces,recursive,omitempty",
		},
	}
}

func (f *fieldsDef) VPCRouterInterfaces() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Interfaces",
		Type: models.vpcRouterInterfaceModel(),
		Tags: &schema.FieldTags{
			JSON:    ",omitempty",
			MapConv: "[]Interfaces,recursive,omitempty",
		},
	}
}

func (f *fieldsDef) NoteClass() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Class",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) NoteContent() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Content",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) Description() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Description",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "min=0,max=512",
		},
	}
}

func (f *fieldsDef) Availability() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Availability",
		Type: meta.TypeAvailability,
	}
}

func (f *fieldsDef) Scope() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Scope",
		Type: meta.TypeScope,
	}
}

func (f *fieldsDef) BandWidthMbps() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "BandWidthMbps",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) DiskConnection() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Connection",
		Type: meta.TypeDiskConnection,
		Tags: &schema.FieldTags{
			JSON:    ",omitempty",
			MapConv: ",omitempty",
		},
	}
}

func (f *fieldsDef) DiskConnectionOrder() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ConnectionOrder",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) DiskReinstallCount() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ReinstallCount",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) SizeMB() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SizeMB",
		Type: meta.TypeInt,
		ExtendAccessors: []*schema.ExtendAccessor{
			{Name: "SizeGB"},
		},
	}
}

func (f *fieldsDef) MigratedMB() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "MigratedMB",
		Type: meta.TypeInt,
		ExtendAccessors: []*schema.ExtendAccessor{
			{Name: "MigratedGB"},
		},
	}
}

func (f *fieldsDef) DefaultRoute() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DefaultRoute",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "ipv4",
		},
	}
}

func (f *fieldsDef) NextHop() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "NextHop",
		Type: meta.TypeString,
		Description: `
			スイッチ+ルータでの追加IPアドレスブロックを示すSubnetの中でのみ設定される項目。
			この場合DefaultRouteの値は設定されないためNextHopを代用する。
			StaticRouteと同じ値が設定される。`,
		Tags: &schema.FieldTags{
			Validate: "ipv4",
		},
	}
}

func (f *fieldsDef) StaticRoute() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "StaticRoute",
		Type: meta.TypeString,
		Description: `
			スイッチ+ルータでの追加IPアドレスブロックを示すSubnetの中でのみ設定される項目。
			この場合DefaultRouteの値は設定されないためNextHopを代用する。
			NextHopと同じ値が設定される。`,
		Tags: &schema.FieldTags{
			Validate: "ipv4",
		},
	}
}

func (f *fieldsDef) NetworkMaskLen() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "NetworkMaskLen",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			Validate: "min=24,max=28",
		},
	}
}

func (f *fieldsDef) NetworkAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "NetworkAddress",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "ipv4",
		},
	}
}

func (f *fieldsDef) UserSubnetNetworkMaskLen() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "NetworkMaskLen",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			Validate: "min=1,max=32", // TODO
			MapConv:  "UserSubnet.NetworkMaskLen",
		},
	}
}

func (f *fieldsDef) UserSubnetDefaultRoute() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DefaultRoute",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "ipv4", // TODO
			MapConv:  "UserSubnet.DefaultRoute",
		},
	}
}

func (f *fieldsDef) RemarkNetworkMaskLen() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "NetworkMaskLen",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			Validate: "min=1,max=32", // TODO
			MapConv:  "Remark.Network.NetworkMaskLen",
		},
	}
}

func (f *fieldsDef) RemarkZoneID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ZoneID",
		Type: meta.TypeID,
		Tags: &schema.FieldTags{
			MapConv: "Remark.Zone.ID",
		},
	}
}

func (f *fieldsDef) RemarkDefaultRoute() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DefaultRoute",
		Type: meta.TypeString,
		Tags: &schema.FieldTags{
			Validate: "ipv4", // TODO
			MapConv:  "Remark.Network.DefaultRoute",
		},
	}
}

func (f *fieldsDef) RemarkServerIPAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddresses",
		Type: meta.TypeStringSlice,
		Tags: &schema.FieldTags{
			MapConv: "Remark.[]Servers.IPAddress",
		},
	}
}

func (f *fieldsDef) RemarkVRID() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "VRID",
		Type: meta.TypeInt,
		Tags: &schema.FieldTags{
			MapConv: "Remark.VRRP.VRID",
		},
	}
}

func (f *fieldsDef) RequiredHostVersion() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "RequiredHostVersion",
		Type: meta.TypeStringNumber,
	}
}

func (f *fieldsDef) PacketFilterExpressions() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Expression",
		Type: models.packetFilterExpressions(),
		Tags: &schema.FieldTags{
			MapConv: "[]Expression,recursive",
		},
	}
}

func (f *fieldsDef) ExpressionHash() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ExpressionHash",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) DisplayOrder() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DisplayOrder",
		Type: meta.TypeInt,
	}
}

func (f *fieldsDef) IsDummy() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IsDummy",
		Type: meta.TypeFlag,
	}
}

func (f *fieldsDef) HostName() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "HostName",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) IPAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "IPAddress",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) UserIPAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "UserIPAddress",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) MACAddress() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "MACAddress",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) User() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "User",
		Type: meta.TypeString,
	}
}
func (f *fieldsDef) Password() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Password",
		Type: meta.TypeString,
	}
}

func (f *fieldsDef) SourceInfo() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "SourceInfo",
		Type: models.sourceArchiveInfo(),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty,recursive",
		},
	}
}

func (f *fieldsDef) VNCProxy() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "VNCProxy",
		Type: models.vncProxyModel(),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) FTPServer() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "FTPServer",
		Type: models.ftpServerInfo(),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Region() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Region",
		Type: models.region(),
		Tags: &schema.FieldTags{
			JSON: ",omitempty",
		},
	}
}

func (f *fieldsDef) Zone() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Zone",
		Type: models.zoneInfoModel(),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty,recursive",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) Storage() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Storage",
		Type: models.storageModel(),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty,recursive",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) BundleInfo() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "BundleInfo",
		Type: models.bundleInfoModel(),
		Tags: &schema.FieldTags{
			MapConv: ",omitempty,recursive",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) CreatedAt() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CreatedAt",
		Type: meta.TypeTime,
	}
}

func (f *fieldsDef) ModifiedAt() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ModifiedAt",
		Type: meta.TypeTime,
	}
}

/*
 for monitor
*/
func (f *fieldsDef) MonitorTime() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Time",
		Type: meta.TypeTime,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorCPUTime() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "CPUTime",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDiskRead() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Read",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDiskWrite() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Write",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorRouterIn() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "In",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorRouterOut() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Out",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorInterfaceSend() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Send",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorInterfaceReceive() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "Receive",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorFreeDiskSize() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "FreeDiskSize",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseTotalMemorySize() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "TotalMemorySize ",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseUsedMemorySize() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "UsedMemorySize",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseTotalDisk1Size() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "TotalDisk1Size",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseUsedDisk1Size() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "UsedDisk1Size",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseTotalDisk2Size() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "TotalDisk2Size",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseUsedDisk2Size() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "UsedDisk2Size",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseBinlogUsedSizeKiB() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "BinlogUsedSizeKiB",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDatabaseDelayTimeSec() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DelayTimeSec",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorResponseTimeSec() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ResponseTimeSec",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorUplinkBPS() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "UplinkBPS",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorDownlinkBPS() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "DownlinkBPS",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorActiveConnections() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ActiveConnections",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}

func (f *fieldsDef) MonitorConnectionsPerSec() *schema.FieldDesc {
	return &schema.FieldDesc{
		Name: "ConnectionsPerSec",
		Type: meta.TypeFloat64,
		Tags: &schema.FieldTags{
			MapConv: ",omitempty",
			JSON:    ",omitempty",
		},
	}
}
