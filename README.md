# makex

An opinionated tool for streamlined deployment of containers with secrets

Note: Unfinished, as I found already available tools, primarily [task](https://github.com/go-task/task) and [sopstool](https://github.com/Ibotta/sopstool).

## Log

- [1.3.2024] Initially wanted to use [urfave/cli/v3](https://github.com/urfave/cli), but could not get it load config from files.
  Currently using [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper).
  Some alternatives:

  - [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea)
  - [charmbracelet/bubbles](https://github.com/charmbracelet/bubbles)
  - [charmbracelet/lipgloss](https://github.com/charmbracelet/lipgloss)
