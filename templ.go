package main

var network = `
config interface 'loopback'
	option ifname 'lo'
	option proto 'static'
	option ipaddr '127.0.0.1'
	option netmask '255.0.0.0'

config interface 'lan'
	option proto 'static'
	option netmask '255.255.255.0'
	option _orig_ifname 'eth0'
	option _orig_bridge 'false'
	option ipaddr '{{.ipAddr}}'
	option ipv6 '0'

config globals 'globals'

config interface 'wan'
	option ifname 'eth0'
	option proto 'dhcp'
	option mtu '1452'
	option peerdns '0'
	option ipv6 '0'
`
var secrets = `
# /etc/ipsec.secrets - strongSwan IPsec secrets file
: RSA /etc/letsencrypt/live/home1.3fire.org/privkey.pem
#: RSA serverKey.pem

{{.username}} : EAP "{{.password}}" 
`

var conf = `
# ipsec.conf - strongSwan IPsec configuration file

# basic configuration

config setup
	# plutodebug=all
	# crlcheckinterval=600
	# strictcrlpolicy=yes
	# cachecrls=yes

conn %default
    keyexchange=ikev2
    ike=aes256-sha256-modp2048!
    esp=aes256-sha256,aes256-sha1,3des-sha1! # Win 7 is aes256-sha1, iOS is aes256-sha256, OS X is 3des-shal1
    dpdaction=clear
    dpddelay=300s
    rekey=no

conn rw
    right={{.server}}
    rightid=%la.tuson.org
    rightsubnet=0.0.0.0/0
    rightauth=pubkey
    leftid={{.username}}.tuson.org
    leftsubnet={{.ipNet}}/24
    leftauth=eap
    leftupdown=/usr/lib/ipsec/_updown
    eap_identity={{.username}}
    compress=yes
    dpdaction=restart
    dpdtimeout=40s
    auto=add
`
