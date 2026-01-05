# 📚 Golang & Vue - Books Manager

A modern, full-stack book management application built with **Vue 3** and **Go**. This project demonstrates best practices for building web applications with RESTful APIs and reactive frontend interfaces.

## ✨ Features

- **Vue 3 Composition API** with TypeScript for type-safe frontend development
- **Go backend** with RESTful architecture
- **Hot-reload development** environment for rapid iteration
- **Docker support** for consistent deployment and development environments

> 💡 This project is based on the [Working with Vue 3 and Go](https://www.udemy.com/course/working-with-vue-3-and-go/) course by Travis Sawler.

## 🛠️ Prerequisites

- [Go Programming Language](https://golang.org/dl/)
- [Node.js (LTS)](https://nodejs.org/en/download/) and [pnpm](https://pnpm.io/installation)
- [Docker](https://www.docker.com/get-started)

## 💻 Recommended Setup

### Recommended IDE Setup

- [VS Code](https://code.visualstudio.com/) + [Vue (Official)](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).
- [Neovim](https://neovim.io/) + [nvim-lspconfig](https://github.com/neovim/nvim-lspconfig) to install the `vue-language-server`.

#### Type Support for `.vue` Imports in TS

TypeScript cannot handle type information for `.vue` imports by default, so we replace the `tsc` CLI with `vue-tsc` for type checking. In editors, we need [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) to make the TypeScript language service aware of `.vue` types.

### Recommended Browser Setup

- Chromium-based browsers (Chrome, Edge, Brave, etc.):
  - [Vue.js devtools](https://chromewebstore.google.com/detail/vuejs-devtools/nhdogjmejiglipccpnnnanhbledajbpd)
  - [Turn on Custom Object Formatter in Chrome DevTools](http://bit.ly/object-formatters)
- Firefox:
  - [Vue.js devtools](https://addons.mozilla.org/en-US/firefox/addon/vue-js-devtools/)
  - [Turn on Custom Object Formatter in Firefox DevTools](https://fxdx.dev/firefox-devtools-custom-object-formatters/)

## ⚙️ Customize configuration

See [Vite Configuration Reference](https://vite.dev/config/).

## 🚀 Project Setup

```sh
pnpm install
```

### Compile and Hot-Reload for Development

```sh
pnpm dev
```

### Type-Check, Compile and Minify for Production

```sh
pnpm build
```

### Lint with [ESLint](https://eslint.org/)

```sh
pnpm lint
```
