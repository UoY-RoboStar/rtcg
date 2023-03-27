package templating

// Generator is a mixin struct for dealing with templated parts of generators.
type Generator struct {
	Files     []File // Files is the list of files this generator will make.
	Templates Set    // Templates is the map of templates to use for files in testFiles.
}

// NewGenerator creates a Generator using files and builder.
func NewGenerator(files []File, builder SetBuilder) (Generator, error) {
	var gen Generator

	tset, err := builder.BuildFiles(files)
	if err != nil {
		return gen, err
	}

	gen.Files = files
	gen.Templates = tset

	return gen, nil
}

// Generate generates each of the templated files in this generator into dir.
// We pass ctx to each template.
func (t *Generator) Generate(dir string, ctx any) error {
	return t.Templates.Generate(t.Files, dir, ctx)
}
