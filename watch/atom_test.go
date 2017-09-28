package watch

import "testing"

func TestAtomParse(t *testing.T) {
	data := `
<?xml version="1.0" encoding="UTF-8"?>
<feed xmlns="http://www.w3.org/2005/Atom" xmlns:media="http://search.yahoo.com/mrss/" xml:lang="en-US">
  <id>tag:github.com,2008:/elves/elvish/releases</id>
  <link type="text/html" rel="alternate" href="/elves/elvish/releases"/>
  <link type="application/atom+xml" rel="self" href="/elves/elvish/releases.atom"/>
  <title>Tags from elvish</title>
  <updated>2017-09-19T05:52:32+08:00</updated>
  <entry>
    <id>tag:github.com,2008:Repository/10717754/0.10.1</id>
    <updated>2017-09-19T05:52:32+08:00</updated>
    <link rel="alternate" type="text/html" href="/elves/elvish/releases/tag/0.10.1"/>
    <title>0.10.1</title>
    <content></content>
    <author>
      <name>xiaq</name>
    </author>
    <media:thumbnail height="30" width="30" url="https://avatars3.githubusercontent.com/u/582021?v=4&amp;s=60"/>
  </entry>
  <entry>
    <id>tag:github.com,2008:Repository/10717754/0.8</id>
    <author>
      <name>xiaq</name>
    </author>
  </entry>
</feed>
`
	entry, err := atomGetLatest([]byte(data))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", entry)
}
