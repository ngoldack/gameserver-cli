FROM gitpod/workspace-full

# Install zsh
# RUN sudo apt-get update && sudo apt-get install -y zsh

# Install .dotfiles
RUN git clone https://github.com/ngoldack/.dotfiles && cd .dotfiles && ./install

# Install gems
RUN gem install colorls
