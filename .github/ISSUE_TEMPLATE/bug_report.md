---
name: 🐞 Bug report
about: Create a report to help us improve
title: '[Bug Report] <title>'
labels: 'bug'
assignees: ''

---

<!--
Please, check the FAQ and Known Problems pages before creating the bug report:

https://github.com/evilsocket/opensnitch/wiki/FAQs

GUI related issues:
https://github.com/evilsocket/opensnitch/wiki/GUI-known-problems

Daemon related issues:
 - Run `opensnitchd -check-requirements` to see if your kernel is compatible.
 - https://github.com/evilsocket/opensnitch/wiki/daemon-known-problems
-->

### Describe the bug:
<!-- A clear and concise description of what the bug is. -->

Include the following information:
 - OpenSnitch version: 
 - OS: [e.g. Debian GNU/Linux, ArchLinux, Slackware, ...]
 - OS version: [e.g. Buster, 10.3, 20.04]
 - Window Manager: [e.g. GNOME Shell, KDE, enlightenment, i3wm, ...]
 - Kernel version: [`uname -a`]

### To Reproduce:
<!-- Describe in detail as much as you can what happened. -->

Steps to reproduce the behavior:
1. Go to '...'
2. Click on '....'
3. Scroll down to '....'
4. See error

### Post error logs:
<!--
If it's a crash of the GUI: 
 - Launch it from a terminal and reproduce the issue.
 - Post the errors logged to the terminal.

If the daemon doesn't start or doesn't intercept connections:
 - Run `opensnitchd -check-requirements` to see if your kernel is compatible.
 - Post last 15 lines of the log file `/var/log/opensnitchd.log`
 - Or launch it from a terminal as root (`# /usr/bin/opensnitchd -rules-path /etc/opensnitchd/rules`) and post the errors logged to the terminal.

If the deb or rpm packages fail to install:
 - Install them from a terminal (`$ sudo dpkg -i opensnitch*` / `$ sudo yum install opensnitch*`), and post the errors logged to stdout.
-->

### Expected behavior (optional):
<!-- A clear and concise description of what you expected to happen. -->

### Screenshots:
<!-- If applicable, add screenshots or videos to help explain your problem. It may help to understand the issue much better. -->

### Additional context:
<!-- Add any other context about the problem here. -->
