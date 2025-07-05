package config

const (
	AppName        = "frg"
	AppShortDesc   = "Forge – project dashboard for devs"
	AppLongDesc    = "Forge is a CLI + TUI tool to manage your local dev projects."
	DefaultDevPath = "dev"

	PackageJSONFile = "package.json"

	NodeModulesDir  = "node_modules"
	GitDir          = ".git"
	HiddenDirPrefix = "."

	GitCommand         = "git"
	GitChangeDir       = "-C"
	GitLogCommand      = "log"
	GitLogOneCommit    = "-1"
	GitLogFormat       = "--format=%ct"
	GitRevParseCommand = "rev-parse"
	GitRevParseGitDir  = "--git-dir"

	TUITitle         = "forge-projects"
	KeyQuit          = "ctrl+c"
	ProjectSeparator = " • "

	SkipNoPackageJSON = "⚠️ Skipping %s: no package.json\n"
	SkipGeneralError  = "⚠️ Skipping %s: %v\n"

	DocMarginH        = 2
	DocMarginV        = 1
	TitleBarPaddingH  = 0
	TitleBarPaddingV  = 1
	TitlePadding      = 0
	StatusPaddingH    = 1
	StatusPaddingV    = 0
	ItemPaddingLeft   = 1
	HelpPaddingTop    = 0
	HelpPaddingLeft   = 1
	BorderPaddingLeft = 2
)
