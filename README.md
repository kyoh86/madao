# madao

Markdown refactoring tool

This tool is still in the PoC phase, and there's no stuff that can be used in practice.

## Requirements

### Setup `Pandoc`

Install [`Pandoc`](https://pandoc.org/).

### Set `LUA_PATH`

You should set `LUA_PATH` by `luarocks` with a lua-version which `Pandoc` has.

For example, using `Pandoc 2.17.1.1`, you should call `eval "$(luarocks --lua-version=5.3 path)"`.

# LICENSE

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://www.opensource.org/licenses/MIT)

This is distributed under the [MIT License](http://www.opensource.org/licenses/MIT).
