package perm

type Perm string

type FarmPerm Perm

const (
	FarmOwner   FarmPerm = "owner"
	FarmManager FarmPerm = "manager"
	FarmViewer  FarmPerm = "viewer"
	FarmNone    FarmPerm = "none"
)

type MinerPerm Perm

const (
	MinerOwner   MinerPerm = "owner"
	MinerManager MinerPerm = "manager"
	MinerViewer  MinerPerm = "viewer"
	MinerNone    MinerPerm = "none"
)

type FsPerm Perm

const (
	FsOwner   FsPerm = "owner"
	FsManager FsPerm = "manager"
	FsViewer  FsPerm = "viewer"
	FsNone    FsPerm = "none"
)
