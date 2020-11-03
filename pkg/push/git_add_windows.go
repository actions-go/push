package push

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-git/go-git/v5"
)

// work around https://github.com/go-git/go-git/issues/55
// aka https://github.com/src-d/go-git/issues/1155
func gitadd(wt *git.Worktree, path string) error {
	c := exec.Command("git", "-C", wt.Filesystem.Root(), "add", path)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

func gitcommit(wt *git.Worktree, commitMessage string, opts *git.CommitOptions) error {
	if opts == nil {
		return fmt.Errorf("options must be provided to use this work-around")
	}
	if opts.All {
		return fmt.Errorf("commit all is not supported by this work-around")
	}
	if opts.Committer != nil {
		return fmt.Errorf("committer must not be provided to use this work-around")
	}
	if opts.Parents != nil {
		return fmt.Errorf("parent commits must not be provided to use this work-around")
	}
	if opts.SignKey != nil {
		return fmt.Errorf("signing key must not be provided to use this work-around")
	}
	if opts.Author == nil {
		return fmt.Errorf("author must be provided to use this work-around")
	}
	fmt.Println("git", "-C", wt.Filesystem.Root(), "commit", fmt.Sprintf("--author=%s <%s>", opts.Author.Name, opts.Author.Email), "--date", opts.Author.When.Format(time.RFC3339), "-m", commitMessage)
	os.Setenv("GIT_COMMITTER_NAME", opts.Author.Name)
	os.Setenv("GIT_COMMITTER_EMAIL", opts.Author.Email)
	os.Setenv("GIT_COMMITTER_DATE", opts.Author.When.Format(time.RFC3339))
	c := exec.Command("git", "-C", wt.Filesystem.Root(), "commit", fmt.Sprintf("--author=%s <%s>", opts.Author.Name, opts.Author.Email), "--date", opts.Author.When.Format(time.RFC3339), "-m", commitMessage)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
