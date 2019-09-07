# Gogios

Simple system to check important services are on remote machines.

## Build Instructions

I use go modules to track project dependencies. They should all be included in the vendor/ folder, but if not
you can use:

```bash
go mod download
```

or

```bash
go mod vendor
```

to collect them all (in theory, go get -d ./... will do the same thing).

Once you have those, you can build the project with:

```bash
make
```

For linting, you have to install golangci-lint. This can be done with:

```bash
go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
```

To install:

```bash
make install
```

This will make folders and put files where they need to be.

## Installing from Binary Packages

### Ubuntu/Debian

Download and install the latest release deb file and then start the service. For example:

```bash
wget https://github.com/BKasin/Gogios/releases/download/VERSION/gogios-VERSION.deb
sudo dpkg -i gogios-VERSION.deb
sudo systemctl start gogios
sudo systemctl enable gogios
```

### Arch

I have made an AUR package that can be installed with something like yay using:

```bash
yay -S gogios-bin
```

And then start and enable the service.

### All OSes

You will need to install a webserver and have it point to

```bash
/opt/gingertechengine
```

as its webroot.

The file that checks are pulled from is in /etc/gingertechengine, as well as an example nginx website config file.

After installing Gogios, you can configure it at /etc/gingertechengine/gogios.toml.

### Telegram

If you so desire, you can follow the development of this project through Telegram here:

https://t.me/bkasin_gogios
