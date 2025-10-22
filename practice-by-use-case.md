# **Use Case: Student Portfolio Website Project**

## **Part 1: Starting the Project (Local Setup)**

**Task 1: Initialize a New Repository**
- Create a new folder called `portfolio-website`
- Use `git init` to initialize it as a Git repository
- Create an `index.html` file with basic HTML structure

**Task 2: First Commit**
- Use `git status` to see untracked files
- Use `git add index.html` to stage the file
- Use `git commit -m "Initial commit: Add index.html"` to commit

**Task 3: Oops, Wrong File!**
- Create a file called `temp.txt` (a file you don't want to commit)
- Use `git add temp.txt` to stage it
- Use `git restore --staged temp.txt` to unstage it
- Verify with `git status`

---

## **Part 2: Working with Remote Repository**

**Task 4: Connect to Remote**
- Create a repository on GitHub (or GitLab/Bitbucket)
- Use `git push` to push your local commits to the remote repository
  - Add the remote first: `git remote add origin <url>`
  - Then: `git push -u origin main`

**Task 5: Clone a Partner's Repository**
- Pair up with another student
- Use `git clone <partner's-repo-url>` to clone your partner's repository
- This simulates joining an existing project

---

## **Part 3: Branching and Features**

**Task 6: Create a Feature Branch**
- Use `git branch about-page` to create a new branch
- Use `git checkout about-page` to switch to it
- Verify you're on the new branch with `git status`

**Task 7: Develop the Feature**
- Create `about.html` with some content
- Use `git add about.html`
- Use `git commit -m "Add about page"`
- Use `git push origin about-page` to push the branch to remote

**Task 8: Switch Between Branches**
- Use `git checkout main` to go back to main branch
- Notice that `about.html` disappears
- Use `git checkout about-page` to go back
- Notice that `about.html` reappears

---

## **Part 4: Merging and Collaboration**

**Task 9: Merge the Feature**
- Switch to main branch: `git checkout main`
- Use `git merge about-page` to merge the feature into main
- Use `git push` to push the merged changes

**Task 10: Pull Partner's Changes**
- Your partner pushes a new file (e.g., `contact.html`) to their repository
- Use `git pull` to download and merge your partner's changes
- Verify the new file appears with `git status`

---

## **Part 5: Handling Conflicts (Advanced)**

**Task 11: Create a Conflict**
- Both partners edit the same line in `index.html`
- Both commit and try to push
- The second person will need to `git pull` first
- Git will show a merge conflict
- Manually resolve the conflict
- Use `git add` to mark as resolved
- Use `git commit` to complete the merge
- Use `git push` to share the resolution
