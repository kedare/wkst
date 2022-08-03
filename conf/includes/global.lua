local asdf = require("asdf")

local asdf_plugins = {
    "python",
    "ruby",
    "golang",
    "nodejs",
    "helm"
}

local asdf_versions = {
    python = {"3.10.5"},
    ruby = {"3.1.0"},
    nodejs = {"17.6.0"},
    golang = {"1.18.4"}
}

asdf:setup()
asdf:ensure_plugins(asdf_plugins)
asdf:ensure_packages(asdf_versions)
