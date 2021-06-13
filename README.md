# AWSW: AWS Switch

Interactively switch between AWS Profiles

![image](https://user-images.githubusercontent.com/12487206/121818691-4f2e3d80-cca6-11eb-98c6-2e522b8f4025.png)


## Setup

    curl -O https://raw.githubusercontent.com/bharathalleni/awsw/main/_awswitch
    curl -O https://raw.githubusercontent.com/bharathalleni/awsw/main/_awsw
    chmod +x ./_awsw*
    sudo mv ./_awsw* /usr/local/bin/

Add the following line to your `.bashrc` or `.zshrc` file

    alias awsw="source _awsw" && source _awsw reset
or

    echo 'alias awsw="source _awsw" && source _awsw reset' >> ~/.bashrc
## Usage

    awsw

