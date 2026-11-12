#!/usr/bin/env bash
set -e
grep -q "max_idle_conn_secs = 40" conf/server.toml
