package catkin

import (
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// Package encodes a subset of the package.xml format used by Catkin.
type Package struct {
	XMLName xml.Name `xml:"package"`               // XMLName names the root tag "package".
	Format  int      `xml:"format,attr,omitempty"` // Format is the manifest format (we use format 2).

	Name        string     `xml:"name"`        // Name is the name of the package.
	Version     string     `xml:"version"`     // Version is the version of the package.
	Description string     `xml:"description"` // Description is the description of the package.
	Maintainer  Maintainer `xml:"maintainer"`  // Maintainer is the maintainer of the package.
	License     string     `xml:"license"`     // License is the license of the package.

	BuildtoolDepends []string `xml:"buildtool_depend"` // BuildtoolDepends holds build-tool depends.
	Depends          []string `xml:"depend"`           // Depends holds normal depends.
}

// Maintainer encodes a package maintainer.
type Maintainer struct {
	Email string `xml:"email,attr"` // Email is the email address of the maintainer.
	Name  string `xml:",chardata"`  // Name is the name of the maintainer.
}

// Write writes package XML to w.
func (p *Package) Write(w io.Writer) error {
	if _, err := io.WriteString(w, xml.Header); err != nil {
		return fmt.Errorf("couldn't write package XML header: %w", err)
	}

	xenc := xml.NewEncoder(w)
	xenc.Indent("", "\t")

	if err := xenc.Encode(p); err != nil {
		return fmt.Errorf("couldn't write package XML body: %w", err)
	}

	return nil
}

// Autofill fills the missing fields of p with default templates.
//
// These will need to be expanded using Expand.
func (p *Package) Autofill() {
	// We don't support other formats.
	p.Format = 2

	p.Name = autofillOne(p.Name, "${NAME}")
	p.Description = autofillOne(p.Description, "Package for ${NAME}")
	p.Maintainer.Name = autofillOne(p.Maintainer.Name, "${NAME} Maintainers")
	p.Maintainer.Email = autofillOne(p.Maintainer.Email, "email@example.com")
	p.License = autofillOne(p.License, "Proprietary")
}

func autofillOne(input, def string) string {
	if input == "" {
		return def
	}

	return input
}

// Expand replaces all instances of '${NAME}' in the metadata of package p with name.
func (p *Package) Expand(name string) {
	p.Name = expandOne(p.Name, name)
	p.Description = expandOne(p.Description, name)
	p.Maintainer.Name = expandOne(p.Maintainer.Name, name)
	p.Maintainer.Email = expandOne(p.Maintainer.Email, name)
	p.License = expandOne(p.License, name)
}

func expandOne(template, name string) string {
	return strings.ReplaceAll(template, "${NAME}", name)
}
