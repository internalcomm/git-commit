# git-commit 🚀



A simple, cross-platform CLI tool to standardize git commit messages by prompting for:



-  **Task ID** (numbers only)

-  **Commit message**

-  **Release type** (major / minor / patch)



The tool automatically:

- Validates input

- Ensures you are inside a git repository

- Creates a formatted git commit



---



## ✨ Features



- ✅ Works on **Linux, macOS, and Windows**

- ✅ Binary-based (no Go required for users)

- ✅ Interactive CLI

- ✅ Enforces numeric Task IDs

- ✅ Consistent commit message format

- ✅ Uses GitHub Releases for distribution

- ✅ Installable via `curl`



---



## 📦 Commit Message Format



<commit_message> - <release_type> refs #<task_id>



---



## 🛠 Installation


### Linux / macOS (Recommended)


```bash

curl  -fsSL  https://github.com/internalcomm/git-commit/releases/latest/download/install.sh | bash

```


### Windows


```bash

[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; irm "https://github.com/internalcomm/git-commit/releases/latest/download/install.ps1" |  iex

```