package push

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func writeTestFile(t *testing.T, folder, file, content string) {
	p := filepath.Join(folder, file)
	d := filepath.Dir(p)
	assert.NoError(t, os.MkdirAll(d, 0755))
	assert.NoError(t, ioutil.WriteFile(p, []byte(content), 0755))
}

func runGit(t *testing.T, args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	assert.NoError(t, cmd.Run())
}

func runCommit(t *testing.T, workdir, message string) {
	runGit(t, "-C", workdir, "commit", "--allow-empty", "--author", "ActionsGo test <actions-go@users.noreply.github.com>", "-m", message)
}

func status(t *testing.T, workdir string) string {
	b := bytes.NewBuffer(nil)
	cmd := exec.Command("git", "-C", workdir, "status", "-s")
	cmd.Stdout = b
	cmd.Stderr = os.Stderr
	assert.NoError(t, cmd.Run())
	return b.String()
}

func TestCommit(t *testing.T) {

	testID := uuid.New().String()

	workdir, err := ioutil.TempDir("", "test-workdir-"+testID)
	assert.NoError(t, err)
	defer os.RemoveAll(workdir)
	assert.NoError(t, os.Chdir(workdir))
	runGit(t, "init", workdir)
	runGit(t, "-C", workdir, "config", "user.email", "actions-go@users.noreply.github.com")
	runGit(t, "-C", workdir, "config", "user.name", "ActionsGo test")
	runCommit(t, workdir, "first commit")
	writeTestFile(t, workdir, "README.md", "version 1")
	runGit(t, "-C", workdir, "add", "README.md")

	runGit(t, "-C", workdir, "status", "-s", "-uno")
	assert.Equal(t, "A  README.md\n", status(t, workdir))

	writeTestFile(t, workdir, "README.md", "version 2")

	assert.Equal(t, "AM README.md\n", status(t, workdir))

	writeTestFile(t, workdir, "untracked.md", "version 1")

	repo, err := git.PlainOpen(workdir)
	assert.NoError(t, err)
	assert.NoError(t, commit(repo, "some-other-file.go", "test commit message", "FALSE"))
	assert.Equal(t, " M README.md\n?? untracked.md\n", status(t, workdir))

	assert.NoError(t, commit(repo, "some-other-file.go", "test commit message", "FALSE"))
	assert.Equal(t, " M README.md\n?? untracked.md\n", status(t, workdir))

	assert.Error(t, commit(repo, "some-other-file.go", "test commit message", "TRUE"))

	assert.NoError(t, commit(repo, "some-other-file.go,*d.md", "test commit message", "TRUE"))
	assert.Equal(t, " M README.md\n", status(t, workdir))

	writeTestFile(t, workdir, "other-untracked.md", "version 1")
	assert.NoError(t, commit(repo, "", "test commit message", "TRUE"))
	assert.Equal(t, "?? other-untracked.md\n", status(t, workdir))
}

func TestPush(t *testing.T) {
	testID := uuid.New().String()

	remote, err := ioutil.TempDir("", "test-remote-"+testID)
	assert.NoError(t, err)
	defer os.RemoveAll(remote)
	runGit(t, "init", "--bare", remote)

	workdir, err := ioutil.TempDir("", "test-workdir-"+testID)
	assert.NoError(t, err)
	defer os.RemoveAll(workdir)
	assert.NoError(t, os.Chdir(workdir))
	runGit(t, "init", workdir)
	runGit(t, "-C", workdir, "config", "user.email", "actions-go@users.noreply.github.com")
	runGit(t, "-C", workdir, "config", "user.name", "ActionsGo test")
	runCommit(t, workdir, "first commit")
	runGit(t, "-C", workdir, "remote", "add", "some-remote", remote)
	runGit(t, "-C", workdir, "push", "-u", "some-remote", "master")
	runCommit(t, workdir, "first commit")

	repo, err := git.PlainOpen(workdir)
	assert.NoError(t, err)

	assert.Error(t, push(repo, "non-existing-remote", plumbing.ReferenceName("refs/heads/master"), plumbing.ReferenceName("refs/heads/master"), "FALSE"))
	assert.NoError(t, push(repo, "some-remote", plumbing.ReferenceName("refs/heads/master"), plumbing.ReferenceName("refs/heads/master"), "FALSE"))
	assert.NoError(t, push(repo, "some-remote", plumbing.ReferenceName("refs/heads/master"), plumbing.ReferenceName("refs/heads/master"), "FALSE"))
	assert.Error(t, push(repo, "some-remote", plumbing.ReferenceName("refs/heads/master"), plumbing.ReferenceName("refs/heads/master"), "TRUE"))

	runGit(t, "-C", workdir, "remote", "add", "origin", remote)
	assert.NoError(t, Push(workdir))
}
