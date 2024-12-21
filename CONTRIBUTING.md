# Contributing to Contentstack Bridge

[![CI Status](https://github.com/Dobefu/csb/actions/workflows/ci.yml/badge.svg)](https://github.com/Dobefu/csb/actions/workflows/ci.yml)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_contentstack-bridge&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=Dobefu_contentstack-bridge)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=Dobefu_contentstack-bridge&metric=coverage)](https://sonarcloud.io/summary/new_code?id=Dobefu_contentstack-bridge)
[![Go Report Card](https://goreportcard.com/badge/github.com/Dobefu/csb)](https://goreportcard.com/report/github.com/Dobefu/csb)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Thank you for your interest in contributing to Contentstack Bridge! This document outlines the process for contributing to our CLI application.

## Development Environment

- Ensure you have [Go](https://golang.org/dl/) installed (version 1.23 or higher).

## Building and Running

- Navigate to the project directory:
  ```bash
  cd csb
  ```
- Build the project:
  ```bash
  go build -ldflags="-s -w"
  ```
- Run the CLI application:
  ```bash
  ./csb
  ```

## Coding Standards

- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use [gofmt](https://golang.org/cmd/gofmt/) to format your code
- Write meaningful variable and function names
- Maintain a consistent code style throughout the project

## Testing

- Write unit tests for new features or bug fixes using the built-in `testing` package
- Assertions can be made with the `github.com/stretchr/testify/assert` package
- Total test coverage must at least be 80%
- Ensure all existing tests pass before submitting a pull request:
  ```bash
  ./scripts/test.sh
  ```

## Submitting Changes

- Fork the repository in GitHub
- Clone the fork locally:
  ```bash
  git clone git@github.com:[USERNAME]/csb.git
  ```
- Create a new branch for your feature or bug fix:
  ```bash
  git checkout -b feature/your-feature-name
  ```
- Commit your changes with a clear and descriptive commit message:
  ```bash
  git commit -m "feat(code): your feature description"
  ```
- Push your changes to your fork:
  ```bash
  git push origin feature/your-feature-name
  ```
- Create a pull request on GitHub from your fork to the main repository

## Pull Request Guidelines

- Provide a clear description of the changes in your pull request
- Include any relevant issue numbers in the PR description
- Ensure your PR passes all CI checks
- Ideally link the PR to an open issue
- Be prepared to make changes if requested by maintainers

## Reporting Issues

- Use the GitHub issue tracker to report bugs or suggest enhancements
- Provide as much detail as possible, including steps to reproduce for bugs

Thank you for contributing to Contentstack Bridge!
