// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-fake-store'; DO NOT EDIT

package fake

import (
	"github.com/sacloud/libsacloud-v2/sacloud"
	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

func (s *store) getArchive(zone string) []*sacloud.Archive {
	values := s.get(ResourceArchive, zone)
	var ret []*sacloud.Archive
	for _, v := range values {
		if v, ok := v.(*sacloud.Archive); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getArchiveByID(zone string, id types.ID) *sacloud.Archive {
	v := s.getByID(ResourceArchive, zone, id)
	if v, ok := v.(*sacloud.Archive); ok {
		return v
	}
	return nil
}

func (s *store) setArchive(zone string, value *sacloud.Archive) {
	s.set(ResourceArchive, zone, value)
}

func (s *store) getBridge(zone string) []*sacloud.Bridge {
	values := s.get(ResourceBridge, zone)
	var ret []*sacloud.Bridge
	for _, v := range values {
		if v, ok := v.(*sacloud.Bridge); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getBridgeByID(zone string, id types.ID) *sacloud.Bridge {
	v := s.getByID(ResourceBridge, zone, id)
	if v, ok := v.(*sacloud.Bridge); ok {
		return v
	}
	return nil
}

func (s *store) setBridge(zone string, value *sacloud.Bridge) {
	s.set(ResourceBridge, zone, value)
}

func (s *store) getCDROM(zone string) []*sacloud.CDROM {
	values := s.get(ResourceCDROM, zone)
	var ret []*sacloud.CDROM
	for _, v := range values {
		if v, ok := v.(*sacloud.CDROM); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getCDROMByID(zone string, id types.ID) *sacloud.CDROM {
	v := s.getByID(ResourceCDROM, zone, id)
	if v, ok := v.(*sacloud.CDROM); ok {
		return v
	}
	return nil
}

func (s *store) setCDROM(zone string, value *sacloud.CDROM) {
	s.set(ResourceCDROM, zone, value)
}

func (s *store) getDisk(zone string) []*sacloud.Disk {
	values := s.get(ResourceDisk, zone)
	var ret []*sacloud.Disk
	for _, v := range values {
		if v, ok := v.(*sacloud.Disk); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getDiskByID(zone string, id types.ID) *sacloud.Disk {
	v := s.getByID(ResourceDisk, zone, id)
	if v, ok := v.(*sacloud.Disk); ok {
		return v
	}
	return nil
}

func (s *store) setDisk(zone string, value *sacloud.Disk) {
	s.set(ResourceDisk, zone, value)
}

func (s *store) getGSLB(zone string) []*sacloud.GSLB {
	values := s.get(ResourceGSLB, zone)
	var ret []*sacloud.GSLB
	for _, v := range values {
		if v, ok := v.(*sacloud.GSLB); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getGSLBByID(zone string, id types.ID) *sacloud.GSLB {
	v := s.getByID(ResourceGSLB, zone, id)
	if v, ok := v.(*sacloud.GSLB); ok {
		return v
	}
	return nil
}

func (s *store) setGSLB(zone string, value *sacloud.GSLB) {
	s.set(ResourceGSLB, zone, value)
}

func (s *store) getInterface(zone string) []*sacloud.Interface {
	values := s.get(ResourceInterface, zone)
	var ret []*sacloud.Interface
	for _, v := range values {
		if v, ok := v.(*sacloud.Interface); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getInterfaceByID(zone string, id types.ID) *sacloud.Interface {
	v := s.getByID(ResourceInterface, zone, id)
	if v, ok := v.(*sacloud.Interface); ok {
		return v
	}
	return nil
}

func (s *store) setInterface(zone string, value *sacloud.Interface) {
	s.set(ResourceInterface, zone, value)
}

func (s *store) getInternet(zone string) []*sacloud.Internet {
	values := s.get(ResourceInternet, zone)
	var ret []*sacloud.Internet
	for _, v := range values {
		if v, ok := v.(*sacloud.Internet); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getInternetByID(zone string, id types.ID) *sacloud.Internet {
	v := s.getByID(ResourceInternet, zone, id)
	if v, ok := v.(*sacloud.Internet); ok {
		return v
	}
	return nil
}

func (s *store) setInternet(zone string, value *sacloud.Internet) {
	s.set(ResourceInternet, zone, value)
}

func (s *store) getLoadBalancer(zone string) []*sacloud.LoadBalancer {
	values := s.get(ResourceLoadBalancer, zone)
	var ret []*sacloud.LoadBalancer
	for _, v := range values {
		if v, ok := v.(*sacloud.LoadBalancer); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getLoadBalancerByID(zone string, id types.ID) *sacloud.LoadBalancer {
	v := s.getByID(ResourceLoadBalancer, zone, id)
	if v, ok := v.(*sacloud.LoadBalancer); ok {
		return v
	}
	return nil
}

func (s *store) setLoadBalancer(zone string, value *sacloud.LoadBalancer) {
	s.set(ResourceLoadBalancer, zone, value)
}

func (s *store) getNFS(zone string) []*sacloud.NFS {
	values := s.get(ResourceNFS, zone)
	var ret []*sacloud.NFS
	for _, v := range values {
		if v, ok := v.(*sacloud.NFS); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getNFSByID(zone string, id types.ID) *sacloud.NFS {
	v := s.getByID(ResourceNFS, zone, id)
	if v, ok := v.(*sacloud.NFS); ok {
		return v
	}
	return nil
}

func (s *store) setNFS(zone string, value *sacloud.NFS) {
	s.set(ResourceNFS, zone, value)
}

func (s *store) getNote(zone string) []*sacloud.Note {
	values := s.get(ResourceNote, zone)
	var ret []*sacloud.Note
	for _, v := range values {
		if v, ok := v.(*sacloud.Note); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getNoteByID(zone string, id types.ID) *sacloud.Note {
	v := s.getByID(ResourceNote, zone, id)
	if v, ok := v.(*sacloud.Note); ok {
		return v
	}
	return nil
}

func (s *store) setNote(zone string, value *sacloud.Note) {
	s.set(ResourceNote, zone, value)
}

func (s *store) getPacketFilter(zone string) []*sacloud.PacketFilter {
	values := s.get(ResourcePacketFilter, zone)
	var ret []*sacloud.PacketFilter
	for _, v := range values {
		if v, ok := v.(*sacloud.PacketFilter); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getPacketFilterByID(zone string, id types.ID) *sacloud.PacketFilter {
	v := s.getByID(ResourcePacketFilter, zone, id)
	if v, ok := v.(*sacloud.PacketFilter); ok {
		return v
	}
	return nil
}

func (s *store) setPacketFilter(zone string, value *sacloud.PacketFilter) {
	s.set(ResourcePacketFilter, zone, value)
}

func (s *store) getServer(zone string) []*sacloud.Server {
	values := s.get(ResourceServer, zone)
	var ret []*sacloud.Server
	for _, v := range values {
		if v, ok := v.(*sacloud.Server); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getServerByID(zone string, id types.ID) *sacloud.Server {
	v := s.getByID(ResourceServer, zone, id)
	if v, ok := v.(*sacloud.Server); ok {
		return v
	}
	return nil
}

func (s *store) setServer(zone string, value *sacloud.Server) {
	s.set(ResourceServer, zone, value)
}

func (s *store) getSIM(zone string) []*sacloud.SIM {
	values := s.get(ResourceSIM, zone)
	var ret []*sacloud.SIM
	for _, v := range values {
		if v, ok := v.(*sacloud.SIM); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getSIMByID(zone string, id types.ID) *sacloud.SIM {
	v := s.getByID(ResourceSIM, zone, id)
	if v, ok := v.(*sacloud.SIM); ok {
		return v
	}
	return nil
}

func (s *store) setSIM(zone string, value *sacloud.SIM) {
	s.set(ResourceSIM, zone, value)
}

func (s *store) getSwitch(zone string) []*sacloud.Switch {
	values := s.get(ResourceSwitch, zone)
	var ret []*sacloud.Switch
	for _, v := range values {
		if v, ok := v.(*sacloud.Switch); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getSwitchByID(zone string, id types.ID) *sacloud.Switch {
	v := s.getByID(ResourceSwitch, zone, id)
	if v, ok := v.(*sacloud.Switch); ok {
		return v
	}
	return nil
}

func (s *store) setSwitch(zone string, value *sacloud.Switch) {
	s.set(ResourceSwitch, zone, value)
}

func (s *store) getVPCRouter(zone string) []*sacloud.VPCRouter {
	values := s.get(ResourceVPCRouter, zone)
	var ret []*sacloud.VPCRouter
	for _, v := range values {
		if v, ok := v.(*sacloud.VPCRouter); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getVPCRouterByID(zone string, id types.ID) *sacloud.VPCRouter {
	v := s.getByID(ResourceVPCRouter, zone, id)
	if v, ok := v.(*sacloud.VPCRouter); ok {
		return v
	}
	return nil
}

func (s *store) setVPCRouter(zone string, value *sacloud.VPCRouter) {
	s.set(ResourceVPCRouter, zone, value)
}

func (s *store) getZone(zone string) []*sacloud.Zone {
	values := s.get(ResourceZone, zone)
	var ret []*sacloud.Zone
	for _, v := range values {
		if v, ok := v.(*sacloud.Zone); ok {
			ret = append(ret, v)
		}
	}
	return ret
}

func (s *store) getZoneByID(zone string, id types.ID) *sacloud.Zone {
	v := s.getByID(ResourceZone, zone, id)
	if v, ok := v.(*sacloud.Zone); ok {
		return v
	}
	return nil
}

func (s *store) setZone(zone string, value *sacloud.Zone) {
	s.set(ResourceZone, zone, value)
}
