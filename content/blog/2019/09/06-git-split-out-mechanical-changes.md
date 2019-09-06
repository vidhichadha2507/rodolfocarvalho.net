---
title: "Git: Split Out Mechanical Changes"
slug: "git-split-out-mechanical-changes"
description: >-
  Clean up your commits and make them easier to review by separating mechanical changes from manually crafted changes.
categories: []
tags: []
series: []
date: 2019-09-06T10:45:15+02:00
---

Sometimes a Git commit gets "contaminated" with mechanical changes applied by
a code editor or some other tool, and that often makes reviewing the code
unnecessarily more onerous as the reviewers need to split in their heads what
was an automatic operation and what are the main changes meant to be in a
commit.

With a few Git commands you can quickly separate those mechanical changes into a
separate commit that can be verified independently. Here is how:

1. Look at the previous commit and see that it has both mechanical changes and
manual changes:

    ```
    git show
    ```

2. (optional) Create a branch named "backup" (or any other name) to help verify
your changes later:

    ```
    git branch backup
    ```

3. Temporarily undo the previous commit:

    ```
    git revert HEAD --no-edit
    ```

    To get a list of the files that changed in your original commit (and
    consequently also in the commit that reverts the changes):

    ```
    git diff-tree --no-commit-id --name-only -r HEAD
    ```

    We'll use that to feed file paths to a tool. The tool may be a code
    formatter, or whatever else.

    As an example, consider a Python code base using `yapf` for formatting and
    `isort` for sorting imports. Let's continue.

4. Run tool to make mechanical changes to files:

    ```
    isort -y $(git diff-tree --no-commit-id --name-only -r HEAD)
    yapf -i $(git diff-tree --no-commit-id --name-only -r HEAD)
    ```

    Note that depending on the tool and on the list of files, you may need to
    accommodate for situations like a new file was added in your original commit
    and thus it doesn't exist once the original commit is reverted.

    There are several ways to handle it, one simple way is to use `grep` or
    `grep -v` to filter the file names.

    ```
    yapf -i $(git diff-tree --no-commit-id --name-only -r HEAD | grep only_this_file.py)
    yapf -i $(git diff-tree --no-commit-id --name-only -r HEAD | grep -v not_this_file.py)
    ```

5. Time to [review your changes]({{< relref
"06-git-review-changes-before-commit" >}}) and make a new commit.


6. Now comes the final step: reorder the commits and remove the revert of the
original change.

    This step is accomplished using `git rebase -i` (interactive rebase).

    ```
    git rebase -i @{upstream} -X theirs
    ```

    This command will open your default editor with a list of commits. There
    should be 3 commits, first your original commit, then the revert, and
    finally the commit with the mechanical changes.

    ```
    1 pick 52fe10a Original commit
    2 pick 324fa22 Revert "Original commit"
    3 pick 4f024b0 Run yapf and isort
    ```

    Use your editor to move the mechanical commit to the top and delete the line
    with the revert commit.

    ```
    1 pick 4f024b0 Run yapf and isort
    2 pick 52fe10a Original commit
    ```

    Save and close the editor.

    You may wonder what the `-X theirs` is for. Sure, what is that for?  
    While applying the original commit on top of the commit with mechanical
    changes, there might be merge conflicts. In this particular case, it is
    clear that the final state should match exactly what was in the original
    commit. The `-X theirs` tell Git to prefer the original commit whenever a conflict would appear.

    The rebase operation should succeed, no manual edits required.

**Bonus**

1. Before submitting your two commits, check that they match the original single
commit:

    ```
    git diff backup..HEAD
    ```

    There should be no output.

2. Review your last two commits:

    ```
    git log -p -2
    ```

3. Remove your backup branch:

    ```
    git branch -D backup
    ```
