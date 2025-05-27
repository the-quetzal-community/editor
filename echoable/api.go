package echoable

import (
	"runtime.link/api"

	"graphics.gd/variant/Color"
	"graphics.gd/variant/Euler"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector3"
)

// Interactions specification for a serialisable instruction set that records the composition for a particular 3D creation.
type API struct {
	_ struct{}

	api.Specification

	ShareAvatar func(is Asset, at Transform3D.BasisOrigin) error `isle:"5facb3a2-2662-4511-8429-ea5756a79f6d"`

	ImportAsset func(uri string, as Asset) error                           `isle:"205a7ab4-9899-4ff8-b33c-524f606bce94"`
	InsertAsset func(do Asset, as Thing, at Transform3D.BasisOrigin) error `isle:"c8fd6e2b-2c22-470a-9e24-100168b843dd"`

	AttachThing func(to Thing, do Thing) error `isle:"870ecfc2-73e9-4f48-abf6-a4c5296cd490"`
	DetachThing func(id Thing) error           `isle:"1a29ae6a-9c2e-4192-bbc7-35042d28f66d"`

	BranchThing func(do Thing, as Thing, at Transform3D.BasisOrigin) error `isle:"25ee2cfc-0b94-42b8-a2b8-0b0eadaf2688"`
	OffsetThing func(to Thing, do Vector3.XYZ) error                       `isle:"dc1bdd1d-a7dd-416b-ab48-028dc7b28321"`
	RotateThing func(to Thing, do Euler.Radians) error                     `isle:"9c5d62dd-887c-45f7-9848-9cc55c948e00"`
	ResizeThing func(to Thing, do Vector3.XYZ) error                       `isle:"acbfd951-6645-4774-a65a-2c3b511e6e2d"`
	ColourThing func(to Thing, do Color.RGBA) error                        `isle:"340ce211-25d0-4e44-ac67-335f664efc60"`
	RemoveThing func(id Thing) error                                       `isle:"28f3c549-efe2-46a8-9064-f7667a3780b7"`

	FollowCurve func(id Thing, at, hz Ticks, p1, p2, p3, p4 Vector3.XYZ) error `isle:"68c2f41c-a118-47a2-9be0-d943e6314b2b"`
	WalkTerrain func(id Thing, at, hz Ticks, do Vector3.XYZ) error             `isle:"b0670d9b-7020-43b0-8c18-c81f750198eb"`

	ResizeSpace func(do Vector3.XYZ) error `isle:"19892c2a-f594-41a3-b167-474c2c833f6b"`

	LiftTerrain func(at Vector3.XYZ, in float32, do float32, xy float32) error `isle:"56803d02-6a93-44d4-a7f1-3c1d3ec4578f"`
	AdjustWater func(at Vector3.XYZ, in float32, do float32) error             `isle:"f8b9651d-0342-4b55-996c-5436c6618569"`

	DrawTexture func(at Vector3.XYZ, in float32, do Asset) error `isle:"c0809c7f-4499-4222-a2e8-a2ce0ec4e86d"`
}

type Asset uint32
type Group uint32
type Frame uint32
type Thing uint64
type Ticks uint32
