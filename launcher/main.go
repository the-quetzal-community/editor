// Package launcher provides a way to launch the editor when imported into a graphics.gd project.
//
// Use an underscore import if you just want to import the engine classes.
package launcher

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/startup"

	"the.quetzal.community/editor/internal"
)

func init() {
	classdb.Register[internal.Tree]()
	classdb.Register[internal.Rock]()
	classdb.Register[internal.TerrainTile]()
	classdb.Register[internal.World]()
	classdb.Register[internal.UI]()
	classdb.Register[internal.PreviewRenderer]()
	classdb.Register[internal.Renderer]()
	classdb.Register[internal.EditorPlugin]()
	classdb.Register[internal.ModelLoader]()
}

func Run() {
	startup.LoadingScene()
	SceneTree.Add(new(internal.World))
	startup.Scene()
}
