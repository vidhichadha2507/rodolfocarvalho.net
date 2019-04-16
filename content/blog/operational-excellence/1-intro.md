---
title: "Operational Excellence"
slug: "operational-excellence"
description: >-
  Seeking for Operational Excellence means constantly improving your people and
  your processes. This is the introduction to a series of posts where I discuss
  Operational Excellence in the context of building and running a Software as a
  Service product.
categories: []
tags: []
series: ["operational-excellence"]
date: 2019-04-22T10:00:00+01:00
blackfriday:
  fractions: false
---

I had the opportunity and privilege to be part of [AWS Elemental
MediaStore](https://aws.amazon.com/mediastore/), a top-notch
[AWS](https://aws.amazon.com/) team building and operating a service to deliver
high performance storage for video workflows.

With the objective of sharing for greater good, and perhaps for my own future
reference, here begins a series of posts diving into some of the interrelated
topics on achieving Operational Excellence.

These notes are insights into development and operational best practices that
are critical to large scale software projects, specially those in which
customers depend on your availability and performance 24/7.

*Disclaimer: the contents of this and related posts are personal notes written
by me and not endorsed by AWS or any other company.*

## Defining Operational Excellence

[AWS Well-Architected](https://aws.amazon.com/architecture/well-architected/) is
one official venue where AWS shares best practice knowledge with the general
public, allowing other companies to follow lessons learned through decades of
solving hard problems.

Here is what it says
about [Operational
Excellence](https://wa.aws.amazon.com/wat.pillar.operationalexcellence.en.html):

> The Operational Excellence pillar includes the ability to run and monitor
> systems to deliver business value and to continually improve supporting
> processes and procedures.

My take is that Operational Excellence is the air Amazon <abbr
title="Software Development Engineer">SDE</abbr>s breathe day after day. It
consists of practices and principles discovered and developed while building and
operating one of Earth's biggest websites and marketplaces,
[Amazon.com](https://smile.amazon.com/), as well as providing the world's
largest portfolio of infrastructure and software as a service offerings,
[AWS](https://aws.amazon.com/).

## Let's begin

The ideas that will be discussed in this series apply directly to <abbr
title="Software as a Service">SaaS</abbr> products, and they may be as well
applicable to other contexts.

This series is by no means all-encompassing. It is, however, a personally
organized collection that may help you improve your existing processes and
procedures. Without further ado, let's introduce the opening act: [Automated
Tests as the Basis for Operational Excellence]({{< relref "2-testing.md" >}}).
