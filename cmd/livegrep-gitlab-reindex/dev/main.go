package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/jessevdk/go-flags"
)

var (
	flagCodesearch   = flag.String("codesearch", "", "Path to the `codesearch` binary")
	flagFetchReindex = flag.String("fetch-reindex", "", "Path to the `livegrep-fetch-reindex` binary")
	flagApiBaseUrl   = flag.String("api-base-url", "https://api.github.com/", "Github API base url")
	flagGitlabToken  = flag.String("gitlab-token", os.Getenv("GITLAB_TOKEN"), "Gitlab API token")
	flagRepoDir      = flag.String("dir", "repos", "Directory to store repos")
	flagIgnorelist   = flag.String("ignorelist", "", "File containing a list of repositories to ignore when indexing")
	flagIndexPath    = dynamicDefault{
		display: "${dir}/livegrep.idx",
		fn:      func() string { return path.Join(*flagRepoDir, "livegrep.idx") },
	}
	flagName        = flag.String("name", "livegrep index", "The name to be stored in the index file")
	flagForks       = flag.Bool("forks", true, "whether to index repositories that are github forks, and not original repos")
	flagDepth       = flag.Int("depth", 0, "clone repository with specify --depth=N depth.")
	flagSkipMissing = flag.Bool("skip-missing", false, "skip repositories where the specified revision is missing")
	flagRepos       = stringList{}
	flagOrgs        = stringList{}
	flagUsers       = stringList{}
)

func init() {
	flag.Var(&flagIndexPath, "out", "Path to write the index")
	flag.Var(&flagRepos, "repo", "Specify a repo to index (may be passed multiple times)")
	flag.Var(&flagOrgs, "org", "Specify a github organization to index (may be passed multiple times)")
	flag.Var(&flagUsers, "user", "Specify a github user to index (may be passed multiple times)")
}

func main() {
	flags.Parse()
	log.SetFlags(0)

	if flagRepos.strings == nil &&
		flagGroups.strings == nil &&
		flagUsers.strings == nil {
		log.Fatal("You must specify at least one repo or group to index")
	}
}
