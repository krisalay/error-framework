# Contributing Guide

## Setup

```bash
git clone https://github.com/yourorg/errorframework
cd errorframework
go mod tidy


Commit using semantic commit messages

These trigger automatic version bump:

feat: add mysql adapter     → MINOR bump
fix: resolve panic issue    → PATCH bump
BREAKING CHANGE: redesign API → MAJOR bump