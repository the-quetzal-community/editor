package internal

import (
	"math/rand/v2"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/classdb/PackedScene"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Texture2D"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/Vector3"
	"the.quetzal.community/editor/echoable"
)

// PreviewRenderer is responsible for rendering items when the user
// is planning where to place it. As such, these items will follow
// the cursor and will be submitted to the Vulture API on click.
type PreviewRenderer struct {
	classdb.Extension[PreviewRenderer, Node3D.Instance]

	mouseOver chan Vector3.XYZ

	preview chan Path.ToResource // resource name

	edits   echoable.API
	terrain *Renderer

	current string
}

func (pr *PreviewRenderer) AsNode() Node.Instance { return pr.Super().AsNode() }

func (pr *PreviewRenderer) Process(dt Float.X) {
	for {
		select {
		case resource := <-pr.preview:
			scene := Resource.Load[PackedScene.Instance](resource)
			instance, ok := Object.As[Node3D.Instance](scene.Instantiate())
			if ok {
				if pr.Super().AsNode().GetChildCount() > 0 {
					Node.Instance(pr.Super().AsNode().GetChild(0)).QueueFree()
				}
				instance.AsNode3D().SetScale(Vector3.MulX(instance.AsNode3D().Scale(), 0.1))
				pr.Super().AsNode().AddChild(instance.AsNode())
			}
			pr.current = resource.String()
		case pos := <-pr.mouseOver:
			pr.Super().AsNode3D().SetPosition(pos)
			continue
		default:

		}
		break
	}
	if Input.IsMouseButtonPressed(Input.MouseButtonLeft) {
		if pr.Super().AsNode().GetChildCount() > 0 {
			Node.Instance(pr.Super().AsNode().GetChild(0)).QueueFree()
			if err := pr.edits.InsertAsset(echoable.Asset(rand.Uint64()), echoable.Thing(rand.Uint64()), pr.Super().AsNode3D().AsNode3D().GlobalTransform()); err != nil {
				Engine.Raise(err)
			}
		}
	}
	pos := pr.Super().AsNode3D().Position()
	pos.Y = (pr.terrain.HeightAt(pos))
	pr.Super().AsNode3D().SetPosition(pos)
}

func (pr *PreviewRenderer) Ready() {

}

func (pr *PreviewRenderer) GenerateTexture2D(scene PackedScene.Instance) Texture2D.Instance {
	return Texture2D.Instance{}
}
