# commit-traductor

A tool to translate commit messages to different languages.
At the moment, the library only translates from Spanish to English.

## Installation

You can install `commit-traductor` using bun, npm, yarn, or pnpm (recommended: bun):

```bash
bun add commit-traductor
```

## Configuration

Need husky to run the pre-commit hook.

```bash
bun add husky -D # or npm, yarn, pnpm
```

Then, you can set up husky in your project:

```bash
# initialize husky
npx husky init

# delete the default pre-commit hook
rm .husky/pre-commit
```

Add the pre-commit hook:

```bash
echo 'bunx commit-traductor "$1"' > .husky/commit-msg
```

Now, every time you make a commit, the commit message will be translated to English.

## Contribution

Feel free to open issues or submit pull requests for improvements or bug fixes.
Contributions are welcome!
