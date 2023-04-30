go install github.com/glebarez/cero@latest
sudo mv $(go env GOPATH)/bin/cero /usr/local/bin

go install github.com/kmskrishna/gdn@latest
sudo mv $(go env GOPATH)/bin/gdn /usr/local/bin


go install github.com/hakluke/hakip2host@latest
sudo mv $(go env GOPATH)/bin/hakip2host /usr/local/bin

go install github.com/hakluke/hakrevdns@latest
sudo mv $(go env GOPATH)/bin/hakrevdns /usr/local/bin

 go install github.com/tomnomnom/httprobe@latest
sudo mv $(go env GOPATH)/bin/httprobe /usr/local/bin
