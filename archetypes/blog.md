---
{{/* $name is the file name without the "DD-" day prefix */ -}}
{{- $name := substr .Name 3 -}}
{{- $title := replace $name "-" " " | title -}}
title: "{{ $title }}"
slug: "{{ $name }}"
description: ""
categories: []
tags: []
series: []
date: {{ .Date }}{{/* dateFormat "2006-01-02" .Date */}}
---

Content for "{{ $title }}" goes here.

Start with a good summary...

<!--more-->

... and continue with the full content.

Or, the preferred, write a good summary in the front matter description.
