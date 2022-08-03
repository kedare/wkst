local brew = require("brew")
local xcode = require("xcode")

local brew_packages = {
    -- Universal tools
    "zsh", -- shell
    "ncdu", -- Better du
    "tmux", -- Terminal multiplexer
    "wget", -- Downloader
    "k9s", -- k9s management tool
    "dropbox", -- Google drive but better
    "iterm2", -- Better terminal for MacOs
    "slack", -- Blabla
    "visual-studio-code", -- Ok'ish code editor
    "neovim", -- Better editor
    "gh", -- Github CLI
    "mtr", -- Better traceroute
    "gpg", -- Safety first
    "sipcalc", -- Subnet calculator
    "jq", -- JSON ETL

    -- Dependencies to build Python
    "openssl",
    "readline",
    "sqlite3",
    "xz",
    "zlib",
    "tcl-tk"
}

xcode:setup()
brew:setup()
brew:ensure_packages(brew_packages)
