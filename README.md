# srcurl

Generate a URL from the filename

## Installation

Unix:

```sh
go install https://github.com/hexium310/srcurl/cmd
```

Windows:

```sh
go install -ldflags '-H windowsgui' https://github.com/hexium310/srcurl/cmd
```

## Configuration

You must create the configuration file somewhere below:

- The path passed to the `--config` option
- `$XDG_CONFIG_HOME/srcurl/srcurl.toml`
- `$HOME/.config/srcurl/srcurl.toml` (Unix)
- `%LOCALAPPDATA%\srcurl\srcurl.toml` (Windows)

Configuration file exapmple:

```toml
#:schema https://raw.githubusercontent.com/hexium310/srcurl/refs/heads/master/schemas/srcurl.json

[[sites]]
# Optional
name = "pixiv"
patterns = [
  # Regexp accepted in Go. See https://pkg.go.dev/regexp/syntax.
  # You must include the capture named `id` as in `(?P<id>[0-9]+)`.
  "^(?P<id>[0-9]+)_p[0-9]+",
]
# Text template of Go. See  https://pkg.go.dev/text/template.
# You must include `{{.Id}}` to be replaced with a value matching the `id` capture.
url = "https://pixiv.net/artworks/{{.Id}}"
```

## Usage

You can open a URL

```sh
# If you use above configuration file example, this opens https://pixiv.net/artworks/0000
srcurl --open ~/pictures/0000_p0.png
```
or copy one into the clipboard

```sh
# If you use above configuration file example, this copy https://pixiv.net/artworks/0000 into the clipboard
srcurl --copy ~/pictures/0000_p0.png
```

When you pass neither `--open` nor `--copy`, this command prints an URL.
