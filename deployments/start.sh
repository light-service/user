#!/bin/sh

printenv | sed 's/^\(.*\)$/export \1/g' >> /etc/profile.d/cron_env.sh
chmod +x /etc/profile.d/cron_env.sh
# 每天执行
(crontab -l ; echo "0 0 * * * . /etc/profile.d/cron_env.sh; sh /app/clean_log.sh") | crontab -

/init.sh
a=$1
nohup fluentd&
GOTRACEBACK=crash /app/main ${a}
