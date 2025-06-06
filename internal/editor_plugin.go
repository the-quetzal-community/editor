package internal

import (
	"graphics.gd/classdb"
	BaseEditorPlugin "graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/GLTFDocument"
)

type EditorPlugin struct {
	BaseEditorPlugin.Extension[EditorPlugin]
	classdb.Tool
}

func (ml *EditorPlugin) EnterTree() {
	GLTFDocument.RegisterGltfDocumentExtension(new(ModelLoader).AsGLTFDocumentExtension(), false)
}

func (ml *EditorPlugin) ExitTree() {
	GLTFDocument.UnregisterGltfDocumentExtension(new(ModelLoader).AsGLTFDocumentExtension())
}
