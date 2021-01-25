Generate ssh key and add git account(one of them github, bitbucket, gitlab)


mkdir ~/.ssh
cd ~/.ssh
ssh-keygen -t rsa -b 4096

eval `ssh-agent -s`

cat ~/.ssh/id_rsa.pub

