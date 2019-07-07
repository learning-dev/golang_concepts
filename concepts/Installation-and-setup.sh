# Download Install and go lang onto the machine 

https://golang.org/dl/

# Download the mac os variant 
Open and follow the instructions
# choose install for all users and DON'T change the install location
#enter your password


# check the go version 
go version

#show the go command menu
go

#install sublime text, if not installed 
# install package control in sublime, 
# just go to Tools > install package control
# Open Command pallete - search golang build


#setup the gopath onto mac - this is required for sublime text in our case 
# It is required in general too 
echo 'export GOPATH=$HOME/go' >> $HOME/.profile

$ source $HOME/.profile
$ go env | grep GOPATH


# update the same settings in sublime text 
sublime-text > preferences > package settings > Golang config
#paste in the following 

{
    "PATH": "/Users/loonycorn/go/bin",
    "GOPATH": "/Users/loonycorn/go"
}

#save the file,
# now you can run go files using cmd +b
=================
** Getting the formatter and syntax recommender ** 

# get into sublime text dir 
cd ~/Library/Application\ Support/Sublime\ Text\ 3/Packages/

# clone the repo 
git clone https://github.com/DisposaBoy/GoSublime.git

open sublime-text

you will see a change log 

cmd + . and cmd + x

then you see a file called  margo.go

close the file and open a new file.go and you can see the changes



