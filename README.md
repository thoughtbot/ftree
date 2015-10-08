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

❯ find ~/code/jsbox/{app,spec}/models
/Users/caleb/code/jsbox/app/models
/Users/caleb/code/jsbox/app/models/.keep
/Users/caleb/code/jsbox/app/models/analytics.rb
/Users/caleb/code/jsbox/app/models/concerns
/Users/caleb/code/jsbox/app/models/concerns/.keep
/Users/caleb/code/jsbox/app/models/execution_record.rb
/Users/caleb/code/jsbox/app/models/guest.rb
/Users/caleb/code/jsbox/app/models/organization.rb
/Users/caleb/code/jsbox/app/models/script.rb
/Users/caleb/code/jsbox/app/models/script_runner.rb
/Users/caleb/code/jsbox/app/models/subscription.rb
/Users/caleb/code/jsbox/app/models/user.rb
/Users/caleb/code/jsbox/spec/models
/Users/caleb/code/jsbox/spec/models/script_runner_spec.rb
/Users/caleb/code/jsbox/spec/models/script_spec.rb
/Users/caleb/code/jsbox/spec/models/subscription_spec.rb
/Users/caleb/code/jsbox/spec/models/user_spec.rb

❯ !! | ftree
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

Installation
------------

See [/dist](/dist).
