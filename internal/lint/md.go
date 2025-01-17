package lint

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/errata-ai/vale/v2/internal/core"
	"github.com/jdkato/regexp"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	grh "github.com/yuin/goldmark/renderer/html"
)

// Markdown configuration.
var goldMd = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM,
	),
	goldmark.WithRendererOptions(
		grh.WithUnsafe(),
	),
)

// Convert extended info strings -- e.g., ```callout{'title': 'NOTE'} -- that
// might confuse Blackfriday into normal "```".
var reExInfo = regexp.MustCompile("`{3,}" + `.+`)
var reLinks = regexp.MustCompile(`\[(.+?)\]\((.+?)\)`)
var reCode = regexp.MustCompile(`<code>|</code>`)

func (l Linter) lintMarkdown(f *core.File) error {
	var buf bytes.Buffer

	s, err := l.prep(f.Content, "<pre><code>$1</code></pre>", "<code>$1</code>", ".md")
	if err != nil {
		return err
	}

	// Convert all links to `[${1}](<...>)`, so that spaces (common in
	// templating) are supported.
	//
	// The key assumption here is that we don't care what the URL is -- we just
	// want to ensure that we ignore its content while maintaining valid link
	// syntax.
	s = core.ReplaceAllStringSubmatchFunc(reLinks, s, func(g []string) string {
		return fmt.Sprintf("[%s](<%s>)", g[1], reCode.ReplaceAllString(g[2], ""))
	})

	if err := goldMd.Convert([]byte(s), &buf); err != nil {
		return core.NewE100(f.Path, err)
	}

	// NOTE: This is required to avoid finding matches inside info strings. For
	// example, if we're looking for 'json' we many incorrectly report the
	// location as being in an infostring like '```json'.
	//
	// See https://github.com/errata-ai/vale/v2/issues/248.
	body := reExInfo.ReplaceAllStringFunc(f.Content, func(m string) string {
		parts := strings.Split(m, "`")

		// This ensures that we respect the number of openning backticks, which
		// could be more than 3.
		//
		// See https://github.com/errata-ai/vale/v2/issues/271.
		tags := strings.Repeat("`", len(parts)-1)
		span := strings.Repeat("*", len(parts[len(parts)-1]))

		return tags + span
	})

	f.Content = body
	return l.lintHTMLTokens(f, buf.Bytes(), 0)
}
