# dnsmasq

[DNS Masquerade](http://www.thekelleys.org.uk/dnsmasq/doc.html) is a tiny DNS subsystem ideal for use with Docker for simple "Service Discovery" and dynamic host modification for local development/testing.

We use it to manage DNS within a Docker Host environment to avoid environment-related hacks within Docker Containers themselves. It is also
handy to use as a Resolver for Nginx, which does not use `/etc/hosts` so can't be configured with Docker's `--link` flag, and helps fix Nginx's infamous [DNS poisoning bug](http://comments.gmane.org/gmane.comp.web.nginx.english/38738).
