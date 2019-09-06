---
title: "Git: Reviewing Changes Before Commit"
slug: "git-review-changes-before-commit"
description: >-
  Get familiar with this small set of Git commands as they will help you review
  your changes before committing them.
categories: []
tags: []
series: []
date: 2019-09-06T09:42:43+02:00
---

Crafting good commits is perhaps almost an art, and along my career I've come to
realize that not all developers care equally about making commits that have a
clear and targeted goal and an accompanying description that will help future
project maintainers understand the changes that were made and why they were
made.

I believe this is not always due to lack of care, but might be due to lack of
familiarity with the tools, in this case namely Git.

Git is fast and powerful, but not always obvious and user-friendly, though it
has been improving over the years. Still, there is a basic set of commands that
once you master will allow you to improve your handling of commits.

When preparing to commit changes, you can review what goes into a commit with
these commands:

```
git diff           # shows the difference between your working area and
                   # the staging area¹
```

```
git add -p         # interactively add changes to the staging area hunk by hunk²
```

```
git diff --cached  # shows the difference between the staging area and the last
                   # commit, this is what will be committed by 'git commit'
```

```
git commit -v      # commit changes and show the diff of what is being committed
                   # in your default text editor
```

If those are not yet second nature to you, start practicing them!

You can use those and other Git commands to craft wonderful commits.
Keep learning!

---

¹ More about the different areas where your code lives in Git in [the Git Book:
The Three States][git-book-states].  
² In this context, a "hunk" is a (relatively) small piece of code change.

[git-book-states]: https://git-scm.com/book/en/v1/Getting-Started-Git-Basics#The-Three-States
