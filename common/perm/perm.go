package perm

type Perm string

type FarmPerm Perm

const (
	FarmOwner   FarmPerm = "owner"
	FarmManager FarmPerm = "manager"
	FarmViewer  FarmPerm = "viewer"
)

type MinerPerm Perm

const (
	MinerOwner   MinerPerm = "owner"
	MinerManager MinerPerm = "manager"
	MinerViewer  MinerPerm = "viewer"
)

type FsPerm Perm

const (
	FsOwner   FsPerm = "owner"
	FsManager FsPerm = "manager"
	FsViewer  FsPerm = "viewer"
)
