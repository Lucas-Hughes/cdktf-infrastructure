## Running infrastructure code in CDKTF against checkov utilizing .gitlab-ci.yml.

This repo is designed to show the somewhat janky workaround I found when trying to use checkov against cdktf. When utilizing data sources in cdktf, checkov cannot parse that file properly. This Go file removes relevant data calls within the cdktf.out directory (why it's not in .gitignore).

Since data calls do not effect the result of security scans while in the cdktf.out, it is safe to remove them. 
