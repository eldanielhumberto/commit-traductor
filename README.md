# ct - Commit Translator

[![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

**ct** (commit translator) is a command-line tool that helps you write clean, consistent, and readable git commit messages across different languages using Google's Gemini AI.

## ✨ Features

- 🌍 **Multilingual Support**: Translate commit messages to any language
- 🤖 **AI-Powered**: Uses Google Gemini 2.5 Flash for accurate translations
- ⚡ **Fast & Efficient**: Built with Go for optimal performance
- 🔒 **Secure**: API keys managed through environment variables
- 🎯 **Simple CLI**: Easy-to-use command-line interface powered by Cobra

## 📋 Prerequisites

- Go 1.25.5 or higher
- Google Gemini API key ([Get one here](https://ai.google.dev/))

## 🚀 Installation

```bash
go install github.com/eldanielhumberto/ct@latest
```

## ⚙️ Configuration

1. Create a `.env` file in the project root:

```bash
cp .env.example .env
```

2. Add your Gemini API key to the `.env` file:

```env
GEMINI_API_KEY=your_api_key_here
```

> **Note**: You can also set the environment variable directly in your shell:
> ```bash
> export GEMINI_API_KEY=your_api_key_here
> ```

## 📖 Usage

### Basic Usage

Translate a commit message to a specific language:

```bash
ct commit "Initial commit" --lang es
```

**Output:**
```
Original message: Initial commit
Translated message (es): Commit inicial
```

### Available Options

```bash
ct commit [message] [flags]
```

**Flags:**
- `-l, --lang string`: Target language for translation (default: "en")
- `-h, --help`: Help for commit command

### Examples

```bash
# Translate to Spanish
ct commit "Add user authentication" --lang es

# Translate to French
ct commit "Fix database connection bug" --lang fr

# Translate to Japanese
ct commit "Update documentation" --lang ja

# Translate to German
ct commit "Refactor payment module" --lang de
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

##  Contact

Daniel Humberto - [@eldanielhumberto](https://github.com/eldanielhumberto)

Project Link: [https://github.com/eldanielhumberto/ct](https://github.com/eldanielhumberto/ct)

---

Made with ❤️ by Daniel Humberto
