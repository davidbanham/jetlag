## What?

This little utility is intended to make it easier to switch between VPNs on a router. I'm using it to target [one of these](https://samgentle.com/posts/2015-05-08-the-amazing-chinarouter).

It assumes:

You are using openvpn
You have passwordless SSH from the box running this server to your router
There are scripts in `/root/` on the router that conform to the naming convention {[country}}-vpn.sh eg: `AU-vpn.sh`

## How?

Download the binary you need for your OS, execute it. The environment variables you can set are:

`ROUTER` required. The hostname of your router as seen from your server.
`PORT` default 3000
`COUNTRIES` default AU,US a comma delimeted list of country prefixes as per your scripts in `/root/`

## Why?

I want it to be simple to change my network's egress point from any device in the house.
