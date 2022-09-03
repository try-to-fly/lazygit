package git_commands

type WorktreeCommands struct {
	*GitCommon
}

func NewWorktreeCommands(gitCommon *GitCommon) *WorktreeCommands {
	return &WorktreeCommands{
		GitCommon: gitCommon,
	}
}

func (self *WorktreeCommands) New(worktreePath string) error {
	cmdArgs := NewGitCmd("worktree").Arg("add", worktreePath).ToArgv()

	return self.cmd.New(cmdArgs).Run()
}

func (self *WorktreeCommands) Delete(worktreePath string, force bool) error {
	cmdArgs := NewGitCmd("worktree").Arg("remove").ArgIf(force, "-f").Arg(worktreePath).ToArgv()

	return self.cmd.New(cmdArgs).Run()
}
