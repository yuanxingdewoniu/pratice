#/bin/bash
sudo rm /var/log/audit/* -rf 
sudo touch /var/log/audit/audit.log
systemctl restart auditd.service
