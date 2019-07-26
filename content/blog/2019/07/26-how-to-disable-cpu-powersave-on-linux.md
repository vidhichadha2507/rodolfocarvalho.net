---
title: "How to Disable CPU Power Save on Linux"
slug: "how-to-disable-cpu-powersave-on-linux"
description: ""
categories: []
tags: []
series: []
date: 2019-07-26T16:36:29+02:00
---

This is a mental note. Ensuring all CPUs are set to provide maximum performance is useful for example when running benchmarks.

<!--more-->

## Step 1

Check current CPU frequency:

```
grep -E '^model name|^cpu MHz' /proc/cpuinfo
```

## Step 2

Check current [CPU scaling governor](https://wiki.archlinux.org/index.php/CPU_frequency_scaling) setting:

```
cat /sys/devices/system/cpu/cpu[0-9]*/cpufreq/scaling_governor
```

## Step 3

Set CPU scaling governor to `performance`:

```
echo performance | sudo tee /sys/devices/system/cpu/cpu[0-9]*/cpufreq/scaling_governor
```

Repeat step 1 and/or step 2 to check that the new setting is in effect.

When done with benchmarks or whatever else, set CPU scaling governor (back) to `powersave`:

```
echo powersave | sudo tee /sys/devices/system/cpu/cpu[0-9]*/cpufreq/scaling_governor
```

---

### Reference

[The Go Blog: Profiling Go Programs](https://blog.golang.org/profiling-go-programs)
