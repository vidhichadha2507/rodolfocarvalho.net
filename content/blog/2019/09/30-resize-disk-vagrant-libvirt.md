---
title: "Resize Disk of a Vagrant/Libvirt VM"
slug: "resize-disk-vagrant-libvirt"
description: >-
  This is a reference for how to increase the disk size of an existing virtual
  machine.
categories: []
tags: []
series: []
date: 2019-09-30T15:06:10+02:00
---

Today I had to resize a disk on a Vagrant VM that uses the libvirt provider.
I was surprised how easy it was, and that no data has been lost.

As a reference for my future self and others, here's what I did.

My VM uses the `fedora/30-cloud-base` base box.

1. Stop the VM:
    ```
    vagrant halt
    ```

2. Resize the disk image file:
    ```
    qemu-img resize ~/.local/share/libvirt/images/vagrant_default.img +100G
    ```

    The image name depends on the VM. In the example above, we're adding 100 GB to
    the existing size.

3. Boot the VM and update the partition table and filesystem:
    ```
    vagrant up && vagrant ssh
    ```

    ```
    echo ", +" | sudo sfdisk -N 1 /dev/vda --no-reread
    sudo partprobe
    sudo resize2fs /dev/vda1
    ```

    The command above will be different depending on the existing partition
    scheme. In this example, we increase the first (and only) partition size to
    the maximum available size, and then resize the filesystem to match the new
    size.

---

My steps were adapted from the post [Increasing a libvirt/KVM virtual machine
disk capacity](https://nullr0ute.com/2018/08/increasing-a-libvirt-kvm-virtual-machine-disk-capacity/).
