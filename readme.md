# Munchain deployment script

```
sudo apt update
sudo apt upgrade -y
sudo apt install build-essential jq -y
```

## Install Golang:

## Install latest go version https://golang.org/doc/install
```
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.18
source ~/.profile
```

## to verify that Golang installed
```
go version
```
// Should return go version go1.18 linux/amd64

## Install the executables

```
sudo rm -rf ~/.mun
make install

clear

mkdir -p ~/.mun/upgrade_manager/upgrades
mkdir -p ~/.mun/upgrade_manager/genesis/bin
```

## symlink genesis binary to upgrade
```
cp $(which mund) ~/.mun/upgrade_manager/genesis/bin
sudo cp $(which mund-manager) /usr/bin
```

## Initialize the validator, where "validator" is a moniker name
```
mund init validator --chain-id test
```
 
## Validator
// mun17zc58s96rxj79jtqqsnzt3wtx3tern6areu43g
```
echo "pet apart myth reflect stuff force attract taste caught fit exact ice slide sheriff state since unusual gaze practice course mesh magnet ozone purchase" | mund keys add validator --keyring-backend test --recover
```

## Validator1
// mun14u53eghrurpeyx5cm47vm3qwugtmhcpnstfx9t
```
echo "bottom soccer blue sniff use improve rough use amateur senior transfer quarter" | mund keys add validator1 --keyring-backend test --recover
```

## Test 1
// mun1dfjns5lk748pzrd79z4zp9k22mrchm2a5t2f6u
```
echo "betray theory cargo way left cricket doll room donkey wire reunion fall left surprise hamster corn village happy bulb token artist twelve whisper expire" | mund keys add test1 --keyring-backend test --recover
```

## Add genesis accounts
```
mund add-genesis-account $(mund keys show validator -a --keyring-backend test) 90000000000000utmun
mund add-genesis-account $(mund keys show validator1 -a --keyring-backend test) 30000000000000utmun
mund add-genesis-account $(mund keys show test1 -a --keyring-backend test) 50000000000000utmun
```

## Generate CreateValidator signed transaction
```
mund gentx validator 20000000000000utmun --keyring-backend test --chain-id test
```

## Collect genesis transactions
```
mund collect-gentxs
```

## replace stake to TMUN
```
sed -i 's/stake/utmun/g' ~/.mun/config/genesis.json
```

## Create the service file "/etc/systemd/system/mund.service" with the following content
```
sudo nano /etc/systemd/system/mund.service
```

## paste following content
```
[Unit]
Description=mund
Requires=network-online.target
After=network-online.target

[Service]
Restart=on-failure
RestartSec=3
User=root
Group=root
Environment=DAEMON_NAME=mund
Environment=DAEMON_HOME=/root/.mun
Environment=DAEMON_ALLOW_DOWNLOAD_BINARIES=on
Environment=DAEMON_RESTART_AFTER_UPGRADE=on
PermissionsStartOnly=true
ExecStart=/usr/bin/mund-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657"
StandardOutput=file:/var/log/mund/mund.log
StandardError=file:/var/log/mund/mund_error.log
ExecReload=/bin/kill -HUP $MAINPID
KillSignal=SIGTERM
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```


## Create log files for mund
```
make log-files

sudo systemctl enable mund
sudo systemctl start mund
```
# How to add a new node to munchain
```
sudo apt update
sudo apt upgrade -y
sudo apt install build-essential jq -y
```

## Install Golang:

## Install latest go version https://golang.org/doc/install
```
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh | bash -s -- --version 1.18
source ~/.profile
```

## to verify that Golang installed
```
go version
```
// Should return go version go1.18 linux/amd64

## Install the executables

```
sudo rm -rf ~/.mun
make install

clear

mkdir -p ~/.mun/upgrade_manager/upgrades
mkdir -p ~/.mun/upgrade_manager/genesis/bin
```

## symlink genesis binary to upgrade
```
cp $(which mund) ~/.mun/upgrade_manager/genesis/bin
sudo cp $(which mund-manager) /usr/bin
```

## Initialize the validator, where "validator3" is a moniker name
```
mund init validator3 --chain-id test
```

## Validator3
```
mund keys add validator3 --keyring-backend test
```

## update genesis.json from the first node
curl http://167.99.6.48:26657/genesis? | jq ".result.genesis" > ~/.mun/config/genesis.json

## update seed in config.json to make p2p connection
seeds = "90b438c39d742bea82178e0b4548eca4dcf83395@167.99.6.48:26656"

## replace stake to TMUN
```
sed -i 's/stake/utmun/g' ~/.mun/config/genesis.json
```

## Create the service file "/etc/systemd/system/mund.service" with the following content
```
sudo nano /etc/systemd/system/mund.service
```

## paste following content
```
[Unit]
Description=mund
Requires=network-online.target
After=network-online.target

[Service]
Restart=on-failure
RestartSec=3
User=root
Group=root
Environment=DAEMON_NAME=mund
Environment=DAEMON_HOME=/root/.mun
Environment=DAEMON_ALLOW_DOWNLOAD_BINARIES=on
Environment=DAEMON_RESTART_AFTER_UPGRADE=on
PermissionsStartOnly=true
ExecStart=/usr/bin/mund-manager start --pruning="nothing" --rpc.laddr "tcp://0.0.0.0:26657"
StandardOutput=file:/var/log/mund/mund.log
StandardError=file:/var/log/mund/mund_error.log
ExecReload=/bin/kill -HUP $MAINPID
KillSignal=SIGTERM
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```


## Create log files for mund
```
make log-files

sudo systemctl enable mund
sudo systemctl start mund
```
## After buying TMUN, he can become a validator by staking it.
mund tx staking create-validator --from validator3 --moniker validator3 --pubkey $(mund tendermint show-validator) --chain-id test --keyring-backend test --amount 2000000000000000utmun --commission-max-change-rate 0.01 --commission-max-rate 0.2 --commission-rate 0.1 --min-self-delegation 1 --fees 20000utmun -y
