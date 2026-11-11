#!/usr/bin/env bash
set -e
grep -q "max_idle_conn_secs = 20" conf/server.toml
