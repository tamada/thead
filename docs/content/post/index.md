---
title: "Index"
date: 2021-07-12
tags: []
draft: false
---

[![build](https://github.com/akanatr/thead/actions/workflows/build.yml/badge.svg)](https://github.com/akanatr/thead/actions/workflows/build.yml)
[![Coverage Status](https://coveralls.io/repos/github/akanatr/thead/badge.svg?branch=main)](https://coveralls.io/github/akanatr/thead?branch=main)
[![codebeat badge](https://codebeat.co/badges/dd29b1a2-7c3c-4c2a-b08e-87e152043f57)](https://codebeat.co/projects/github-com-akanatr-thead-main)
[![Go Report Card](https://goreportcard.com/badge/github.com/akanatr/thead)](https://goreportcard.com/report/github.com/akanatr/thead)
[![GitHub license](https://img.shields.io/github/license/akanatr/thead)](https://github.com/akanatr/thead/blob/main/LICENSE)
[![Dockerfile](https://img.shields.io/badge/Docker-ghcr.io%2Fakanatr%2Fthead%3A1.0.0-green?logo=docker)](https://hub.docker.com/repository/docker/akanatr/thead_image)

# thead
headコマンドの機能拡張 

<img src="./theadicon.svg" width="100">

## Description
既存のheadコマンドは複数のファイル指定し，各ファイルごとの先頭10行分の内容を出力することが可能である．
しかし，特定のディレクトリ直下のファイル全てを指定する際には，正規表現を工夫する必要があり面倒である．
そこでtheadではheadコマンドの機能の一部に加え，ディレクトリを指定するだけでその直下のファイル全てを対象とする機能を追加する．
