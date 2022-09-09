#!/bin/sh

# check the user is root
if [ $(id -u) != "0" ]; then
    echo "Error: You must be root to run this script, please use root to configure redis"
    exit 1
fi

# check of read config file exists
if [ ! -f /etc/redis/redis.conf ]; then
    echo "Redis config file not found at /etc/redis/redis.conf. Exiting..."
    exit 1
fi

# edit the config file to set the maxmemory value and policy
sed -i "s/^maxmemory .*/maxmemory 100000000/" /etc/redis/redis.conf
sed -i "s/^maxmemory-policy .*/maxmemory-policy allkeys-lru/" /etc/redis/redis.conf

# restart redis server
/etc/init.d/redis-server restart