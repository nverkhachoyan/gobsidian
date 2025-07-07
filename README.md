# Gobsidian

The simplest static site generator for Obsidian notes written in Go.

### Run with this command

```bash
go run ./cmd
```

### For Live Reload

Install [air](https://github.com/air-verse/air)

```bash
go install github.com/air-verse/air@latest
```

And run it at the root directory.

```bash
air
```

### Flags

`--serve` serve the generated website after building.

`--port` port to serve the website on.

`--vhealth` prints out health diagnostics on your vault, and saves it to `vault-diagnostics.json` in the output directory.

`--clear` removes the public directory before generating the site.

`--cpuprofile <filename>` to generate CPU profiler file.

`--memprofile <filename>` to generate memory profiler file.

### Config

`.config.toml` is the config file.

```toml
input_directory = "posts"
output_directory = "public"
site_title = "Nver Khachoyan's Blog"
site_subtitle = "My Personal Blog - Built with Go and Obsidian Notes"
notes_per_page = 5
base_url = "/"
env = "production"
```

`input_directory` is the obsidian vault directory.

### Current state

1. [x] Backlinks
2. [x] Backlink embeds
3. [x] Obsidian style images
4. [x] Frontmatter parsing (yaml)
5. [x] Folders and Subfolders
6. [x] Paginated content (set # pages in config)

### Frontmatter

Supported frontmatter fields are `title`, `date`, `author`, and `updatedAt`.

You can demo it at [blog.nverk.me](https://blog.nverk.me/).
