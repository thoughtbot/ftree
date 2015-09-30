ftree
=====

Convert lists of files into `tree(1)`-like output.

Usage
-----

```shell
❯ git ls-files | ftree
├── README.md
└── ftree.go

❯ lsrc | cut -f 1 -d : | ftree
└── Users/gordon
    ├── .agignore
    ├── .config
    │   ├── git
    │   │   ├── config
    │   │   └── ignore
    │   └── zsh
    │       ├── aliases.zsh
    │       ├── completion.zsh
    │       └── prompt.zsh
    ├── .gemrc
    ├── .local/bin
    │   ├── addcontact
    │   ├── git-review
    │   ├── keychain.py
    │   └── tat
    ├── .muttrc
    ├── .rcrc
    ├── .vimrc
    └── .zshrc

❯ find ~/code/jsbox/{app,spec}/models | ftree
└── Users/caleb/code/jsbox
    ├── app/models
    │   ├── .keep
    │   ├── analytics.rb
    │   ├── concerns/.keep
    │   ├── execution_record.rb
    │   ├── guest.rb
    │   ├── organization.rb
    │   ├── script.rb
    │   ├── script_runner.rb
    │   ├── subscription.rb
    │   └── user.rb
    └── spec/models
        ├── script_runner_spec.rb
        ├── script_spec.rb
        ├── subscription_spec.rb
        └── user_spec.rb
```
