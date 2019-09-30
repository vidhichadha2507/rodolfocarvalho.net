---
title: "Floating Point Number Comparison"
slug: "floating-point-number-comparison"
description: >-
  In short, floating point math is hard. Peer reviewing code that involved the
  comparison of floats made me look at the interesting things that we need to
  keep in mind or, rather surprisingly, maybe not.
categories: []
tags: []
series: []
date: 2019-09-10T11:16:25+02:00
---

This post summarizes my thoughts after doing a code review which involved
testing outputs containing several small floating point numbers.

The code was written in Python 2/3, and used floats stored in both regular
builtin collection types as well as NumPy arrays.

In the course of the code review, I got interested on how we were comparing
results to expected values. Equality comparison needs to take into account the
inherent errors in floating point arithmetic and numerical computations.
Therefore, we do not compare for direct equality, but instead for some
definition of "almost equal".

The interesting part was determining what definition of "almost equal" to use. A
good reading on the topic was [Comparing Floating Point Numbers, 2012
Edition](https://randomascii.wordpress.com/2012/02/25/comparing-floating-point-numbers-2012-edition/),
by Bruce Dawson.




In  comments on numpy issue tracker: https://github.com/numpy/numpy/issues/10161#issuecomment-350021545
