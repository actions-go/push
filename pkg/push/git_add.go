// +build !windows

package push

import "github.com/go-git/go-git/v5"

func gitadd(wt *git.Worktree, path string) error {
	_, err := wt.Add(path)
	return err
}

func gitcommit(wt *git.Worktree, commitMessage string, opts *git.CommitOptions) error {
	_, err := wt.Commit(commitMessage, opts)
	return err
}
