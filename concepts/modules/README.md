# 📦 Building a Go Module

A **reusable, well-tested module** in Go for _web-related functionality_.
Designed for clarity, modularity, and ease of maintenance.

---

## 📚 What are Go Modules

Modules in Go provide a way to organize and separate code, making it easier to
maintain and understand. They allow developers to share code across different
projects or with the wider community.

## ❓ About this Project

> This project is based on the
> [Building a module in Go](https://www.udemy.com/course/building-a-module-in-go-golang/)
> course of Travis Sawler.

### 🧠 Topics & Features

| Topics Covered                            | Features of the Module                     |
| ----------------------------------------- | ------------------------------------------ |
| 🏗️ Go module for web development tasks    | 📄 JSON read/write operations              |
| ✅ Implementing tests for each feature    | ⚠️ Human-readable JSON error responses     |
| 🧰 Utilizing Go's workspace functionality | 📤 File upload handling                    |
| 📚 Stick to the standard library          | 📥 Static file downloads                   |
|                                           | 🎲 Random number generation                |
|                                           | 🌐 Posting JSON payloads to remote servers |
|                                           | 📁 Directory creation                      |
|                                           | 🔗 Slug generation from text               |

### 📁 Structure

```bash
modules/
├── app/            # Demo app for the 'core' functionality
├── app-download/   # Demo app for file download functionality
├── app-json/       # Demo app for handling json functionality
├── app-upload/     # Demo app for file upload functionality
└── toolkit/        # The web development toolkit itself
```

### 🚀 Running the Example Applications

Each `app`-folder is a Go project. To start one, run:

```bash
cd app && go run .        # "app" is the core functionality demo
cd app-upload && go run . # "app-upload" is one of the example applications
```
