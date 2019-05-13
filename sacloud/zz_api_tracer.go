// generated by 'github.com/sacloud/libsacloud/internal/tools/gen-api-tracer'; DO NOT EDIT

package sacloud

import (
	"context"
	"log"

	"github.com/sacloud/libsacloud-v2/sacloud/types"
)

/*************************************************
* CDROMTracer
*************************************************/

// CDROMTracer is for trace CDROMOp operations
type CDROMTracer struct {
	Internal CDROMAPI
}

// NewCDROMTracer creates new CDROMTracer instance
func NewCDROMTracer(in CDROMAPI) *CDROMTracer {
	return &CDROMTracer{
		Internal: in,
	}
}

// Find is API call with trace log
func (t *CDROMTracer) Find(ctx context.Context, zone string, conditions *FindCondition) ([]*CDROM, error) {
	log.Println("[TRACE] CDROMTracer.Find start:	args => [", "zone=", zone, "conditions=", conditions, "]")
	defer func() {
		log.Println("[TRACE] CDROMTracer.Find: end")
	}()

	return t.Internal.Find(ctx, zone, conditions)
}

// Create is API call with trace log
func (t *CDROMTracer) Create(ctx context.Context, zone string, param *CDROMCreateRequest) (*CDROM, *FTPServer, error) {
	log.Println("[TRACE] CDROMTracer.Create start:	args => [", "zone=", zone, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] CDROMTracer.Create: end")
	}()

	return t.Internal.Create(ctx, zone, param)
}

// Read is API call with trace log
func (t *CDROMTracer) Read(ctx context.Context, zone string, id types.ID) (*CDROM, error) {
	log.Println("[TRACE] CDROMTracer.Read start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] CDROMTracer.Read: end")
	}()

	return t.Internal.Read(ctx, zone, id)
}

// Update is API call with trace log
func (t *CDROMTracer) Update(ctx context.Context, zone string, id types.ID, param *CDROMUpdateRequest) (*CDROM, error) {
	log.Println("[TRACE] CDROMTracer.Update start:	args => [", "zone=", zone, "id=", id, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] CDROMTracer.Update: end")
	}()

	return t.Internal.Update(ctx, zone, id, param)
}

// Delete is API call with trace log
func (t *CDROMTracer) Delete(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] CDROMTracer.Delete start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] CDROMTracer.Delete: end")
	}()

	return t.Internal.Delete(ctx, zone, id)
}

// OpenFTP is API call with trace log
func (t *CDROMTracer) OpenFTP(ctx context.Context, zone string, id types.ID, openOption *OpenFTPParam) (*FTPServer, error) {
	log.Println("[TRACE] CDROMTracer.OpenFTP start:	args => [", "zone=", zone, "id=", id, "openOption=", openOption, "]")
	defer func() {
		log.Println("[TRACE] CDROMTracer.OpenFTP: end")
	}()

	return t.Internal.OpenFTP(ctx, zone, id, openOption)
}

// CloseFTP is API call with trace log
func (t *CDROMTracer) CloseFTP(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] CDROMTracer.CloseFTP start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] CDROMTracer.CloseFTP: end")
	}()

	return t.Internal.CloseFTP(ctx, zone, id)
}

/*************************************************
* GSLBTracer
*************************************************/

// GSLBTracer is for trace GSLBOp operations
type GSLBTracer struct {
	Internal GSLBAPI
}

// NewGSLBTracer creates new GSLBTracer instance
func NewGSLBTracer(in GSLBAPI) *GSLBTracer {
	return &GSLBTracer{
		Internal: in,
	}
}

// Find is API call with trace log
func (t *GSLBTracer) Find(ctx context.Context, zone string, conditions *FindCondition) ([]*GSLB, error) {
	log.Println("[TRACE] GSLBTracer.Find start:	args => [", "zone=", zone, "conditions=", conditions, "]")
	defer func() {
		log.Println("[TRACE] GSLBTracer.Find: end")
	}()

	return t.Internal.Find(ctx, zone, conditions)
}

// Create is API call with trace log
func (t *GSLBTracer) Create(ctx context.Context, zone string, param *GSLBCreateRequest) (*GSLB, error) {
	log.Println("[TRACE] GSLBTracer.Create start:	args => [", "zone=", zone, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] GSLBTracer.Create: end")
	}()

	return t.Internal.Create(ctx, zone, param)
}

// Read is API call with trace log
func (t *GSLBTracer) Read(ctx context.Context, zone string, id types.ID) (*GSLB, error) {
	log.Println("[TRACE] GSLBTracer.Read start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] GSLBTracer.Read: end")
	}()

	return t.Internal.Read(ctx, zone, id)
}

// Update is API call with trace log
func (t *GSLBTracer) Update(ctx context.Context, zone string, id types.ID, param *GSLBUpdateRequest) (*GSLB, error) {
	log.Println("[TRACE] GSLBTracer.Update start:	args => [", "zone=", zone, "id=", id, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] GSLBTracer.Update: end")
	}()

	return t.Internal.Update(ctx, zone, id, param)
}

// Delete is API call with trace log
func (t *GSLBTracer) Delete(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] GSLBTracer.Delete start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] GSLBTracer.Delete: end")
	}()

	return t.Internal.Delete(ctx, zone, id)
}

/*************************************************
* NFSTracer
*************************************************/

// NFSTracer is for trace NFSOp operations
type NFSTracer struct {
	Internal NFSAPI
}

// NewNFSTracer creates new NFSTracer instance
func NewNFSTracer(in NFSAPI) *NFSTracer {
	return &NFSTracer{
		Internal: in,
	}
}

// Find is API call with trace log
func (t *NFSTracer) Find(ctx context.Context, zone string, conditions *FindCondition) ([]*NFS, error) {
	log.Println("[TRACE] NFSTracer.Find start:	args => [", "zone=", zone, "conditions=", conditions, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.Find: end")
	}()

	return t.Internal.Find(ctx, zone, conditions)
}

// Create is API call with trace log
func (t *NFSTracer) Create(ctx context.Context, zone string, param *NFSCreateRequest) (*NFS, error) {
	log.Println("[TRACE] NFSTracer.Create start:	args => [", "zone=", zone, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.Create: end")
	}()

	return t.Internal.Create(ctx, zone, param)
}

// Read is API call with trace log
func (t *NFSTracer) Read(ctx context.Context, zone string, id types.ID) (*NFS, error) {
	log.Println("[TRACE] NFSTracer.Read start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.Read: end")
	}()

	return t.Internal.Read(ctx, zone, id)
}

// Update is API call with trace log
func (t *NFSTracer) Update(ctx context.Context, zone string, id types.ID, param *NFSUpdateRequest) (*NFS, error) {
	log.Println("[TRACE] NFSTracer.Update start:	args => [", "zone=", zone, "id=", id, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.Update: end")
	}()

	return t.Internal.Update(ctx, zone, id, param)
}

// Delete is API call with trace log
func (t *NFSTracer) Delete(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] NFSTracer.Delete start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.Delete: end")
	}()

	return t.Internal.Delete(ctx, zone, id)
}

// Boot is API call with trace log
func (t *NFSTracer) Boot(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] NFSTracer.Boot start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.Boot: end")
	}()

	return t.Internal.Boot(ctx, zone, id)
}

// Shutdown is API call with trace log
func (t *NFSTracer) Shutdown(ctx context.Context, zone string, id types.ID, shutdownOption *ShutdownOption) error {
	log.Println("[TRACE] NFSTracer.Shutdown start:	args => [", "zone=", zone, "id=", id, "shutdownOption=", shutdownOption, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.Shutdown: end")
	}()

	return t.Internal.Shutdown(ctx, zone, id, shutdownOption)
}

// Reset is API call with trace log
func (t *NFSTracer) Reset(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] NFSTracer.Reset start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.Reset: end")
	}()

	return t.Internal.Reset(ctx, zone, id)
}

// MonitorFreeDiskSize is API call with trace log
func (t *NFSTracer) MonitorFreeDiskSize(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*FreeDiskSizeActivity, error) {
	log.Println("[TRACE] NFSTracer.MonitorFreeDiskSize start:	args => [", "zone=", zone, "id=", id, "condition=", condition, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.MonitorFreeDiskSize: end")
	}()

	return t.Internal.MonitorFreeDiskSize(ctx, zone, id, condition)
}

// MonitorInterface is API call with trace log
func (t *NFSTracer) MonitorInterface(ctx context.Context, zone string, id types.ID, condition *MonitorCondition) (*InterfaceActivity, error) {
	log.Println("[TRACE] NFSTracer.MonitorInterface start:	args => [", "zone=", zone, "id=", id, "condition=", condition, "]")
	defer func() {
		log.Println("[TRACE] NFSTracer.MonitorInterface: end")
	}()

	return t.Internal.MonitorInterface(ctx, zone, id, condition)
}

/*************************************************
* NoteTracer
*************************************************/

// NoteTracer is for trace NoteOp operations
type NoteTracer struct {
	Internal NoteAPI
}

// NewNoteTracer creates new NoteTracer instance
func NewNoteTracer(in NoteAPI) *NoteTracer {
	return &NoteTracer{
		Internal: in,
	}
}

// Find is API call with trace log
func (t *NoteTracer) Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Note, error) {
	log.Println("[TRACE] NoteTracer.Find start:	args => [", "zone=", zone, "conditions=", conditions, "]")
	defer func() {
		log.Println("[TRACE] NoteTracer.Find: end")
	}()

	return t.Internal.Find(ctx, zone, conditions)
}

// Create is API call with trace log
func (t *NoteTracer) Create(ctx context.Context, zone string, param *NoteCreateRequest) (*Note, error) {
	log.Println("[TRACE] NoteTracer.Create start:	args => [", "zone=", zone, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] NoteTracer.Create: end")
	}()

	return t.Internal.Create(ctx, zone, param)
}

// Read is API call with trace log
func (t *NoteTracer) Read(ctx context.Context, zone string, id types.ID) (*Note, error) {
	log.Println("[TRACE] NoteTracer.Read start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] NoteTracer.Read: end")
	}()

	return t.Internal.Read(ctx, zone, id)
}

// Update is API call with trace log
func (t *NoteTracer) Update(ctx context.Context, zone string, id types.ID, param *NoteUpdateRequest) (*Note, error) {
	log.Println("[TRACE] NoteTracer.Update start:	args => [", "zone=", zone, "id=", id, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] NoteTracer.Update: end")
	}()

	return t.Internal.Update(ctx, zone, id, param)
}

// Delete is API call with trace log
func (t *NoteTracer) Delete(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] NoteTracer.Delete start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] NoteTracer.Delete: end")
	}()

	return t.Internal.Delete(ctx, zone, id)
}

/*************************************************
* SwitchTracer
*************************************************/

// SwitchTracer is for trace SwitchOp operations
type SwitchTracer struct {
	Internal SwitchAPI
}

// NewSwitchTracer creates new SwitchTracer instance
func NewSwitchTracer(in SwitchAPI) *SwitchTracer {
	return &SwitchTracer{
		Internal: in,
	}
}

// Find is API call with trace log
func (t *SwitchTracer) Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Switch, error) {
	log.Println("[TRACE] SwitchTracer.Find start:	args => [", "zone=", zone, "conditions=", conditions, "]")
	defer func() {
		log.Println("[TRACE] SwitchTracer.Find: end")
	}()

	return t.Internal.Find(ctx, zone, conditions)
}

// Create is API call with trace log
func (t *SwitchTracer) Create(ctx context.Context, zone string, param *SwitchCreateRequest) (*Switch, error) {
	log.Println("[TRACE] SwitchTracer.Create start:	args => [", "zone=", zone, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] SwitchTracer.Create: end")
	}()

	return t.Internal.Create(ctx, zone, param)
}

// Read is API call with trace log
func (t *SwitchTracer) Read(ctx context.Context, zone string, id types.ID) (*Switch, error) {
	log.Println("[TRACE] SwitchTracer.Read start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] SwitchTracer.Read: end")
	}()

	return t.Internal.Read(ctx, zone, id)
}

// Update is API call with trace log
func (t *SwitchTracer) Update(ctx context.Context, zone string, id types.ID, param *SwitchUpdateRequest) (*Switch, error) {
	log.Println("[TRACE] SwitchTracer.Update start:	args => [", "zone=", zone, "id=", id, "param=", param, "]")
	defer func() {
		log.Println("[TRACE] SwitchTracer.Update: end")
	}()

	return t.Internal.Update(ctx, zone, id, param)
}

// Delete is API call with trace log
func (t *SwitchTracer) Delete(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] SwitchTracer.Delete start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] SwitchTracer.Delete: end")
	}()

	return t.Internal.Delete(ctx, zone, id)
}

// ConnectToBridge is API call with trace log
func (t *SwitchTracer) ConnectToBridge(ctx context.Context, zone string, id types.ID, bridgeID types.ID) error {
	log.Println("[TRACE] SwitchTracer.ConnectToBridge start:	args => [", "zone=", zone, "id=", id, "bridgeID=", bridgeID, "]")
	defer func() {
		log.Println("[TRACE] SwitchTracer.ConnectToBridge: end")
	}()

	return t.Internal.ConnectToBridge(ctx, zone, id, bridgeID)
}

// DisconnectFromBridge is API call with trace log
func (t *SwitchTracer) DisconnectFromBridge(ctx context.Context, zone string, id types.ID) error {
	log.Println("[TRACE] SwitchTracer.DisconnectFromBridge start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] SwitchTracer.DisconnectFromBridge: end")
	}()

	return t.Internal.DisconnectFromBridge(ctx, zone, id)
}

/*************************************************
* ZoneTracer
*************************************************/

// ZoneTracer is for trace ZoneOp operations
type ZoneTracer struct {
	Internal ZoneAPI
}

// NewZoneTracer creates new ZoneTracer instance
func NewZoneTracer(in ZoneAPI) *ZoneTracer {
	return &ZoneTracer{
		Internal: in,
	}
}

// Find is API call with trace log
func (t *ZoneTracer) Find(ctx context.Context, zone string, conditions *FindCondition) ([]*Zone, error) {
	log.Println("[TRACE] ZoneTracer.Find start:	args => [", "zone=", zone, "conditions=", conditions, "]")
	defer func() {
		log.Println("[TRACE] ZoneTracer.Find: end")
	}()

	return t.Internal.Find(ctx, zone, conditions)
}

// Read is API call with trace log
func (t *ZoneTracer) Read(ctx context.Context, zone string, id types.ID) (*Zone, error) {
	log.Println("[TRACE] ZoneTracer.Read start:	args => [", "zone=", zone, "id=", id, "]")
	defer func() {
		log.Println("[TRACE] ZoneTracer.Read: end")
	}()

	return t.Internal.Read(ctx, zone, id)
}
