#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/bytedance-mall"
exec "$CURDIR/bin/bytedance-mall"
